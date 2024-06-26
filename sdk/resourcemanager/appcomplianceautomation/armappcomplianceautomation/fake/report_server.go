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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// ReportServer is a fake server for instances of the armappcomplianceautomation.ReportClient type.
type ReportServer struct {
	// BeginCreateOrUpdate is the fake for method ReportClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, reportName string, properties armappcomplianceautomation.ReportResource, options *armappcomplianceautomation.ReportClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armappcomplianceautomation.ReportClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ReportClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, reportName string, options *armappcomplianceautomation.ReportClientBeginDeleteOptions) (resp azfake.PollerResponder[armappcomplianceautomation.ReportClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginFix is the fake for method ReportClient.BeginFix
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginFix func(ctx context.Context, reportName string, options *armappcomplianceautomation.ReportClientBeginFixOptions) (resp azfake.PollerResponder[armappcomplianceautomation.ReportClientFixResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ReportClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, reportName string, options *armappcomplianceautomation.ReportClientGetOptions) (resp azfake.Responder[armappcomplianceautomation.ReportClientGetResponse], errResp azfake.ErrorResponder)

	// GetScopingQuestions is the fake for method ReportClient.GetScopingQuestions
	// HTTP status codes to indicate success: http.StatusOK
	GetScopingQuestions func(ctx context.Context, reportName string, options *armappcomplianceautomation.ReportClientGetScopingQuestionsOptions) (resp azfake.Responder[armappcomplianceautomation.ReportClientGetScopingQuestionsResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ReportClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armappcomplianceautomation.ReportClientListOptions) (resp azfake.PagerResponder[armappcomplianceautomation.ReportClientListResponse])

	// NestedResourceCheckNameAvailability is the fake for method ReportClient.NestedResourceCheckNameAvailability
	// HTTP status codes to indicate success: http.StatusOK
	NestedResourceCheckNameAvailability func(ctx context.Context, reportName string, body armappcomplianceautomation.CheckNameAvailabilityRequest, options *armappcomplianceautomation.ReportClientNestedResourceCheckNameAvailabilityOptions) (resp azfake.Responder[armappcomplianceautomation.ReportClientNestedResourceCheckNameAvailabilityResponse], errResp azfake.ErrorResponder)

	// BeginSyncCertRecord is the fake for method ReportClient.BeginSyncCertRecord
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginSyncCertRecord func(ctx context.Context, reportName string, body armappcomplianceautomation.SyncCertRecordRequest, options *armappcomplianceautomation.ReportClientBeginSyncCertRecordOptions) (resp azfake.PollerResponder[armappcomplianceautomation.ReportClientSyncCertRecordResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method ReportClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, reportName string, properties armappcomplianceautomation.ReportResourcePatch, options *armappcomplianceautomation.ReportClientBeginUpdateOptions) (resp azfake.PollerResponder[armappcomplianceautomation.ReportClientUpdateResponse], errResp azfake.ErrorResponder)

	// BeginVerify is the fake for method ReportClient.BeginVerify
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginVerify func(ctx context.Context, reportName string, options *armappcomplianceautomation.ReportClientBeginVerifyOptions) (resp azfake.PollerResponder[armappcomplianceautomation.ReportClientVerifyResponse], errResp azfake.ErrorResponder)
}

// NewReportServerTransport creates a new instance of ReportServerTransport with the provided implementation.
// The returned ReportServerTransport instance is connected to an instance of armappcomplianceautomation.ReportClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewReportServerTransport(srv *ReportServer) *ReportServerTransport {
	return &ReportServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armappcomplianceautomation.ReportClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armappcomplianceautomation.ReportClientDeleteResponse]](),
		beginFix:            newTracker[azfake.PollerResponder[armappcomplianceautomation.ReportClientFixResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armappcomplianceautomation.ReportClientListResponse]](),
		beginSyncCertRecord: newTracker[azfake.PollerResponder[armappcomplianceautomation.ReportClientSyncCertRecordResponse]](),
		beginUpdate:         newTracker[azfake.PollerResponder[armappcomplianceautomation.ReportClientUpdateResponse]](),
		beginVerify:         newTracker[azfake.PollerResponder[armappcomplianceautomation.ReportClientVerifyResponse]](),
	}
}

