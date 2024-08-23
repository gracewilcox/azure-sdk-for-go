//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package exported

import (
	"encoding/base64"
)

// Base64Encoding is usesd to specify which base-64 encoder/decoder to use when
// encoding/decoding a slice of bytes to/from a string.
// Exported as runtime.Base64Encoding
type Base64Encoding int

const (
	// Base64StdFormat uses base64.StdEncoding for encoding and decoding payloads.
	Base64StdFormat Base64Encoding = 0

	// Base64URLFormat uses base64.RawURLEncoding for encoding and decoding payloads.
	Base64URLFormat Base64Encoding = 1
)

// EncodeByteArray will base-64 encode the byte slice v.
// Exported as runtime.EncodeByteArray()
func EncodeByteArray(v []byte, format Base64Encoding) string {
	if format == Base64URLFormat {
		return base64.RawURLEncoding.EncodeToString(v)
	}
	return base64.StdEncoding.EncodeToString(v)
}
