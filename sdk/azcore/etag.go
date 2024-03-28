//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// TODO maybe we move all the wrappers to single go file?? too many files and aliases, for each package
package azcore

import (
	"github.com/Azure/azure-sdk-for-go/sdk/tscore"
)

// KEEP
// ETag is a property used for optimistic concurrency during updates
// ETag is a validator based on https://tools.ietf.org/html/rfc7232#section-2.3.2
// An ETag can be empty ("").
type ETag = tscore.ETag

// ETagAny is an ETag that represents everything, the value is "*"
const ETagAny = tscore.ETagAny
