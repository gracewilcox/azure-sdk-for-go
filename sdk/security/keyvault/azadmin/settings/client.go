//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package settings

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// Client contains the methods for the Client group.
// Don't use this type directly, use a constructor function instead.
type Client struct {
	internal *azcore.Client
	endpoint string
}

// GetSetting - Retrieves the setting object of a specified setting name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 7.6-preview.1
//   - settingName - The name of the account setting. Must be a valid settings option.
//   - options - GetSettingOptions contains the optional parameters for the Client.GetSetting method.
func (client *Client) GetSetting(ctx context.Context, settingName string, options *GetSettingOptions) (GetSettingResponse, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "settings.Client.GetSetting", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getSettingCreateRequest(ctx, settingName, options)
	if err != nil {
		return GetSettingResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GetSettingResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GetSettingResponse{}, err
	}
	resp, err := client.getSettingHandleResponse(httpResp)
	return resp, err
}

// getSettingCreateRequest creates the GetSetting request.
func (client *Client) getSettingCreateRequest(ctx context.Context, settingName string, options *GetSettingOptions) (*policy.Request, error) {
	urlPath := "/settings/{setting-name}"
	if settingName == "" {
		return nil, errors.New("parameter settingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{setting-name}", url.PathEscape(settingName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "7.6-preview.1")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSettingHandleResponse handles the GetSetting response.
func (client *Client) getSettingHandleResponse(resp *http.Response) (GetSettingResponse, error) {
	result := GetSettingResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Setting); err != nil {
		return GetSettingResponse{}, err
	}
	return result, nil
}

// GetSettings - Retrieves a list of all the available account settings that can be configured.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 7.6-preview.1
//   - options - GetSettingsOptions contains the optional parameters for the Client.GetSettings method.
func (client *Client) GetSettings(ctx context.Context, options *GetSettingsOptions) (GetSettingsResponse, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "settings.Client.GetSettings", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getSettingsCreateRequest(ctx, options)
	if err != nil {
		return GetSettingsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GetSettingsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GetSettingsResponse{}, err
	}
	resp, err := client.getSettingsHandleResponse(httpResp)
	return resp, err
}

// getSettingsCreateRequest creates the GetSettings request.
func (client *Client) getSettingsCreateRequest(ctx context.Context, options *GetSettingsOptions) (*policy.Request, error) {
	urlPath := "/settings"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "7.6-preview.1")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSettingsHandleResponse handles the GetSettings response.
func (client *Client) getSettingsHandleResponse(resp *http.Response) (GetSettingsResponse, error) {
	result := GetSettingsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return GetSettingsResponse{}, err
	}
	return result, nil
}

// UpdateSetting - Description of the pool setting to be updated
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 7.6-preview.1
//   - settingName - The name of the account setting. Must be a valid settings option.
//   - parameters - The parameters to update an account setting.
//   - options - UpdateSettingOptions contains the optional parameters for the Client.UpdateSetting method.
func (client *Client) UpdateSetting(ctx context.Context, settingName string, parameters UpdateSettingRequest, options *UpdateSettingOptions) (UpdateSettingResponse, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "settings.Client.UpdateSetting", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateSettingCreateRequest(ctx, settingName, parameters, options)
	if err != nil {
		return UpdateSettingResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UpdateSettingResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return UpdateSettingResponse{}, err
	}
	resp, err := client.updateSettingHandleResponse(httpResp)
	return resp, err
}

// updateSettingCreateRequest creates the UpdateSetting request.
func (client *Client) updateSettingCreateRequest(ctx context.Context, settingName string, parameters UpdateSettingRequest, options *UpdateSettingOptions) (*policy.Request, error) {
	urlPath := "/settings/{setting-name}"
	if settingName == "" {
		return nil, errors.New("parameter settingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{setting-name}", url.PathEscape(settingName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "7.6-preview.1")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateSettingHandleResponse handles the UpdateSetting response.
func (client *Client) updateSettingHandleResponse(resp *http.Response) (UpdateSettingResponse, error) {
	result := UpdateSettingResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Setting); err != nil {
		return UpdateSettingResponse{}, err
	}
	return result, nil
}
