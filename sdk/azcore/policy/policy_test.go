//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package policy

// func TestWithCaptureResponse(t *testing.T) {
// 	var httpResp *http.Response
// 	ctx := WithCaptureResponse(context.Background(), &httpResp)
// 	require.NotNil(t, ctx)
// 	raw := ctx.Value(tscontext.CtxWithCaptureResponse{})
// 	resp, ok := raw.(**http.Response)
// 	require.True(t, ok)
// 	require.Same(t, &httpResp, resp)
// }

// func TestWithHTTPHeader(t *testing.T) {
// 	const (
// 		key = "some"
// 		val = "thing"
// 	)
// 	input := http.Header{}
// 	input.Set(key, val)
// 	ctx := WithHTTPHeader(context.Background(), input)
// 	require.NotNil(t, ctx)
// 	raw := ctx.Value(tscontext.CtxWithHTTPHeaderKey{})
// 	header, ok := raw.(http.Header)
// 	require.True(t, ok)
// 	require.EqualValues(t, val, header.Get(key))
// }

// func TestWithRetryOptions(t *testing.T) {
// 	ctx := WithRetryOptions(context.Background(), RetryOptions{
// 		MaxRetries: math.MaxInt32,
// 	})
// 	require.NotNil(t, ctx)
// 	raw := ctx.Value(tscontext.CtxWithRetryOptionsKey{})
// 	opts, ok := raw.(RetryOptions)
// 	require.True(t, ok)
// 	require.EqualValues(t, math.MaxInt32, opts.MaxRetries)
// }
