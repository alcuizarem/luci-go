// Copyright 2016 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package tokenclient

import (
	"crypto/x509"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/luci/luci-go/common/api/tokenserver/v1"
	"github.com/luci/luci-go/common/clock"
	"github.com/luci/luci-go/common/clock/testclock"
	"github.com/luci/luci-go/common/proto/google"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTokenClient(t *testing.T) {
	Convey("works", t, func() {
		ctx := context.Background()
		ctx, _ = testclock.UseTime(ctx, time.Date(2015, time.February, 3, 4, 5, 6, 7, time.UTC))

		expectedResp := &tokenserver.TokenResponse{
			TokenType: &tokenserver.TokenResponse_GoogleOauth2AccessToken{
				GoogleOauth2AccessToken: &tokenserver.OAuth2AccessToken{
					AccessToken: "blah",
				},
			},
		}

		c := Client{
			Client: &fakeRPCClient{
				Out: tokenserver.MintTokenResponse{TokenResponse: expectedResp},
			},
			Signer: &fakeSigner{},
		}

		resp, err := c.MintToken(ctx, &tokenserver.TokenRequest{
			TokenType:    tokenserver.TokenRequest_GOOGLE_OAUTH2_ACCESS_TOKEN,
			Oauth2Scopes: []string{"scope1", "scope2"},
		})
		So(err, ShouldBeNil)
		So(resp, ShouldResemble, expectedResp)

		rpc := c.Client.(*fakeRPCClient).In
		So(rpc.Signature, ShouldResemble, []byte("fake signature"))

		tokReq := tokenserver.TokenRequest{}
		So(proto.Unmarshal(rpc.SerializedTokenRequest, &tokReq), ShouldBeNil)
		So(tokReq, ShouldResemble, tokenserver.TokenRequest{
			Certificate:        []byte("fake certificate"),
			SignatureAlgorithm: tokenserver.TokenRequest_SHA256_RSA_ALGO,
			IssuedAt:           google.NewTimestamp(clock.Now(ctx)),
			TokenType:          tokenserver.TokenRequest_GOOGLE_OAUTH2_ACCESS_TOKEN,
			Oauth2Scopes:       []string{"scope1", "scope2"},
		})
	})

	Convey("handles error", t, func() {
		ctx := context.Background()

		c := Client{
			Client: &fakeRPCClient{
				Out: tokenserver.MintTokenResponse{
					ErrorCode:    1234,
					ErrorMessage: "blah",
				},
			},
			Signer: &fakeSigner{},
		}

		_, err := c.MintToken(ctx, &tokenserver.TokenRequest{
			TokenType:    tokenserver.TokenRequest_GOOGLE_OAUTH2_ACCESS_TOKEN,
			Oauth2Scopes: []string{"scope1", "scope2"},
		})
		So(err.Error(), ShouldEqual, "token server error 1234 - blah")
	})
}

// fakeRPCClient implements tokenserver.TokenMinterClient.
type fakeRPCClient struct {
	In  tokenserver.MintTokenRequest
	Out tokenserver.MintTokenResponse
}

func (f *fakeRPCClient) MintToken(ctx context.Context, in *tokenserver.MintTokenRequest, opts ...grpc.CallOption) (*tokenserver.MintTokenResponse, error) {
	f.In = *in
	return &f.Out, nil
}

// fakeSigner implements Signer.
type fakeSigner struct{}

func (f *fakeSigner) Algo(ctx context.Context) (x509.SignatureAlgorithm, error) {
	return x509.SHA256WithRSA, nil
}

func (f *fakeSigner) Certificate(ctx context.Context) ([]byte, error) {
	return []byte("fake certificate"), nil
}

func (f *fakeSigner) Sign(ctx context.Context, blob []byte) ([]byte, error) {
	return []byte("fake signature"), nil
}