package operationalinsights

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

// LinkedServicesClient is the operational Insights Client
type LinkedServicesClient struct {
	BaseClient
}

// NewLinkedServicesClient creates an instance of the LinkedServicesClient client.
func NewLinkedServicesClient(subscriptionID string) LinkedServicesClient {
	return NewLinkedServicesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLinkedServicesClientWithBaseURI creates an instance of the LinkedServicesClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewLinkedServicesClientWithBaseURI(baseURI string, subscriptionID string) LinkedServicesClient {
	return LinkedServicesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update a linked service.
// Parameters:
// resourceGroupName - the name of the resource group to get. The name is case insensitive.
// workspaceName - name of the Log Analytics Workspace that will contain the linkedServices resource
// linkedServiceName - name of the linkedServices resource
// parameters - the parameters required to create or update a linked service.
func (client LinkedServicesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, linkedServiceName string, parameters LinkedService) (result LinkedService, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LinkedServicesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.LinkedServiceProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.LinkedServiceProperties.ResourceID", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("operationalinsights.LinkedServicesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, workspaceName, linkedServiceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.LinkedServicesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "operationalinsights.LinkedServicesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.LinkedServicesClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client LinkedServicesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, linkedServiceName string, parameters LinkedService) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"linkedServiceName": autorest.Encode("path", linkedServiceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/linkedServices/{linkedServiceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client LinkedServicesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client LinkedServicesClient) CreateOrUpdateResponder(resp *http.Response) (result LinkedService, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a linked service instance.
// Parameters:
// resourceGroupName - the name of the resource group to get. The name is case insensitive.
// workspaceName - name of the Log Analytics Workspace that contains the linkedServices resource
// linkedServiceName - name of the linked service.
func (client LinkedServicesClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, linkedServiceName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LinkedServicesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("operationalinsights.LinkedServicesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, workspaceName, linkedServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.LinkedServicesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "operationalinsights.LinkedServicesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.LinkedServicesClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client LinkedServicesClient) DeletePreparer(ctx context.Context, resourceGroupName string, workspaceName string, linkedServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"linkedServiceName": autorest.Encode("path", linkedServiceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/linkedServices/{linkedServiceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client LinkedServicesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client LinkedServicesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a linked service instance.
// Parameters:
// resourceGroupName - the name of the resource group to get. The name is case insensitive.
// workspaceName - name of the Log Analytics Workspace that contains the linkedServices resource
// linkedServiceName - name of the linked service.
func (client LinkedServicesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, linkedServiceName string) (result LinkedService, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LinkedServicesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("operationalinsights.LinkedServicesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, linkedServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.LinkedServicesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "operationalinsights.LinkedServicesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.LinkedServicesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client LinkedServicesClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, linkedServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"linkedServiceName": autorest.Encode("path", linkedServiceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/linkedServices/{linkedServiceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client LinkedServicesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client LinkedServicesClient) GetResponder(resp *http.Response) (result LinkedService, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByWorkspace gets the linked services instances in a workspace.
// Parameters:
// resourceGroupName - the name of the resource group to get. The name is case insensitive.
// workspaceName - name of the Log Analytics Workspace that contains the linked services.
func (client LinkedServicesClient) ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string) (result LinkedServiceListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LinkedServicesClient.ListByWorkspace")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("operationalinsights.LinkedServicesClient", "ListByWorkspace", err.Error())
	}

	req, err := client.ListByWorkspacePreparer(ctx, resourceGroupName, workspaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.LinkedServicesClient", "ListByWorkspace", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByWorkspaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "operationalinsights.LinkedServicesClient", "ListByWorkspace", resp, "Failure sending request")
		return
	}

	result, err = client.ListByWorkspaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.LinkedServicesClient", "ListByWorkspace", resp, "Failure responding to request")
		return
	}

	return
}

// ListByWorkspacePreparer prepares the ListByWorkspace request.
func (client LinkedServicesClient) ListByWorkspacePreparer(ctx context.Context, resourceGroupName string, workspaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/linkedServices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByWorkspaceSender sends the ListByWorkspace request. The method will close the
// http.Response Body if it receives an error.
func (client LinkedServicesClient) ListByWorkspaceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByWorkspaceResponder handles the response to the ListByWorkspace request. The method always
// closes the http.Response Body.
func (client LinkedServicesClient) ListByWorkspaceResponder(resp *http.Response) (result LinkedServiceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
