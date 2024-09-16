//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

import (
	"net/http"

	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/sdk/runtime"
)

// Payload reads and returns the response body or an error.
// On a successful read, the response body is cached.
// Subsequent reads will access the cached value.
func Payload(resp *http.Response) ([]byte, error) {
	return runtime.Payload(resp)
}

// HasStatusCode returns true if the Response's status code is one of the specified values.
func HasStatusCode(resp *http.Response, statusCodes ...int) bool {
	return runtime.HasStatusCode(resp, statusCodes...)
}

// UnmarshalAsByteArray will base-64 decode the received payload and place the result into the value pointed to by v.
func UnmarshalAsByteArray(resp *http.Response, v *[]byte, format Base64Encoding) error {
	return runtime.UnmarshalAsByteArray(resp, v, format)
}

// UnmarshalAsJSON calls json.Unmarshal() to unmarshal the received payload into the value pointed to by v.
func UnmarshalAsJSON(resp *http.Response, v any) error {
	return runtime.UnmarshalAsJSON(resp, v)
}

// UnmarshalAsXML calls xml.Unmarshal() to unmarshal the received payload into the value pointed to by v.
func UnmarshalAsXML(resp *http.Response, v any) error {
	return runtime.UnmarshalAsXML(resp, v)
}

// Drain reads the response body to completion then closes it.  The bytes read are discarded.
func Drain(resp *http.Response) {
	runtime.Drain(resp)
}

// DecodeByteArray will base-64 decode the provided string into v.
func DecodeByteArray(s string, v *[]byte, format Base64Encoding) error {
	return runtime.DecodeByteArray(s, v, format)
}
