//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

import (
	"time"

	"github.com/gracewilcox/azure-sdk-for-go/sdk/azcore/internal/shared"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/azcore/policy"
	sdkpolicy "github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/sdk/policy"
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

	options = addAzureToRetry(options)
	return sdkpolicy.NewRetryPolicy(&options)
}

func addAzureToRetry(o policy.RetryOptions) policy.RetryOptions {
	// setting default retries for Azure
	// tscore has its own defaults, this overrides them to be Azure specific
	if o.RetryData == nil {
		nop := func(string) time.Duration { return 0 }
		o.RetryData = []policy.RetryData{
			{
				Header: shared.HeaderRetryAfterMS,
				Units:  time.Millisecond,
				Custom: nop,
			},
			{
				Header: shared.HeaderXMSRetryAfterMS,
				Units:  time.Millisecond,
				Custom: nop,
			},
			{
				Header: shared.HeaderRetryAfter,
				Units:  time.Second,

				// retry-after values are expressed in either number of
				// seconds or an HTTP-date indicating when to try again
				Custom: func(ra string) time.Duration {
					t, err := time.Parse(time.RFC1123, ra)
					if err != nil {
						return 0
					}
					return time.Until(t)
				},
			},
		}
	}

	return o
}
