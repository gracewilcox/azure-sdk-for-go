// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"

	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azsecrets"
)

// Server is a fake server for instances of the azsecrets.Client type.
type Server struct {
	// BackupSecret is the fake for method Client.BackupSecret
	// HTTP status codes to indicate success: http.StatusOK
	BackupSecret func(ctx context.Context, name string, options *azsecrets.BackupSecretOptions) (resp azfake.Responder[azsecrets.BackupSecretResponse], errResp azfake.ErrorResponder)

	// DeleteSecret is the fake for method Client.DeleteSecret
	// HTTP status codes to indicate success: http.StatusOK
	DeleteSecret func(ctx context.Context, name string, options *azsecrets.DeleteSecretOptions) (resp azfake.Responder[azsecrets.DeleteSecretResponse], errResp azfake.ErrorResponder)

	// GetDeletedSecret is the fake for method Client.GetDeletedSecret
	// HTTP status codes to indicate success: http.StatusOK
	GetDeletedSecret func(ctx context.Context, name string, options *azsecrets.GetDeletedSecretOptions) (resp azfake.Responder[azsecrets.GetDeletedSecretResponse], errResp azfake.ErrorResponder)

	// GetSecret is the fake for method Client.GetSecret
	// HTTP status codes to indicate success: http.StatusOK
	GetSecret func(ctx context.Context, name string, version string, options *azsecrets.GetSecretOptions) (resp azfake.Responder[azsecrets.GetSecretResponse], errResp azfake.ErrorResponder)

	// NewListDeletedSecretPropertiesPager is the fake for method Client.NewListDeletedSecretPropertiesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListDeletedSecretPropertiesPager func(options *azsecrets.ListDeletedSecretPropertiesOptions) (resp azfake.PagerResponder[azsecrets.ListDeletedSecretPropertiesResponse])

	// NewListSecretPropertiesPager is the fake for method Client.NewListSecretPropertiesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListSecretPropertiesPager func(options *azsecrets.ListSecretPropertiesOptions) (resp azfake.PagerResponder[azsecrets.ListSecretPropertiesResponse])

	// NewListSecretPropertiesVersionsPager is the fake for method Client.NewListSecretPropertiesVersionsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListSecretPropertiesVersionsPager func(name string, options *azsecrets.ListSecretPropertiesVersionsOptions) (resp azfake.PagerResponder[azsecrets.ListSecretPropertiesVersionsResponse])

	// PurgeDeletedSecret is the fake for method Client.PurgeDeletedSecret
	// HTTP status codes to indicate success: http.StatusNoContent
	PurgeDeletedSecret func(ctx context.Context, name string, options *azsecrets.PurgeDeletedSecretOptions) (resp azfake.Responder[azsecrets.PurgeDeletedSecretResponse], errResp azfake.ErrorResponder)

	// RecoverDeletedSecret is the fake for method Client.RecoverDeletedSecret
	// HTTP status codes to indicate success: http.StatusOK
	RecoverDeletedSecret func(ctx context.Context, name string, options *azsecrets.RecoverDeletedSecretOptions) (resp azfake.Responder[azsecrets.RecoverDeletedSecretResponse], errResp azfake.ErrorResponder)

	// RestoreSecret is the fake for method Client.RestoreSecret
	// HTTP status codes to indicate success: http.StatusOK
	RestoreSecret func(ctx context.Context, parameters azsecrets.RestoreSecretParameters, options *azsecrets.RestoreSecretOptions) (resp azfake.Responder[azsecrets.RestoreSecretResponse], errResp azfake.ErrorResponder)

	// SetSecret is the fake for method Client.SetSecret
	// HTTP status codes to indicate success: http.StatusOK
	SetSecret func(ctx context.Context, name string, parameters azsecrets.SetSecretParameters, options *azsecrets.SetSecretOptions) (resp azfake.Responder[azsecrets.SetSecretResponse], errResp azfake.ErrorResponder)

	// UpdateSecretProperties is the fake for method Client.UpdateSecretProperties
	// HTTP status codes to indicate success: http.StatusOK
	UpdateSecretProperties func(ctx context.Context, name string, version string, parameters azsecrets.UpdateSecretPropertiesParameters, options *azsecrets.UpdateSecretPropertiesOptions) (resp azfake.Responder[azsecrets.UpdateSecretPropertiesResponse], errResp azfake.ErrorResponder)
}

// NewServerTransport creates a new instance of ServerTransport with the provided implementation.
// The returned ServerTransport instance is connected to an instance of azsecrets.Client via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerTransport(srv *Server) *ServerTransport {
	return &ServerTransport{
		srv:                                  srv,
		newListDeletedSecretPropertiesPager:  newTracker[azfake.PagerResponder[azsecrets.ListDeletedSecretPropertiesResponse]](),
		newListSecretPropertiesPager:         newTracker[azfake.PagerResponder[azsecrets.ListSecretPropertiesResponse]](),
		newListSecretPropertiesVersionsPager: newTracker[azfake.PagerResponder[azsecrets.ListSecretPropertiesVersionsResponse]](),
	}
}

