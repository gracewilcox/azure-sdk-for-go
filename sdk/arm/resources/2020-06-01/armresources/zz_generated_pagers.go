// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresources

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// DeploymentListResultPager provides iteration over DeploymentListResult pages.
type DeploymentListResultPager interface {
	// NextPage returns true if the pager advanced to the next page.
	// Returns false if there are no more pages or an error occurred.
	NextPage(context.Context) bool

	// Page returns the current DeploymentListResultResponse.
	PageResponse() DeploymentListResultResponse

	// Err returns the last error encountered while paging.
	Err() error
}

type deploymentListResultCreateRequest func(context.Context) (*azcore.Request, error)

type deploymentListResultHandleError func(*azcore.Response) error

type deploymentListResultHandleResponse func(*azcore.Response) (DeploymentListResultResponse, error)

type deploymentListResultAdvancePage func(context.Context, DeploymentListResultResponse) (*azcore.Request, error)

type deploymentListResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester deploymentListResultCreateRequest
	// callback for handling response errors
	errorer deploymentListResultHandleError
	// callback for handling the HTTP response
	responder deploymentListResultHandleResponse
	// callback for advancing to the next page
	advancer deploymentListResultAdvancePage
	// contains the current response
	current DeploymentListResultResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *deploymentListResultPager) Err() error {
	return p.err
}

func (p *deploymentListResultPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeploymentListResult.NextLink == nil || len(*p.current.DeploymentListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
		p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *deploymentListResultPager) PageResponse() DeploymentListResultResponse {
	return p.current
}

// DeploymentOperationsListResultPager provides iteration over DeploymentOperationsListResult pages.
type DeploymentOperationsListResultPager interface {
	// NextPage returns true if the pager advanced to the next page.
	// Returns false if there are no more pages or an error occurred.
	NextPage(context.Context) bool

	// Page returns the current DeploymentOperationsListResultResponse.
	PageResponse() DeploymentOperationsListResultResponse

	// Err returns the last error encountered while paging.
	Err() error
}

type deploymentOperationsListResultCreateRequest func(context.Context) (*azcore.Request, error)

type deploymentOperationsListResultHandleError func(*azcore.Response) error

type deploymentOperationsListResultHandleResponse func(*azcore.Response) (DeploymentOperationsListResultResponse, error)

type deploymentOperationsListResultAdvancePage func(context.Context, DeploymentOperationsListResultResponse) (*azcore.Request, error)

type deploymentOperationsListResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester deploymentOperationsListResultCreateRequest
	// callback for handling response errors
	errorer deploymentOperationsListResultHandleError
	// callback for handling the HTTP response
	responder deploymentOperationsListResultHandleResponse
	// callback for advancing to the next page
	advancer deploymentOperationsListResultAdvancePage
	// contains the current response
	current DeploymentOperationsListResultResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *deploymentOperationsListResultPager) Err() error {
	return p.err
}

func (p *deploymentOperationsListResultPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeploymentOperationsListResult.NextLink == nil || len(*p.current.DeploymentOperationsListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
		p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *deploymentOperationsListResultPager) PageResponse() DeploymentOperationsListResultResponse {
	return p.current
}

// OperationListResultPager provides iteration over OperationListResult pages.
type OperationListResultPager interface {
	// NextPage returns true if the pager advanced to the next page.
	// Returns false if there are no more pages or an error occurred.
	NextPage(context.Context) bool

	// Page returns the current OperationListResultResponse.
	PageResponse() OperationListResultResponse

	// Err returns the last error encountered while paging.
	Err() error
}

type operationListResultCreateRequest func(context.Context) (*azcore.Request, error)

type operationListResultHandleError func(*azcore.Response) error

type operationListResultHandleResponse func(*azcore.Response) (OperationListResultResponse, error)

type operationListResultAdvancePage func(context.Context, OperationListResultResponse) (*azcore.Request, error)

type operationListResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester operationListResultCreateRequest
	// callback for handling response errors
	errorer operationListResultHandleError
	// callback for handling the HTTP response
	responder operationListResultHandleResponse
	// callback for advancing to the next page
	advancer operationListResultAdvancePage
	// contains the current response
	current OperationListResultResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *operationListResultPager) Err() error {
	return p.err
}

func (p *operationListResultPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OperationListResult.NextLink == nil || len(*p.current.OperationListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
		p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *operationListResultPager) PageResponse() OperationListResultResponse {
	return p.current
}

// ProviderListResultPager provides iteration over ProviderListResult pages.
type ProviderListResultPager interface {
	// NextPage returns true if the pager advanced to the next page.
	// Returns false if there are no more pages or an error occurred.
	NextPage(context.Context) bool

	// Page returns the current ProviderListResultResponse.
	PageResponse() ProviderListResultResponse

	// Err returns the last error encountered while paging.
	Err() error
}

type providerListResultCreateRequest func(context.Context) (*azcore.Request, error)

type providerListResultHandleError func(*azcore.Response) error

type providerListResultHandleResponse func(*azcore.Response) (ProviderListResultResponse, error)

type providerListResultAdvancePage func(context.Context, ProviderListResultResponse) (*azcore.Request, error)

type providerListResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester providerListResultCreateRequest
	// callback for handling response errors
	errorer providerListResultHandleError
	// callback for handling the HTTP response
	responder providerListResultHandleResponse
	// callback for advancing to the next page
	advancer providerListResultAdvancePage
	// contains the current response
	current ProviderListResultResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *providerListResultPager) Err() error {
	return p.err
}

