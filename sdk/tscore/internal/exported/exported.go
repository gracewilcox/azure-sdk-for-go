//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package exported

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

type nopCloser struct {
	io.ReadSeeker
}

func (n nopCloser) Close() error {
	return nil
}

// NopCloser returns a ReadSeekCloser with a no-op close method wrapping the provided io.ReadSeeker.
// Exported as streaming.NopCloser().
func NopCloser(rs io.ReadSeeker) io.ReadSeekCloser {
	return nopCloser{rs}
}

// HasStatusCode returns true if the Response's status code is one of the specified values.
// Exported as runtime.HasStatusCode().
func HasStatusCode(resp *http.Response, statusCodes ...int) bool {
	if resp == nil {
		return false
	}
	for _, sc := range statusCodes {
		if resp.StatusCode == sc {
			return true
		}
	}
	return false
}

// DecodeByteArray will base-64 decode the provided string into v.
// Exported as runtime.DecodeByteArray()
func DecodeByteArray(s string, v *[]byte, format Base64Encoding) error {
	if len(s) == 0 {
		return nil
	}
	payload := string(s)
	if payload[0] == '"' {
		// remove surrounding quotes
		payload = payload[1 : len(payload)-1]
	}
	switch format {
	case Base64StdFormat:
		decoded, err := base64.StdEncoding.DecodeString(payload)
		if err == nil {
			*v = decoded
			return nil
		}
		return err
	case Base64URLFormat:
		// use raw encoding as URL format should not contain any '=' characters
		decoded, err := base64.RawURLEncoding.DecodeString(payload)
		if err == nil {
			*v = decoded
			return nil
		}
		return err
	default:
		return fmt.Errorf("unrecognized byte array format: %d", format)
	}
}