// ServerTransport connects instances of azsecrets.Client to instances of Server.
// Don't use this type directly, use NewServerTransport instead.
type ServerTransport struct {
	srv                                  *Server
	newListDeletedSecretPropertiesPager  *tracker[azfake.PagerResponder[azsecrets.ListDeletedSecretPropertiesResponse]]
	newListSecretPropertiesPager         *tracker[azfake.PagerResponder[azsecrets.ListSecretPropertiesResponse]]
	newListSecretPropertiesVersionsPager *tracker[azfake.PagerResponder[azsecrets.ListSecretPropertiesVersionsResponse]]
}

// Do implements the policy.Transporter interface for ServerTransport.
func (s *ServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *ServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if serverTransportInterceptor != nil {
			res.resp, res.err, intercepted = serverTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "Client.BackupSecret":
				res.resp, res.err = s.dispatchBackupSecret(req)
			case "Client.DeleteSecret":
				res.resp, res.err = s.dispatchDeleteSecret(req)
			case "Client.GetDeletedSecret":
				res.resp, res.err = s.dispatchGetDeletedSecret(req)
			case "Client.GetSecret":
				res.resp, res.err = s.dispatchGetSecret(req)
			case "Client.NewListDeletedSecretPropertiesPager":
				res.resp, res.err = s.dispatchNewListDeletedSecretPropertiesPager(req)
			case "Client.NewListSecretPropertiesPager":
				res.resp, res.err = s.dispatchNewListSecretPropertiesPager(req)
			case "Client.NewListSecretPropertiesVersionsPager":
				res.resp, res.err = s.dispatchNewListSecretPropertiesVersionsPager(req)
			case "Client.PurgeDeletedSecret":
				res.resp, res.err = s.dispatchPurgeDeletedSecret(req)
			case "Client.RecoverDeletedSecret":
				res.resp, res.err = s.dispatchRecoverDeletedSecret(req)
			case "Client.RestoreSecret":
				res.resp, res.err = s.dispatchRestoreSecret(req)
			case "Client.SetSecret":
				res.resp, res.err = s.dispatchSetSecret(req)
			case "Client.UpdateSecretProperties":
				res.resp, res.err = s.dispatchUpdateSecretProperties(req)
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

func (s *ServerTransport) dispatchBackupSecret(req *http.Request) (*http.Response, error) {
	if s.srv.BackupSecret == nil {
		return nil, &nonRetriableError{errors.New("fake for method BackupSecret not implemented")}
	}
	const regexStr = `/secrets/(?P<secret_name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backup`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("secret_name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.BackupSecret(req.Context(), nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BackupSecretResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerTransport) dispatchDeleteSecret(req *http.Request) (*http.Response, error) {
	if s.srv.DeleteSecret == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteSecret not implemented")}
	}
	const regexStr = `/secrets/(?P<secret_name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("secret_name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.DeleteSecret(req.Context(), nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DeletedSecret, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerTransport) dispatchGetDeletedSecret(req *http.Request) (*http.Response, error) {
	if s.srv.GetDeletedSecret == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetDeletedSecret not implemented")}
	}
	const regexStr = `/deletedsecrets/(?P<secret_name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("secret_name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.GetDeletedSecret(req.Context(), nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DeletedSecret, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerTransport) dispatchGetSecret(req *http.Request) (*http.Response, error) {
	if s.srv.GetSecret == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetSecret not implemented")}
	}
	const regexStr = `/secrets/(?P<secret_name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/?(?P<secret_version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)?`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("secret_name")])
	if err != nil {
		return nil, err
	}
	versionParam, err := url.PathUnescape(matches[regex.SubexpIndex("secret_version")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.GetSecret(req.Context(), nameParam, versionParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Secret, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerTransport) dispatchNewListDeletedSecretPropertiesPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListDeletedSecretPropertiesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListDeletedSecretPropertiesPager not implemented")}
	}
	newListDeletedSecretPropertiesPager := s.newListDeletedSecretPropertiesPager.get(req)
	if newListDeletedSecretPropertiesPager == nil {
		qp := req.URL.Query()
		maxResultsUnescaped, err := url.QueryUnescape(qp.Get("maxresults"))
		if err != nil {
			return nil, err
		}
		maxResultsParam, err := parseOptional(maxResultsUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *azsecrets.ListDeletedSecretPropertiesOptions
		if maxResultsParam != nil {
			options = &azsecrets.ListDeletedSecretPropertiesOptions{
				MaxResults: maxResultsParam,
			}
		}
		resp := s.srv.NewListDeletedSecretPropertiesPager(options)
		newListDeletedSecretPropertiesPager = &resp
		s.newListDeletedSecretPropertiesPager.add(req, newListDeletedSecretPropertiesPager)
		server.PagerResponderInjectNextLinks(newListDeletedSecretPropertiesPager, req, func(page *azsecrets.ListDeletedSecretPropertiesResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListDeletedSecretPropertiesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListDeletedSecretPropertiesPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListDeletedSecretPropertiesPager) {
		s.newListDeletedSecretPropertiesPager.remove(req)
	}
	return resp, nil
}

func (s *ServerTransport) dispatchNewListSecretPropertiesPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListSecretPropertiesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListSecretPropertiesPager not implemented")}
	}
	newListSecretPropertiesPager := s.newListSecretPropertiesPager.get(req)
	if newListSecretPropertiesPager == nil {
		qp := req.URL.Query()
		maxResultsUnescaped, err := url.QueryUnescape(qp.Get("maxresults"))
		if err != nil {
			return nil, err
		}
		maxResultsParam, err := parseOptional(maxResultsUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *azsecrets.ListSecretPropertiesOptions
		if maxResultsParam != nil {
			options = &azsecrets.ListSecretPropertiesOptions{
				MaxResults: maxResultsParam,
			}
		}
		resp := s.srv.NewListSecretPropertiesPager(options)
		newListSecretPropertiesPager = &resp
		s.newListSecretPropertiesPager.add(req, newListSecretPropertiesPager)
		server.PagerResponderInjectNextLinks(newListSecretPropertiesPager, req, func(page *azsecrets.ListSecretPropertiesResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListSecretPropertiesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListSecretPropertiesPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListSecretPropertiesPager) {
		s.newListSecretPropertiesPager.remove(req)
	}
	return resp, nil
}

func (s *ServerTransport) dispatchNewListSecretPropertiesVersionsPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListSecretPropertiesVersionsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListSecretPropertiesVersionsPager not implemented")}
	}
	newListSecretPropertiesVersionsPager := s.newListSecretPropertiesVersionsPager.get(req)
	if newListSecretPropertiesVersionsPager == nil {
		const regexStr = `/secrets/(?P<secret_name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("secret_name")])
		if err != nil {
			return nil, err
		}
		maxResultsUnescaped, err := url.QueryUnescape(qp.Get("maxresults"))
		if err != nil {
			return nil, err
		}
		maxResultsParam, err := parseOptional(maxResultsUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *azsecrets.ListSecretPropertiesVersionsOptions
		if maxResultsParam != nil {
			options = &azsecrets.ListSecretPropertiesVersionsOptions{
				MaxResults: maxResultsParam,
			}
		}
		resp := s.srv.NewListSecretPropertiesVersionsPager(nameParam, options)
		newListSecretPropertiesVersionsPager = &resp
		s.newListSecretPropertiesVersionsPager.add(req, newListSecretPropertiesVersionsPager)
		server.PagerResponderInjectNextLinks(newListSecretPropertiesVersionsPager, req, func(page *azsecrets.ListSecretPropertiesVersionsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListSecretPropertiesVersionsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListSecretPropertiesVersionsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListSecretPropertiesVersionsPager) {
		s.newListSecretPropertiesVersionsPager.remove(req)
	}
	return resp, nil
}

func (s *ServerTransport) dispatchPurgeDeletedSecret(req *http.Request) (*http.Response, error) {
	if s.srv.PurgeDeletedSecret == nil {
		return nil, &nonRetriableError{errors.New("fake for method PurgeDeletedSecret not implemented")}
	}
	const regexStr = `/deletedsecrets/(?P<secret_name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("secret_name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.PurgeDeletedSecret(req.Context(), nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerTransport) dispatchRecoverDeletedSecret(req *http.Request) (*http.Response, error) {
	if s.srv.RecoverDeletedSecret == nil {
		return nil, &nonRetriableError{errors.New("fake for method RecoverDeletedSecret not implemented")}
	}
	const regexStr = `/deletedsecrets/(?P<secret_name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/recover`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("secret_name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.RecoverDeletedSecret(req.Context(), nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Secret, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerTransport) dispatchRestoreSecret(req *http.Request) (*http.Response, error) {
	if s.srv.RestoreSecret == nil {
		return nil, &nonRetriableError{errors.New("fake for method RestoreSecret not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[azsecrets.RestoreSecretParameters](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.RestoreSecret(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Secret, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerTransport) dispatchSetSecret(req *http.Request) (*http.Response, error) {
	if s.srv.SetSecret == nil {
		return nil, &nonRetriableError{errors.New("fake for method SetSecret not implemented")}
	}
	const regexStr = `/secrets/(?P<secret_name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[azsecrets.SetSecretParameters](req)
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("secret_name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.SetSecret(req.Context(), nameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Secret, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerTransport) dispatchUpdateSecretProperties(req *http.Request) (*http.Response, error) {
	if s.srv.UpdateSecretProperties == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateSecretProperties not implemented")}
	}
	const regexStr = `/secrets/(?P<secret_name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/?(?P<secret_version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)?`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[azsecrets.UpdateSecretPropertiesParameters](req)
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("secret_name")])
	if err != nil {
		return nil, err
	}
	versionParam, err := url.PathUnescape(matches[regex.SubexpIndex("secret_version")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.UpdateSecretProperties(req.Context(), nameParam, versionParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Secret, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ServerTransport
var serverTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
