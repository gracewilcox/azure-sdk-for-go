//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
	"github.com/Azure/azure-sdk-for-go/sdk/tscore/runtime"
)

// TODO where do these constants get used
// TODO possible extensibiblity point
const (
	attrHTTPMethod     = "http.method"
	attrHTTPURL        = "http.url"
	attrHTTPUserAgent  = "http.user_agent"
	attrHTTPStatusCode = "http.status_code"

	// REMOVE
	attrAZClientReqID  = "az.client_request_id"
	attrAZServiceReqID = "az.service_request_id"

	attrNetPeerName = "net.peer.name"
)

// KEEP
// StartSpanOptions contains the optional values for StartSpan.
type StartSpanOptions = runtime.StartSpanOptions

// KEEP
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
