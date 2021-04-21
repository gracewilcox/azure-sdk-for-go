// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// NetworkInterfaceLoadBalancersClient contains the methods for the NetworkInterfaceLoadBalancers group.
// Don't use this type directly, use NewNetworkInterfaceLoadBalancersClient() instead.
type NetworkInterfaceLoadBalancersClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewNetworkInterfaceLoadBalancersClient creates a new instance of NetworkInterfaceLoadBalancersClient with the specified values.
func NewNetworkInterfaceLoadBalancersClient(con *armcore.Connection, subscriptionID string) *NetworkInterfaceLoadBalancersClient {
	return &NetworkInterfaceLoadBalancersClient{con: con, subscriptionID: subscriptionID}
}

// List - List all load balancers in a network interface.
func (client *NetworkInterfaceLoadBalancersClient) List(resourceGroupName string, networkInterfaceName string, options *NetworkInterfaceLoadBalancersListOptions) NetworkInterfaceLoadBalancerListResultPager {
	return &networkInterfaceLoadBalancerListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, networkInterfaceName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp NetworkInterfaceLoadBalancerListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NetworkInterfaceLoadBalancerListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *NetworkInterfaceLoadBalancersClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkInterfaceName string, options *NetworkInterfaceLoadBalancersListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/loadBalancers"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *NetworkInterfaceLoadBalancersClient) listHandleResponse(resp *azcore.Response) (NetworkInterfaceLoadBalancerListResultResponse, error) {
	var val *NetworkInterfaceLoadBalancerListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return NetworkInterfaceLoadBalancerListResultResponse{}, err
	}
	return NetworkInterfaceLoadBalancerListResultResponse{RawResponse: resp.Response, NetworkInterfaceLoadBalancerListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *NetworkInterfaceLoadBalancersClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
