//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

import (
	"context"
	"io"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/gracewilcox/azure-sdk-for-go/sdk/azcore/internal/shared"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/runtime"
)

const (
	defaultMaxRetries = 3
)

func setDefaults(o *policy.RetryOptions) {
	if o.MaxRetries == 0 {
		o.MaxRetries = defaultMaxRetries
	} else if o.MaxRetries < 0 {
		o.MaxRetries = 0
	}

	// SDK guidelines specify the default MaxRetryDelay is 60 seconds
	if o.MaxRetryDelay == 0 {
		o.MaxRetryDelay = 60 * time.Second
	} else if o.MaxRetryDelay < 0 {
		// not really an unlimited cap, but sufficiently large enough to be considered as such
		o.MaxRetryDelay = math.MaxInt64
	}
	if o.RetryDelay == 0 {
		o.RetryDelay = 800 * time.Millisecond
	} else if o.RetryDelay < 0 {
		o.RetryDelay = 0
	}
	if o.StatusCodes == nil {
		// NOTE: if you change this list, you MUST update the docs in policy/policy.go
		o.StatusCodes = []int{
			http.StatusRequestTimeout,      // 408
			http.StatusTooManyRequests,     // 429
			http.StatusInternalServerError, // 500
			http.StatusBadGateway,          // 502
			http.StatusServiceUnavailable,  // 503
			http.StatusGatewayTimeout,      // 504
		}
	}
}

func calcDelay(o policy.RetryOptions, try int32) time.Duration { // try is >=1; never 0
	delay := time.Duration((1<<try)-1) * o.RetryDelay

	// Introduce some jitter:  [0.0, 1.0) / 2 = [0.0, 0.5) + 0.8 = [0.8, 1.3)
	delay = time.Duration(delay.Seconds() * (rand.Float64()/2 + 0.8) * float64(time.Second)) // NOTE: We want math/rand; not crypto/rand
	if delay > o.MaxRetryDelay {
		delay = o.MaxRetryDelay
	}
	return delay
}

// NewRetryPolicy creates a policy object configured using the specified options.
// Pass nil to accept the default values; this is the same as passing a zero-value options.
func NewRetryPolicy(o *policy.RetryOptions) policy.Policy {
	options := policy.RetryOptions{}
	if o != nil {
		options = *o
	}

	options = addAzureToRetry(options)
	return runtime.NewRetryPolicy(&options)
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

// ********** The following type/methods implement the retryableRequestBody (a ReadSeekCloser)

// This struct is used when sending a body to the network
type retryableRequestBody struct {
	body io.ReadSeeker // Seeking is required to support retries
}

// Read reads a block of data from an inner stream and reports progress
func (b *retryableRequestBody) Read(p []byte) (n int, err error) {
	return b.body.Read(p)
}

func (b *retryableRequestBody) Seek(offset int64, whence int) (offsetFromStart int64, err error) {
	return b.body.Seek(offset, whence)
}

func (b *retryableRequestBody) Close() error {
	// We don't want the underlying transport to close the request body on transient failures so this is a nop.
	// The retry policy closes the request body upon success.
	return nil
}

func (b *retryableRequestBody) realClose() error {
	if c, ok := b.body.(io.Closer); ok {
		return c.Close()
	}
	return nil
}

// ********** The following type/methods implement the contextCancelReadCloser

// contextCancelReadCloser combines an io.ReadCloser with a cancel func.
// it ensures the cancel func is invoked once the body has been read and closed.
type contextCancelReadCloser struct {
	cf   context.CancelFunc
	body io.ReadCloser
}

func (rc *contextCancelReadCloser) Read(p []byte) (n int, err error) {
	return rc.body.Read(p)
}

func (rc *contextCancelReadCloser) Close() error {
	err := rc.body.Close()
	rc.cf()
	return err
}
