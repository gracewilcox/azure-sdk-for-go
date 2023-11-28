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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
	"net/http"
	"net/url"
	"regexp"
)

// WorkspaceManagedSQLServerSecurityAlertPolicyServer is a fake server for instances of the armsynapse.WorkspaceManagedSQLServerSecurityAlertPolicyClient type.
type WorkspaceManagedSQLServerSecurityAlertPolicyServer struct {
	// BeginCreateOrUpdate is the fake for method WorkspaceManagedSQLServerSecurityAlertPolicyClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, workspaceName string, securityAlertPolicyName armsynapse.SecurityAlertPolicyNameAutoGenerated, parameters armsynapse.ServerSecurityAlertPolicy, options *armsynapse.WorkspaceManagedSQLServerSecurityAlertPolicyClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armsynapse.WorkspaceManagedSQLServerSecurityAlertPolicyClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method WorkspaceManagedSQLServerSecurityAlertPolicyClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workspaceName string, securityAlertPolicyName armsynapse.SecurityAlertPolicyNameAutoGenerated, options *armsynapse.WorkspaceManagedSQLServerSecurityAlertPolicyClientGetOptions) (resp azfake.Responder[armsynapse.WorkspaceManagedSQLServerSecurityAlertPolicyClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method WorkspaceManagedSQLServerSecurityAlertPolicyClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, workspaceName string, options *armsynapse.WorkspaceManagedSQLServerSecurityAlertPolicyClientListOptions) (resp azfake.PagerResponder[armsynapse.WorkspaceManagedSQLServerSecurityAlertPolicyClientListResponse])
}

// NewWorkspaceManagedSQLServerSecurityAlertPolicyServerTransport creates a new instance of WorkspaceManagedSQLServerSecurityAlertPolicyServerTransport with the provided implementation.
// The returned WorkspaceManagedSQLServerSecurityAlertPolicyServerTransport instance is connected to an instance of armsynapse.WorkspaceManagedSQLServerSecurityAlertPolicyClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWorkspaceManagedSQLServerSecurityAlertPolicyServerTransport(srv *WorkspaceManagedSQLServerSecurityAlertPolicyServer) *WorkspaceManagedSQLServerSecurityAlertPolicyServerTransport {
	return &WorkspaceManagedSQLServerSecurityAlertPolicyServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armsynapse.WorkspaceManagedSQLServerSecurityAlertPolicyClientCreateOrUpdateResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armsynapse.WorkspaceManagedSQLServerSecurityAlertPolicyClientListResponse]](),
	}
}

// WorkspaceManagedSQLServerSecurityAlertPolicyServerTransport connects instances of armsynapse.WorkspaceManagedSQLServerSecurityAlertPolicyClient to instances of WorkspaceManagedSQLServerSecurityAlertPolicyServer.
// Don't use this type directly, use NewWorkspaceManagedSQLServerSecurityAlertPolicyServerTransport instead.
type WorkspaceManagedSQLServerSecurityAlertPolicyServerTransport struct {
	srv                 *WorkspaceManagedSQLServerSecurityAlertPolicyServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armsynapse.WorkspaceManagedSQLServerSecurityAlertPolicyClientCreateOrUpdateResponse]]
	newListPager        *tracker[azfake.PagerResponder[armsynapse.WorkspaceManagedSQLServerSecurityAlertPolicyClientListResponse]]
}

// Do implements the policy.Transporter interface for WorkspaceManagedSQLServerSecurityAlertPolicyServerTransport.
func (w *WorkspaceManagedSQLServerSecurityAlertPolicyServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "WorkspaceManagedSQLServerSecurityAlertPolicyClient.BeginCreateOrUpdate":
		resp, err = w.dispatchBeginCreateOrUpdate(req)
	case "WorkspaceManagedSQLServerSecurityAlertPolicyClient.Get":
		resp, err = w.dispatchGet(req)
	case "WorkspaceManagedSQLServerSecurityAlertPolicyClient.NewListPager":
		resp, err = w.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *WorkspaceManagedSQLServerSecurityAlertPolicyServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if w.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := w.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Synapse/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/securityAlertPolicies/(?P<securityAlertPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsynapse.ServerSecurityAlertPolicy](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		securityAlertPolicyNameParam, err := parseWithCast(matches[regex.SubexpIndex("securityAlertPolicyName")], func(v string) (armsynapse.SecurityAlertPolicyNameAutoGenerated, error) {
			p, unescapeErr := url.PathUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armsynapse.SecurityAlertPolicyNameAutoGenerated(p), nil
		})
		if err != nil {
			return nil, err
		}
		respr, errRespr := w.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, workspaceNameParam, securityAlertPolicyNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		w.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		w.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		w.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (w *WorkspaceManagedSQLServerSecurityAlertPolicyServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if w.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Synapse/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/securityAlertPolicies/(?P<securityAlertPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
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
	securityAlertPolicyNameParam, err := parseWithCast(matches[regex.SubexpIndex("securityAlertPolicyName")], func(v string) (armsynapse.SecurityAlertPolicyNameAutoGenerated, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armsynapse.SecurityAlertPolicyNameAutoGenerated(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Get(req.Context(), resourceGroupNameParam, workspaceNameParam, securityAlertPolicyNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ServerSecurityAlertPolicy, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkspaceManagedSQLServerSecurityAlertPolicyServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := w.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Synapse/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/securityAlertPolicies`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
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
		resp := w.srv.NewListPager(resourceGroupNameParam, workspaceNameParam, nil)
		newListPager = &resp
		w.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsynapse.WorkspaceManagedSQLServerSecurityAlertPolicyClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		w.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		w.newListPager.remove(req)
	}
	return resp, nil
}