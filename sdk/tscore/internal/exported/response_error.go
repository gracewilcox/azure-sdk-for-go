//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package exported

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/tscore/internal/log"
)

// NewResponseError creates a new *ResponseError from the provided HTTP response.
// Exported as runtime.NewResponseError().
func NewResponseError(resp *http.Response) error {
	respErr := &ResponseError{
		StatusCode:  resp.StatusCode,
		RawResponse: resp,
	}
	log.Write(log.EventResponseError, respErr.Error())
	return respErr
}

// ResponseError is returned when a request is made to a service and
// the service returns a non-success HTTP status code.
// Use errors.As() to access this type in the error chain.
// Exported as tscore.ResponseError.
type ResponseError struct {
	// StatusCode is the HTTP status code as defined in https://pkg.go.dev/net/http#pkg-constants.
	StatusCode int

	// RawResponse is the underlying HTTP response.
	RawResponse *http.Response
}

// Error implements the error interface for type ResponseError.
// Note that the message contents are not contractual and can change over time.
func (e *ResponseError) Error() string {
	const separator = "--------------------------------------------------------------------------------"
	// write the request method and URL with response status code
	msg := &bytes.Buffer{}
	if e.RawResponse != nil {
		if e.RawResponse.Request != nil {
			fmt.Fprintf(msg, "%s %s://%s%s\n", e.RawResponse.Request.Method, e.RawResponse.Request.URL.Scheme, e.RawResponse.Request.URL.Host, e.RawResponse.Request.URL.Path)
		} else {
			fmt.Fprintln(msg, "Request information not available")
		}
		fmt.Fprintln(msg, separator)
		fmt.Fprintf(msg, "RESPONSE %d: %s\n", e.RawResponse.StatusCode, e.RawResponse.Status)
	} else {
		fmt.Fprintln(msg, "Missing RawResponse")
		fmt.Fprintln(msg, separator)
	}
	if e.RawResponse != nil {
		fmt.Fprintln(msg, separator)
		body, err := Payload(e.RawResponse, nil)
		if err != nil {
			// this really shouldn't fail at this point as the response
			// body is already cached (it was read in NewResponseError)
			fmt.Fprintf(msg, "Error reading response body: %v", err)
		} else if len(body) > 0 {
			if err := json.Indent(msg, body, "", "  "); err != nil {
				// failed to pretty-print so just dump it verbatim
				fmt.Fprint(msg, string(body))
			}
			// the standard library doesn't have a pretty-printer for XML
			fmt.Fprintln(msg)
		} else {
			fmt.Fprintln(msg, "Response contained no body")
		}
	}
	fmt.Fprintln(msg, separator)

	return msg.String()
}
