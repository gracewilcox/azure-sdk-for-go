//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// JEFF- topic at scrum, let's talk about if we want privacy in tscore
// JEFF we can have a different approach in tscore
// JEFF maybe use slog
// JEFF list to create: brittle areas and (if we broke, things we can do)
package runtime

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/shared"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/tscore/runtime"
)

// TODO test
// KEEP but remove x-ms headers
// JEFF remove allowed headers, more facade refactoring
// JEFF add stuff to azcore client constructor
// NewLogPolicy creates a request/response logging policy object configured using the specified options.
// Pass nil to accept the default values; this is the same as passing a zero-value options.
func NewLogPolicy(o *policy.LogOptions) policy.Policy {
	options := policy.LogOptions{}
	if o != nil {
		options = *o
	}

	options = appendAzureAllowedHeaders(options)
	return runtime.NewLogPolicy(&options)
}

// adding azure specific headers to LogOptions.AllowedHeaders
func appendAzureAllowedHeaders(o policy.LogOptions) policy.LogOptions {
	o.AllowedHeaders = append(o.AllowedHeaders, "x-ms-request-id")
	o.AllowedHeaders = append(o.AllowedHeaders, "x-ms-client-request-id")
	o.AllowedHeaders = append(o.AllowedHeaders, "x-ms-return-client-request-id")

	return o
}

// returns true if the request/response body should be logged.
// this is determined by looking at the content-type header value.
func shouldLogBody(b *bytes.Buffer, contentType string) bool {
	contentType = strings.ToLower(contentType)
	if strings.HasPrefix(contentType, "text") ||
		strings.Contains(contentType, "json") ||
		strings.Contains(contentType, "xml") {
		return true
	}
	fmt.Fprintf(b, "   Skip logging body for %s\n", contentType)
	return false
}

// writes to a buffer, used for logging purposes
func writeReqBody(req *policy.Request, b *bytes.Buffer) error {
	if req.Raw().Body == nil {
		fmt.Fprint(b, "   Request contained no body\n")
		return nil
	}
	if ct := req.Raw().Header.Get(shared.HeaderContentType); !shouldLogBody(b, ct) {
		return nil
	}
	body, err := io.ReadAll(req.Raw().Body)
	if err != nil {
		fmt.Fprintf(b, "   Failed to read request body: %s\n", err.Error())
		return err
	}
	if err := req.RewindBody(); err != nil {
		return err
	}
	logBody(b, body)
	return nil
}

// writes to a buffer, used for logging purposes
func writeRespBody(resp *http.Response, b *bytes.Buffer) error {
	ct := resp.Header.Get(shared.HeaderContentType)
	if ct == "" {
		fmt.Fprint(b, "   Response contained no body\n")
		return nil
	} else if !shouldLogBody(b, ct) {
		return nil
	}
	body, err := Payload(resp)
	if err != nil {
		fmt.Fprintf(b, "   Failed to read response body: %s\n", err.Error())
		return err
	}
	if len(body) > 0 {
		logBody(b, body)
	} else {
		fmt.Fprint(b, "   Response contained no body\n")
	}
	return nil
}

func logBody(b *bytes.Buffer, body []byte) {
	fmt.Fprintln(b, "   --------------------------------------------------------------------------------")
	fmt.Fprintln(b, string(body))
	fmt.Fprintln(b, "   --------------------------------------------------------------------------------")
}
