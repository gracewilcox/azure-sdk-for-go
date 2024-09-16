//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package shared

import (
	"context"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

// NOTE: when adding a new context key type, it likely needs to be
// added to the deny-list of key types in ContextWithDeniedValues

// CtxWithHTTPHeaderKey is used as a context key for adding/retrieving http.Header.
type CtxWithHTTPHeaderKey struct{}

// CtxWithRetryOptionsKey is used as a context key for adding/retrieving RetryOptions.
type CtxWithRetryOptionsKey struct{}

// CtxWithCaptureResponse is used as a context key for retrieving the raw response.
type CtxWithCaptureResponse struct{}

// Delay waits for the duration to elapse or the context to be cancelled.
func Delay(ctx context.Context, delay time.Duration) error {
	select {
	case <-time.After(delay):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

type RetryAfterOptions struct {
	Header string
	Units  time.Duration

	// custom is used when the regular algorithm failed and is optional.
	// the returned duration is used verbatim (units is not applied).
	Custom func(string) time.Duration
}

// RetryAfter returns non-zero if the response contains one of the headers with a "retry after" value.
// Headers are checked in the following order: retry-after
func RetryAfter(resp *http.Response, retryData []RetryAfterOptions) time.Duration {
	if resp == nil {
		return 0
	}

	// the headers are listed in order of preference
	retries := retryData

	for _, retry := range retries {
		v := resp.Header.Get(retry.Header)
		if v == "" {
			continue
		}
		if retryAfter, _ := strconv.Atoi(v); retryAfter > 0 {
			return time.Duration(retryAfter) * retry.Units
		} else if d := retry.Custom(v); d > 0 {
			return d
		}
	}

	return 0
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

// ContextWithDeniedValues wraps an existing [context.Context], denying access to certain context values.
// Pipeline policies that create new requests to be sent down their own pipeline MUST wrap the caller's
// context with an instance of this type. This is to prevent context values from flowing across disjoint
// requests which can have unintended side-effects.
type ContextWithDeniedValues struct {
	context.Context

	deniedValues []any
}

func NewContextWithDeniedValues(ctx context.Context, deniedValues []any) *ContextWithDeniedValues {
	defaults := []any{CtxWithCaptureResponse{}, CtxWithHTTPHeaderKey{}, CtxWithRetryOptionsKey{}}
	return &ContextWithDeniedValues{
		Context:      ctx,
		deniedValues: append(defaults, deniedValues...),
	}
}

// Value implements part of the [context.Context] interface.
// It acts as a deny-list for certain context keys.
func (c ContextWithDeniedValues) Value(key any) any {
	for _, dv := range c.deniedValues {
		switch t1 := dv.(type) {
		default:
			switch t2 := key.(type) {
			default:
				// TODO: is this really better than reflect.TypeOf?
				if t1 == t2 {
					return nil
				}
			}
		}
	}
	return c.Context.Value(key)
}
