//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package policy

import (
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/options"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/sdk/pipeline"
)

type testPipelineOptions struct {
	// Logging configures the built-in logging policy.
	Logging options.LogOptions

	// Retry configures the built-in retry policy.
	Retry RetryPolicyOptions

	// Transport sets the transport for HTTP requests.
	Transport pipeline.Transporter

	// PerCall contains custom policies to inject into the pipeline.
	// Each policy is executed once per request.
	PerCall []pipeline.Policy

	// PerRetry contains custom policies to inject into the pipeline.
	// Each policy is executed once per request, and for each retry of that request.
	PerRetry []pipeline.Policy
}

// testing the pipeline
// always specify the transport
func newTestPipeline(options *testPipelineOptions) pipeline.Pipeline {
	if options == nil {
		options = &testPipelineOptions{}
	}

	// we put the captureResponsePolicy at the very beginning so that the raw response
	// is populated with the final response (some policies might mutate the response)
	policies := []pipeline.Policy{NewCaptureResponsePolicy(nil)}
	policies = append(policies, options.PerCall...)
	policies = append(policies, NewRetryPolicy(&options.Retry))
	policies = append(policies, options.PerRetry...)
	policies = append(policies, NewHTTPHeaderPolicy(nil))
	policies = append(policies, NewLogPolicy(&options.Logging))
	policies = append(policies, NewBodyDownloadPolicy(nil))
	transport := options.Transport

	return pipeline.New(transport, policies...)
}
