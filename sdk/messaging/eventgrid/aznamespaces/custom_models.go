//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package aznamespaces

// Error implements the error interface for type Error.
// Note that the message contents are not contractual and can change over time.
func (e *Error) Error() string {
	if e.Message == nil {
		return ""
	}

	return *e.Message
}
