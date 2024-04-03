//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package policy

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
	"github.com/Azure/azure-sdk-for-go/sdk/tscore/policy"
)

// KEEP
// Policy represents an extensibility point for the Pipeline that can mutate the specified
// Request and react to the received Response.
type Policy = exported.Policy

// KEEP
// Transporter represents an HTTP pipeline transport used to send HTTP requests and receive responses.
type Transporter = exported.Transporter

// KEEP
// Request is an abstraction over the creation of an HTTP request as it passes through the pipeline.
// Don't use this type directly, use runtime.NewRequest() instead.
type Request = exported.Request

// KEEP- remove APIVersion and Cloud
// TODO, remove telemetry options
// ClientOptions contains optional settings for a client's pipeline.
// Instances can be shared across calls to SDK client constructors when uniform configuration is desired.
// Zero-value fields will have their specified default values applied during use.
type ClientOptions struct {
	// APIVersion overrides the default version requested of the service.
	// Set with caution as this package version has not been tested with arbitrary service versions.
	APIVersion string

	// Cloud specifies a cloud for the client. The default is Azure Public Cloud.
	Cloud cloud.Configuration

	// Logging configures the built-in logging policy.
	Logging LogOptions

	// Retry configures the built-in retry policy.
	Retry RetryOptions

	// Telemetry configures the built-in telemetry policy.
	Telemetry TelemetryOptions

	// TracingProvider configures the tracing provider.
	// It defaults to a no-op tracer.
	TracingProvider tracing.Provider

	// Transport sets the transport for HTTP requests.
	Transport Transporter

	// PerCallPolicies contains custom policies to inject into the pipeline.
	// Each policy is executed once per request.
	PerCallPolicies []Policy

	// PerRetryPolicies contains custom policies to inject into the pipeline.
	// Each policy is executed once per request, and for each retry of that request.
	PerRetryPolicies []Policy
}

// KEEP
// LogOptions configures the logging policy's behavior.
type LogOptions = policy.LogOptions

// KEEP
// RetryOptions configures the retry policy's behavior.
// Zero-value fields will have their specified default values applied during use.
// This allows for modification of a subset of fields.
type RetryOptions = policy.RetryOptions

type RetryData = policy.RetryData

// REMOVE
// TelemetryOptions configures the telemetry policy's behavior.
type TelemetryOptions struct {
	// ApplicationID is an application-specific identification string to add to the User-Agent.
	// It has a maximum length of 24 characters and must not contain any spaces.
	ApplicationID string

	// Disabled will prevent the addition of any telemetry data to the User-Agent.
	Disabled bool
}

// MAYBE- duplication between tscore and azcore, need different fields
// JEFF we will talk to charles (all of us in the meeting)
// TokenRequestOptions contain specific parameter that may be used by credentials types when attempting to get a token.
type TokenRequestOptions = exported.TokenRequestOptions

// KEEP
// BearerTokenOptions configures the bearer token policy's behavior.
type BearerTokenOptions struct {
	// AuthorizationHandler allows SDK developers to run client-specific logic when BearerTokenPolicy must authorize a request.
	// When this field isn't set, the policy follows its default behavior of authorizing every request with a bearer token from
	// its given credential.
	AuthorizationHandler AuthorizationHandler
}

// MAYBE (talk to charles)
// AuthorizationHandler allows SDK developers to insert custom logic that runs when BearerTokenPolicy must authorize a request.
type AuthorizationHandler struct {
	// OnRequest is called each time the policy receives a request. Its func parameter authorizes the request with a token
	// from the policy's given credential. Implementations that need to perform I/O should use the Request's context,
	// available from Request.Raw().Context(). When OnRequest returns an error, the policy propagates that error and doesn't
	// send the request. When OnRequest is nil, the policy follows its default behavior, authorizing the request with a
	// token from its credential according to its configuration.
	OnRequest func(*Request, func(TokenRequestOptions) error) error

	// OnChallenge is called when the policy receives a 401 response, allowing the AuthorizationHandler to re-authorize the
	// request according to an authentication challenge (the Response's WWW-Authenticate header). OnChallenge is responsible
	// for parsing parameters from the challenge. Its func parameter will authorize the request with a token from the policy's
	// given credential. Implementations that need to perform I/O should use the Request's context, available from
	// Request.Raw().Context(). When OnChallenge returns nil, the policy will send the request again. When OnChallenge is nil,
	// the policy will return any 401 response to the client.
	OnChallenge func(*Request, *http.Response, func(TokenRequestOptions) error) error
}

// KEEP
// WithCaptureResponse applies the HTTP response retrieval annotation to the parent context.
// The resp parameter will contain the HTTP response after the request has completed.
func WithCaptureResponse(parent context.Context, resp **http.Response) context.Context {
	return policy.WithCaptureResponse(parent, resp)
}

// KEEP
// WithHTTPHeader adds the specified http.Header to the parent context.
// Use this to specify custom HTTP headers at the API-call level.
// Any overlapping headers will have their values replaced with the values specified here.
func WithHTTPHeader(parent context.Context, header http.Header) context.Context {
	return policy.WithHTTPHeader(parent, header)
}

// KEEP
// WithRetryOptions adds the specified RetryOptions to the parent context.
// Use this to specify custom RetryOptions at the API-call level.
func WithRetryOptions(parent context.Context, options RetryOptions) context.Context {
	return policy.WithRetryOptions(parent, options)
}
