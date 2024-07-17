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
// TODO can we remove??
// Pipeline represents a primitive for sending HTTP requests and receiving responses.
// Its behavior can be extended by specifying policies during construction.
// Exported as runtime.Pipeline.
//type Pipeline = runtime.Pipeline

// KEEP
// TODO figure out how to remove
// NewPipeline creates a new Pipeline object from the specified Policies.
// Not directly exported, but used as part of runtime.NewPipeline().
func NewPipeline(transport policy.Transporter, policies ...policy.Policy) runtime.Pipeline {
	return runtime.NewCustomPipeline(transport, policies...)
}
