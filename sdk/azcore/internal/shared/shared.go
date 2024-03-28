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

	"github.com/Azure/azure-sdk-for-go/sdk/tscore/runtime"
)

// NOTE: when adding a new context key type, it likely needs to be
// added to the deny-list of key types in ContextWithDeniedValues

// KEEP
// CtxWithHTTPHeaderKey is used as a context key for adding/retrieving http.Header.
type CtxWithHTTPHeaderKey struct{}

// KEEP
// CtxWithRetryOptionsKey is used as a context key for adding/retrieving RetryOptions.
type CtxWithRetryOptionsKey struct{}

// KEEP
// CtxWithCaptureResponse is used as a context key for retrieving the raw response.
type CtxWithCaptureResponse struct{}

// KEEP
// CtxWithTracingTracer is used as a context key for adding/retrieving tracing.Tracer.
type CtxWithTracingTracer struct{}

// KEEP
// CtxAPINameKey is used as a context key for adding/retrieving the API name.
type CtxAPINameKey = runtime.CtxAPINameKey

// KEEP
// Delay waits for the duration to elapse or the context to be cancelled.
func Delay(ctx context.Context, delay time.Duration) error {
	select {
	case <-time.After(delay):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// JEFF we gotta keep the headers
// JEFF talk with Joel, can we just remove x-ms?, could we keep in tscore
// JEFF probably will replicate retry policy, therefore retryAfter
// KEEP, but removing retry-after-ms??, x-ms-retry-after-ms; removing rety data list
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

// KEEP
// TypeOfT returns the type of the generic type param.
func TypeOfT[T any]() reflect.Type {
	// you can't, at present, obtain the type of
	// a type parameter, so this is the trick
	return reflect.TypeOf((*T)(nil)).Elem()
}

// KEEP
// TransportFunc is a helper to use a first-class func to satisfy the Transporter interface.
type TransportFunc func(*http.Request) (*http.Response, error)

// Do implements the Transporter interface for the TransportFunc type.
func (pf TransportFunc) Do(req *http.Request) (*http.Response, error) {
	return pf(req)
}

// REMOVE
// TODO, remove this. We don't need mod version
// ValidateModVer verifies that moduleVersion is a valid semver 2.0 string.
func ValidateModVer(moduleVersion string) error {
	modVerRegx := regexp.MustCompile(`^v\d+\.\d+\.\d+(?:-[a-zA-Z0-9_.-]+)?$`)
	if !modVerRegx.MatchString(moduleVersion) {
		return fmt.Errorf("malformed moduleVersion param value %s", moduleVersion)
	}
	return nil
}

// KEEP
// ContextWithDeniedValues wraps an existing [context.Context], denying access to certain context values.
// Pipeline policies that create new requests to be sent down their own pipeline MUST wrap the caller's
// context with an instance of this type. This is to prevent context values from flowing across disjoint
// requests which can have unintended side-effects.
type ContextWithDeniedValues struct {
	context.Context
}

// Value implements part of the [context.Context] interface.
// It acts as a deny-list for certain context keys.
func (c *ContextWithDeniedValues) Value(key any) any {
	switch key.(type) {
	case CtxAPINameKey, CtxWithCaptureResponse, CtxWithHTTPHeaderKey, CtxWithRetryOptionsKey, CtxWithTracingTracer:
		return nil
	default:
		return c.Context.Value(key)
	}
}
