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

	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/internal/mock"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/options"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/sdk/pipeline"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/sdk/policy"
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

	pl := NewPipeline(&PipelineOptions{
		PerCall:   []pipeline.Policy{defaultPerCallPolicy, customPerCallPolicy},
		PerRetry:  []pipeline.Policy{defaultPerRetryPolicy, customPerRetryPolicy},
		Transport: srv,
		Retry:     policy.RetryPolicyOptions{RetryOptions: options.RetryOptions{RetryDelay: time.Microsecond, MaxRetries: 1}},
	})
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

	pl := NewPipeline(nil)

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
	pl := NewPipeline(nil)
	client, err := NewClient(pl, nil)
	require.NoError(t, err)
	require.NotNil(t, client)
	require.NotZero(t, client.Pipeline())

	client, err = NewClient(pl, nil)
	require.NoError(t, err)
	require.NotNil(t, client)
}
