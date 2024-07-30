//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

import (
	"context"
	"net/http"
	"testing"

	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/internal/mock"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/internal/shared"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/policy"
	"github.com/stretchr/testify/require"
)

func TestIncludeResponse(t *testing.T) {
	var respFromCtx *http.Response
	ctx := policy.WithCaptureResponse(context.Background(), &respFromCtx)
	require.NotNil(t, ctx)
	raw := ctx.Value(shared.CtxWithCaptureResponse{})
	_, ok := raw.(**http.Response)
	require.Truef(t, ok, "unexpected type %T", raw)
	require.Nil(t, respFromCtx)
}

func TestIncludeResponsePolicy(t *testing.T) {
	srv, close := mock.NewServer()
	defer close()
	// add a generic HTTP 200 response
	srv.SetResponse()
	// include response policy is automatically added during pipeline construction
	pl := newTestPipeline(&policy.ClientOptions{Transport: srv})
	var respFromCtx *http.Response
	ctxWithResp := policy.WithCaptureResponse(context.Background(), &respFromCtx)
	req, err := NewRequest(ctxWithResp, http.MethodGet, srv.URL())
	require.NoError(t, err)
	resp, err := pl.Do(req)
	require.NoError(t, err)
	require.NotNil(t, respFromCtx)
	require.Equal(t, respFromCtx, resp)
}
