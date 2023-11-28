//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
	"net/http"
	"net/url"
	"regexp"
)

// WebTestLocationsServer is a fake server for instances of the armapplicationinsights.WebTestLocationsClient type.
type WebTestLocationsServer struct {
	// NewListPager is the fake for method WebTestLocationsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, resourceName string, options *armapplicationinsights.WebTestLocationsClientListOptions) (resp azfake.PagerResponder[armapplicationinsights.WebTestLocationsClientListResponse])
}

// NewWebTestLocationsServerTransport creates a new instance of WebTestLocationsServerTransport with the provided implementation.
// The returned WebTestLocationsServerTransport instance is connected to an instance of armapplicationinsights.WebTestLocationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWebTestLocationsServerTransport(srv *WebTestLocationsServer) *WebTestLocationsServerTransport {
	return &WebTestLocationsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armapplicationinsights.WebTestLocationsClientListResponse]](),
	}
}

// WebTestLocationsServerTransport connects instances of armapplicationinsights.WebTestLocationsClient to instances of WebTestLocationsServer.
// Don't use this type directly, use NewWebTestLocationsServerTransport instead.
type WebTestLocationsServerTransport struct {
	srv          *WebTestLocationsServer
	newListPager *tracker[azfake.PagerResponder[armapplicationinsights.WebTestLocationsClientListResponse]]
}

// Do implements the policy.Transporter interface for WebTestLocationsServerTransport.
func (w *WebTestLocationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "WebTestLocationsClient.NewListPager":
		resp, err = w.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *WebTestLocationsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := w.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Insights/components/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syntheticmonitorlocations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		resp := w.srv.NewListPager(resourceGroupNameParam, resourceNameParam, nil)
		newListPager = &resp
		w.newListPager.add(req, newListPager)
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