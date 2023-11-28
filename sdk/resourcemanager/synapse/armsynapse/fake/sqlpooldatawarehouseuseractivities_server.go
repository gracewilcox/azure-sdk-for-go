//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
	"net/http"
	"net/url"
	"regexp"
)

// SQLPoolDataWarehouseUserActivitiesServer is a fake server for instances of the armsynapse.SQLPoolDataWarehouseUserActivitiesClient type.
type SQLPoolDataWarehouseUserActivitiesServer struct {
	// Get is the fake for method SQLPoolDataWarehouseUserActivitiesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, dataWarehouseUserActivityName armsynapse.DataWarehouseUserActivityName, options *armsynapse.SQLPoolDataWarehouseUserActivitiesClientGetOptions) (resp azfake.Responder[armsynapse.SQLPoolDataWarehouseUserActivitiesClientGetResponse], errResp azfake.ErrorResponder)
}

// NewSQLPoolDataWarehouseUserActivitiesServerTransport creates a new instance of SQLPoolDataWarehouseUserActivitiesServerTransport with the provided implementation.
// The returned SQLPoolDataWarehouseUserActivitiesServerTransport instance is connected to an instance of armsynapse.SQLPoolDataWarehouseUserActivitiesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSQLPoolDataWarehouseUserActivitiesServerTransport(srv *SQLPoolDataWarehouseUserActivitiesServer) *SQLPoolDataWarehouseUserActivitiesServerTransport {
	return &SQLPoolDataWarehouseUserActivitiesServerTransport{srv: srv}
}

// SQLPoolDataWarehouseUserActivitiesServerTransport connects instances of armsynapse.SQLPoolDataWarehouseUserActivitiesClient to instances of SQLPoolDataWarehouseUserActivitiesServer.
// Don't use this type directly, use NewSQLPoolDataWarehouseUserActivitiesServerTransport instead.
type SQLPoolDataWarehouseUserActivitiesServerTransport struct {
	srv *SQLPoolDataWarehouseUserActivitiesServer
}

// Do implements the policy.Transporter interface for SQLPoolDataWarehouseUserActivitiesServerTransport.
func (s *SQLPoolDataWarehouseUserActivitiesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SQLPoolDataWarehouseUserActivitiesClient.Get":
		resp, err = s.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SQLPoolDataWarehouseUserActivitiesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Synapse/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlPools/(?P<sqlPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dataWarehouseUserActivities/(?P<dataWarehouseUserActivityName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	sqlPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlPoolName")])
	if err != nil {
		return nil, err
	}
	dataWarehouseUserActivityNameParam, err := parseWithCast(matches[regex.SubexpIndex("dataWarehouseUserActivityName")], func(v string) (armsynapse.DataWarehouseUserActivityName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armsynapse.DataWarehouseUserActivityName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, workspaceNameParam, sqlPoolNameParam, dataWarehouseUserActivityNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DataWarehouseUserActivities, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}