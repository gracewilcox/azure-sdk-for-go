//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// KEEP entire file
package exported

import (
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/tscore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/tscore/runtime"
)

// TODO can remove??
// Base64Encoding is usesd to specify which base-64 encoder/decoder to use when
// encoding/decoding a slice of bytes to/from a string.
// Exported as runtime.Base64Encoding
type Base64Encoding = runtime.Base64Encoding

// not exported but dependent on Request

// TODO can remove??
// PolicyFunc is a type that implements the Policy interface.
// Use this type when implementing a stateless policy as a first-class function.
type PolicyFunc func(*policy.Request) (*http.Response, error)

// Do implements the Policy interface on policyFunc.
func (pf PolicyFunc) Do(req *policy.Request) (*http.Response, error) {
	return pf(req)
}
