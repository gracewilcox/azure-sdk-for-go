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

// ReplicationRecoveryServicesProvidersServer is a fake server for instances of the armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient type.
type ReplicationRecoveryServicesProvidersServer struct {
	// BeginCreate is the fake for method ReplicationRecoveryServicesProvidersClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreate func(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, providerName string, addProviderInput armrecoveryservicessiterecovery.AddRecoveryServicesProviderInput, options *armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientBeginCreateOptions) (resp azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ReplicationRecoveryServicesProvidersClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, providerName string, options *armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientBeginDeleteOptions) (resp azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ReplicationRecoveryServicesProvidersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, providerName string, options *armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientGetOptions) (resp azfake.Responder[armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ReplicationRecoveryServicesProvidersClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, resourceName string, options *armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientListOptions) (resp azfake.PagerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientListResponse])

	// NewListByReplicationFabricsPager is the fake for method ReplicationRecoveryServicesProvidersClient.NewListByReplicationFabricsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByReplicationFabricsPager func(resourceGroupName string, resourceName string, fabricName string, options *armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientListByReplicationFabricsOptions) (resp azfake.PagerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientListByReplicationFabricsResponse])

	// BeginPurge is the fake for method ReplicationRecoveryServicesProvidersClient.BeginPurge
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginPurge func(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, providerName string, options *armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientBeginPurgeOptions) (resp azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientPurgeResponse], errResp azfake.ErrorResponder)

	// BeginRefreshProvider is the fake for method ReplicationRecoveryServicesProvidersClient.BeginRefreshProvider
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRefreshProvider func(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, providerName string, options *armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientBeginRefreshProviderOptions) (resp azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientRefreshProviderResponse], errResp azfake.ErrorResponder)
}

// NewReplicationRecoveryServicesProvidersServerTransport creates a new instance of ReplicationRecoveryServicesProvidersServerTransport with the provided implementation.
// The returned ReplicationRecoveryServicesProvidersServerTransport instance is connected to an instance of armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewReplicationRecoveryServicesProvidersServerTransport(srv *ReplicationRecoveryServicesProvidersServer) *ReplicationRecoveryServicesProvidersServerTransport {
	return &ReplicationRecoveryServicesProvidersServerTransport{
		srv:                              srv,
		beginCreate:                      newTracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientCreateResponse]](),
		beginDelete:                      newTracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientDeleteResponse]](),
		newListPager:                     newTracker[azfake.PagerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientListResponse]](),
		newListByReplicationFabricsPager: newTracker[azfake.PagerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientListByReplicationFabricsResponse]](),
		beginPurge:                       newTracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientPurgeResponse]](),
		beginRefreshProvider:             newTracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientRefreshProviderResponse]](),
	}
}

// ReplicationRecoveryServicesProvidersServerTransport connects instances of armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClient to instances of ReplicationRecoveryServicesProvidersServer.
// Don't use this type directly, use NewReplicationRecoveryServicesProvidersServerTransport instead.
type ReplicationRecoveryServicesProvidersServerTransport struct {
	srv                              *ReplicationRecoveryServicesProvidersServer
	beginCreate                      *tracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientCreateResponse]]
	beginDelete                      *tracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientDeleteResponse]]
	newListPager                     *tracker[azfake.PagerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientListResponse]]
	newListByReplicationFabricsPager *tracker[azfake.PagerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientListByReplicationFabricsResponse]]
	beginPurge                       *tracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientPurgeResponse]]
	beginRefreshProvider             *tracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientRefreshProviderResponse]]
}

// Do implements the policy.Transporter interface for ReplicationRecoveryServicesProvidersServerTransport.
func (r *ReplicationRecoveryServicesProvidersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return r.dispatchToMethodFake(req, method)
}

