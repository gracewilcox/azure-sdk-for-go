//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

import (
	"github.com/gracewilcox/azure-sdk-for-go/sdk/azcore/internal/shared"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/azcore/policy"
	tspolicy "github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/sdk/policy"
)

const (
	defaultMaxRetries = 3
)

// NewRetryPolicy creates a policy object configured using the specified options.
// Pass nil to accept the default values; this is the same as passing a zero-value options.
func NewRetryPolicy(o *policy.RetryOptions) policy.Policy {
	options := policy.RetryOptions{}
	if o != nil {
		options = *o
	}

	tsoptions := shared.ConvertRetryOptions(options)
	return tspolicy.NewRetryPolicy(&tsoptions)
}
