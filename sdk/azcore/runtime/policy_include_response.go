//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/tscore/runtime"
)

// KEEP
// includeResponsePolicy creates a policy that retrieves the raw HTTP response upon request
func includeResponsePolicy(req *policy.Request) (*http.Response, error) {
	return runtime.IncludeResponsePolicy(req)
}

// REMOVE
// WithCaptureResponse applies the HTTP response retrieval annotation to the parent context.
// The resp parameter will contain the HTTP response after the request has completed.
// Deprecated: use [policy.WithCaptureResponse] instead.
func WithCaptureResponse(parent context.Context, resp **http.Response) context.Context {
	return policy.WithCaptureResponse(parent, resp)
}