func (p *providerListResultPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProviderListResult.NextLink == nil || len(*p.current.ProviderListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
		p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *providerListResultPager) PageResponse() ProviderListResultResponse {
	return p.current
}

// ResourceGroupListResultPager provides iteration over ResourceGroupListResult pages.
type ResourceGroupListResultPager interface {
	// NextPage returns true if the pager advanced to the next page.
	// Returns false if there are no more pages or an error occurred.
	NextPage(context.Context) bool

	// Page returns the current ResourceGroupListResultResponse.
	PageResponse() ResourceGroupListResultResponse

	// Err returns the last error encountered while paging.
	Err() error
}

type resourceGroupListResultCreateRequest func(context.Context) (*azcore.Request, error)

type resourceGroupListResultHandleError func(*azcore.Response) error

type resourceGroupListResultHandleResponse func(*azcore.Response) (ResourceGroupListResultResponse, error)

type resourceGroupListResultAdvancePage func(context.Context, ResourceGroupListResultResponse) (*azcore.Request, error)

type resourceGroupListResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester resourceGroupListResultCreateRequest
	// callback for handling response errors
	errorer resourceGroupListResultHandleError
	// callback for handling the HTTP response
	responder resourceGroupListResultHandleResponse
	// callback for advancing to the next page
	advancer resourceGroupListResultAdvancePage
	// contains the current response
	current ResourceGroupListResultResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *resourceGroupListResultPager) Err() error {
	return p.err
}

func (p *resourceGroupListResultPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ResourceGroupListResult.NextLink == nil || len(*p.current.ResourceGroupListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
		p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *resourceGroupListResultPager) PageResponse() ResourceGroupListResultResponse {
	return p.current
}

// ResourceListResultPager provides iteration over ResourceListResult pages.
type ResourceListResultPager interface {
	// NextPage returns true if the pager advanced to the next page.
	// Returns false if there are no more pages or an error occurred.
	NextPage(context.Context) bool

	// Page returns the current ResourceListResultResponse.
	PageResponse() ResourceListResultResponse

	// Err returns the last error encountered while paging.
	Err() error
}

type resourceListResultCreateRequest func(context.Context) (*azcore.Request, error)

type resourceListResultHandleError func(*azcore.Response) error

type resourceListResultHandleResponse func(*azcore.Response) (ResourceListResultResponse, error)

type resourceListResultAdvancePage func(context.Context, ResourceListResultResponse) (*azcore.Request, error)

type resourceListResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester resourceListResultCreateRequest
	// callback for handling response errors
	errorer resourceListResultHandleError
	// callback for handling the HTTP response
	responder resourceListResultHandleResponse
	// callback for advancing to the next page
	advancer resourceListResultAdvancePage
	// contains the current response
	current ResourceListResultResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *resourceListResultPager) Err() error {
	return p.err
}

func (p *resourceListResultPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ResourceListResult.NextLink == nil || len(*p.current.ResourceListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
		p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *resourceListResultPager) PageResponse() ResourceListResultResponse {
	return p.current
}

// TagsListResultPager provides iteration over TagsListResult pages.
type TagsListResultPager interface {
	// NextPage returns true if the pager advanced to the next page.
	// Returns false if there are no more pages or an error occurred.
	NextPage(context.Context) bool

	// Page returns the current TagsListResultResponse.
	PageResponse() TagsListResultResponse

	// Err returns the last error encountered while paging.
	Err() error
}

type tagsListResultCreateRequest func(context.Context) (*azcore.Request, error)

type tagsListResultHandleError func(*azcore.Response) error

type tagsListResultHandleResponse func(*azcore.Response) (TagsListResultResponse, error)

type tagsListResultAdvancePage func(context.Context, TagsListResultResponse) (*azcore.Request, error)

type tagsListResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// creates the initial request (non-LRO case)
	requester tagsListResultCreateRequest
	// callback for handling response errors
	errorer tagsListResultHandleError
	// callback for handling the HTTP response
	responder tagsListResultHandleResponse
	// callback for advancing to the next page
	advancer tagsListResultAdvancePage
	// contains the current response
	current TagsListResultResponse
	// status codes for successful retrieval
	statusCodes []int
	// any error encountered
	err error
}

func (p *tagsListResultPager) Err() error {
	return p.err
}

func (p *tagsListResultPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.TagsListResult.NextLink == nil || len(*p.current.TagsListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.pipeline.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(p.statusCodes...) {
		p.err = p.errorer(resp)
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *tagsListResultPager) PageResponse() TagsListResultResponse {
	return p.current
}
