// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
)

// REMOVE
// WithCaptureResponse applies the HTTP response retrieval annotation to the parent context.
// The resp parameter will contain the HTTP response after the request has completed.
// Deprecated: use [policy.WithCaptureResponse] instead.
func WithCaptureResponse(parent context.Context, resp **http.Response) context.Context {
	return policy.WithCaptureResponse(parent, resp)
}

// REMOVE
// WithHTTPHeader adds the specified http.Header to the parent context.
// Use this to specify custom HTTP headers at the API-call level.
// Any overlapping headers will have their values replaced with the values specified here.
// Deprecated: use [policy.WithHTTPHeader] instead.
func WithHTTPHeader(parent context.Context, header http.Header) context.Context {
	return policy.WithHTTPHeader(parent, header)
}

// REMOVE
// WithRetryOptions adds the specified RetryOptions to the parent context.
// Use this to specify custom RetryOptions at the API-call level.
// Deprecated: use [policy.WithRetryOptions] instead.
func WithRetryOptions(parent context.Context, options policy.RetryOptions) context.Context {
	return policy.WithRetryOptions(parent, options)
}
