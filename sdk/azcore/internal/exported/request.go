//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package exported

import (
	"net/http"

	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/sdk/pipeline"
)

// not exported but dependent on Request

// PolicyFunc is a type that implements the Policy interface.
// Use this type when implementing a stateless policy as a first-class function.
type PolicyFunc func(*pipeline.Request) (*http.Response, error)

// Do implements the Policy interface on policyFunc.
func (pf PolicyFunc) Do(req *pipeline.Request) (*http.Response, error) {
	return pf(req)
}
