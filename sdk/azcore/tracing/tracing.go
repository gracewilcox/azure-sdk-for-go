//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// Package tracing contains the definitions needed to support distributed tracing.
package tracing

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/tscore/tracing"
)

// KEEP
// ProviderOptions contains the optional values when creating a Provider.
type ProviderOptions = tracing.ProviderOptions

// KEEP
// NewProvider creates a new Provider with the specified values.
//   - newTracerFn is the underlying implementation for creating Tracer instances
//   - options contains optional values; pass nil to accept the default value
func NewProvider(newTracerFn func(name, version string) Tracer, options *ProviderOptions) Provider {
	return tracing.NewProvider(newTracerFn, options)
}

// KEEP
// Provider is the factory that creates Tracer instances.
// It defaults to a no-op provider.
type Provider = tracing.Provider

/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// KEEP
// TracerOptions contains the optional values when creating a Tracer.
type TracerOptions = tracing.TracerOptions

// KEEP
// NewTracer creates a Tracer with the specified values.
//   - newSpanFn is the underlying implementation for creating Span instances
//   - options contains optional values; pass nil to accept the default value
func NewTracer(newSpanFn func(ctx context.Context, spanName string, options *SpanOptions) (context.Context, Span), options *TracerOptions) Tracer {
	return tracing.NewTracer(newSpanFn, options)
}

// KEEP
// Tracer is the factory that creates Span instances.
type Tracer = tracing.Tracer

// KEEP
// SpanOptions contains optional settings for creating a span.
type SpanOptions = tracing.SpanOptions

/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// KEEP
// SpanImpl abstracts the underlying implementation for Span,
// allowing it to work with various tracing implementations.
// Any zero-values will have their default, no-op behavior.
type SpanImpl = tracing.SpanImpl

// KEEP
// NewSpan creates a Span with the specified implementation.
func NewSpan(impl SpanImpl) Span {
	return tracing.NewSpan(impl)
}

// KEEP
// Span is a single unit of a trace.  A trace can contain multiple spans.
// A zero-value Span provides a no-op implementation.
type Span = tracing.Span

/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// KEEP
// Attribute is a key-value pair.
type Attribute = tracing.Attribute
