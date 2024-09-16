//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package options

import (
	"context"
	"net/http"
	"time"

	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/internal/shared"
)

// LogOptions configures the logging policy's behavior.
type LogOptions struct {
	// IncludeBody indicates if request and response bodies should be included in logging.
	// The default value is false.
	// NOTE: enabling this can lead to disclosure of sensitive information, use with care.
	IncludeBody bool

	// AllowedHeaders is the slice of headers to log with their values intact.
	// All headers not in the slice will have their values REDACTED.
	// Applies to request and response headers.
	AllowedHeaders []string

	// AllowedQueryParams is the slice of query parameters to log with their values intact.
	// All query parameters not in the slice will have their values REDACTED.
	AllowedQueryParams []string
}

// RetryOptions configures the retry policy's behavior.
// Zero-value fields will have their specified default values applied during use.
// This allows for modification of a subset of fields.
type RetryOptions struct {
	// MaxRetries specifies the maximum number of attempts a failed operation will be retried
	// before producing an error.
	// The default value is three.  A value less than zero means one try and no retries.
	MaxRetries int32

	// TryTimeout indicates the maximum time allowed for any single try of an HTTP request.
	// This is disabled by default.  Specify a value greater than zero to enable.
	// NOTE: Setting this to a small value might cause premature HTTP request time-outs.
	TryTimeout time.Duration

	// RetryDelay specifies the initial amount of delay to use before retrying an operation.
	// The value is used only if the HTTP response does not contain a Retry-After header.
	// The delay increases exponentially with each retry up to the maximum specified by MaxRetryDelay.
	// The default value is four seconds.  A value less than zero means no delay between retries.
	RetryDelay time.Duration

	// MaxRetryDelay specifies the maximum delay allowed before retrying an operation.
	// Typically the value is greater than or equal to the value specified in RetryDelay.
	// The default Value is 60 seconds.  A value less than zero means there is no cap.
	MaxRetryDelay time.Duration

	// StatusCodes specifies the HTTP status codes that indicate the operation should be retried.
	// A nil slice will use the following values.
	//   http.StatusRequestTimeout      408
	//   http.StatusTooManyRequests     429
	//   http.StatusInternalServerError 500
	//   http.StatusBadGateway          502
	//   http.StatusServiceUnavailable  503
	//   http.StatusGatewayTimeout      504
	// Specifying values will replace the default values.
	// Specifying an empty slice will disable retries for HTTP status codes.
	StatusCodes []int
}

type RetryAfterOptions = shared.RetryAfterOptions

// WithCaptureResponse applies the HTTP response retrieval annotation to the parent context.
// The resp parameter will contain the HTTP response after the request has completed.
func WithCaptureResponse(parent context.Context, resp **http.Response) context.Context {
	return context.WithValue(parent, shared.CtxWithCaptureResponse{}, resp)
}

// WithHTTPHeader adds the specified http.Header to the parent context.
// Use this to specify custom HTTP headers at the API-call level.
// Any overlapping headers will have their values replaced with the values specified here.
func WithHTTPHeader(parent context.Context, header http.Header) context.Context {
	return context.WithValue(parent, shared.CtxWithHTTPHeaderKey{}, header)
}

// WithRetryOptions adds the specified RetryOptions to the parent context.
// Use this to specify custom RetryOptions at the API-call level.
func WithRetryOptions(parent context.Context, options RetryOptions) context.Context {
	return context.WithValue(parent, shared.CtxWithRetryOptionsKey{}, options)
}
