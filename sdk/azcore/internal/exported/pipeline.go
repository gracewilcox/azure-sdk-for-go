//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package exported

import (
	"github.com/Azure/azure-sdk-for-go/sdk/tscore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/tscore/runtime"
)

// KEEP
// Policy represents an extensibility point for the Pipeline that can mutate the specified
// Request and react to the received Response.
// Exported as policy.Policy.
type Policy = policy.Policy

// KEEP
// Pipeline represents a primitive for sending HTTP requests and receiving responses.
// Its behavior can be extended by specifying policies during construction.
// Exported as runtime.Pipeline.
type Pipeline = runtime.Pipeline

// KEEP
// Transporter represents an HTTP pipeline transport used to send HTTP requests and receive responses.
// Exported as policy.Transporter.
type Transporter = policy.Transporter

// KEEP
// NewPipeline creates a new Pipeline object from the specified Policies.
// Not directly exported, but used as part of runtime.NewPipeline().
func NewPipeline(transport Transporter, policies ...Policy) Pipeline {
	return runtime.NewCustomPipeline(transport, policies...)
}
