//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package sdk

import (
	"context"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/internal/mock"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/internal/shared"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/policy"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/sdk/pipeline"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/tracing"
	"github.com/stretchr/testify/require"
)

// policy that tracks the number of times it was invoked
type countingPolicy struct {
	count    int
	callback func()
}

func (p *countingPolicy) Do(req *pipeline.Request) (*http.Response, error) {
	p.count++
	if p.callback != nil {
		p.callback()
	}
	return req.Next()
}

func TestNewPipelineCustomPolicies(t *testing.T) {
	srv, close := mock.NewServer()
	defer close()
	srv.AppendResponse(mock.WithStatusCode(http.StatusInternalServerError))
	srv.AppendResponse(mock.WithStatusCode(http.StatusOK))
	req, err := pipeline.NewRequest(context.Background(), http.MethodGet, srv.URL())
	require.NoError(t, err)

	// NewPipeline should place policies from ClientOptions (i.e. application-specified policies)
	// after policies from PipelineOptions (i.e. client default policies)
	defaultPerCallPolicy := &countingPolicy{}
	defaultPerRetryPolicy := &countingPolicy{}
	customPerCallPolicy := &countingPolicy{}
	customPerCallPolicy.callback = func() {
		require.Equal(t, 1, defaultPerCallPolicy.count)
	}
	customPerRetryPolicy := &countingPolicy{}
	customPerRetryPolicy.callback = func() {
		require.Equal(t, 1, defaultPerCallPolicy.count)
		require.Equal(t, 1, customPerCallPolicy.count)
		require.GreaterOrEqual(t, defaultPerRetryPolicy.count, 1)
	}

	pl := NewPipeline(PipelineOptions{PerCall: []pipeline.Policy{defaultPerCallPolicy}, PerRetry: []pipeline.Policy{defaultPerRetryPolicy}},
		&policy.ClientOptions{
			Transport:        srv,
			Retry:            policy.RetryOptions{RetryDelay: time.Microsecond, MaxRetries: 1},
			PerCallPolicies:  []policy.Policy{customPerCallPolicy},
			PerRetryPolicies: []policy.Policy{customPerRetryPolicy},
		},
	)
	_, err = pl.Do(req)
	require.NoError(t, err)
	require.Equal(t, 1, defaultPerCallPolicy.count)
	require.Equal(t, 1, customPerCallPolicy.count)
	require.Equal(t, 2, defaultPerRetryPolicy.count)
	require.Equal(t, 2, customPerRetryPolicy.count)
}

func TestPipelineDoConcurrent(t *testing.T) {
	srv, close := mock.NewServer()
	defer close()
	srv.SetResponse()

	pl := NewPipeline(PipelineOptions{}, nil)

	plErr := make(chan error, 1)
	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			req, err := pipeline.NewRequest(context.Background(), http.MethodGet, srv.URL())
			if err != nil {
				// test bug
				panic(err)
			}
			_, err = pl.Do(req)
			if err != nil {
				select {
				case plErr <- err:
					// set error
				default:
					// pending error
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()

	select {
	case err := <-plErr:
		t.Fatal(err)
	default:
		// no error
	}
}

func TestNewClient(t *testing.T) {
	client, err := NewClient("package.Client", "v1.0.0", runtime.PipelineOptions{}, nil)
	require.NoError(t, err)
	require.NotNil(t, client)
	require.NotZero(t, client.Pipeline())
	require.Zero(t, client.Tracer())

	client, err = NewClient("package.Client", "", runtime.PipelineOptions{}, &policy.ClientOptions{})
	require.NoError(t, err)
	require.NotNil(t, client)
}

func TestNewClientTracingEnabled(t *testing.T) {
	srv, close := mock.NewServer()
	defer close()

	var attrString string
	client, err := NewClient("package.Client", "v1.0.0", runtime.PipelineOptions{
		Tracing: runtime.TracingOptions{
			Namespace: "Widget.Factory",
		},
	}, &policy.ClientOptions{
		TracingProvider: tracing.NewProvider(func(name, version string) tracing.Tracer {
			return tracing.NewTracer(func(ctx context.Context, spanName string, options *tracing.SpanOptions) (context.Context, tracing.Span) {
				require.NotNil(t, options)
				for _, attr := range options.Attributes {
					if attr.Key == shared.TracingNamespaceAttrName {
						v, ok := attr.Value.(string)
						require.True(t, ok)
						attrString = attr.Key + ":" + v
					}
				}
				return ctx, tracing.Span{}
			}, nil)
		}, nil),
		Transport: srv,
	})
	require.NoError(t, err)
	require.NotNil(t, client)
	require.NotZero(t, client.Pipeline())
	require.NotZero(t, client.Tracer())

	const requestEndpoint = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/fakeResourceGroupo/providers/Microsoft.Storage/storageAccounts/fakeAccountName"
	req, err := pipeline.NewRequest(context.WithValue(context.Background(), shared.CtxWithTracingTracer{}, client.Tracer()), http.MethodGet, srv.URL()+requestEndpoint)
	require.NoError(t, err)
	srv.AppendResponse()
	_, err = client.Pipeline().Do(req)
	require.NoError(t, err)
	require.EqualValues(t, "namespace:Widget.Factory", attrString)
}

func TestClientWithClientName(t *testing.T) {
	srv, close := mock.NewServer()
	defer close()

	var clientName string
	var modVersion string
	var attrString string
	client, err := NewClient("module", "v1.0.0", runtime.PipelineOptions{
		Tracing: runtime.TracingOptions{
			Namespace: "Widget.Factory",
		},
	}, &policy.ClientOptions{
		TracingProvider: tracing.NewProvider(func(name, version string) tracing.Tracer {
			clientName = name
			modVersion = version
			return tracing.NewTracer(func(ctx context.Context, spanName string, options *tracing.SpanOptions) (context.Context, tracing.Span) {
				require.NotNil(t, options)
				for _, attr := range options.Attributes {
					if attr.Key == shared.TracingNamespaceAttrName {
						v, ok := attr.Value.(string)
						require.True(t, ok)
						attrString = attr.Key + ":" + v
					}
				}
				return ctx, tracing.Span{}
			}, nil)
		}, nil),
		Transport: srv,
	})
	require.NoError(t, err)
	require.NotNil(t, client)
	require.NotZero(t, client.Pipeline())
	require.NotZero(t, client.Tracer())
	require.EqualValues(t, "module", clientName)
	require.EqualValues(t, "v1.0.0", modVersion)

	const requestEndpoint = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/fakeResourceGroupo/providers/Microsoft.Storage/storageAccounts/fakeAccountName"
	req, err := pipeline.NewRequest(context.WithValue(context.Background(), shared.CtxWithTracingTracer{}, client.Tracer()), http.MethodGet, srv.URL()+requestEndpoint)
	require.NoError(t, err)
	srv.SetResponse()
	_, err = client.Pipeline().Do(req)
	require.NoError(t, err)
	require.EqualValues(t, "namespace:Widget.Factory", attrString)
}
