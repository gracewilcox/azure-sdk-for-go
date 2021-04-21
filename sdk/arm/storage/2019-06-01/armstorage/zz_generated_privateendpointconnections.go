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

// PrivateEndpointConnectionsClient contains the methods for the PrivateEndpointConnections group.
// Don't use this type directly, use NewPrivateEndpointConnectionsClient() instead.
type PrivateEndpointConnectionsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewPrivateEndpointConnectionsClient creates a new instance of PrivateEndpointConnectionsClient with the specified values.
func NewPrivateEndpointConnectionsClient(con *armcore.Connection, subscriptionID string) *PrivateEndpointConnectionsClient {
	return &PrivateEndpointConnectionsClient{con: con, subscriptionID: subscriptionID}
}

// Delete - Deletes the specified private endpoint connection associated with the storage account.
func (client *PrivateEndpointConnectionsClient) Delete(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, privateEndpointConnectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp.Response, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PrivateEndpointConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/privateEndpointConnections/{privateEndpointConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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

// deleteHandleError handles the Delete error response.
func (client *PrivateEndpointConnectionsClient) deleteHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified private endpoint connection associated with the storage account.
func (client *PrivateEndpointConnectionsClient) Get(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsGetOptions) (PrivateEndpointConnectionResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, privateEndpointConnectionName, options)
	if err != nil {
		return PrivateEndpointConnectionResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PrivateEndpointConnectionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PrivateEndpointConnectionResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateEndpointConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/privateEndpointConnections/{privateEndpointConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
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
func (client *PrivateEndpointConnectionsClient) getHandleResponse(resp *azcore.Response) (PrivateEndpointConnectionResponse, error) {
	var val *PrivateEndpointConnection
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PrivateEndpointConnectionResponse{}, err
	}
	return PrivateEndpointConnectionResponse{RawResponse: resp.Response, PrivateEndpointConnection: val}, nil
}

// getHandleError handles the Get error response.
func (client *PrivateEndpointConnectionsClient) getHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - List all the private endpoint connections associated with the storage account.
func (client *PrivateEndpointConnectionsClient) List(ctx context.Context, resourceGroupName string, accountName string, options *PrivateEndpointConnectionsListOptions) (PrivateEndpointConnectionListResultResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return PrivateEndpointConnectionListResultResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PrivateEndpointConnectionListResultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PrivateEndpointConnectionListResultResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *PrivateEndpointConnectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *PrivateEndpointConnectionsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/privateEndpointConnections"
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
func (client *PrivateEndpointConnectionsClient) listHandleResponse(resp *azcore.Response) (PrivateEndpointConnectionListResultResponse, error) {
	var val *PrivateEndpointConnectionListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PrivateEndpointConnectionListResultResponse{}, err
	}
	return PrivateEndpointConnectionListResultResponse{RawResponse: resp.Response, PrivateEndpointConnectionListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *PrivateEndpointConnectionsClient) listHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Put - Update the state of specified private endpoint connection associated with the storage account.
func (client *PrivateEndpointConnectionsClient) Put(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string, properties PrivateEndpointConnection, options *PrivateEndpointConnectionsPutOptions) (PrivateEndpointConnectionResponse, error) {
	req, err := client.putCreateRequest(ctx, resourceGroupName, accountName, privateEndpointConnectionName, properties, options)
	if err != nil {
		return PrivateEndpointConnectionResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PrivateEndpointConnectionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PrivateEndpointConnectionResponse{}, client.putHandleError(resp)
	}
	return client.putHandleResponse(resp)
}

// putCreateRequest creates the Put request.
func (client *PrivateEndpointConnectionsClient) putCreateRequest(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string, properties PrivateEndpointConnection, options *PrivateEndpointConnectionsPutOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/privateEndpointConnections/{privateEndpointConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(properties)
}

// putHandleResponse handles the Put response.
func (client *PrivateEndpointConnectionsClient) putHandleResponse(resp *azcore.Response) (PrivateEndpointConnectionResponse, error) {
	var val *PrivateEndpointConnection
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PrivateEndpointConnectionResponse{}, err
	}
	return PrivateEndpointConnectionResponse{RawResponse: resp.Response, PrivateEndpointConnection: val}, nil
}

// putHandleError handles the Put error response.
func (client *PrivateEndpointConnectionsClient) putHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
