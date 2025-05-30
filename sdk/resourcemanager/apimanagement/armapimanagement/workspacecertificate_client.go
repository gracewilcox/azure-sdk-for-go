// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// WorkspaceCertificateClient contains the methods for the WorkspaceCertificate group.
// Don't use this type directly, use NewWorkspaceCertificateClient() instead.
type WorkspaceCertificateClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWorkspaceCertificateClient creates a new instance of WorkspaceCertificateClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkspaceCertificateClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspaceCertificateClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WorkspaceCertificateClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates the certificate being used for authentication with the backend.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - certificateID - Identifier of the certificate entity. Must be unique in the current API Management service instance.
//   - parameters - Create or Update parameters.
//   - options - WorkspaceCertificateClientCreateOrUpdateOptions contains the optional parameters for the WorkspaceCertificateClient.CreateOrUpdate
//     method.
func (client *WorkspaceCertificateClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, certificateID string, parameters CertificateCreateOrUpdateParameters, options *WorkspaceCertificateClientCreateOrUpdateOptions) (WorkspaceCertificateClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "WorkspaceCertificateClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, certificateID, parameters, options)
	if err != nil {
		return WorkspaceCertificateClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceCertificateClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceCertificateClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkspaceCertificateClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, certificateID string, parameters CertificateCreateOrUpdateParameters, options *WorkspaceCertificateClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/certificates/{certificateId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if certificateID == "" {
		return nil, errors.New("parameter certificateID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateId}", url.PathEscape(certificateID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *WorkspaceCertificateClient) createOrUpdateHandleResponse(resp *http.Response) (WorkspaceCertificateClientCreateOrUpdateResponse, error) {
	result := WorkspaceCertificateClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateContract); err != nil {
		return WorkspaceCertificateClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes specific certificate.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - certificateID - Identifier of the certificate entity. Must be unique in the current API Management service instance.
//   - ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
//     it should be * for unconditional update.
//   - options - WorkspaceCertificateClientDeleteOptions contains the optional parameters for the WorkspaceCertificateClient.Delete
//     method.
func (client *WorkspaceCertificateClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, certificateID string, ifMatch string, options *WorkspaceCertificateClientDeleteOptions) (WorkspaceCertificateClientDeleteResponse, error) {
	var err error
	const operationName = "WorkspaceCertificateClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, certificateID, ifMatch, options)
	if err != nil {
		return WorkspaceCertificateClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceCertificateClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceCertificateClientDeleteResponse{}, err
	}
	return WorkspaceCertificateClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WorkspaceCertificateClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, certificateID string, ifMatch string, _ *WorkspaceCertificateClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/certificates/{certificateId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if certificateID == "" {
		return nil, errors.New("parameter certificateID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateId}", url.PathEscape(certificateID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["If-Match"] = []string{ifMatch}
	return req, nil
}

// Get - Gets the details of the certificate specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - certificateID - Identifier of the certificate entity. Must be unique in the current API Management service instance.
//   - options - WorkspaceCertificateClientGetOptions contains the optional parameters for the WorkspaceCertificateClient.Get
//     method.
func (client *WorkspaceCertificateClient) Get(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, certificateID string, options *WorkspaceCertificateClientGetOptions) (WorkspaceCertificateClientGetResponse, error) {
	var err error
	const operationName = "WorkspaceCertificateClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, certificateID, options)
	if err != nil {
		return WorkspaceCertificateClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceCertificateClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceCertificateClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *WorkspaceCertificateClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, certificateID string, _ *WorkspaceCertificateClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/certificates/{certificateId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if certificateID == "" {
		return nil, errors.New("parameter certificateID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateId}", url.PathEscape(certificateID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkspaceCertificateClient) getHandleResponse(resp *http.Response) (WorkspaceCertificateClientGetResponse, error) {
	result := WorkspaceCertificateClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateContract); err != nil {
		return WorkspaceCertificateClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Gets the entity state (Etag) version of the certificate specified by its identifier.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - certificateID - Identifier of the certificate entity. Must be unique in the current API Management service instance.
//   - options - WorkspaceCertificateClientGetEntityTagOptions contains the optional parameters for the WorkspaceCertificateClient.GetEntityTag
//     method.
func (client *WorkspaceCertificateClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, certificateID string, options *WorkspaceCertificateClientGetEntityTagOptions) (WorkspaceCertificateClientGetEntityTagResponse, error) {
	var err error
	const operationName = "WorkspaceCertificateClient.GetEntityTag"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, certificateID, options)
	if err != nil {
		return WorkspaceCertificateClientGetEntityTagResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceCertificateClientGetEntityTagResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceCertificateClientGetEntityTagResponse{}, err
	}
	resp, err := client.getEntityTagHandleResponse(httpResp)
	return resp, err
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *WorkspaceCertificateClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, certificateID string, _ *WorkspaceCertificateClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/certificates/{certificateId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if certificateID == "" {
		return nil, errors.New("parameter certificateID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateId}", url.PathEscape(certificateID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *WorkspaceCertificateClient) getEntityTagHandleResponse(resp *http.Response) (WorkspaceCertificateClientGetEntityTagResponse, error) {
	result := WorkspaceCertificateClientGetEntityTagResponse{Success: resp.StatusCode >= 200 && resp.StatusCode < 300}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	return result, nil
}

// NewListByWorkspacePager - Lists a collection of all certificates in the specified workspace.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - options - WorkspaceCertificateClientListByWorkspaceOptions contains the optional parameters for the WorkspaceCertificateClient.NewListByWorkspacePager
//     method.
func (client *WorkspaceCertificateClient) NewListByWorkspacePager(resourceGroupName string, serviceName string, workspaceID string, options *WorkspaceCertificateClientListByWorkspaceOptions) *runtime.Pager[WorkspaceCertificateClientListByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkspaceCertificateClientListByWorkspaceResponse]{
		More: func(page WorkspaceCertificateClientListByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkspaceCertificateClientListByWorkspaceResponse) (WorkspaceCertificateClientListByWorkspaceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "WorkspaceCertificateClient.NewListByWorkspacePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByWorkspaceCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, options)
			}, nil)
			if err != nil {
				return WorkspaceCertificateClientListByWorkspaceResponse{}, err
			}
			return client.listByWorkspaceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *WorkspaceCertificateClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, options *WorkspaceCertificateClientListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/certificates"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-05-01")
	if options != nil && options.IsKeyVaultRefreshFailed != nil {
		reqQP.Set("isKeyVaultRefreshFailed", strconv.FormatBool(*options.IsKeyVaultRefreshFailed))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *WorkspaceCertificateClient) listByWorkspaceHandleResponse(resp *http.Response) (WorkspaceCertificateClientListByWorkspaceResponse, error) {
	result := WorkspaceCertificateClientListByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateCollection); err != nil {
		return WorkspaceCertificateClientListByWorkspaceResponse{}, err
	}
	return result, nil
}

// RefreshSecret - From KeyVault, Refresh the certificate being used for authentication with the backend.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - certificateID - Identifier of the certificate entity. Must be unique in the current API Management service instance.
//   - options - WorkspaceCertificateClientRefreshSecretOptions contains the optional parameters for the WorkspaceCertificateClient.RefreshSecret
//     method.
func (client *WorkspaceCertificateClient) RefreshSecret(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, certificateID string, options *WorkspaceCertificateClientRefreshSecretOptions) (WorkspaceCertificateClientRefreshSecretResponse, error) {
	var err error
	const operationName = "WorkspaceCertificateClient.RefreshSecret"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.refreshSecretCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, certificateID, options)
	if err != nil {
		return WorkspaceCertificateClientRefreshSecretResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceCertificateClientRefreshSecretResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceCertificateClientRefreshSecretResponse{}, err
	}
	resp, err := client.refreshSecretHandleResponse(httpResp)
	return resp, err
}

// refreshSecretCreateRequest creates the RefreshSecret request.
func (client *WorkspaceCertificateClient) refreshSecretCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, certificateID string, _ *WorkspaceCertificateClientRefreshSecretOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/certificates/{certificateId}/refreshSecret"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if certificateID == "" {
		return nil, errors.New("parameter certificateID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateId}", url.PathEscape(certificateID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// refreshSecretHandleResponse handles the RefreshSecret response.
func (client *WorkspaceCertificateClient) refreshSecretHandleResponse(resp *http.Response) (WorkspaceCertificateClientRefreshSecretResponse, error) {
	result := WorkspaceCertificateClientRefreshSecretResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateContract); err != nil {
		return WorkspaceCertificateClientRefreshSecretResponse{}, err
	}
	return result, nil
}
