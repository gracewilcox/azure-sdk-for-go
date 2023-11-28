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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
	"net/http"
	"net/url"
	"regexp"
)

// UsageServer is a fake server for instances of the armtestbase.UsageClient type.
type UsageServer struct {
	// NewListPager is the fake for method UsageClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, testBaseAccountName string, options *armtestbase.UsageClientListOptions) (resp azfake.PagerResponder[armtestbase.UsageClientListResponse])
}

// NewUsageServerTransport creates a new instance of UsageServerTransport with the provided implementation.
// The returned UsageServerTransport instance is connected to an instance of armtestbase.UsageClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewUsageServerTransport(srv *UsageServer) *UsageServerTransport {
	return &UsageServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armtestbase.UsageClientListResponse]](),
	}
}

// UsageServerTransport connects instances of armtestbase.UsageClient to instances of UsageServer.
// Don't use this type directly, use NewUsageServerTransport instead.
type UsageServerTransport struct {
	srv          *UsageServer
	newListPager *tracker[azfake.PagerResponder[armtestbase.UsageClientListResponse]]
}

// Do implements the policy.Transporter interface for UsageServerTransport.
func (u *UsageServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "UsageClient.NewListPager":
		resp, err = u.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (u *UsageServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if u.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := u.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.TestBase/testBaseAccounts/(?P<testBaseAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/usages`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		testBaseAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("testBaseAccountName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armtestbase.UsageClientListOptions
		if filterParam != nil {
			options = &armtestbase.UsageClientListOptions{
				Filter: filterParam,
			}
		}
		resp := u.srv.NewListPager(resourceGroupNameParam, testBaseAccountNameParam, options)
		newListPager = &resp
		u.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armtestbase.UsageClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		u.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		u.newListPager.remove(req)
	}
	return resp, nil
}