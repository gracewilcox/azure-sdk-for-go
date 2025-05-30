// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ReplicationAlertSettingsClient contains the methods for the ReplicationAlertSettings group.
// Don't use this type directly, use NewReplicationAlertSettingsClient() instead.
type ReplicationAlertSettingsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewReplicationAlertSettingsClient creates a new instance of ReplicationAlertSettingsClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewReplicationAlertSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ReplicationAlertSettingsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ReplicationAlertSettingsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Create or update an email notification(alert) configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - resourceName - The name of the recovery services vault.
//   - alertSettingName - The name of the email notification(alert) configuration.
//   - request - The input to configure the email notification(alert).
//   - options - ReplicationAlertSettingsClientCreateOptions contains the optional parameters for the ReplicationAlertSettingsClient.Create
//     method.
func (client *ReplicationAlertSettingsClient) Create(ctx context.Context, resourceGroupName string, resourceName string, alertSettingName string, request ConfigureAlertRequest, options *ReplicationAlertSettingsClientCreateOptions) (ReplicationAlertSettingsClientCreateResponse, error) {
	var err error
	const operationName = "ReplicationAlertSettingsClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, resourceName, alertSettingName, request, options)
	if err != nil {
		return ReplicationAlertSettingsClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReplicationAlertSettingsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ReplicationAlertSettingsClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *ReplicationAlertSettingsClient) createCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, alertSettingName string, request ConfigureAlertRequest, _ *ReplicationAlertSettingsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationAlertSettings/{alertSettingName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if alertSettingName == "" {
		return nil, errors.New("parameter alertSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alertSettingName}", url.PathEscape(alertSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, request); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *ReplicationAlertSettingsClient) createHandleResponse(resp *http.Response) (ReplicationAlertSettingsClientCreateResponse, error) {
	result := ReplicationAlertSettingsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Alert); err != nil {
		return ReplicationAlertSettingsClientCreateResponse{}, err
	}
	return result, nil
}

// Get - Gets the details of the specified email notification(alert) configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - resourceName - The name of the recovery services vault.
//   - alertSettingName - The name of the email notification configuration.
//   - options - ReplicationAlertSettingsClientGetOptions contains the optional parameters for the ReplicationAlertSettingsClient.Get
//     method.
func (client *ReplicationAlertSettingsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, alertSettingName string, options *ReplicationAlertSettingsClientGetOptions) (ReplicationAlertSettingsClientGetResponse, error) {
	var err error
	const operationName = "ReplicationAlertSettingsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, alertSettingName, options)
	if err != nil {
		return ReplicationAlertSettingsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReplicationAlertSettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ReplicationAlertSettingsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ReplicationAlertSettingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, alertSettingName string, _ *ReplicationAlertSettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationAlertSettings/{alertSettingName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if alertSettingName == "" {
		return nil, errors.New("parameter alertSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alertSettingName}", url.PathEscape(alertSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReplicationAlertSettingsClient) getHandleResponse(resp *http.Response) (ReplicationAlertSettingsClientGetResponse, error) {
	result := ReplicationAlertSettingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Alert); err != nil {
		return ReplicationAlertSettingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the list of email notification(alert) configurations for the vault.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - resourceName - The name of the recovery services vault.
//   - options - ReplicationAlertSettingsClientListOptions contains the optional parameters for the ReplicationAlertSettingsClient.NewListPager
//     method.
func (client *ReplicationAlertSettingsClient) NewListPager(resourceGroupName string, resourceName string, options *ReplicationAlertSettingsClientListOptions) *runtime.Pager[ReplicationAlertSettingsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicationAlertSettingsClientListResponse]{
		More: func(page ReplicationAlertSettingsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReplicationAlertSettingsClientListResponse) (ReplicationAlertSettingsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ReplicationAlertSettingsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
			}, nil)
			if err != nil {
				return ReplicationAlertSettingsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ReplicationAlertSettingsClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, _ *ReplicationAlertSettingsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationAlertSettings"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReplicationAlertSettingsClient) listHandleResponse(resp *http.Response) (ReplicationAlertSettingsClientListResponse, error) {
	result := ReplicationAlertSettingsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertCollection); err != nil {
		return ReplicationAlertSettingsClientListResponse{}, err
	}
	return result, nil
}
