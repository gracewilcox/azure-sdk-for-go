//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package pipeline

import (
	"context"
	"net/http"

	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/internal/exported"
)

// Policy represents an extensibility point for the Pipeline that can mutate the specified
// Request and react to the received Response.
type Policy = exported.Policy

// Transporter represents an HTTP pipeline transport used to send HTTP requests and receive responses.
// Exported as policy.Transporter.
type Transporter = exported.Transporter

// Pipeline represents a primitive for sending HTTP requests and receiving responses.
// Its behavior can be extended by specifying policies during construction.
type Pipeline = exported.Pipeline

// New creates a new Pipeline object from the specified Policies.
func New(transport Transporter, policies ...Policy) Pipeline {
	return exported.NewPipeline(transport, policies...)
}

// Request is an abstraction over the creation of an HTTP request as it passes through the pipeline.
// Don't use this type directly, use NewRequest() instead.
type Request = exported.Request

// NewRequestFromRequest creates a new policy.Request with an existing *http.Request
// Exported as runtime.NewRequestFromRequest().
func NewRequestFromRequest(req *http.Request) (*Request, error) {
	return exported.NewRequestFromRequest(req)
}

// NewRequest creates a new Request with the specified input.
// Exported as runtime.NewRequest().
func NewRequest(ctx context.Context, httpMethod string, endpoint string) (*Request, error) {
	return exported.NewRequest(ctx, httpMethod, endpoint)
}
