//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package exported

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/tscore"
	"github.com/stretchr/testify/require"
)

type widget struct {
	Name string
}

func TestResponder(t *testing.T) {
	respr := Responder[widget]{}
	header := http.Header{}
	header.Set("one", "1")
	header.Set("two", "2")
	thing := widget{Name: "foo"}
	respr.SetResponse(http.StatusOK, thing, &SetResponseOptions{Header: header})
	require.EqualValues(t, thing, respr.GetResponse())
	require.EqualValues(t, http.StatusOK, respr.GetResponseContent().HTTPStatus)
	require.EqualValues(t, header, respr.GetResponseContent().Header)
}

func TestErrorResponder(t *testing.T) {
	req := &http.Request{}

	errResp := ErrorResponder{}
	require.NoError(t, errResp.GetError(req))

	myErr := errors.New("failed")
	errResp.SetError(myErr)
	require.ErrorIs(t, errResp.GetError(req), myErr)

	errResp.SetResponseError(http.StatusBadRequest)
	var respErr *tscore.ResponseError
	require.ErrorAs(t, errResp.GetError(req), &respErr)
	require.Equal(t, http.StatusBadRequest, respErr.StatusCode)
	require.NotNil(t, respErr.RawResponse)
	require.Equal(t, req, respErr.RawResponse.Request)
}

func unmarshal[T any](resp *http.Response) (T, error) {
	var t T
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return t, err
	}
	resp.Body.Close()

	err = json.Unmarshal(body, &t)
	return t, err
}

func TestNewResponse(t *testing.T) {
	resp, err := NewResponse(ResponseContent{}, nil)
	require.Error(t, err)
	require.Nil(t, resp)

	resp, err = NewResponse(ResponseContent{HTTPStatus: http.StatusNoContent}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.EqualValues(t, http.StatusNoContent, resp.StatusCode)
	require.Empty(t, resp.Header)
}

func TestNewErrorResponse(t *testing.T) {
	resp, err := newErrorResponse(0, nil)
	require.Error(t, err)
	require.Nil(t, resp)
	resp, err = newErrorResponse(http.StatusForbidden, nil)
	require.NoError(t, err)
	require.EqualValues(t, http.StatusForbidden, resp.StatusCode)
}
