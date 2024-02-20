//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package fake_test

import (
	"errors"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/tscore/fake"
)

// Widget is a hypothetical type used in the following examples.
type Widget struct {
	ID    int
	Shape string
}

// WidgetResponse is a hypothetical type used in the following examples.
type WidgetResponse struct {
	Widget
}

// WidgetListResponse is a hypothetical type used in the following examples.
type WidgetListResponse struct {
	Widgets []Widget
}

func ExampleTokenCredential_SetError() {
	cred := fake.TokenCredential{}

	// set an error to be returned during authentication
	cred.SetError(errors.New("failed to authenticate"))
}

func ExampleResponder() {
	// for a hypothetical API GetNextWidget(context.Context) (WidgetResponse, error)

	// a Responder is used to build a scalar response
	resp := fake.Responder[WidgetResponse]{}

	// optional HTTP headers can be included in the raw response
	header := http.Header{}
	header.Set("custom-header1", "value1")
	header.Set("custom-header2", "value2")

	// here we set the instance of Widget the Responder is to return
	resp.SetResponse(http.StatusOK, WidgetResponse{
		Widget{ID: 123, Shape: "triangle"},
	}, &fake.SetResponseOptions{
		Header: header,
	})
}

func ExampleErrorResponder() {
	// an ErrorResponder is used to build an error response
	errResp := fake.ErrorResponder{}

	// use SetError to return a generic error
	errResp.SetError(errors.New("the system is down"))

	// to return an *tscore.ResponseError, use SetResponseError
	errResp.SetResponseError(http.StatusConflict, "ErrorCodeConflict")

	// ErrorResponder returns a singular error, so calling Set* APIs overwrites any previous value
}
