package peering

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// PrefixesClient is the peering Client
type PrefixesClient struct {
	BaseClient
}

// NewPrefixesClient creates an instance of the PrefixesClient client.
func NewPrefixesClient(subscriptionID string) PrefixesClient {
	return NewPrefixesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPrefixesClientWithBaseURI creates an instance of the PrefixesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPrefixesClientWithBaseURI(baseURI string, subscriptionID string) PrefixesClient {
	return PrefixesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a new prefix with the specified name under the given subscription, resource group and peering
// service.
// Parameters:
// resourceGroupName - the name of the resource group.
// peeringServiceName - the name of the peering service.
// prefixName - the name of the prefix.
// peeringServicePrefix - the properties needed to create a prefix.
func (client PrefixesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, peeringServiceName string, prefixName string, peeringServicePrefix ServicePrefix) (result ServicePrefix, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrefixesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, peeringServiceName, prefixName, peeringServicePrefix)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.PrefixesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "peering.PrefixesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.PrefixesClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client PrefixesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, peeringServiceName string, prefixName string, peeringServicePrefix ServicePrefix) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"peeringServiceName": autorest.Encode("path", peeringServiceName),
		"prefixName":         autorest.Encode("path", prefixName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}/prefixes/{prefixName}", pathParameters),
		autorest.WithJSON(peeringServicePrefix),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client PrefixesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client PrefixesClient) CreateOrUpdateResponder(resp *http.Response) (result ServicePrefix, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an existing prefix with the specified name under the given subscription, resource group and peering
// service.
// Parameters:
// resourceGroupName - the name of the resource group.
// peeringServiceName - the name of the peering service.
// prefixName - the name of the prefix.
func (client PrefixesClient) Delete(ctx context.Context, resourceGroupName string, peeringServiceName string, prefixName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrefixesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, peeringServiceName, prefixName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.PrefixesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "peering.PrefixesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.PrefixesClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PrefixesClient) DeletePreparer(ctx context.Context, resourceGroupName string, peeringServiceName string, prefixName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"peeringServiceName": autorest.Encode("path", peeringServiceName),
		"prefixName":         autorest.Encode("path", prefixName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}/prefixes/{prefixName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PrefixesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PrefixesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets an existing prefix with the specified name under the given subscription, resource group and peering
// service.
// Parameters:
// resourceGroupName - the name of the resource group.
// peeringServiceName - the name of the peering service.
// prefixName - the name of the prefix.
// expand - the properties to be expanded.
func (client PrefixesClient) Get(ctx context.Context, resourceGroupName string, peeringServiceName string, prefixName string, expand string) (result ServicePrefix, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrefixesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, peeringServiceName, prefixName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.PrefixesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "peering.PrefixesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.PrefixesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client PrefixesClient) GetPreparer(ctx context.Context, resourceGroupName string, peeringServiceName string, prefixName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"peeringServiceName": autorest.Encode("path", peeringServiceName),
		"prefixName":         autorest.Encode("path", prefixName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}/prefixes/{prefixName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PrefixesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PrefixesClient) GetResponder(resp *http.Response) (result ServicePrefix, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByPeeringService lists all prefixes under the given subscription, resource group and peering service.
// Parameters:
// resourceGroupName - the name of the resource group.
// peeringServiceName - the name of the peering service.
// expand - the properties to be expanded.
func (client PrefixesClient) ListByPeeringService(ctx context.Context, resourceGroupName string, peeringServiceName string, expand string) (result ServicePrefixListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrefixesClient.ListByPeeringService")
		defer func() {
			sc := -1
			if result.splr.Response.Response != nil {
				sc = result.splr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByPeeringServiceNextResults
	req, err := client.ListByPeeringServicePreparer(ctx, resourceGroupName, peeringServiceName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.PrefixesClient", "ListByPeeringService", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByPeeringServiceSender(req)
	if err != nil {
		result.splr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "peering.PrefixesClient", "ListByPeeringService", resp, "Failure sending request")
		return
	}

	result.splr, err = client.ListByPeeringServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.PrefixesClient", "ListByPeeringService", resp, "Failure responding to request")
		return
	}
	if result.splr.hasNextLink() && result.splr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByPeeringServicePreparer prepares the ListByPeeringService request.
func (client PrefixesClient) ListByPeeringServicePreparer(ctx context.Context, resourceGroupName string, peeringServiceName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"peeringServiceName": autorest.Encode("path", peeringServiceName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}/prefixes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByPeeringServiceSender sends the ListByPeeringService request. The method will close the
// http.Response Body if it receives an error.
func (client PrefixesClient) ListByPeeringServiceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByPeeringServiceResponder handles the response to the ListByPeeringService request. The method always
// closes the http.Response Body.
func (client PrefixesClient) ListByPeeringServiceResponder(resp *http.Response) (result ServicePrefixListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByPeeringServiceNextResults retrieves the next set of results, if any.
func (client PrefixesClient) listByPeeringServiceNextResults(ctx context.Context, lastResults ServicePrefixListResult) (result ServicePrefixListResult, err error) {
	req, err := lastResults.servicePrefixListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "peering.PrefixesClient", "listByPeeringServiceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByPeeringServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "peering.PrefixesClient", "listByPeeringServiceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByPeeringServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.PrefixesClient", "listByPeeringServiceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByPeeringServiceComplete enumerates all values, automatically crossing page boundaries as required.
func (client PrefixesClient) ListByPeeringServiceComplete(ctx context.Context, resourceGroupName string, peeringServiceName string, expand string) (result ServicePrefixListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrefixesClient.ListByPeeringService")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByPeeringService(ctx, resourceGroupName, peeringServiceName, expand)
	return
}
