// Copyright 2017 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package serviceaccounts

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/luci/luci-go/server/auth/identity"
	"github.com/luci/luci-go/server/auth/signing"

	"github.com/luci/luci-go/tokenserver/api"
	"github.com/luci/luci-go/tokenserver/api/admin/v1"
)

// InspectOAuthTokenGrantRPC implements admin.InspectOAuthTokenGrant method.
type InspectOAuthTokenGrantRPC struct {
	// Signer is mocked in tests.
	//
	// In prod it is gaesigner.Signer.
	Signer signing.Signer

	// Rules returns service account rules to use for the request.
	//
	// In prod it is GlobalRulesCache.Rules.
	Rules func(context.Context) (*Rules, error)
}

// InspectOAuthTokenGrant decodes the given OAuth token grant.
func (r *InspectOAuthTokenGrantRPC) InspectOAuthTokenGrant(c context.Context, req *admin.InspectOAuthTokenGrantRequest) (*admin.InspectOAuthTokenGrantResponse, error) {
	inspection, err := InspectGrant(c, r.Signer, req.Token)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}

	resp := &admin.InspectOAuthTokenGrantResponse{
		Valid:            inspection.Signed && inspection.NonExpired,
		Signed:           inspection.Signed,
		NonExpired:       inspection.NonExpired,
		InvalidityReason: inspection.InvalidityReason,
	}

	if env, _ := inspection.Envelope.(*tokenserver.OAuthTokenGrantEnvelope); env != nil {
		resp.SigningKeyId = env.KeyId
	}

	// Examine the body, even if the token is expired or unsigned. This helps to
	// debug expired or unsigned tokens...
	resp.TokenBody, _ = inspection.Body.(*tokenserver.OAuthTokenGrantBody)
	if resp.TokenBody != nil {
		rules, err := r.Rules(c)
		if err != nil {
			return nil, grpc.Errorf(codes.Internal, "failed to load service accounts rules")
		}

		// Always return the rule that matches the service account, even if the
		// token itself is not allowed by it (we check it separately below).
		if rule := rules.Rule(resp.TokenBody.ServiceAccount); rule != nil {
			resp.MatchingRule = rule.Rule
		}

		q := &RulesQuery{
			ServiceAccount: resp.TokenBody.ServiceAccount,
			Proxy:          identity.Identity(resp.TokenBody.Proxy),
			EndUser:        identity.Identity(resp.TokenBody.EndUser),
		}
		switch _, err = rules.Check(c, q); {
		case err == nil:
			resp.AllowedByRules = true
		case grpc.Code(err) == codes.Internal:
			return nil, err // a transient error when checking rules
		default: // fatal gRPC error => the rules forbid the token
			if resp.Valid {
				resp.Valid = false
				resp.InvalidityReason = "not allowed by the rules"
			}
		}
	}

	return resp, nil
}
