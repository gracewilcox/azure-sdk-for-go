//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package streaming

import (
	"io"

	"github.com/Azure/azure-sdk-for-go/sdk/tscore/streaming"
)

// NopCloser returns a ReadSeekCloser with a no-op close method wrapping the provided io.ReadSeeker.
// In addition to adding a Close method to an io.ReadSeeker, this can also be used to wrap an
// io.ReadSeekCloser with a no-op Close method to allow explicit control of when the io.ReedSeekCloser
// has its underlying stream closed.
func NopCloser(rs io.ReadSeeker) io.ReadSeekCloser {
	return streaming.NopCloser(rs)
}

// NewRequestProgress adds progress reporting to an HTTP request's body stream.
func NewRequestProgress(body io.ReadSeekCloser, pr func(bytesTransferred int64)) io.ReadSeekCloser {
	return streaming.NewRequestProgress(body, pr)
}

// NewResponseProgress adds progress reporting to an HTTP response's body stream.
func NewResponseProgress(body io.ReadCloser, pr func(bytesTransferred int64)) io.ReadCloser {
	return streaming.NewResponseProgress(body, pr)
}

// MultipartContent contains streaming content used in multipart/form payloads.
type MultipartContent = streaming.MultipartContent
