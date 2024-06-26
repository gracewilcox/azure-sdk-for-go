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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ThreatIntelligenceIndicatorServer is a fake server for instances of the armsecurityinsights.ThreatIntelligenceIndicatorClient type.
type ThreatIntelligenceIndicatorServer struct {
	// AppendTags is the fake for method ThreatIntelligenceIndicatorClient.AppendTags
	// HTTP status codes to indicate success: http.StatusOK
	AppendTags func(ctx context.Context, resourceGroupName string, workspaceName string, name string, threatIntelligenceAppendTags armsecurityinsights.ThreatIntelligenceAppendTags, options *armsecurityinsights.ThreatIntelligenceIndicatorClientAppendTagsOptions) (resp azfake.Responder[armsecurityinsights.ThreatIntelligenceIndicatorClientAppendTagsResponse], errResp azfake.ErrorResponder)

	// Create is the fake for method ThreatIntelligenceIndicatorClient.Create
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Create func(ctx context.Context, resourceGroupName string, workspaceName string, name string, threatIntelligenceProperties armsecurityinsights.ThreatIntelligenceIndicatorModel, options *armsecurityinsights.ThreatIntelligenceIndicatorClientCreateOptions) (resp azfake.Responder[armsecurityinsights.ThreatIntelligenceIndicatorClientCreateResponse], errResp azfake.ErrorResponder)

	// CreateIndicator is the fake for method ThreatIntelligenceIndicatorClient.CreateIndicator
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateIndicator func(ctx context.Context, resourceGroupName string, workspaceName string, threatIntelligenceProperties armsecurityinsights.ThreatIntelligenceIndicatorModel, options *armsecurityinsights.ThreatIntelligenceIndicatorClientCreateIndicatorOptions) (resp azfake.Responder[armsecurityinsights.ThreatIntelligenceIndicatorClientCreateIndicatorResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ThreatIntelligenceIndicatorClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *armsecurityinsights.ThreatIntelligenceIndicatorClientDeleteOptions) (resp azfake.Responder[armsecurityinsights.ThreatIntelligenceIndicatorClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ThreatIntelligenceIndicatorClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *armsecurityinsights.ThreatIntelligenceIndicatorClientGetOptions) (resp azfake.Responder[armsecurityinsights.ThreatIntelligenceIndicatorClientGetResponse], errResp azfake.ErrorResponder)

	// NewQueryIndicatorsPager is the fake for method ThreatIntelligenceIndicatorClient.NewQueryIndicatorsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewQueryIndicatorsPager func(resourceGroupName string, workspaceName string, threatIntelligenceFilteringCriteria armsecurityinsights.ThreatIntelligenceFilteringCriteria, options *armsecurityinsights.ThreatIntelligenceIndicatorClientQueryIndicatorsOptions) (resp azfake.PagerResponder[armsecurityinsights.ThreatIntelligenceIndicatorClientQueryIndicatorsResponse])

	// ReplaceTags is the fake for method ThreatIntelligenceIndicatorClient.ReplaceTags
	// HTTP status codes to indicate success: http.StatusOK
	ReplaceTags func(ctx context.Context, resourceGroupName string, workspaceName string, name string, threatIntelligenceReplaceTags armsecurityinsights.ThreatIntelligenceIndicatorModel, options *armsecurityinsights.ThreatIntelligenceIndicatorClientReplaceTagsOptions) (resp azfake.Responder[armsecurityinsights.ThreatIntelligenceIndicatorClientReplaceTagsResponse], errResp azfake.ErrorResponder)
}

// NewThreatIntelligenceIndicatorServerTransport creates a new instance of ThreatIntelligenceIndicatorServerTransport with the provided implementation.
// The returned ThreatIntelligenceIndicatorServerTransport instance is connected to an instance of armsecurityinsights.ThreatIntelligenceIndicatorClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewThreatIntelligenceIndicatorServerTransport(srv *ThreatIntelligenceIndicatorServer) *ThreatIntelligenceIndicatorServerTransport {
	return &ThreatIntelligenceIndicatorServerTransport{
		srv:                     srv,
		newQueryIndicatorsPager: newTracker[azfake.PagerResponder[armsecurityinsights.ThreatIntelligenceIndicatorClientQueryIndicatorsResponse]](),
	}
}

// ThreatIntelligenceIndicatorServerTransport connects instances of armsecurityinsights.ThreatIntelligenceIndicatorClient to instances of ThreatIntelligenceIndicatorServer.
// Don't use this type directly, use NewThreatIntelligenceIndicatorServerTransport instead.
type ThreatIntelligenceIndicatorServerTransport struct {
	srv                     *ThreatIntelligenceIndicatorServer
	newQueryIndicatorsPager *tracker[azfake.PagerResponder[armsecurityinsights.ThreatIntelligenceIndicatorClientQueryIndicatorsResponse]]
}

// Do implements the policy.Transporter interface for ThreatIntelligenceIndicatorServerTransport.
func (t *ThreatIntelligenceIndicatorServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ThreatIntelligenceIndicatorClient.AppendTags":
		resp, err = t.dispatchAppendTags(req)
	case "ThreatIntelligenceIndicatorClient.Create":
		resp, err = t.dispatchCreate(req)
	case "ThreatIntelligenceIndicatorClient.CreateIndicator":
		resp, err = t.dispatchCreateIndicator(req)
	case "ThreatIntelligenceIndicatorClient.Delete":
		resp, err = t.dispatchDelete(req)
	case "ThreatIntelligenceIndicatorClient.Get":
		resp, err = t.dispatchGet(req)
	case "ThreatIntelligenceIndicatorClient.NewQueryIndicatorsPager":
		resp, err = t.dispatchNewQueryIndicatorsPager(req)
	case "ThreatIntelligenceIndicatorClient.ReplaceTags":
		resp, err = t.dispatchReplaceTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *ThreatIntelligenceIndicatorServerTransport) dispatchAppendTags(req *http.Request) (*http.Response, error) {
	if t.srv.AppendTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method AppendTags not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/threatIntelligence/main/indicators/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/appendTags`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsecurityinsights.ThreatIntelligenceAppendTags](req)
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
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.AppendTags(req.Context(), resourceGroupNameParam, workspaceNameParam, nameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *ThreatIntelligenceIndicatorServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if t.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/threatIntelligence/main/indicators/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsecurityinsights.ThreatIntelligenceIndicatorModel](req)
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
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Create(req.Context(), resourceGroupNameParam, workspaceNameParam, nameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ThreatIntelligenceInformationClassification, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *ThreatIntelligenceIndicatorServerTransport) dispatchCreateIndicator(req *http.Request) (*http.Response, error) {
	if t.srv.CreateIndicator == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateIndicator not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/threatIntelligence/main/createIndicator`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsecurityinsights.ThreatIntelligenceIndicatorModel](req)
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
	respr, errRespr := t.srv.CreateIndicator(req.Context(), resourceGroupNameParam, workspaceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ThreatIntelligenceInformationClassification, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *ThreatIntelligenceIndicatorServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if t.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/threatIntelligence/main/indicators/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Delete(req.Context(), resourceGroupNameParam, workspaceNameParam, nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *ThreatIntelligenceIndicatorServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if t.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/threatIntelligence/main/indicators/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Get(req.Context(), resourceGroupNameParam, workspaceNameParam, nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ThreatIntelligenceInformationClassification, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *ThreatIntelligenceIndicatorServerTransport) dispatchNewQueryIndicatorsPager(req *http.Request) (*http.Response, error) {
	if t.srv.NewQueryIndicatorsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewQueryIndicatorsPager not implemented")}
	}
	newQueryIndicatorsPager := t.newQueryIndicatorsPager.get(req)
	if newQueryIndicatorsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/threatIntelligence/main/queryIndicators`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsecurityinsights.ThreatIntelligenceFilteringCriteria](req)
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
		resp := t.srv.NewQueryIndicatorsPager(resourceGroupNameParam, workspaceNameParam, body, nil)
		newQueryIndicatorsPager = &resp
		t.newQueryIndicatorsPager.add(req, newQueryIndicatorsPager)
		server.PagerResponderInjectNextLinks(newQueryIndicatorsPager, req, func(page *armsecurityinsights.ThreatIntelligenceIndicatorClientQueryIndicatorsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newQueryIndicatorsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		t.newQueryIndicatorsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newQueryIndicatorsPager) {
		t.newQueryIndicatorsPager.remove(req)
	}
	return resp, nil
}

func (t *ThreatIntelligenceIndicatorServerTransport) dispatchReplaceTags(req *http.Request) (*http.Response, error) {
	if t.srv.ReplaceTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method ReplaceTags not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/threatIntelligence/main/indicators/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replaceTags`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsecurityinsights.ThreatIntelligenceIndicatorModel](req)
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
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.ReplaceTags(req.Context(), resourceGroupNameParam, workspaceNameParam, nameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ThreatIntelligenceInformationClassification, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