func (r *ReplicationRecoveryServicesProvidersServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if replicationRecoveryServicesProvidersServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = replicationRecoveryServicesProvidersServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ReplicationRecoveryServicesProvidersClient.BeginCreate":
				res.resp, res.err = r.dispatchBeginCreate(req)
			case "ReplicationRecoveryServicesProvidersClient.BeginDelete":
				res.resp, res.err = r.dispatchBeginDelete(req)
			case "ReplicationRecoveryServicesProvidersClient.Get":
				res.resp, res.err = r.dispatchGet(req)
			case "ReplicationRecoveryServicesProvidersClient.NewListPager":
				res.resp, res.err = r.dispatchNewListPager(req)
			case "ReplicationRecoveryServicesProvidersClient.NewListByReplicationFabricsPager":
				res.resp, res.err = r.dispatchNewListByReplicationFabricsPager(req)
			case "ReplicationRecoveryServicesProvidersClient.BeginPurge":
				res.resp, res.err = r.dispatchBeginPurge(req)
			case "ReplicationRecoveryServicesProvidersClient.BeginRefreshProvider":
				res.resp, res.err = r.dispatchBeginRefreshProvider(req)
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

func (r *ReplicationRecoveryServicesProvidersServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := r.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationRecoveryServicesProviders/(?P<providerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armrecoveryservicessiterecovery.AddRecoveryServicesProviderInput](req)
		if err != nil {
			return nil, err
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
		providerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("providerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginCreate(req.Context(), resourceGroupNameParam, resourceNameParam, fabricNameParam, providerNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		r.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		r.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		r.beginCreate.remove(req)
	}

	return resp, nil
}

func (r *ReplicationRecoveryServicesProvidersServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if r.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := r.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationRecoveryServicesProviders/(?P<providerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/remove`
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
		providerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("providerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginDelete(req.Context(), resourceGroupNameParam, resourceNameParam, fabricNameParam, providerNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		r.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		r.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		r.beginDelete.remove(req)
	}

	return resp, nil
}

func (r *ReplicationRecoveryServicesProvidersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationRecoveryServicesProviders/(?P<providerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	providerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("providerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameParam, resourceNameParam, fabricNameParam, providerNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RecoveryServicesProvider, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *ReplicationRecoveryServicesProvidersServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationRecoveryServicesProviders`
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
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientListResponse, createLink func() string) {
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

func (r *ReplicationRecoveryServicesProvidersServerTransport) dispatchNewListByReplicationFabricsPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListByReplicationFabricsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByReplicationFabricsPager not implemented")}
	}
	newListByReplicationFabricsPager := r.newListByReplicationFabricsPager.get(req)
	if newListByReplicationFabricsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationRecoveryServicesProviders`
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
		server.PagerResponderInjectNextLinks(newListByReplicationFabricsPager, req, func(page *armrecoveryservicessiterecovery.ReplicationRecoveryServicesProvidersClientListByReplicationFabricsResponse, createLink func() string) {
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

func (r *ReplicationRecoveryServicesProvidersServerTransport) dispatchBeginPurge(req *http.Request) (*http.Response, error) {
	if r.srv.BeginPurge == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPurge not implemented")}
	}
	beginPurge := r.beginPurge.get(req)
	if beginPurge == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationRecoveryServicesProviders/(?P<providerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
		providerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("providerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginPurge(req.Context(), resourceGroupNameParam, resourceNameParam, fabricNameParam, providerNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPurge = &respr
		r.beginPurge.add(req, beginPurge)
	}

	resp, err := server.PollerResponderNext(beginPurge, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		r.beginPurge.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPurge) {
		r.beginPurge.remove(req)
	}

	return resp, nil
}

func (r *ReplicationRecoveryServicesProvidersServerTransport) dispatchBeginRefreshProvider(req *http.Request) (*http.Response, error) {
	if r.srv.BeginRefreshProvider == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRefreshProvider not implemented")}
	}
	beginRefreshProvider := r.beginRefreshProvider.get(req)
	if beginRefreshProvider == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationRecoveryServicesProviders/(?P<providerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/refreshProvider`
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
		providerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("providerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginRefreshProvider(req.Context(), resourceGroupNameParam, resourceNameParam, fabricNameParam, providerNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRefreshProvider = &respr
		r.beginRefreshProvider.add(req, beginRefreshProvider)
	}

	resp, err := server.PollerResponderNext(beginRefreshProvider, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		r.beginRefreshProvider.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRefreshProvider) {
		r.beginRefreshProvider.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to ReplicationRecoveryServicesProvidersServerTransport
var replicationRecoveryServicesProvidersServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