// ReportServerTransport connects instances of armappcomplianceautomation.ReportClient to instances of ReportServer.
// Don't use this type directly, use NewReportServerTransport instead.
type ReportServerTransport struct {
	srv                 *ReportServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armappcomplianceautomation.ReportClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armappcomplianceautomation.ReportClientDeleteResponse]]
	beginFix            *tracker[azfake.PollerResponder[armappcomplianceautomation.ReportClientFixResponse]]
	newListPager        *tracker[azfake.PagerResponder[armappcomplianceautomation.ReportClientListResponse]]
	beginSyncCertRecord *tracker[azfake.PollerResponder[armappcomplianceautomation.ReportClientSyncCertRecordResponse]]
	beginUpdate         *tracker[azfake.PollerResponder[armappcomplianceautomation.ReportClientUpdateResponse]]
	beginVerify         *tracker[azfake.PollerResponder[armappcomplianceautomation.ReportClientVerifyResponse]]
}

// Do implements the policy.Transporter interface for ReportServerTransport.
func (r *ReportServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ReportClient.BeginCreateOrUpdate":
		resp, err = r.dispatchBeginCreateOrUpdate(req)
	case "ReportClient.BeginDelete":
		resp, err = r.dispatchBeginDelete(req)
	case "ReportClient.BeginFix":
		resp, err = r.dispatchBeginFix(req)
	case "ReportClient.Get":
		resp, err = r.dispatchGet(req)
	case "ReportClient.GetScopingQuestions":
		resp, err = r.dispatchGetScopingQuestions(req)
	case "ReportClient.NewListPager":
		resp, err = r.dispatchNewListPager(req)
	case "ReportClient.NestedResourceCheckNameAvailability":
		resp, err = r.dispatchNestedResourceCheckNameAvailability(req)
	case "ReportClient.BeginSyncCertRecord":
		resp, err = r.dispatchBeginSyncCertRecord(req)
	case "ReportClient.BeginUpdate":
		resp, err = r.dispatchBeginUpdate(req)
	case "ReportClient.BeginVerify":
		resp, err = r.dispatchBeginVerify(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *ReportServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := r.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/providers/Microsoft\.AppComplianceAutomation/reports/(?P<reportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armappcomplianceautomation.ReportResource](req)
		if err != nil {
			return nil, err
		}
		reportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reportName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginCreateOrUpdate(req.Context(), reportNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		r.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		r.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		r.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (r *ReportServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if r.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := r.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/providers/Microsoft\.AppComplianceAutomation/reports/(?P<reportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		reportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reportName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginDelete(req.Context(), reportNameParam, nil)
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

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		r.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		r.beginDelete.remove(req)
	}

	return resp, nil
}

func (r *ReportServerTransport) dispatchBeginFix(req *http.Request) (*http.Response, error) {
	if r.srv.BeginFix == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginFix not implemented")}
	}
	beginFix := r.beginFix.get(req)
	if beginFix == nil {
		const regexStr = `/providers/Microsoft\.AppComplianceAutomation/reports/(?P<reportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fix`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		reportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reportName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginFix(req.Context(), reportNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginFix = &respr
		r.beginFix.add(req, beginFix)
	}

	resp, err := server.PollerResponderNext(beginFix, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		r.beginFix.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginFix) {
		r.beginFix.remove(req)
	}

	return resp, nil
}

func (r *ReportServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/providers/Microsoft\.AppComplianceAutomation/reports/(?P<reportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	reportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reportName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), reportNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ReportResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *ReportServerTransport) dispatchGetScopingQuestions(req *http.Request) (*http.Response, error) {
	if r.srv.GetScopingQuestions == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetScopingQuestions not implemented")}
	}
	const regexStr = `/providers/Microsoft\.AppComplianceAutomation/reports/(?P<reportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getScopingQuestions`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	reportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reportName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.GetScopingQuestions(req.Context(), reportNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ScopingQuestions, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *ReportServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		qp := req.URL.Query()
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
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
		selectUnescaped, err := url.QueryUnescape(qp.Get("$select"))
		if err != nil {
			return nil, err
		}
		selectParam := getOptional(selectUnescaped)
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		orderbyUnescaped, err := url.QueryUnescape(qp.Get("$orderby"))
		if err != nil {
			return nil, err
		}
		orderbyParam := getOptional(orderbyUnescaped)
		offerGUIDUnescaped, err := url.QueryUnescape(qp.Get("offerGuid"))
		if err != nil {
			return nil, err
		}
		offerGUIDParam := getOptional(offerGUIDUnescaped)
		reportCreatorTenantIDUnescaped, err := url.QueryUnescape(qp.Get("reportCreatorTenantId"))
		if err != nil {
			return nil, err
		}
		reportCreatorTenantIDParam := getOptional(reportCreatorTenantIDUnescaped)
		var options *armappcomplianceautomation.ReportClientListOptions
		if skipTokenParam != nil || topParam != nil || selectParam != nil || filterParam != nil || orderbyParam != nil || offerGUIDParam != nil || reportCreatorTenantIDParam != nil {
			options = &armappcomplianceautomation.ReportClientListOptions{
				SkipToken:             skipTokenParam,
				Top:                   topParam,
				Select:                selectParam,
				Filter:                filterParam,
				Orderby:               orderbyParam,
				OfferGUID:             offerGUIDParam,
				ReportCreatorTenantID: reportCreatorTenantIDParam,
			}
		}
		resp := r.srv.NewListPager(options)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armappcomplianceautomation.ReportClientListResponse, createLink func() string) {
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

func (r *ReportServerTransport) dispatchNestedResourceCheckNameAvailability(req *http.Request) (*http.Response, error) {
	if r.srv.NestedResourceCheckNameAvailability == nil {
		return nil, &nonRetriableError{errors.New("fake for method NestedResourceCheckNameAvailability not implemented")}
	}
	const regexStr = `/providers/Microsoft\.AppComplianceAutomation/reports/(?P<reportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/checkNameAvailability`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armappcomplianceautomation.CheckNameAvailabilityRequest](req)
	if err != nil {
		return nil, err
	}
	reportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reportName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.NestedResourceCheckNameAvailability(req.Context(), reportNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CheckNameAvailabilityResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *ReportServerTransport) dispatchBeginSyncCertRecord(req *http.Request) (*http.Response, error) {
	if r.srv.BeginSyncCertRecord == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginSyncCertRecord not implemented")}
	}
	beginSyncCertRecord := r.beginSyncCertRecord.get(req)
	if beginSyncCertRecord == nil {
		const regexStr = `/providers/Microsoft\.AppComplianceAutomation/reports/(?P<reportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncCertRecord`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armappcomplianceautomation.SyncCertRecordRequest](req)
		if err != nil {
			return nil, err
		}
		reportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reportName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginSyncCertRecord(req.Context(), reportNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginSyncCertRecord = &respr
		r.beginSyncCertRecord.add(req, beginSyncCertRecord)
	}

	resp, err := server.PollerResponderNext(beginSyncCertRecord, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		r.beginSyncCertRecord.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginSyncCertRecord) {
		r.beginSyncCertRecord.remove(req)
	}

	return resp, nil
}

func (r *ReportServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := r.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/providers/Microsoft\.AppComplianceAutomation/reports/(?P<reportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armappcomplianceautomation.ReportResourcePatch](req)
		if err != nil {
			return nil, err
		}
		reportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reportName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginUpdate(req.Context(), reportNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		r.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		r.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		r.beginUpdate.remove(req)
	}

	return resp, nil
}

func (r *ReportServerTransport) dispatchBeginVerify(req *http.Request) (*http.Response, error) {
	if r.srv.BeginVerify == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginVerify not implemented")}
	}
	beginVerify := r.beginVerify.get(req)
	if beginVerify == nil {
		const regexStr = `/providers/Microsoft\.AppComplianceAutomation/reports/(?P<reportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/verify`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		reportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reportName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginVerify(req.Context(), reportNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginVerify = &respr
		r.beginVerify.add(req, beginVerify)
	}

	resp, err := server.PollerResponderNext(beginVerify, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		r.beginVerify.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginVerify) {
		r.beginVerify.remove(req)
	}

	return resp, nil
}
