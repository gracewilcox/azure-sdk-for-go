//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

import (
	"context"

	"github.com/gracewilcox/azure-sdk-for-go/sdk/azcore/internal/shared"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/azcore/tracing"
	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/runtime"
)

const (
	attrAZClientReqID  = "az.client_request_id"
	attrAZServiceReqID = "az.service_request_id"
)

// TODO REVERT TO ORIGINAL
// newHTTPTracePolicy creates a new instance of the httpTracePolicy.
//   - allowedQueryParams contains the user-specified query parameters that don't need to be redacted from the trace
func newHTTPTracePolicy(allowedQueryParams []string, o *TracingOptions) policy.Policy {
	options := TracingOptions{}
	if o != nil {
		options = *o
	}
	options = addAzureToTracing(options)
	return runtime.NewHTTPTracePolicy(allowedQueryParams, &options)
}

func addAzureToTracing(o TracingOptions) TracingOptions {
	o.RequestAttributes = map[string]string{shared.HeaderXMSClientRequestID: attrAZClientReqID}
	o.ResponseAttributes = map[string]string{shared.HeaderXMSRequestID: attrAZServiceReqID}

	return o
}

// StartSpanOptions contains the optional values for StartSpan.
type StartSpanOptions = runtime.StartSpanOptions

// StartSpan starts a new tracing span.
// You must call the returned func to terminate the span. Pass the applicable error
// if the span will exit with an error condition.
//   - ctx is the parent context of the newly created context
//   - name is the name of the span. this is typically the fully qualified name of an API ("Client.Method")
//   - tracer is the client's Tracer for creating spans
//   - options contains optional values. pass nil to accept any default values
func StartSpan(ctx context.Context, name string, tracer tracing.Tracer, options *StartSpanOptions) (context.Context, func(error)) {
	return runtime.StartSpan(ctx, name, tracer, options)
}
