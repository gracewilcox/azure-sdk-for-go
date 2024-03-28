//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	tspolicy "github.com/Azure/azure-sdk-for-go/sdk/tscore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/tscore/runtime"
)

// JEFF maybe use reflection for tests where structs are brittle
// JEFF brittle structs ClientOptions, PipelineOptions, ResponseError,

// KEEP- remove APIVersion? double check with Jeff
// JEFF replicate pipeline options, removing APIVersionOptions
// PipelineOptions contains Pipeline options for SDK developers
type PipelineOptions struct {
	// AllowedHeaders is the slice of headers to log with their values intact.
	// All headers not in the slice will have their values REDACTED.
	// Applies to request and response headers.
	AllowedHeaders []string

	// AllowedQueryParameters is the slice of query parameters to log with their values intact.
	// All query parameters not in the slice will have their values REDACTED.
	AllowedQueryParameters []string

	// APIVersion overrides the default version requested of the service.
	// Set with caution as this package version has not been tested with arbitrary service versions.
	APIVersion APIVersionOptions

	// PerCall contains custom policies to inject into the pipeline.
	// Each policy is executed once per request.
	PerCall []policy.Policy

	// PerRetry contains custom policies to inject into the pipeline.
	// Each policy is executed once per request, and for each retry of that request.
	PerRetry []policy.Policy

	// Tracing contains options used to configure distributed tracing.
	Tracing TracingOptions
}

// KEEP
// TracingOptions contains tracing options for SDK developers.
type TracingOptions = runtime.TracingOptions

// KEEP
// Pipeline represents a primitive for sending HTTP requests and receiving responses.
// Its behavior can be extended by specifying policies during construction.
type Pipeline = exported.Pipeline

// TODO get retry and logging policy to work
// FIX RETRY POLICY
// TODO- remove private policies that are exported in tscore
// KEEP
// JEFF in azcore, we will call tscore.NewPipeline (but tscore.NewPipeline will not contain the apiversion code)
//
//		JEFF We will instead insert the apiversion policy before calling tscore.NewPipeline
//
//	 JEFF ClientOptions will be duplicate between azcore and tscore
//
// Note, we also removed telemetry
// NewPipeline creates a pipeline from connection options, with any additional policies as specified.
// Policies from ClientOptions are placed after policies from PipelineOptions.
// The module and version parameters are used by the telemetry policy, when enabled.
func NewPipeline(module, version string, plOpts PipelineOptions, options *policy.ClientOptions) Pipeline {
	cp := policy.ClientOptions{}
	if options != nil {
		cp = *options
	}

	// add in azcore specific policies
	azcorePolicies := []policy.Policy{}
	if cp.APIVersion != "" {
		azcorePolicies = append(azcorePolicies, newAPIVersionPolicy(cp.APIVersion, &plOpts.APIVersion))
	}
	if !cp.Telemetry.Disabled {
		azcorePolicies = append(azcorePolicies, NewTelemetryPolicy(module, version, &cp.Telemetry))
	}

	tsPlOpts := runtime.PipelineOptions{
		AllowedHeaders:         plOpts.AllowedHeaders,
		AllowedQueryParameters: plOpts.AllowedQueryParameters,
		PerCall:                append(azcorePolicies, plOpts.PerCall...),
		PerRetry:               plOpts.PerRetry,
		Tracing:                plOpts.Tracing,
	}
	tsOptions := &tspolicy.ClientOptions{
		Logging:          cp.Logging,
		Retry:            cp.Retry,
		TracingProvider:  options.TracingProvider,
		Transport:        cp.Transport,
		PerCallPolicies:  cp.PerCallPolicies,
		PerRetryPolicies: cp.PerRetryPolicies,
	}
	return runtime.NewPipeline(tsPlOpts, tsOptions)
}
