// Copyright 2017 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package buildstore

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/context"

	"go.chromium.org/luci/buildbucket"
	bbapi "go.chromium.org/luci/common/api/buildbucket/buildbucket/v1"
	"go.chromium.org/luci/common/data/strpair"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/logdog/common/types"
	"go.chromium.org/luci/milo/api/buildbot"
	"go.chromium.org/luci/milo/common"
	"go.chromium.org/luci/server/auth"
)

// This file implements conversion of buildbucket builds to buildbot builds.

// buildFromBuildbucket converts a buildbucket build to a buildbot build.
// If details is false, steps, text and properties are not guaranteed to be
// loaded.
//
// Does not populate OSFamily, OSVersion, Blame or SourceStamp.Changes
// fields.
func buildFromBuildbucket(c context.Context, master string, msg *bbapi.ApiCommonBuildMessage, fetchAnnotations bool) (*buildbot.Build, error) {
	var b buildbucket.Build
	if err := b.ParseMessage(msg); err != nil {
		return nil, err
	}
	num, err := buildNumber(&b)
	if err != nil {
		return nil, errors.Annotate(err, "parsing buildnumber").Err()
	}

	res := &buildbot.Build{
		Emulated:    true,
		Master:      master,
		Buildername: b.Builder,
		Number:      num,
		Results:     statusResult(b.Status),
		TimeStamp:   buildbot.Time{b.UpdateTime},
		Times: buildbot.TimeRange{
			Start:  buildbot.Time{b.StartTime},
			Finish: buildbot.Time{b.CompletionTime},
		},
		// TODO(nodir): use buildbucket access API when it is ready,
		// or, perhaps, just delete this code.
		// Note that this function should not be called on builds
		// that the requester does not have access to, in the first place.
		Internal:    b.Bucket != "luci.chromium.try",
		Finished:    b.Status.Completed(),
		Sourcestamp: &buildbot.SourceStamp{},
	}

	for _, bs := range b.BuildSets {
		if commit, ok := bs.(*buildbucket.GitilesCommit); ok {
			res.Sourcestamp.Repository = commit.RepoURL()
			res.Sourcestamp.Revision = commit.Revision
			break
		}
	}

	if fetchAnnotations && b.Status != buildbucket.StatusScheduled {
		addr, err := logLocation(&b)
		if err != nil {
			return nil, err
		}

		ann, err := fetchAnnotationProto(c, addr)
		switch {
		case err == errAnnotationNotFound:
		case err != nil:
			return nil, errors.Annotate(err, "could not load annotation proto").Err()
		default:
			res.Properties = extractProperties(ann)

			prefix, _ := addr.Path.Split()
			conv := annotationConverter{
				logdogServer:       addr.Host,
				logdogPrefix:       fmt.Sprintf("%s/%s", addr.Project, prefix),
				buildCompletedTime: b.CompletionTime,
			}
			convCtx := logging.SetField(c, "build_id", b.ID)
			conv.addSteps(convCtx, &res.Steps, ann.Substep, "")
			for i := range res.Steps {
				s := &res.Steps[i]
				s.StepNumber = i + 1
				res.Logs = append(res.Logs, s.Logs...)
				if !s.IsFinished {
					res.Currentstep = s.Name
				}
				if s.Results.Result != buildbot.Success {
					res.Text = append(res.Text, fmt.Sprintf("%s %s", s.Results.Result, s.Name))
				}
			}
		}
	}

	if len(res.Text) == 0 && b.Status == buildbucket.StatusSuccess {
		res.Text = []string{"Build successful"}
	}

	return res, nil
}

func buildbucketClient(c context.Context) (*bbapi.Service, error) {
	t, err := auth.GetRPCTransport(c, auth.AsUser)
	if err != nil {
		return nil, err
	}
	client, err := bbapi.New(&http.Client{Transport: t})
	if err != nil {
		return nil, err
	}

	settings := common.GetSettings(c)
	if settings.GetBuildbucket().GetHost() == "" {
		return nil, errors.New("missing buildbucket host in settings")
	}

	client.BasePath = fmt.Sprintf("https://%s/api/buildbucket/v1/", settings.Buildbucket.Host)
	return client, nil
}

func buildNumber(b *buildbucket.Build) (int, error) {
	address := b.Tags.Get("build_address")
	if address == "" {
		return 0, errors.Reason("no build_address", b.ID).Err()
	}

	// address format is "<bucket>/<builder>/<buildnumber>"
	parts := strings.Split(address, "/")
	if len(parts) != 3 {
		return 0, errors.Reason("unexpected build_address format, %q", address).Err()
	}

	return strconv.Atoi(parts[2])
}

func statusResult(status buildbucket.Status) buildbot.Result {
	switch status {
	case buildbucket.StatusScheduled, buildbucket.StatusStarted:
		return buildbot.NoResult
	case buildbucket.StatusSuccess:
		return buildbot.Success
	case buildbucket.StatusFailure:
		return buildbot.Failure
	case buildbucket.StatusError, buildbucket.StatusTimeout, buildbucket.StatusCancelled:
		return buildbot.Exception
	default:
		panic(errors.Reason("unexpected buildbucket status %q", status).Err())
	}
}

func logLocation(b *buildbucket.Build) (*types.StreamAddr, error) {
	swarmingTags := strpair.ParseMap(b.Tags["swarming_tag"])
	logLocation := swarmingTags.Get("log_location")
	if logLocation == "" {
		return nil, errors.New("log_location not found")
	}

	// Parse LogDog URL
	addr, err := types.ParseURL(logLocation)
	if err != nil {
		return nil, errors.Annotate(err, "could not parse LogDog stream from location").Err()
	}
	return addr, nil
}