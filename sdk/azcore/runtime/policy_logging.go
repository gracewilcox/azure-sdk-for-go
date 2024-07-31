//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

import (
	"github.com/gracewilcox/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/runtime"
)

// NewLogPolicy creates a request/response logging policy object configured using the specified options.
// Pass nil to accept the default values; this is the same as passing a zero-value options.
func NewLogPolicy(o *policy.LogOptions) policy.Policy {
	options := policy.LogOptions{}
	if o != nil {
		options = *o
	}

	options = addAzureToLogging(options)
	return runtime.NewLogPolicy(&options)
}

// adding azure specific headers to LogOptions.AllowedHeaders
func addAzureToLogging(o policy.LogOptions) policy.LogOptions {
	o.AllowedHeaders = append(o.AllowedHeaders, "x-ms-request-id")
	o.AllowedHeaders = append(o.AllowedHeaders, "x-ms-client-request-id")
	o.AllowedHeaders = append(o.AllowedHeaders, "x-ms-return-client-request-id")

	return o
}
