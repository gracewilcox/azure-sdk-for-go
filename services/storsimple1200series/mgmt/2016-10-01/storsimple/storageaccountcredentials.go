package storsimple

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

// StorageAccountCredentialsClient is the client for the StorageAccountCredentials methods of the Storsimple service.
type StorageAccountCredentialsClient struct {
	BaseClient
}

// NewStorageAccountCredentialsClient creates an instance of the StorageAccountCredentialsClient client.
func NewStorageAccountCredentialsClient(subscriptionID string) StorageAccountCredentialsClient {
	return NewStorageAccountCredentialsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewStorageAccountCredentialsClientWithBaseURI creates an instance of the StorageAccountCredentialsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewStorageAccountCredentialsClientWithBaseURI(baseURI string, subscriptionID string) StorageAccountCredentialsClient {
	return StorageAccountCredentialsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates the storage account credential
// Parameters:
// credentialName - the credential name.
// storageAccount - the storage account credential to be added or updated.
// resourceGroupName - the resource group name
// managerName - the manager name
func (client StorageAccountCredentialsClient) CreateOrUpdate(ctx context.Context, credentialName string, storageAccount StorageAccountCredential, resourceGroupName string, managerName string) (result StorageAccountCredentialsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageAccountCredentialsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: storageAccount,
			Constraints: []validation.Constraint{{Target: "storageAccount.StorageAccountCredentialProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "storageAccount.StorageAccountCredentialProperties.EndPoint", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "storageAccount.StorageAccountCredentialProperties.Login", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "storageAccount.StorageAccountCredentialProperties.AccessKey", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "storageAccount.StorageAccountCredentialProperties.AccessKey.Value", Name: validation.Null, Rule: true, Chain: nil}}},
				}}}},
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.StorageAccountCredentialsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, credentialName, storageAccount, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.StorageAccountCredentialsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.StorageAccountCredentialsClient", "CreateOrUpdate", nil, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client StorageAccountCredentialsClient) CreateOrUpdatePreparer(ctx context.Context, credentialName string, storageAccount StorageAccountCredential, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"credentialName":    autorest.Encode("path", credentialName),
		"managerName":       autorest.Encode("path", managerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials/{credentialName}", pathParameters),
		autorest.WithJSON(storageAccount),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client StorageAccountCredentialsClient) CreateOrUpdateSender(req *http.Request) (future StorageAccountCredentialsCreateOrUpdateFuture, err error) {
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

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client StorageAccountCredentialsClient) CreateOrUpdateResponder(resp *http.Response) (result StorageAccountCredential, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the storage account credential
// Parameters:
// credentialName - the name of the storage account credential.
// resourceGroupName - the resource group name
// managerName - the manager name
func (client StorageAccountCredentialsClient) Delete(ctx context.Context, credentialName string, resourceGroupName string, managerName string) (result StorageAccountCredentialsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageAccountCredentialsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.StorageAccountCredentialsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, credentialName, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.StorageAccountCredentialsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.StorageAccountCredentialsClient", "Delete", nil, "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client StorageAccountCredentialsClient) DeletePreparer(ctx context.Context, credentialName string, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"credentialName":    autorest.Encode("path", credentialName),
		"managerName":       autorest.Encode("path", managerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials/{credentialName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client StorageAccountCredentialsClient) DeleteSender(req *http.Request) (future StorageAccountCredentialsDeleteFuture, err error) {
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

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client StorageAccountCredentialsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get returns the properties of the specified storage account credential name.
// Parameters:
// credentialName - the name of storage account credential to be fetched.
// resourceGroupName - the resource group name
// managerName - the manager name
func (client StorageAccountCredentialsClient) Get(ctx context.Context, credentialName string, resourceGroupName string, managerName string) (result StorageAccountCredential, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageAccountCredentialsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.StorageAccountCredentialsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, credentialName, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.StorageAccountCredentialsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storsimple.StorageAccountCredentialsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.StorageAccountCredentialsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client StorageAccountCredentialsClient) GetPreparer(ctx context.Context, credentialName string, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"credentialName":    autorest.Encode("path", credentialName),
		"managerName":       autorest.Encode("path", managerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials/{credentialName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client StorageAccountCredentialsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client StorageAccountCredentialsClient) GetResponder(resp *http.Response) (result StorageAccountCredential, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByManager retrieves all the storage account credentials in a manager.
// Parameters:
// resourceGroupName - the resource group name
// managerName - the manager name
func (client StorageAccountCredentialsClient) ListByManager(ctx context.Context, resourceGroupName string, managerName string) (result StorageAccountCredentialList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageAccountCredentialsClient.ListByManager")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.StorageAccountCredentialsClient", "ListByManager", err.Error())
	}

	req, err := client.ListByManagerPreparer(ctx, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.StorageAccountCredentialsClient", "ListByManager", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByManagerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storsimple.StorageAccountCredentialsClient", "ListByManager", resp, "Failure sending request")
		return
	}

	result, err = client.ListByManagerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.StorageAccountCredentialsClient", "ListByManager", resp, "Failure responding to request")
		return
	}

	return
}

// ListByManagerPreparer prepares the ListByManager request.
func (client StorageAccountCredentialsClient) ListByManagerPreparer(ctx context.Context, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managerName":       autorest.Encode("path", managerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByManagerSender sends the ListByManager request. The method will close the
// http.Response Body if it receives an error.
func (client StorageAccountCredentialsClient) ListByManagerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByManagerResponder handles the response to the ListByManager request. The method always
// closes the http.Response Body.
func (client StorageAccountCredentialsClient) ListByManagerResponder(resp *http.Response) (result StorageAccountCredentialList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
