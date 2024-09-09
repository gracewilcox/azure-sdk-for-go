//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package options_test

import (
	"context"
	"net/http"

	"github.com/gracewilcox/azure-sdk-for-go/sdk/tscore/options"
)

func ExampleWithCaptureResponse() {
	// options.WithCaptureResponse provides a mechanism for obtaining an API's underlying *http.Response
	var respFromCtx *http.Response
	ctx := options.WithCaptureResponse(context.TODO(), &respFromCtx)
	// make some client method call using the updated context
	// resp, err := client.SomeMethod(ctx, ...)
	// *respFromCtx contains the raw *http.Response returned during the client method call.
	// if the HTTP transport didn't return a response due to an error then *respFromCtx will be nil.
	_ = ctx
}

func ExampleWithHTTPHeader() {
	// options.WithHTTPHeader allows callers to augment API calls with custom headers
	customHeaders := http.Header{}
	customHeaders.Add("key", "value")
	ctx := options.WithHTTPHeader(context.TODO(), customHeaders)
	// make some client method call using the updated context
	// resp, err := client.SomeMethod(ctx, ...)
	// the underlying HTTP request will include the values in customHeaders
	_ = ctx
}

func ExampleWithRetryOptions() {
	// options.WithRetryOptions contains a [options.RetryOptions] that can be used to customize the retry policy on a per-API call basis
	customRetryOptions := options.RetryOptions{ /* populate with custom values */ }
	ctx := options.WithRetryOptions(context.TODO(), customRetryOptions)
	// make some client method call using the updated context
	// resp, err := client.SomeMethod(ctx, ...)
	// the behavior of the retry policy will correspond to the values provided in customRetryPolicy
	_ = ctx
}
