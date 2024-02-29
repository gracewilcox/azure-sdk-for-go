//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// KEEP entire file
package exported

import (
	"context"
	"net/http"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/tscore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/tscore/runtime"
)

// Base64Encoding is usesd to specify which base-64 encoder/decoder to use when
// encoding/decoding a slice of bytes to/from a string.
// Exported as runtime.Base64Encoding
type Base64Encoding = runtime.Base64Encoding

const (
	// Base64StdFormat uses base64.StdEncoding for encoding and decoding payloads.
	Base64StdFormat Base64Encoding = runtime.Base64StdFormat

	// Base64URLFormat uses base64.RawURLEncoding for encoding and decoding payloads.
	Base64URLFormat Base64Encoding = runtime.Base64URLFormat
)

// EncodeByteArray will base-64 encode the byte slice v.
// Exported as runtime.EncodeByteArray()
func EncodeByteArray(v []byte, format Base64Encoding) string {
	return runtime.EncodeByteArray(v, format)
}

// Request is an abstraction over the creation of an HTTP request as it passes through the pipeline.
// Don't use this type directly, use NewRequest() instead.
// Exported as policy.Request.
type Request = policy.Request

type opValues map[reflect.Type]interface{}

// Set adds/changes a value
func (ov opValues) set(value interface{}) {
	ov[reflect.TypeOf(value)] = value
}

// Get looks for a value set by SetValue first
func (ov opValues) get(value interface{}) bool {
	v, ok := ov[reflect.ValueOf(value).Elem().Type()]
	if ok {
		reflect.ValueOf(value).Elem().Set(reflect.ValueOf(v))
	}
	return ok
}

// NewRequest creates a new Request with the specified input.
// Exported as runtime.NewRequest().
func NewRequest(ctx context.Context, httpMethod string, endpoint string) (*Request, error) {
	return runtime.NewRequest(ctx, httpMethod, endpoint)
}

// not exported but dependent on Request

// PolicyFunc is a type that implements the Policy interface.
// Use this type when implementing a stateless policy as a first-class function.
type PolicyFunc func(*Request) (*http.Response, error)

// Do implements the Policy interface on policyFunc.
func (pf PolicyFunc) Do(req *Request) (*http.Response, error) {
	return pf(req)
}
