//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package tracing_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
	"github.com/stretchr/testify/require"
)

func TestProviderZeroValues(t *testing.T) {
	pr := tracing.Provider{}
	tr := pr.NewTracer("name", "version")
	require.Zero(t, tr)
	require.False(t, tr.Enabled())
	tr.SetAttributes()
	ctx, sp := tr.Start(context.Background(), "spanName", nil)
	require.Equal(t, context.Background(), ctx)
	require.Zero(t, sp)
	sp.AddEvent("event")
	sp.End()
	sp.SetAttributes(tracing.Attribute{})
	sp.SetStatus(tracing.SpanStatusError, "boom")
	spCtx := tr.SpanFromContext(ctx)
	require.Zero(t, spCtx)
}

func TestProvider(t *testing.T) {
	var addEventCalled bool
	var endCalled bool
	var setAttributesCalled bool
	var setStatusCalled bool
	var spanFromContextCalled bool

	pr := tracing.NewProvider(func(name, version string) tracing.Tracer {
		return tracing.NewTracer(func(context.Context, string, *tracing.SpanOptions) (context.Context, tracing.Span) {
			return nil, tracing.NewSpan(tracing.SpanImpl{
				AddEvent:      func(string, ...tracing.Attribute) { addEventCalled = true },
				End:           func() { endCalled = true },
				SetAttributes: func(...tracing.Attribute) { setAttributesCalled = true },
				SetStatus:     func(tracing.SpanStatus, string) { setStatusCalled = true },
			})
		}, &tracing.TracerOptions{
			SpanFromContext: func(context.Context) tracing.Span {
				spanFromContextCalled = true
				return tracing.Span{}
			},
		})
	}, nil)
	tr := pr.NewTracer("name", "version")
	require.NotZero(t, tr)
	require.True(t, tr.Enabled())
	sp := tr.SpanFromContext(context.Background())
	require.Zero(t, sp)
	tr.SetAttributes(tracing.Attribute{Key: "some", Value: "attribute"})

	ctx, sp := tr.Start(context.Background(), "name", nil)
	require.NotEqual(t, context.Background(), ctx)
	require.NotZero(t, sp)

	sp.AddEvent("event")
	sp.End()
	sp.SetAttributes()
	sp.SetStatus(tracing.SpanStatusError, "desc")
	require.True(t, addEventCalled)
	require.True(t, endCalled)
	require.True(t, setAttributesCalled)
	require.True(t, setStatusCalled)
	require.True(t, spanFromContextCalled)
}
