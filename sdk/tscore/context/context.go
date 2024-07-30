//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package context

import (
	"context"
	"reflect"
	"time"

	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/internal/shared"
)

// NOTE: when adding a new context key type, it likely needs to be
// added to the deny-list of key types in ContextWithDeniedValues

// CtxWithHTTPHeaderKey is used as a context key for adding/retrieving http.Header.
type CtxWithHTTPHeaderKey = shared.CtxWithHTTPHeaderKey

// CtxWithRetryOptionsKey is used as a context key for adding/retrieving RetryOptions.
type CtxWithRetryOptionsKey = shared.CtxWithRetryOptionsKey

// CtxWithCaptureResponse is used as a context key for retrieving the raw response.
type CtxWithCaptureResponse = shared.CtxWithCaptureResponse

// CtxWithTracingTracer is used as a context key for adding/retrieving tracing.Tracer.
type CtxWithTracingTracer = shared.CtxWithTracingTracer

// CtxAPINameKey is used as a context key for adding/retrieving the API name.
type CtxAPINameKey = shared.CtxAPINameKey

// Delay waits for the duration to elapse or the context to be cancelled.
func Delay(ctx context.Context, delay time.Duration) error {
	return shared.Delay(ctx, delay)
}

// TypeOfT returns the type of the generic type param.
func TypeOfT[T any]() reflect.Type {
	return shared.TypeOfT[T]()
}

// TransportFunc is a helper to use a first-class func to satisfy the Transporter interface.
type TransportFunc = shared.TransportFunc

// ContextWithDeniedValues wraps an existing [context.Context], denying access to certain context values.
// Pipeline policies that create new requests to be sent down their own pipeline MUST wrap the caller's
// context with an instance of this type. This is to prevent context values from flowing across disjoint
// requests which can have unintended side-effects.
type ContextWithDeniedValues = shared.ContextWithDeniedValues
