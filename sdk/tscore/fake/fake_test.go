//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package fake_test

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/tscore"
	"github.com/Azure/azure-sdk-for-go/sdk/tscore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/tscore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/tscore/policy"
	"github.com/stretchr/testify/require"
)

type widget struct {
	Name string
}

func TestNewTokenCredential(t *testing.T) {
	cred := fake.TokenCredential{}

	tk, err := cred.GetToken(context.Background(), policy.TokenRequestOptions{})
	require.NoError(t, err)
	require.NotZero(t, tk)

	myErr := errors.New("failed")
	cred.SetError(myErr)
	tk, err = cred.GetToken(context.Background(), policy.TokenRequestOptions{})
	require.ErrorIs(t, err, myErr)
	require.Zero(t, tk)
}

func TestResponder(t *testing.T) {
	respr := fake.Responder[widget]{}
	header := http.Header{}
	header.Set("one", "1")
	header.Set("two", "2")
	respr.SetResponse(http.StatusOK, widget{Name: "foo"}, &fake.SetResponseOptions{Header: header})

	req := &http.Request{}
	resp, err := server.MarshalResponseAsJSON(server.GetResponseContent(respr), server.GetResponse(respr), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, req, resp.Request)
	require.Equal(t, "1", resp.Header.Get("one"))
	require.Equal(t, "2", resp.Header.Get("two"))
	require.EqualValues(t, http.StatusOK, resp.StatusCode)
	require.EqualValues(t, "200 OK", resp.Status)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	w := widget{}
	require.NoError(t, json.Unmarshal(body, &w))
	require.Equal(t, "foo", w.Name)
}

func TestErrorResponder(t *testing.T) {
	req := &http.Request{}

	errResp := fake.ErrorResponder{}
	require.NoError(t, server.GetError(errResp, req))

	myErr := errors.New("failed")
	errResp.SetError(myErr)
	require.ErrorIs(t, server.GetError(errResp, req), myErr)

	errResp.SetResponseError(http.StatusBadRequest, "ErrorInvalidWidget")
	var respErr *tscore.ResponseError
	require.ErrorAs(t, server.GetError(errResp, req), &respErr)
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
