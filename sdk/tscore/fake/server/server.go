//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// Package server provides runtime functionality for fake servers.
// Application code won't need to import this package.
package server

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/tscore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/tscore/fake/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/tscore/internal/errorinfo"
	azexported "github.com/Azure/azure-sdk-for-go/sdk/tscore/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/tscore/internal/shared"
)

// ResponseContent is used when building the *http.Response.
// This type is used by the fake server internals.
type ResponseContent = exported.ResponseContent

// ResponseOptions contains the optional values for NewResponse().
type ResponseOptions = exported.ResponseOptions

// NewResponse returns a *http.Response.
// This function is called by the fake server internals.
func NewResponse(content ResponseContent, req *http.Request, opts *ResponseOptions) (*http.Response, error) {
	resp, err := exported.NewResponse(content, req)
	if err != nil {
		return nil, err
	}
	if opts != nil {
		if opts.Body != nil {
			resp.Body = opts.Body
		}
		if opts.ContentType != "" {
			resp.Header.Set(shared.HeaderContentType, opts.ContentType)
		}
	}
	return resp, nil
}

// MarshalResponseAsByteArray base-64 encodes the body with the specified format and returns it in a *http.Response.
// This function is called by the fake server internals.
func MarshalResponseAsByteArray(content ResponseContent, body []byte, format azexported.Base64Encoding, req *http.Request) (*http.Response, error) {
	resp, err := exported.NewResponse(content, req)
	if err != nil {
		return nil, err
	}
	if body != nil {
		resp = exported.SetResponseBody(resp, []byte(azexported.EncodeByteArray(body, format)), shared.ContentTypeAppJSON)
	}
	return resp, nil
}

// MarshalResponseAsJSON converts the body into JSON and returns it in a *http.Response.
// This function is called by the fake server internals.
func MarshalResponseAsJSON(content ResponseContent, v any, req *http.Request) (*http.Response, error) {
	body, err := json.Marshal(v)
	if err != nil {
		return nil, errorinfo.NonRetriableError(err)
	}
	resp, err := exported.NewResponse(content, req)
	if err != nil {
		return nil, err
	}
	resp = exported.SetResponseBody(resp, body, shared.ContentTypeAppJSON)
	return resp, nil
}

// MarshalResponseAsText converts the body into text and returns it in a *http.Response.
// This function is called by the fake server internals.
func MarshalResponseAsText(content ResponseContent, body *string, req *http.Request) (*http.Response, error) {
	resp, err := exported.NewResponse(content, req)
	if err != nil {
		return nil, err
	}
	var bodyAsBytes []byte
	if body != nil {
		bodyAsBytes = []byte(*body)
	}
	resp = exported.SetResponseBody(resp, bodyAsBytes, shared.ContentTypeTextPlain)
	return resp, nil
}

// MarshalResponseAsXML converts the body into XML and returns it in a *http.Response.
// This function is called by the fake server internals.
func MarshalResponseAsXML(content ResponseContent, v any, req *http.Request) (*http.Response, error) {
	body, err := xml.Marshal(v)
	if err != nil {
		return nil, errorinfo.NonRetriableError(err)
	}
	resp, err := exported.NewResponse(content, req)
	if err != nil {
		return nil, err
	}
	resp = exported.SetResponseBody(resp, body, shared.ContentTypeAppXML)
	return resp, nil
}

// UnmarshalRequestAsByteArray base-64 decodes the body in the specified format.
// This function is called by the fake server internals.
func UnmarshalRequestAsByteArray(req *http.Request, format azexported.Base64Encoding) ([]byte, error) {
	if req.Body == nil {
		return nil, nil
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, errorinfo.NonRetriableError(err)
	}
	req.Body.Close()
	var val []byte
	if err := azexported.DecodeByteArray(string(body), &val, format); err != nil {
		return nil, errorinfo.NonRetriableError(err)
	}
	return val, nil
}

// UnmarshalRequestAsJSON unmarshalls the request body into an instance of T.
// This function is called by the fake server internals.
func UnmarshalRequestAsJSON[T any](req *http.Request) (T, error) {
	tt := *new(T)
	if req.Body == nil {
		return tt, nil
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return tt, errorinfo.NonRetriableError(err)
	}
	req.Body.Close()
	if err = json.Unmarshal(body, &tt); err != nil {
		err = errorinfo.NonRetriableError(err)
	}
	return tt, err
}

// UnmarshalRequestAsText unmarshalls the request body into a string.
// This function is called by the fake server internals.
func UnmarshalRequestAsText(req *http.Request) (string, error) {
	if req.Body == nil {
		return "", nil
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return "", errorinfo.NonRetriableError(err)
	}
	req.Body.Close()
	return string(body), nil
}

// UnmarshalRequestAsXML unmarshalls the request body into an instance of T.
// This function is called by the fake server internals.
func UnmarshalRequestAsXML[T any](req *http.Request) (T, error) {
	tt := *new(T)
	if req.Body == nil {
		return tt, nil
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return tt, errorinfo.NonRetriableError(err)
	}
	req.Body.Close()
	if err = xml.Unmarshal(body, &tt); err != nil {
		err = errorinfo.NonRetriableError(err)
	}
	return tt, err
}

// GetResponse returns the response associated with the Responder.
// This function is called by the fake server internals.
func GetResponse[T any](r fake.Responder[T]) T {
	return exported.Responder[T](r).GetResponse()
}

// GetResponseContent returns the ResponseContent associated with the Responder.
// This function is called by the fake server internals.
func GetResponseContent[T any](r fake.Responder[T]) ResponseContent {
	return exported.Responder[T](r).GetResponseContent()
}

// GetError returns the error for this responder.
// This function is called by the fake server internals.
func GetError(e fake.ErrorResponder, req *http.Request) error {
	return exported.ErrorResponder(e).GetError(req)
}
