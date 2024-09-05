//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/gracewilcox/azure-sdk-for-go/sdk/azcore/internal/exported"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/azcore/tracing"
	sdkpolicy "github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/sdk/policy"
	"github.com/stretchr/testify/require"
)

func TestStartSpan(t *testing.T) {
	// tracing disabled
	ctxIn := context.Background()
	ctx, end := StartSpan(ctxIn, "TestStartSpan", tracing.Tracer{}, nil)
	end(nil)
	require.Equal(t, ctxIn, ctx)

	// span no error
	var startCalled bool
	var endCalled bool
	tr := tracing.NewTracer(func(ctx context.Context, spanName string, options *tracing.SpanOptions) (context.Context, tracing.Span) {
		startCalled = true
		require.EqualValues(t, "TestStartSpan", spanName)
		require.NotNil(t, options)
		require.EqualValues(t, tracing.SpanKindInternal, options.Kind)
		spanImpl := tracing.SpanImpl{
			End: func() { endCalled = true },
		}
		return ctx, tracing.NewSpan(spanImpl)
	}, nil)
	ctx, end = StartSpan(context.Background(), "TestStartSpan", tr, nil)
	end(nil)
	_, _ = startCalled, endCalled
	ctxTr := ctx.Value(sdkpolicy.CtxWithTracingTracer{})
	require.NotNil(t, ctxTr)
	_, ok := ctxTr.(tracing.Tracer)
	require.True(t, ok)
	require.True(t, startCalled)
	require.True(t, endCalled)

	// with error
	var spanStatus tracing.SpanStatus
	var errStr string
	tr = tracing.NewTracer(func(ctx context.Context, spanName string, options *tracing.SpanOptions) (context.Context, tracing.Span) {
		spanImpl := tracing.SpanImpl{
			End: func() { endCalled = true },
			SetStatus: func(ss tracing.SpanStatus, s string) {
				spanStatus = ss
				errStr = s
			},
		}
		return ctx, tracing.NewSpan(spanImpl)
	}, nil)
	_, end = StartSpan(context.Background(), "TestStartSpan", tr, nil)
	u, err := url.Parse("https://contoso.com")
	require.NoError(t, err)
	resp := &http.Response{
		Status:     "the operation failed",
		StatusCode: http.StatusBadRequest,
		Body:       io.NopCloser(strings.NewReader(`{ "error": { "code": "ErrorItFailed", "message": "it's not working" } }`)),
		Request: &http.Request{
			Method: http.MethodGet,
			URL:    u,
		},
	}
	end(exported.NewResponseError(resp))
	require.EqualValues(t, tracing.SpanStatusError, spanStatus)

	// TODO is it ok that it's not longer azcore??
	//require.Contains(t, errStr, "*azcore.ResponseError")
	require.Contains(t, errStr, "ERROR CODE: ErrorItFailed")
}
