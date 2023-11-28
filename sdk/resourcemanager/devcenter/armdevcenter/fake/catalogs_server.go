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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// CatalogsServer is a fake server for instances of the armdevcenter.CatalogsClient type.
type CatalogsServer struct {
	// BeginConnect is the fake for method CatalogsClient.BeginConnect
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginConnect func(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, options *armdevcenter.CatalogsClientBeginConnectOptions) (resp azfake.PollerResponder[armdevcenter.CatalogsClientConnectResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method CatalogsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, body armdevcenter.Catalog, options *armdevcenter.CatalogsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armdevcenter.CatalogsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method CatalogsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, options *armdevcenter.CatalogsClientBeginDeleteOptions) (resp azfake.PollerResponder[armdevcenter.CatalogsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method CatalogsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, options *armdevcenter.CatalogsClientGetOptions) (resp azfake.Responder[armdevcenter.CatalogsClientGetResponse], errResp azfake.ErrorResponder)

	// GetSyncErrorDetails is the fake for method CatalogsClient.GetSyncErrorDetails
	// HTTP status codes to indicate success: http.StatusOK
	GetSyncErrorDetails func(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, options *armdevcenter.CatalogsClientGetSyncErrorDetailsOptions) (resp azfake.Responder[armdevcenter.CatalogsClientGetSyncErrorDetailsResponse], errResp azfake.ErrorResponder)

	// NewListByDevCenterPager is the fake for method CatalogsClient.NewListByDevCenterPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByDevCenterPager func(resourceGroupName string, devCenterName string, options *armdevcenter.CatalogsClientListByDevCenterOptions) (resp azfake.PagerResponder[armdevcenter.CatalogsClientListByDevCenterResponse])

	// BeginSync is the fake for method CatalogsClient.BeginSync
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginSync func(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, options *armdevcenter.CatalogsClientBeginSyncOptions) (resp azfake.PollerResponder[armdevcenter.CatalogsClientSyncResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method CatalogsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, body armdevcenter.CatalogUpdate, options *armdevcenter.CatalogsClientBeginUpdateOptions) (resp azfake.PollerResponder[armdevcenter.CatalogsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewCatalogsServerTransport creates a new instance of CatalogsServerTransport with the provided implementation.
// The returned CatalogsServerTransport instance is connected to an instance of armdevcenter.CatalogsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCatalogsServerTransport(srv *CatalogsServer) *CatalogsServerTransport {
	return &CatalogsServerTransport{
		srv:                     srv,
		beginConnect:            newTracker[azfake.PollerResponder[armdevcenter.CatalogsClientConnectResponse]](),
		beginCreateOrUpdate:     newTracker[azfake.PollerResponder[armdevcenter.CatalogsClientCreateOrUpdateResponse]](),
		beginDelete:             newTracker[azfake.PollerResponder[armdevcenter.CatalogsClientDeleteResponse]](),
		newListByDevCenterPager: newTracker[azfake.PagerResponder[armdevcenter.CatalogsClientListByDevCenterResponse]](),
		beginSync:               newTracker[azfake.PollerResponder[armdevcenter.CatalogsClientSyncResponse]](),
		beginUpdate:             newTracker[azfake.PollerResponder[armdevcenter.CatalogsClientUpdateResponse]](),
	}
}

// CatalogsServerTransport connects instances of armdevcenter.CatalogsClient to instances of CatalogsServer.
// Don't use this type directly, use NewCatalogsServerTransport instead.
type CatalogsServerTransport struct {
	srv                     *CatalogsServer
	beginConnect            *tracker[azfake.PollerResponder[armdevcenter.CatalogsClientConnectResponse]]
	beginCreateOrUpdate     *tracker[azfake.PollerResponder[armdevcenter.CatalogsClientCreateOrUpdateResponse]]
	beginDelete             *tracker[azfake.PollerResponder[armdevcenter.CatalogsClientDeleteResponse]]
	newListByDevCenterPager *tracker[azfake.PagerResponder[armdevcenter.CatalogsClientListByDevCenterResponse]]
	beginSync               *tracker[azfake.PollerResponder[armdevcenter.CatalogsClientSyncResponse]]
	beginUpdate             *tracker[azfake.PollerResponder[armdevcenter.CatalogsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for CatalogsServerTransport.
func (c *CatalogsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CatalogsClient.BeginConnect":
		resp, err = c.dispatchBeginConnect(req)
	case "CatalogsClient.BeginCreateOrUpdate":
		resp, err = c.dispatchBeginCreateOrUpdate(req)
	case "CatalogsClient.BeginDelete":
		resp, err = c.dispatchBeginDelete(req)
	case "CatalogsClient.Get":
		resp, err = c.dispatchGet(req)
	case "CatalogsClient.GetSyncErrorDetails":
		resp, err = c.dispatchGetSyncErrorDetails(req)
	case "CatalogsClient.NewListByDevCenterPager":
		resp, err = c.dispatchNewListByDevCenterPager(req)
	case "CatalogsClient.BeginSync":
		resp, err = c.dispatchBeginSync(req)
	case "CatalogsClient.BeginUpdate":
		resp, err = c.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CatalogsServerTransport) dispatchBeginConnect(req *http.Request) (*http.Response, error) {
	if c.srv.BeginConnect == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginConnect not implemented")}
	}
	beginConnect := c.beginConnect.get(req)
	if beginConnect == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/devcenters/(?P<devCenterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/connect`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		devCenterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("devCenterName")])
		if err != nil {
			return nil, err
		}
		catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginConnect(req.Context(), resourceGroupNameParam, devCenterNameParam, catalogNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginConnect = &respr
		c.beginConnect.add(req, beginConnect)
	}

	resp, err := server.PollerResponderNext(beginConnect, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		c.beginConnect.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginConnect) {
		c.beginConnect.remove(req)
	}

	return resp, nil
}

func (c *CatalogsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := c.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/devcenters/(?P<devCenterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdevcenter.Catalog](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		devCenterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("devCenterName")])
		if err != nil {
			return nil, err
		}
		catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, devCenterNameParam, catalogNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		c.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusCreated}, resp.StatusCode) {
		c.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		c.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (c *CatalogsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/devcenters/(?P<devCenterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		devCenterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("devCenterName")])
		if err != nil {
			return nil, err
		}
		catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, devCenterNameParam, catalogNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		c.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		c.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		c.beginDelete.remove(req)
	}

	return resp, nil
}

func (c *CatalogsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/devcenters/(?P<devCenterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	devCenterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("devCenterName")])
	if err != nil {
		return nil, err
	}
	catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, devCenterNameParam, catalogNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Catalog, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CatalogsServerTransport) dispatchGetSyncErrorDetails(req *http.Request) (*http.Response, error) {
	if c.srv.GetSyncErrorDetails == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetSyncErrorDetails not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/devcenters/(?P<devCenterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getSyncErrorDetails`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	devCenterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("devCenterName")])
	if err != nil {
		return nil, err
	}
	catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.GetSyncErrorDetails(req.Context(), resourceGroupNameParam, devCenterNameParam, catalogNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SyncErrorDetails, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CatalogsServerTransport) dispatchNewListByDevCenterPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListByDevCenterPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByDevCenterPager not implemented")}
	}
	newListByDevCenterPager := c.newListByDevCenterPager.get(req)
	if newListByDevCenterPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/devcenters/(?P<devCenterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/catalogs`
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
		devCenterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("devCenterName")])
		if err != nil {
			return nil, err
		}
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armdevcenter.CatalogsClientListByDevCenterOptions
		if topParam != nil {
			options = &armdevcenter.CatalogsClientListByDevCenterOptions{
				Top: topParam,
			}
		}
		resp := c.srv.NewListByDevCenterPager(resourceGroupNameParam, devCenterNameParam, options)
		newListByDevCenterPager = &resp
		c.newListByDevCenterPager.add(req, newListByDevCenterPager)
		server.PagerResponderInjectNextLinks(newListByDevCenterPager, req, func(page *armdevcenter.CatalogsClientListByDevCenterResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByDevCenterPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListByDevCenterPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByDevCenterPager) {
		c.newListByDevCenterPager.remove(req)
	}
	return resp, nil
}

func (c *CatalogsServerTransport) dispatchBeginSync(req *http.Request) (*http.Response, error) {
	if c.srv.BeginSync == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginSync not implemented")}
	}
	beginSync := c.beginSync.get(req)
	if beginSync == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/devcenters/(?P<devCenterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sync`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		devCenterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("devCenterName")])
		if err != nil {
			return nil, err
		}
		catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginSync(req.Context(), resourceGroupNameParam, devCenterNameParam, catalogNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginSync = &respr
		c.beginSync.add(req, beginSync)
	}

	resp, err := server.PollerResponderNext(beginSync, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		c.beginSync.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginSync) {
		c.beginSync.remove(req)
	}

	return resp, nil
}

func (c *CatalogsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := c.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/devcenters/(?P<devCenterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdevcenter.CatalogUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		devCenterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("devCenterName")])
		if err != nil {
			return nil, err
		}
		catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginUpdate(req.Context(), resourceGroupNameParam, devCenterNameParam, catalogNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		c.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		c.beginUpdate.remove(req)
	}

	return resp, nil
}