//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

import (
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/tscore/runtime"
)

// KEEP
// bodyDownloadPolicy creates a policy object that downloads the response's body to a []byte.
func bodyDownloadPolicy(req *policy.Request) (*http.Response, error) {
	return runtime.BodyDownloadPolicy(req)
}
