//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package shared

const (
	// KEEP
	ContentTypeAppJSON   = "application/json"
	ContentTypeAppXML    = "application/xml"
	ContentTypeTextPlain = "text/plain"
)

const (
	HeaderAuthorization          = "Authorization"
	HeaderAuxiliaryAuthorization = "x-ms-authorization-auxiliary" // REMOVE
	HeaderAzureAsync             = "Azure-AsyncOperation"         // REMOVE
	HeaderContentLength          = "Content-Length"
	HeaderContentType            = "Content-Type"
	HeaderFakePollerStatus       = "Fake-Poller-Status" // REMOVE
	HeaderLocation               = "Location"
	HeaderOperationLocation      = "Operation-Location"
	HeaderRetryAfter             = "Retry-After"
	HeaderRetryAfterMS           = "Retry-After-Ms" // REMOVE
	HeaderUserAgent              = "User-Agent"
	HeaderWWWAuthenticate        = "WWW-Authenticate"
	HeaderXMSClientRequestID     = "x-ms-client-request-id" // REMOVE
	HeaderXMSRequestID           = "x-ms-request-id"        // REMOVE
	HeaderXMSErrorCode           = "x-ms-error-code"        // REMOVE
	HeaderXMSRetryAfterMS        = "x-ms-retry-after-ms"    // REMOVE
)

// KEEP
const BearerTokenPrefix = "Bearer "

// KEEP but rename
// JEFF ask joel what to rename to, maybe ts.namespace??
const TracingNamespaceAttrName = "az.namespace"

// KEEP but rename
const (
	// Module is the name of the calling module used in telemetry data.
	Module = "azcore"

	// Version is the semantic version (see http://semver.org) of this module.
	Version = "v1.9.3"
)
