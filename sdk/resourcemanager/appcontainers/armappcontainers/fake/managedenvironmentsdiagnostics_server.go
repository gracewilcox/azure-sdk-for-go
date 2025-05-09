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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
	"net/http"
	"net/url"
	"regexp"
)

// ManagedEnvironmentsDiagnosticsServer is a fake server for instances of the armappcontainers.ManagedEnvironmentsDiagnosticsClient type.
type ManagedEnvironmentsDiagnosticsServer struct {
	// GetRoot is the fake for method ManagedEnvironmentsDiagnosticsClient.GetRoot
	// HTTP status codes to indicate success: http.StatusOK
	GetRoot func(ctx context.Context, resourceGroupName string, environmentName string, options *armappcontainers.ManagedEnvironmentsDiagnosticsClientGetRootOptions) (resp azfake.Responder[armappcontainers.ManagedEnvironmentsDiagnosticsClientGetRootResponse], errResp azfake.ErrorResponder)
}

// NewManagedEnvironmentsDiagnosticsServerTransport creates a new instance of ManagedEnvironmentsDiagnosticsServerTransport with the provided implementation.
// The returned ManagedEnvironmentsDiagnosticsServerTransport instance is connected to an instance of armappcontainers.ManagedEnvironmentsDiagnosticsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagedEnvironmentsDiagnosticsServerTransport(srv *ManagedEnvironmentsDiagnosticsServer) *ManagedEnvironmentsDiagnosticsServerTransport {
	return &ManagedEnvironmentsDiagnosticsServerTransport{srv: srv}
}

// ManagedEnvironmentsDiagnosticsServerTransport connects instances of armappcontainers.ManagedEnvironmentsDiagnosticsClient to instances of ManagedEnvironmentsDiagnosticsServer.
// Don't use this type directly, use NewManagedEnvironmentsDiagnosticsServerTransport instead.
type ManagedEnvironmentsDiagnosticsServerTransport struct {
	srv *ManagedEnvironmentsDiagnosticsServer
}

// Do implements the policy.Transporter interface for ManagedEnvironmentsDiagnosticsServerTransport.
func (m *ManagedEnvironmentsDiagnosticsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return m.dispatchToMethodFake(req, method)
}

func (m *ManagedEnvironmentsDiagnosticsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if managedEnvironmentsDiagnosticsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = managedEnvironmentsDiagnosticsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ManagedEnvironmentsDiagnosticsClient.GetRoot":
				res.resp, res.err = m.dispatchGetRoot(req)
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

func (m *ManagedEnvironmentsDiagnosticsServerTransport) dispatchGetRoot(req *http.Request) (*http.Response, error) {
	if m.srv.GetRoot == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetRoot not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments/(?P<environmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/detectorProperties/rootApi/`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	environmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.GetRoot(req.Context(), resourceGroupNameParam, environmentNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagedEnvironment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ManagedEnvironmentsDiagnosticsServerTransport
var managedEnvironmentsDiagnosticsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
