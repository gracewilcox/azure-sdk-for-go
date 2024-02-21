//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package shared

const (
	ContentTypeAppJSON   = "application/json"
	ContentTypeAppXML    = "application/xml"
	ContentTypeTextPlain = "text/plain"
)

const (
	HeaderAuthorization     = "Authorization"
	HeaderContentLength     = "Content-Length"
	HeaderContentType       = "Content-Type"
	HeaderLocation          = "Location"
	HeaderOperationLocation = "Operation-Location"
	HeaderRetryAfter        = "Retry-After"
	HeaderUserAgent         = "User-Agent"
	HeaderWWWAuthenticate   = "WWW-Authenticate"
)

const BearerTokenPrefix = "Bearer "

const TracingNamespaceAttrName = "az.namespace"

const (
	// Module is the name of the calling module used in telemetry data.
	Module = "tscore"

	// Version is the semantic version (see http://semver.org) of this module.
	Version = "v1.9.3"
)
