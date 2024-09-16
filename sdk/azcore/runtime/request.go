//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

import (
	"context"
	"net/http"

	"github.com/gracewilcox/azure-sdk-for-go/sdk/azcore/internal/shared"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/sdk/pipeline"
	sdkpolicy "github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/sdk/policy"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/sdk/runtime"
)

// Base64Encoding is usesd to specify which base-64 encoder/decoder to use when
// encoding/decoding a slice of bytes to/from a string.
type Base64Encoding = runtime.Base64Encoding

const (
	// Base64StdFormat uses base64.StdEncoding for encoding and decoding payloads.
	Base64StdFormat Base64Encoding = runtime.Base64StdFormat

	// Base64URLFormat uses base64.RawURLEncoding for encoding and decoding payloads.
	Base64URLFormat Base64Encoding = runtime.Base64URLFormat
)

// NewRequest creates a new policy.Request with the specified input.
// The endpoint MUST be properly encoded before calling this function.
func NewRequest(ctx context.Context, httpMethod string, endpoint string) (*policy.Request, error) {
	return pipeline.NewRequest(ctx, httpMethod, endpoint)
}

// NewRequestFromRequest creates a new policy.Request with an existing *http.Request
func NewRequestFromRequest(req *http.Request) (*policy.Request, error) {
	return pipeline.NewRequestFromRequest(req)
}

// EncodeQueryParams will parse and encode any query parameters in the specified URL.
// Any semicolons will automatically be escaped.
func EncodeQueryParams(u string) (string, error) {
	return runtime.EncodeQueryParams(u)
}

// JoinPaths concatenates multiple URL path segments into one path,
// inserting path separation characters as required. JoinPaths will preserve
// query parameters in the root path
func JoinPaths(root string, paths ...string) string {
	return runtime.JoinPaths(root, paths...)
}

// EncodeByteArray will base-64 encode the byte slice v.
func EncodeByteArray(v []byte, format Base64Encoding) string {
	return runtime.EncodeByteArray(v, format)
}

// MarshalAsByteArray will base-64 encode the byte slice v, then calls SetBody.
// The encoded value is treated as a JSON string.
func MarshalAsByteArray(req *policy.Request, v []byte, format Base64Encoding) error {
	return runtime.MarshalAsByteArray(req, v, format)
}

// MarshalAsJSON calls json.Marshal() to get the JSON encoding of v then calls SetBody.
func MarshalAsJSON(req *policy.Request, v any) error {
	return runtime.MarshalAsJSON(req, v)
}

// MarshalAsXML calls xml.Marshal() to get the XML encoding of v then calls SetBody.
func MarshalAsXML(req *policy.Request, v any) error {
	return runtime.MarshalAsXML(req, v)
}

// SetMultipartFormData writes the specified keys/values as multi-part form fields with the specified value.
// File content must be specified as an [io.ReadSeekCloser] or [streaming.MultipartContent].
// Byte slices will be treated as JSON. All other values are treated as string values.
func SetMultipartFormData(req *policy.Request, formData map[string]any) error {
	return runtime.SetMultipartFormData(req, formData)
}

// SkipBodyDownload will disable automatic downloading of the response body.
func SkipBodyDownload(req *policy.Request) {
	sdkpolicy.SkipBodyDownload(req)
}

// CtxAPINameKey is used as a context key for adding/retrieving the API name.
type CtxAPINameKey = shared.CtxAPINameKey

// NewUUID returns a new UUID using the RFC4122 algorithm.
func NewUUID() (string, error) {
	return runtime.NewUUID()
}
