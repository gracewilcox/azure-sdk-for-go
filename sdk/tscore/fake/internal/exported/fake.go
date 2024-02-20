//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package exported

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/internal/errorinfo"
	"github.com/Azure/azure-sdk-for-go/sdk/tscore"
	"github.com/Azure/azure-sdk-for-go/sdk/tscore/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/tscore/internal/shared"
)

// Responder represents a scalar response.
type Responder[T any] struct {
	httpStatus int
	resp       T
	opts       SetResponseOptions
}

// SetResponse sets the specified value to be returned.
//   - httpStatus is the HTTP status code to be returned
//   - resp is the response to be returned
//   - o contains optional values, pass nil to accept the defaults
func (r *Responder[T]) SetResponse(httpStatus int, resp T, o *SetResponseOptions) {
	r.httpStatus = httpStatus
	r.resp = resp
	if o != nil {
		r.opts = *o
	}
}

// SetResponseOptions contains the optional values for Responder[T].SetResponse.
type SetResponseOptions struct {
	// Header contains optional HTTP headers to include in the response.
	Header http.Header
}

// GetResponse returns the response associated with the Responder.
// This function is called by the fake server internals.
func (r Responder[T]) GetResponse() T {
	return r.resp
}

// GetResponseContent returns the ResponseContent associated with the Responder.
// This function is called by the fake server internals.
func (r Responder[T]) GetResponseContent() ResponseContent {
	return ResponseContent{HTTPStatus: r.httpStatus, Header: r.opts.Header}
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// ErrorResponder represents a scalar error response.
type ErrorResponder struct {
	err error
}

// SetError sets the specified error to be returned.
// Use SetResponseError for returning an *tscore.ResponseError.
func (e *ErrorResponder) SetError(err error) {
	e.err = errorinfo.NonRetriableError(err)
}

// SetResponseError sets an *tscore.ResponseError with the specified values to be returned.
//   - httpStatus is the HTTP status code
func (e *ErrorResponder) SetResponseError(httpStatus int) {
	e.err = errorinfo.NonRetriableError(&exported.ResponseError{StatusCode: httpStatus})
}

// GetError returns the error for this responder.
// This function is called by the fake server internals.
func (e ErrorResponder) GetError(req *http.Request) error {
	if e.err == nil {
		return nil
	}

	var respErr *tscore.ResponseError
	if errors.As(e.err, &respErr) {
		// fix up the raw response
		rawResp, err := newErrorResponse(respErr.StatusCode, req)
		if err != nil {
			return errorinfo.NonRetriableError(err)
		}
		respErr.RawResponse = rawResp
	}
	return errorinfo.NonRetriableError(e.err)
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// ResponseContent is used when building the *http.Response.
// This type is used by the fake server internals.
type ResponseContent struct {
	// HTTPStatus is the HTTP status code to use in the response.
	HTTPStatus int

	// Header contains the headers from SetResponseOptions.Header to include in the HTTP response.
	Header http.Header
}

// ResponseOptions contains the optional values for NewResponse().
type ResponseOptions struct {
	// Body is the HTTP response body.
	Body io.ReadCloser

	// ContentType is the value for the Content-Type HTTP header.
	ContentType string
}

// SetResponseBody wraps body in a nop-closing bytes reader and assigned it to resp.Body.
// The Content-Type header will be added with the specified value.
func SetResponseBody(resp *http.Response, body []byte, contentType string) *http.Response {
	if l := int64(len(body)); l > 0 {
		resp.Header.Set(shared.HeaderContentType, contentType)
		resp.ContentLength = l
		resp.Body = io.NopCloser(bytes.NewReader(body))
	}
	return resp
}

// NewResponse creates a new *http.Response with the specified content and req as the response's request.
func NewResponse(content ResponseContent, req *http.Request) (*http.Response, error) {
	if content.HTTPStatus == 0 {
		return nil, errors.New("fake: no HTTP status code was specified")
	} else if content.Header == nil {
		content.Header = http.Header{}
	}
	return &http.Response{
		Body:       http.NoBody,
		Header:     content.Header,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Request:    req,
		Status:     fmt.Sprintf("%d %s", content.HTTPStatus, http.StatusText(content.HTTPStatus)),
		StatusCode: content.HTTPStatus,
	}, nil
}

func newErrorResponse(statusCode int, req *http.Request) (*http.Response, error) {
	content := ResponseContent{
		HTTPStatus: statusCode,
		Header:     http.Header{},
	}
	resp, err := NewResponse(content, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
