//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package shared

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"time"

	tspolicy "github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/sdk/policy"
)

// NOTE: when adding a new context key type, it likely needs to be
// added to the deny-list of key types in ContextWithDeniedValues

// CtxWithTracingTracer is used as a context key for adding/retrieving tracing.Tracer.
type CtxWithTracingTracer struct{}

// CtxAPINameKey is used as a context key for adding/retrieving the API name.
type CtxAPINameKey struct{}

// TODO connect with tscore
func NewContextWithDeniedValues(ctx context.Context) {
	// defaults := []any{CtxWithCaptureResponse{}, CtxWithHTTPHeaderKey{}, CtxWithRetryOptionsKey{}}
	// return &ContextWithDeniedValues{
	// 	Context:      ctx,
	// 	deniedValues: append(defaults, deniedValues...),
	// }
}

// Delay waits for the duration to elapse or the context to be cancelled.
func Delay(ctx context.Context, delay time.Duration) error {
	select {
	case <-time.After(delay):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RetryAfter returns non-zero if the response contains one of the headers with a "retry after" value.
// Headers are checked in the following order: retry-after-ms, x-ms-retry-after-ms, retry-after
func RetryAfter(resp *http.Response) time.Duration {
	if resp == nil {
		return 0
	}

	type retryData struct {
		header string
		units  time.Duration

		// custom is used when the regular algorithm failed and is optional.
		// the returned duration is used verbatim (units is not applied).
		custom func(string) time.Duration
	}

	nop := func(string) time.Duration { return 0 }

	// the headers are listed in order of preference
	retries := []retryData{
		{
			header: HeaderRetryAfterMS,
			units:  time.Millisecond,
			custom: nop,
		},
		{
			header: HeaderXMSRetryAfterMS,
			units:  time.Millisecond,
			custom: nop,
		},
		{
			header: HeaderRetryAfter,
			units:  time.Second,

			// retry-after values are expressed in either number of
			// seconds or an HTTP-date indicating when to try again
			custom: func(ra string) time.Duration {
				t, err := time.Parse(time.RFC1123, ra)
				if err != nil {
					return 0
				}
				return time.Until(t)
			},
		},
	}

	for _, retry := range retries {
		v := resp.Header.Get(retry.header)
		if v == "" {
			continue
		}
		if retryAfter, _ := strconv.Atoi(v); retryAfter > 0 {
			return time.Duration(retryAfter) * retry.units
		} else if d := retry.custom(v); d > 0 {
			return d
		}
	}

	return 0
}

// RetryOptions configures the retry policy's behavior.
// Zero-value fields will have their specified default values applied during use.
// This allows for modification of a subset of fields.
type RetryOptions struct {
	// MaxRetries specifies the maximum number of attempts a failed operation will be retried
	// before producing an error.
	// The default value is three.  A value less than zero means one try and no retries.
	MaxRetries int32

	// TryTimeout indicates the maximum time allowed for any single try of an HTTP request.
	// This is disabled by default.  Specify a value greater than zero to enable.
	// NOTE: Setting this to a small value might cause premature HTTP request time-outs.
	TryTimeout time.Duration

	// RetryDelay specifies the initial amount of delay to use before retrying an operation.
	// The value is used only if the HTTP response does not contain a Retry-After header.
	// The delay increases exponentially with each retry up to the maximum specified by MaxRetryDelay.
	// The default value is four seconds.  A value less than zero means no delay between retries.
	RetryDelay time.Duration

	// MaxRetryDelay specifies the maximum delay allowed before retrying an operation.
	// Typically the value is greater than or equal to the value specified in RetryDelay.
	// The default Value is 60 seconds.  A value less than zero means there is no cap.
	MaxRetryDelay time.Duration

	// StatusCodes specifies the HTTP status codes that indicate the operation should be retried.
	// A nil slice will use the following values.
	//   http.StatusRequestTimeout      408
	//   http.StatusTooManyRequests     429
	//   http.StatusInternalServerError 500
	//   http.StatusBadGateway          502
	//   http.StatusServiceUnavailable  503
	//   http.StatusGatewayTimeout      504
	// Specifying values will replace the default values.
	// Specifying an empty slice will disable retries for HTTP status codes.
	StatusCodes []int

	// ShouldRetry evaluates if the retry policy should retry the request.
	// When specified, the function overrides comparison against the list of
	// HTTP status codes and error checking within the retry policy. Context
	// and NonRetriable errors remain evaluated before calling ShouldRetry.
	// The *http.Response and error parameters are mutually exclusive, i.e.
	// if one is nil, the other is not nil.
	// A return value of true means the retry policy should retry.
	ShouldRetry func(*http.Response, error) bool
}

func ConvertRetryOptions(o RetryOptions) tspolicy.RetryPolicyOptions {
	var options tspolicy.RetryPolicyOptions

	// setting default retries for Azure
	// tscore has its own defaults, this overrides them to be Azure specific
	nop := func(string) time.Duration { return 0 }
	options.RetryAfterOptions = []tspolicy.RetryAfterOptions{
		{
			Header: HeaderRetryAfterMS,
			Units:  time.Millisecond,
			Custom: nop,
		},
		{
			Header: HeaderXMSRetryAfterMS,
			Units:  time.Millisecond,
			Custom: nop,
		},
		{
			Header: HeaderRetryAfter,
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

	// converting azcore options to tscore options
	options.ShouldRetry = o.ShouldRetry
	options.MaxRetries = o.MaxRetries
	options.TryTimeout = o.TryTimeout
	options.RetryDelay = o.RetryDelay
	options.MaxRetryDelay = o.MaxRetryDelay
	options.StatusCodes = o.StatusCodes

	return options
}

// TypeOfT returns the type of the generic type param.
func TypeOfT[T any]() reflect.Type {
	// you can't, at present, obtain the type of
	// a type parameter, so this is the trick
	return reflect.TypeOf((*T)(nil)).Elem()
}

// TransportFunc is a helper to use a first-class func to satisfy the Transporter interface.
type TransportFunc func(*http.Request) (*http.Response, error)

// Do implements the Transporter interface for the TransportFunc type.
func (pf TransportFunc) Do(req *http.Request) (*http.Response, error) {
	return pf(req)
}

// ValidateModVer verifies that moduleVersion is a valid semver 2.0 string.
func ValidateModVer(moduleVersion string) error {
	modVerRegx := regexp.MustCompile(`^v\d+\.\d+\.\d+(?:-[a-zA-Z0-9_.-]+)?$`)
	if !modVerRegx.MatchString(moduleVersion) {
		return fmt.Errorf("malformed moduleVersion param value %s", moduleVersion)
	}
	return nil
}
