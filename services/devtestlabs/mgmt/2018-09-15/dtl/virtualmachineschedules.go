package dtl

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

// VirtualMachineSchedulesClient is the the DevTest Labs Client.
type VirtualMachineSchedulesClient struct {
	BaseClient
}

// NewVirtualMachineSchedulesClient creates an instance of the VirtualMachineSchedulesClient client.
func NewVirtualMachineSchedulesClient(subscriptionID string) VirtualMachineSchedulesClient {
	return NewVirtualMachineSchedulesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualMachineSchedulesClientWithBaseURI creates an instance of the VirtualMachineSchedulesClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewVirtualMachineSchedulesClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineSchedulesClient {
	return VirtualMachineSchedulesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or replace an existing schedule.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// virtualMachineName - the name of the virtual machine.
// name - the name of the schedule.
// schedule - a schedule.
func (client VirtualMachineSchedulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, name string, schedule Schedule) (result Schedule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineSchedulesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: schedule,
			Constraints: []validation.Constraint{{Target: "schedule.ScheduleProperties", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("dtl.VirtualMachineSchedulesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, labName, virtualMachineName, name, schedule)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.VirtualMachineSchedulesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.VirtualMachineSchedulesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.VirtualMachineSchedulesClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualMachineSchedulesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, name string, schedule Schedule) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":            autorest.Encode("path", labName),
		"name":               autorest.Encode("path", name),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualMachineName": autorest.Encode("path", virtualMachineName),
	}

	const APIVersion = "2018-09-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{virtualMachineName}/schedules/{name}", pathParameters),
		autorest.WithJSON(schedule),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineSchedulesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VirtualMachineSchedulesClient) CreateOrUpdateResponder(resp *http.Response) (result Schedule, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete schedule.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// virtualMachineName - the name of the virtual machine.
// name - the name of the schedule.
func (client VirtualMachineSchedulesClient) Delete(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, name string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineSchedulesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, labName, virtualMachineName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.VirtualMachineSchedulesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "dtl.VirtualMachineSchedulesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.VirtualMachineSchedulesClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VirtualMachineSchedulesClient) DeletePreparer(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":            autorest.Encode("path", labName),
		"name":               autorest.Encode("path", name),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualMachineName": autorest.Encode("path", virtualMachineName),
	}

	const APIVersion = "2018-09-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{virtualMachineName}/schedules/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineSchedulesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VirtualMachineSchedulesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Execute execute a schedule. This operation can take a while to complete.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// virtualMachineName - the name of the virtual machine.
// name - the name of the schedule.
func (client VirtualMachineSchedulesClient) Execute(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, name string) (result VirtualMachineSchedulesExecuteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineSchedulesClient.Execute")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ExecutePreparer(ctx, resourceGroupName, labName, virtualMachineName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.VirtualMachineSchedulesClient", "Execute", nil, "Failure preparing request")
		return
	}

	result, err = client.ExecuteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.VirtualMachineSchedulesClient", "Execute", nil, "Failure sending request")
		return
	}

	return
}

// ExecutePreparer prepares the Execute request.
func (client VirtualMachineSchedulesClient) ExecutePreparer(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":            autorest.Encode("path", labName),
		"name":               autorest.Encode("path", name),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualMachineName": autorest.Encode("path", virtualMachineName),
	}

	const APIVersion = "2018-09-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{virtualMachineName}/schedules/{name}/execute", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ExecuteSender sends the Execute request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineSchedulesClient) ExecuteSender(req *http.Request) (future VirtualMachineSchedulesExecuteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// ExecuteResponder handles the response to the Execute request. The method always
// closes the http.Response Body.
func (client VirtualMachineSchedulesClient) ExecuteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get schedule.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// virtualMachineName - the name of the virtual machine.
// name - the name of the schedule.
// expand - specify the $expand query. Example: 'properties($select=status)'
func (client VirtualMachineSchedulesClient) Get(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, name string, expand string) (result Schedule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineSchedulesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, labName, virtualMachineName, name, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.VirtualMachineSchedulesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.VirtualMachineSchedulesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.VirtualMachineSchedulesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualMachineSchedulesClient) GetPreparer(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, name string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":            autorest.Encode("path", labName),
		"name":               autorest.Encode("path", name),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualMachineName": autorest.Encode("path", virtualMachineName),
	}

	const APIVersion = "2018-09-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{virtualMachineName}/schedules/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineSchedulesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualMachineSchedulesClient) GetResponder(resp *http.Response) (result Schedule, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list schedules in a given virtual machine.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// virtualMachineName - the name of the virtual machine.
// expand - specify the $expand query. Example: 'properties($select=status)'
// filter - the filter to apply to the operation. Example: '$filter=contains(name,'myName')
// top - the maximum number of resources to return from the operation. Example: '$top=10'
// orderby - the ordering expression for the results, using OData notation. Example: '$orderby=name desc'
func (client VirtualMachineSchedulesClient) List(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, expand string, filter string, top *int32, orderby string) (result ScheduleListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineSchedulesClient.List")
		defer func() {
			sc := -1
			if result.sl.Response.Response != nil {
				sc = result.sl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, labName, virtualMachineName, expand, filter, top, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.VirtualMachineSchedulesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.sl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.VirtualMachineSchedulesClient", "List", resp, "Failure sending request")
		return
	}

	result.sl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.VirtualMachineSchedulesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.sl.hasNextLink() && result.sl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualMachineSchedulesClient) ListPreparer(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, expand string, filter string, top *int32, orderby string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":            autorest.Encode("path", labName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualMachineName": autorest.Encode("path", virtualMachineName),
	}

	const APIVersion = "2018-09-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{virtualMachineName}/schedules", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineSchedulesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualMachineSchedulesClient) ListResponder(resp *http.Response) (result ScheduleList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client VirtualMachineSchedulesClient) listNextResults(ctx context.Context, lastResults ScheduleList) (result ScheduleList, err error) {
	req, err := lastResults.scheduleListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dtl.VirtualMachineSchedulesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dtl.VirtualMachineSchedulesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.VirtualMachineSchedulesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client VirtualMachineSchedulesClient) ListComplete(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, expand string, filter string, top *int32, orderby string) (result ScheduleListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineSchedulesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, labName, virtualMachineName, expand, filter, top, orderby)
	return
}

// Update allows modifying tags of schedules. All other properties will be ignored.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// virtualMachineName - the name of the virtual machine.
// name - the name of the schedule.
// schedule - a schedule.
func (client VirtualMachineSchedulesClient) Update(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, name string, schedule ScheduleFragment) (result Schedule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineSchedulesClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, labName, virtualMachineName, name, schedule)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.VirtualMachineSchedulesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.VirtualMachineSchedulesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.VirtualMachineSchedulesClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client VirtualMachineSchedulesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, name string, schedule ScheduleFragment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":            autorest.Encode("path", labName),
		"name":               autorest.Encode("path", name),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualMachineName": autorest.Encode("path", virtualMachineName),
	}

	const APIVersion = "2018-09-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{virtualMachineName}/schedules/{name}", pathParameters),
		autorest.WithJSON(schedule),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineSchedulesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client VirtualMachineSchedulesClient) UpdateResponder(resp *http.Response) (result Schedule, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
