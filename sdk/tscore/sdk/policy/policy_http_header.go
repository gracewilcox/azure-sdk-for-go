//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package policy

import (
	"net/http"

	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/internal/shared"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/policy"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/sdk/pipeline"
)

func NewHTTPHeaderPolicy(options *policy.HTTPHeaderOptions) pipeline.Policy {
	return policyFunc(httpHeaderPolicy)
}

// newHTTPHeaderPolicy creates a policy object that adds custom HTTP headers to a request
func httpHeaderPolicy(req *pipeline.Request) (*http.Response, error) {
	// check if any custom HTTP headers have been specified
	if header := req.Raw().Context().Value(shared.CtxWithHTTPHeaderKey{}); header != nil {
		for k, v := range header.(http.Header) {
			// use Set to replace any existing value
			// it also canonicalizes the header key
			req.Raw().Header.Set(k, v[0])
			// add any remaining values
			for i := 1; i < len(v); i++ {
				req.Raw().Header.Add(k, v[i])
			}
		}
	}
	return req.Next()
}
