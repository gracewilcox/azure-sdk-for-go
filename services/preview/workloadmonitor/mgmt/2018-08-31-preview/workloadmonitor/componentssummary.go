package workloadmonitor

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ComponentsSummaryClient is the aPIs for workload monitoring
type ComponentsSummaryClient struct {
	BaseClient
}

// NewComponentsSummaryClient creates an instance of the ComponentsSummaryClient client.
func NewComponentsSummaryClient(subscriptionID string) ComponentsSummaryClient {
	return NewComponentsSummaryClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewComponentsSummaryClientWithBaseURI creates an instance of the ComponentsSummaryClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewComponentsSummaryClientWithBaseURI(baseURI string, subscriptionID string) ComponentsSummaryClient {
	return ComponentsSummaryClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List sends the list request.
// Parameters:
// selectParameter - properties to be returned in the response.
// filter - filter to be applied on the operation.
// apply - apply aggregation.
// orderby - sort the result on one or more properties.
// expand - include properties inline in the response.
// top - limit the result to the specified number of rows.
// skiptoken - the page-continuation token to use with a paged version of this API.
func (client ComponentsSummaryClient) List(ctx context.Context, selectParameter string, filter string, apply string, orderby string, expand string, top string, skiptoken string) (result ComponentsCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ComponentsSummaryClient.List")
		defer func() {
			sc := -1
			if result.cc.Response.Response != nil {
				sc = result.cc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("workloadmonitor.ComponentsSummaryClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, selectParameter, filter, apply, orderby, expand, top, skiptoken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadmonitor.ComponentsSummaryClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.cc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "workloadmonitor.ComponentsSummaryClient", "List", resp, "Failure sending request")
		return
	}

	result.cc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadmonitor.ComponentsSummaryClient", "List", resp, "Failure responding to request")
		return
	}
	if result.cc.hasNextLink() && result.cc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ComponentsSummaryClient) ListPreparer(ctx context.Context, selectParameter string, filter string, apply string, orderby string, expand string, top string, skiptoken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-08-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(apply) > 0 {
		queryParameters["$apply"] = autorest.Encode("query", apply)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(top) > 0 {
		queryParameters["$top"] = autorest.Encode("query", top)
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.WorkloadMonitor/componentsSummary", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ComponentsSummaryClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ComponentsSummaryClient) ListResponder(resp *http.Response) (result ComponentsCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ComponentsSummaryClient) listNextResults(ctx context.Context, lastResults ComponentsCollection) (result ComponentsCollection, err error) {
	req, err := lastResults.componentsCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "workloadmonitor.ComponentsSummaryClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "workloadmonitor.ComponentsSummaryClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadmonitor.ComponentsSummaryClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ComponentsSummaryClient) ListComplete(ctx context.Context, selectParameter string, filter string, apply string, orderby string, expand string, top string, skiptoken string) (result ComponentsCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ComponentsSummaryClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, selectParameter, filter, apply, orderby, expand, top, skiptoken)
	return
}
