// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// EncryptionScopesClient contains the methods for the EncryptionScopes group.
// Don't use this type directly, use NewEncryptionScopesClient() instead.
type EncryptionScopesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewEncryptionScopesClient creates a new instance of EncryptionScopesClient with the specified values.
func NewEncryptionScopesClient(con *armcore.Connection, subscriptionID string) *EncryptionScopesClient {
	return &EncryptionScopesClient{con: con, subscriptionID: subscriptionID}
}

// Get - Returns the properties for the specified encryption scope.
func (client *EncryptionScopesClient) Get(ctx context.Context, resourceGroupName string, accountName string, encryptionScopeName string, options *EncryptionScopesGetOptions) (EncryptionScopeResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, encryptionScopeName, options)
	if err != nil {
		return EncryptionScopeResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return EncryptionScopeResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return EncryptionScopeResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *EncryptionScopesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, encryptionScopeName string, options *EncryptionScopesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/encryptionScopes/{encryptionScopeName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{encryptionScopeName}", url.PathEscape(encryptionScopeName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EncryptionScopesClient) getHandleResponse(resp *azcore.Response) (EncryptionScopeResponse, error) {
	var val *EncryptionScope
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return EncryptionScopeResponse{}, err
	}
	return EncryptionScopeResponse{RawResponse: resp.Response, EncryptionScope: val}, nil
}

// getHandleError handles the Get error response.
func (client *EncryptionScopesClient) getHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Lists all the encryption scopes available under the specified storage account.
func (client *EncryptionScopesClient) List(resourceGroupName string, accountName string, options *EncryptionScopesListOptions) EncryptionScopeListResultPager {
	return &encryptionScopeListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, accountName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp EncryptionScopeListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.EncryptionScopeListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *EncryptionScopesClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *EncryptionScopesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/encryptionScopes"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *EncryptionScopesClient) listHandleResponse(resp *azcore.Response) (EncryptionScopeListResultResponse, error) {
	var val *EncryptionScopeListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return EncryptionScopeListResultResponse{}, err
	}
	return EncryptionScopeListResultResponse{RawResponse: resp.Response, EncryptionScopeListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *EncryptionScopesClient) listHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Patch - Update encryption scope properties as specified in the request body. Update fails if the specified encryption scope does not already exist.
func (client *EncryptionScopesClient) Patch(ctx context.Context, resourceGroupName string, accountName string, encryptionScopeName string, encryptionScope EncryptionScope, options *EncryptionScopesPatchOptions) (EncryptionScopeResponse, error) {
	req, err := client.patchCreateRequest(ctx, resourceGroupName, accountName, encryptionScopeName, encryptionScope, options)
	if err != nil {
		return EncryptionScopeResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return EncryptionScopeResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return EncryptionScopeResponse{}, client.patchHandleError(resp)
	}
	return client.patchHandleResponse(resp)
}

// patchCreateRequest creates the Patch request.
func (client *EncryptionScopesClient) patchCreateRequest(ctx context.Context, resourceGroupName string, accountName string, encryptionScopeName string, encryptionScope EncryptionScope, options *EncryptionScopesPatchOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/encryptionScopes/{encryptionScopeName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{encryptionScopeName}", url.PathEscape(encryptionScopeName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(encryptionScope)
}

// patchHandleResponse handles the Patch response.
func (client *EncryptionScopesClient) patchHandleResponse(resp *azcore.Response) (EncryptionScopeResponse, error) {
	var val *EncryptionScope
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return EncryptionScopeResponse{}, err
	}
	return EncryptionScopeResponse{RawResponse: resp.Response, EncryptionScope: val}, nil
}

// patchHandleError handles the Patch error response.
func (client *EncryptionScopesClient) patchHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Put - Synchronously creates or updates an encryption scope under the specified storage account. If an encryption scope is already created and a subsequent
// request is issued with different properties, the
// encryption scope properties will be updated per the specified request.
func (client *EncryptionScopesClient) Put(ctx context.Context, resourceGroupName string, accountName string, encryptionScopeName string, encryptionScope EncryptionScope, options *EncryptionScopesPutOptions) (EncryptionScopeResponse, error) {
	req, err := client.putCreateRequest(ctx, resourceGroupName, accountName, encryptionScopeName, encryptionScope, options)
	if err != nil {
		return EncryptionScopeResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return EncryptionScopeResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return EncryptionScopeResponse{}, client.putHandleError(resp)
	}
	return client.putHandleResponse(resp)
}

// putCreateRequest creates the Put request.
func (client *EncryptionScopesClient) putCreateRequest(ctx context.Context, resourceGroupName string, accountName string, encryptionScopeName string, encryptionScope EncryptionScope, options *EncryptionScopesPutOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/encryptionScopes/{encryptionScopeName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{encryptionScopeName}", url.PathEscape(encryptionScopeName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(encryptionScope)
}

// putHandleResponse handles the Put response.
func (client *EncryptionScopesClient) putHandleResponse(resp *azcore.Response) (EncryptionScopeResponse, error) {
	var val *EncryptionScope
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return EncryptionScopeResponse{}, err
	}
	return EncryptionScopeResponse{RawResponse: resp.Response, EncryptionScope: val}, nil
}

// putHandleError handles the Put error response.
func (client *EncryptionScopesClient) putHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
