//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

import (
	"net/http"

	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/internal/exported"
)

// NewResponseError creates an *tscore.ResponseError from the provided HTTP response.
// Call this when a service request returns a non-successful status code.
func NewResponseError(resp *http.Response) error {
	return exported.NewResponseError(resp)
}
