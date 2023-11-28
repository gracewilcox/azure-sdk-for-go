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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"net/http"
	"net/url"
	"regexp"
)

// CapabilitiesServer is a fake server for instances of the armsql.CapabilitiesClient type.
type CapabilitiesServer struct {
	// ListByLocation is the fake for method CapabilitiesClient.ListByLocation
	// HTTP status codes to indicate success: http.StatusOK
	ListByLocation func(ctx context.Context, locationName string, options *armsql.CapabilitiesClientListByLocationOptions) (resp azfake.Responder[armsql.CapabilitiesClientListByLocationResponse], errResp azfake.ErrorResponder)
}

// NewCapabilitiesServerTransport creates a new instance of CapabilitiesServerTransport with the provided implementation.
// The returned CapabilitiesServerTransport instance is connected to an instance of armsql.CapabilitiesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCapabilitiesServerTransport(srv *CapabilitiesServer) *CapabilitiesServerTransport {
	return &CapabilitiesServerTransport{srv: srv}
}

// CapabilitiesServerTransport connects instances of armsql.CapabilitiesClient to instances of CapabilitiesServer.
// Don't use this type directly, use NewCapabilitiesServerTransport instead.
type CapabilitiesServerTransport struct {
	srv *CapabilitiesServer
}

// Do implements the policy.Transporter interface for CapabilitiesServerTransport.
func (c *CapabilitiesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CapabilitiesClient.ListByLocation":
		resp, err = c.dispatchListByLocation(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CapabilitiesServerTransport) dispatchListByLocation(req *http.Request) (*http.Response, error) {
	if c.srv.ListByLocation == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListByLocation not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/locations/(?P<locationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/capabilities`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	locationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationName")])
	if err != nil {
		return nil, err
	}
	includeUnescaped, err := url.QueryUnescape(qp.Get("include"))
	if err != nil {
		return nil, err
	}
	includeParam := getOptional(armsql.CapabilityGroup(includeUnescaped))
	var options *armsql.CapabilitiesClientListByLocationOptions
	if includeParam != nil {
		options = &armsql.CapabilitiesClientListByLocationOptions{
			Include: includeParam,
		}
	}
	respr, errRespr := c.srv.ListByLocation(req.Context(), locationNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LocationCapabilities, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}