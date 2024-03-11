//go:build go1.18
// +build go1.18

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
// newHTTPHeaderPolicy creates a policy object that adds custom HTTP headers to a request
func httpHeaderPolicy(req *policy.Request) (*http.Response, error) {
	return runtime.HttpHeaderPolicy(req)
}

// REMOVE
// WithHTTPHeader adds the specified http.Header to the parent context.
// Use this to specify custom HTTP headers at the API-call level.
// Any overlapping headers will have their values replaced with the values specified here.
// Deprecated: use [policy.WithHTTPHeader] instead.
func WithHTTPHeader(parent context.Context, header http.Header) context.Context {
	return policy.WithHTTPHeader(parent, header)
}
