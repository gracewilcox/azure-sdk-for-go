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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ReplicationNetworksServer is a fake server for instances of the armrecoveryservicessiterecovery.ReplicationNetworksClient type.
type ReplicationNetworksServer struct {
	// Get is the fake for method ReplicationNetworksClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, networkName string, options *armrecoveryservicessiterecovery.ReplicationNetworksClientGetOptions) (resp azfake.Responder[armrecoveryservicessiterecovery.ReplicationNetworksClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ReplicationNetworksClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, resourceName string, options *armrecoveryservicessiterecovery.ReplicationNetworksClientListOptions) (resp azfake.PagerResponder[armrecoveryservicessiterecovery.ReplicationNetworksClientListResponse])

	// NewListByReplicationFabricsPager is the fake for method ReplicationNetworksClient.NewListByReplicationFabricsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByReplicationFabricsPager func(resourceGroupName string, resourceName string, fabricName string, options *armrecoveryservicessiterecovery.ReplicationNetworksClientListByReplicationFabricsOptions) (resp azfake.PagerResponder[armrecoveryservicessiterecovery.ReplicationNetworksClientListByReplicationFabricsResponse])
}

// NewReplicationNetworksServerTransport creates a new instance of ReplicationNetworksServerTransport with the provided implementation.
// The returned ReplicationNetworksServerTransport instance is connected to an instance of armrecoveryservicessiterecovery.ReplicationNetworksClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewReplicationNetworksServerTransport(srv *ReplicationNetworksServer) *ReplicationNetworksServerTransport {
	return &ReplicationNetworksServerTransport{
		srv:                              srv,
		newListPager:                     newTracker[azfake.PagerResponder[armrecoveryservicessiterecovery.ReplicationNetworksClientListResponse]](),
		newListByReplicationFabricsPager: newTracker[azfake.PagerResponder[armrecoveryservicessiterecovery.ReplicationNetworksClientListByReplicationFabricsResponse]](),
	}
}

// ReplicationNetworksServerTransport connects instances of armrecoveryservicessiterecovery.ReplicationNetworksClient to instances of ReplicationNetworksServer.
// Don't use this type directly, use NewReplicationNetworksServerTransport instead.
type ReplicationNetworksServerTransport struct {
	srv                              *ReplicationNetworksServer
	newListPager                     *tracker[azfake.PagerResponder[armrecoveryservicessiterecovery.ReplicationNetworksClientListResponse]]
	newListByReplicationFabricsPager *tracker[azfake.PagerResponder[armrecoveryservicessiterecovery.ReplicationNetworksClientListByReplicationFabricsResponse]]
}

// Do implements the policy.Transporter interface for ReplicationNetworksServerTransport.
func (r *ReplicationNetworksServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return r.dispatchToMethodFake(req, method)
}

func (r *ReplicationNetworksServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if replicationNetworksServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = replicationNetworksServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ReplicationNetworksClient.Get":
				res.resp, res.err = r.dispatchGet(req)
			case "ReplicationNetworksClient.NewListPager":
				res.resp, res.err = r.dispatchNewListPager(req)
			case "ReplicationNetworksClient.NewListByReplicationFabricsPager":
				res.resp, res.err = r.dispatchNewListByReplicationFabricsPager(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (r *ReplicationNetworksServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationNetworks/(?P<networkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	fabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fabricName")])
	if err != nil {
		return nil, err
	}
	networkNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameParam, resourceNameParam, fabricNameParam, networkNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Network, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *ReplicationNetworksServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationNetworks`
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
		resp := r.srv.NewListPager(resourceGroupNameParam, resourceNameParam, nil)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armrecoveryservicessiterecovery.ReplicationNetworksClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		r.newListPager.remove(req)
	}
	return resp, nil
}

func (r *ReplicationNetworksServerTransport) dispatchNewListByReplicationFabricsPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListByReplicationFabricsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByReplicationFabricsPager not implemented")}
	}
	newListByReplicationFabricsPager := r.newListByReplicationFabricsPager.get(req)
	if newListByReplicationFabricsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationNetworks`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
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
		fabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fabricName")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListByReplicationFabricsPager(resourceGroupNameParam, resourceNameParam, fabricNameParam, nil)
		newListByReplicationFabricsPager = &resp
		r.newListByReplicationFabricsPager.add(req, newListByReplicationFabricsPager)
		server.PagerResponderInjectNextLinks(newListByReplicationFabricsPager, req, func(page *armrecoveryservicessiterecovery.ReplicationNetworksClientListByReplicationFabricsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByReplicationFabricsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListByReplicationFabricsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByReplicationFabricsPager) {
		r.newListByReplicationFabricsPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ReplicationNetworksServerTransport
var replicationNetworksServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
