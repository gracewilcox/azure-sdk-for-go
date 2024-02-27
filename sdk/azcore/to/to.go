//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package to

import "github.com/Azure/azure-sdk-for-go/sdk/tscore/to"

// KEEP
// Ptr returns a pointer to the provided value.
func Ptr[T any](v T) *T {
	return to.Ptr(v)
}

// KEEP
// SliceOfPtrs returns a slice of *T from the specified values.
func SliceOfPtrs[T any](vv ...T) []*T {
	return to.SliceOfPtrs(vv...)
}
