//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/shared"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/tscore/runtime"
)

// KEEP
// Base64Encoding is usesd to specify which base-64 encoder/decoder to use when
// encoding/decoding a slice of bytes to/from a string.
type Base64Encoding = runtime.Base64Encoding

// KEEP
const (
	// Base64StdFormat uses base64.StdEncoding for encoding and decoding payloads.
	Base64StdFormat Base64Encoding = runtime.Base64StdFormat

	// Base64URLFormat uses base64.RawURLEncoding for encoding and decoding payloads.
	Base64URLFormat Base64Encoding = runtime.Base64URLFormat
)

// KEEP
// NewRequest creates a new policy.Request with the specified input.
// The endpoint MUST be properly encoded before calling this function.
func NewRequest(ctx context.Context, httpMethod string, endpoint string) (*policy.Request, error) {
	return runtime.NewRequest(ctx, httpMethod, endpoint)
}

// KEEP
// EncodeQueryParams will parse and encode any query parameters in the specified URL.
func EncodeQueryParams(u string) (string, error) {
	return runtime.EncodeQueryParams(u)
}

// KEEP
// JoinPaths concatenates multiple URL path segments into one path,
// inserting path separation characters as required. JoinPaths will preserve
// query parameters in the root path
func JoinPaths(root string, paths ...string) string {
	return runtime.JoinPaths(root, paths...)
}

// KEEP
// EncodeByteArray will base-64 encode the byte slice v.
func EncodeByteArray(v []byte, format Base64Encoding) string {
	return runtime.EncodeByteArray(v, format)
}

// KEEP
// MarshalAsByteArray will base-64 encode the byte slice v, then calls SetBody.
// The encoded value is treated as a JSON string.
func MarshalAsByteArray(req *policy.Request, v []byte, format Base64Encoding) error {
	return runtime.MarshalAsByteArray(req, v, format)
}

// KEEP
// MarshalAsJSON calls json.Marshal() to get the JSON encoding of v then calls SetBody.
func MarshalAsJSON(req *policy.Request, v interface{}) error {
	return runtime.MarshalAsJSON(req, v)
}

// KEEP
// MarshalAsXML calls xml.Marshal() to get the XML encoding of v then calls SetBody.
func MarshalAsXML(req *policy.Request, v interface{}) error {
	return runtime.MarshalAsXML(req, v)
}

// KEEP
// SetMultipartFormData writes the specified keys/values as multi-part form
// fields with the specified value.  File content must be specified as a ReadSeekCloser.
// All other values are treated as string values.
func SetMultipartFormData(req *policy.Request, formData map[string]interface{}) error {
	return runtime.SetMultipartFormData(req, formData)
}

// KEEP
// SkipBodyDownload will disable automatic downloading of the response body.
func SkipBodyDownload(req *policy.Request) {
	runtime.SkipBodyDownload(req)
}

// KEEP
// CtxAPINameKey is used as a context key for adding/retrieving the API name.
type CtxAPINameKey = shared.CtxAPINameKey
