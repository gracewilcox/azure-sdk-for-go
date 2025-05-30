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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ClusterRecoveryPointServer is a fake server for instances of the armrecoveryservicessiterecovery.ClusterRecoveryPointClient type.
type ClusterRecoveryPointServer struct {
	// Get is the fake for method ClusterRecoveryPointClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, protectionContainerName string, replicationProtectionClusterName string, recoveryPointName string, options *armrecoveryservicessiterecovery.ClusterRecoveryPointClientGetOptions) (resp azfake.Responder[armrecoveryservicessiterecovery.ClusterRecoveryPointClientGetResponse], errResp azfake.ErrorResponder)
}

// NewClusterRecoveryPointServerTransport creates a new instance of ClusterRecoveryPointServerTransport with the provided implementation.
// The returned ClusterRecoveryPointServerTransport instance is connected to an instance of armrecoveryservicessiterecovery.ClusterRecoveryPointClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewClusterRecoveryPointServerTransport(srv *ClusterRecoveryPointServer) *ClusterRecoveryPointServerTransport {
	return &ClusterRecoveryPointServerTransport{srv: srv}
}

// ClusterRecoveryPointServerTransport connects instances of armrecoveryservicessiterecovery.ClusterRecoveryPointClient to instances of ClusterRecoveryPointServer.
// Don't use this type directly, use NewClusterRecoveryPointServerTransport instead.
type ClusterRecoveryPointServerTransport struct {
	srv *ClusterRecoveryPointServer
}

// Do implements the policy.Transporter interface for ClusterRecoveryPointServerTransport.
func (c *ClusterRecoveryPointServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *ClusterRecoveryPointServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if clusterRecoveryPointServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = clusterRecoveryPointServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ClusterRecoveryPointClient.Get":
				res.resp, res.err = c.dispatchGet(req)
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

func (c *ClusterRecoveryPointServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationProtectionContainers/(?P<protectionContainerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationProtectionClusters/(?P<replicationProtectionClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/recoveryPoints/(?P<recoveryPointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 7 {
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
	protectionContainerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("protectionContainerName")])
	if err != nil {
		return nil, err
	}
	replicationProtectionClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("replicationProtectionClusterName")])
	if err != nil {
		return nil, err
	}
	recoveryPointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("recoveryPointName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, resourceNameParam, fabricNameParam, protectionContainerNameParam, replicationProtectionClusterNameParam, recoveryPointNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ClusterRecoveryPoint, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ClusterRecoveryPointServerTransport
var clusterRecoveryPointServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
