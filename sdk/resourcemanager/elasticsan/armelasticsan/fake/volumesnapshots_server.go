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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
	"net/http"
	"net/url"
	"regexp"
)

// VolumeSnapshotsServer is a fake server for instances of the armelasticsan.VolumeSnapshotsClient type.
type VolumeSnapshotsServer struct {
	// BeginCreate is the fake for method VolumeSnapshotsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, elasticSanName string, volumeGroupName string, snapshotName string, parameters armelasticsan.Snapshot, options *armelasticsan.VolumeSnapshotsClientBeginCreateOptions) (resp azfake.PollerResponder[armelasticsan.VolumeSnapshotsClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VolumeSnapshotsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, elasticSanName string, volumeGroupName string, snapshotName string, options *armelasticsan.VolumeSnapshotsClientBeginDeleteOptions) (resp azfake.PollerResponder[armelasticsan.VolumeSnapshotsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VolumeSnapshotsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, elasticSanName string, volumeGroupName string, snapshotName string, options *armelasticsan.VolumeSnapshotsClientGetOptions) (resp azfake.Responder[armelasticsan.VolumeSnapshotsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByVolumeGroupPager is the fake for method VolumeSnapshotsClient.NewListByVolumeGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByVolumeGroupPager func(resourceGroupName string, elasticSanName string, volumeGroupName string, options *armelasticsan.VolumeSnapshotsClientListByVolumeGroupOptions) (resp azfake.PagerResponder[armelasticsan.VolumeSnapshotsClientListByVolumeGroupResponse])
}

// NewVolumeSnapshotsServerTransport creates a new instance of VolumeSnapshotsServerTransport with the provided implementation.
// The returned VolumeSnapshotsServerTransport instance is connected to an instance of armelasticsan.VolumeSnapshotsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVolumeSnapshotsServerTransport(srv *VolumeSnapshotsServer) *VolumeSnapshotsServerTransport {
	return &VolumeSnapshotsServerTransport{
		srv:                       srv,
		beginCreate:               newTracker[azfake.PollerResponder[armelasticsan.VolumeSnapshotsClientCreateResponse]](),
		beginDelete:               newTracker[azfake.PollerResponder[armelasticsan.VolumeSnapshotsClientDeleteResponse]](),
		newListByVolumeGroupPager: newTracker[azfake.PagerResponder[armelasticsan.VolumeSnapshotsClientListByVolumeGroupResponse]](),
	}
}

// VolumeSnapshotsServerTransport connects instances of armelasticsan.VolumeSnapshotsClient to instances of VolumeSnapshotsServer.
// Don't use this type directly, use NewVolumeSnapshotsServerTransport instead.
type VolumeSnapshotsServerTransport struct {
	srv                       *VolumeSnapshotsServer
	beginCreate               *tracker[azfake.PollerResponder[armelasticsan.VolumeSnapshotsClientCreateResponse]]
	beginDelete               *tracker[azfake.PollerResponder[armelasticsan.VolumeSnapshotsClientDeleteResponse]]
	newListByVolumeGroupPager *tracker[azfake.PagerResponder[armelasticsan.VolumeSnapshotsClientListByVolumeGroupResponse]]
}

// Do implements the policy.Transporter interface for VolumeSnapshotsServerTransport.
func (v *VolumeSnapshotsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return v.dispatchToMethodFake(req, method)
}

func (v *VolumeSnapshotsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if volumeSnapshotsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = volumeSnapshotsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "VolumeSnapshotsClient.BeginCreate":
				res.resp, res.err = v.dispatchBeginCreate(req)
			case "VolumeSnapshotsClient.BeginDelete":
				res.resp, res.err = v.dispatchBeginDelete(req)
			case "VolumeSnapshotsClient.Get":
				res.resp, res.err = v.dispatchGet(req)
			case "VolumeSnapshotsClient.NewListByVolumeGroupPager":
				res.resp, res.err = v.dispatchNewListByVolumeGroupPager(req)
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

func (v *VolumeSnapshotsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := v.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ElasticSan/elasticSans/(?P<elasticSanName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/volumegroups/(?P<volumeGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/snapshots/(?P<snapshotName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armelasticsan.Snapshot](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		elasticSanNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("elasticSanName")])
		if err != nil {
			return nil, err
		}
		volumeGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("volumeGroupName")])
		if err != nil {
			return nil, err
		}
		snapshotNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("snapshotName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginCreate(req.Context(), resourceGroupNameParam, elasticSanNameParam, volumeGroupNameParam, snapshotNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		v.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		v.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		v.beginCreate.remove(req)
	}

	return resp, nil
}

func (v *VolumeSnapshotsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := v.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ElasticSan/elasticSans/(?P<elasticSanName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/volumegroups/(?P<volumeGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/snapshots/(?P<snapshotName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		elasticSanNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("elasticSanName")])
		if err != nil {
			return nil, err
		}
		volumeGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("volumeGroupName")])
		if err != nil {
			return nil, err
		}
		snapshotNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("snapshotName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), resourceGroupNameParam, elasticSanNameParam, volumeGroupNameParam, snapshotNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		v.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		v.beginDelete.remove(req)
	}

	return resp, nil
}

func (v *VolumeSnapshotsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ElasticSan/elasticSans/(?P<elasticSanName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/volumegroups/(?P<volumeGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/snapshots/(?P<snapshotName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	elasticSanNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("elasticSanName")])
	if err != nil {
		return nil, err
	}
	volumeGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("volumeGroupName")])
	if err != nil {
		return nil, err
	}
	snapshotNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("snapshotName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameParam, elasticSanNameParam, volumeGroupNameParam, snapshotNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Snapshot, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VolumeSnapshotsServerTransport) dispatchNewListByVolumeGroupPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListByVolumeGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByVolumeGroupPager not implemented")}
	}
	newListByVolumeGroupPager := v.newListByVolumeGroupPager.get(req)
	if newListByVolumeGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ElasticSan/elasticSans/(?P<elasticSanName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/volumegroups/(?P<volumeGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/snapshots`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		elasticSanNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("elasticSanName")])
		if err != nil {
			return nil, err
		}
		volumeGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("volumeGroupName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armelasticsan.VolumeSnapshotsClientListByVolumeGroupOptions
		if filterParam != nil {
			options = &armelasticsan.VolumeSnapshotsClientListByVolumeGroupOptions{
				Filter: filterParam,
			}
		}
		resp := v.srv.NewListByVolumeGroupPager(resourceGroupNameParam, elasticSanNameParam, volumeGroupNameParam, options)
		newListByVolumeGroupPager = &resp
		v.newListByVolumeGroupPager.add(req, newListByVolumeGroupPager)
		server.PagerResponderInjectNextLinks(newListByVolumeGroupPager, req, func(page *armelasticsan.VolumeSnapshotsClientListByVolumeGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByVolumeGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListByVolumeGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByVolumeGroupPager) {
		v.newListByVolumeGroupPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to VolumeSnapshotsServerTransport
var volumeSnapshotsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
