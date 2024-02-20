//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// Package fake provides the building blocks for fake servers.
// This includes fakes for authentication, API responses, and more.
package fake

import (
	"context"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/internal/errorinfo"
	"github.com/Azure/azure-sdk-for-go/sdk/tscore"
	"github.com/Azure/azure-sdk-for-go/sdk/tscore/fake/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/tscore/policy"
)

// TokenCredential is a fake credential that implements the tscore.TokenCredential interface.
type TokenCredential struct {
	err error
}

// SetError sets the specified error to be returned from GetToken().
// Use this to simulate an error during authentication.
func (t *TokenCredential) SetError(err error) {
	t.err = errorinfo.NonRetriableError(err)
}

// GetToken implements the tscore.TokenCredential for the TokenCredential type.
func (t *TokenCredential) GetToken(ctx context.Context, opts policy.TokenRequestOptions) (tscore.AccessToken, error) {
	if t.err != nil {
		return tscore.AccessToken{}, errorinfo.NonRetriableError(t.err)
	}
	return tscore.AccessToken{Token: "fake_token", ExpiresOn: time.Now().Add(24 * time.Hour)}, nil
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Responder represents a scalar response.
type Responder[T any] exported.Responder[T]

// SetResponse sets the specified value to be returned.
//   - httpStatus is the HTTP status code to be returned
//   - resp is the response to be returned
//   - o contains optional values, pass nil to accept the defaults
func (r *Responder[T]) SetResponse(httpStatus int, resp T, o *SetResponseOptions) {
	(*exported.Responder[T])(r).SetResponse(httpStatus, resp, o)
}

// SetResponseOptions contains the optional values for Responder[T].SetResponse.
type SetResponseOptions = exported.SetResponseOptions

/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// ErrorResponder represents a scalar error response.
type ErrorResponder exported.ErrorResponder

// SetError sets the specified error to be returned.
// Use SetResponseError for returning an *tscore.ResponseError.
func (e *ErrorResponder) SetError(err error) {
	(*exported.ErrorResponder)(e).SetError(err)
}

// SetResponseError sets an *tscore.ResponseError with the specified values to be returned.
//   - errorCode is the value to be used in the ResponseError.Code field
//   - httpStatus is the HTTP status code
func (e *ErrorResponder) SetResponseError(httpStatus int, errorCode string) {
	(*exported.ErrorResponder)(e).SetResponseError(httpStatus)
}
