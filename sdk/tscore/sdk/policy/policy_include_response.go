//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package policy

import (
	"net/http"

	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/internal/shared"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/policy"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/sdk/pipeline"
)

func NewIncludeResponsePolicy(options *policy.IncludeResponseOptions) pipeline.Policy {
	return policyFunc(includeResponsePolicy)
}

// includeResponsePolicy creates a policy that retrieves the raw HTTP response upon request
func includeResponsePolicy(req *pipeline.Request) (*http.Response, error) {
	resp, err := req.Next()
	if resp == nil {
		return resp, err
	}
	if httpOutRaw := req.Raw().Context().Value(shared.CtxWithCaptureResponse{}); httpOutRaw != nil {
		httpOut := httpOutRaw.(**http.Response)
		*httpOut = resp
	}
	return resp, err
}
