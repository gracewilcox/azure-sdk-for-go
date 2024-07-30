//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

import (
	"net/http"

	"github.com/gracewilcox/azure-sdk-for-go/sdk/azcore/policy"
	tscontext "github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/context"
)

// includeResponsePolicy creates a policy that retrieves the raw HTTP response upon request
func includeResponsePolicy(req *policy.Request) (*http.Response, error) {
	resp, err := req.Next()
	if resp == nil {
		return resp, err
	}
	if httpOutRaw := req.Raw().Context().Value(tscontext.CtxWithCaptureResponse{}); httpOutRaw != nil {
		httpOut := httpOutRaw.(**http.Response)
		*httpOut = resp
	}
	return resp, err
}
