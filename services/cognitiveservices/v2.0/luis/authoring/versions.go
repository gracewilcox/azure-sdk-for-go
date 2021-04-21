package authoring

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
	"github.com/gofrs/uuid"
	"net/http"
)

// VersionsClient is the client for the Versions methods of the Authoring service.
type VersionsClient struct {
	BaseClient
}

// NewVersionsClient creates an instance of the VersionsClient client.
func NewVersionsClient(endpoint string) VersionsClient {
	return VersionsClient{New(endpoint)}
}

// Clone creates a new version from the selected version.
// Parameters:
// appID - the application ID.
// versionID - the version ID.
// versionCloneObject - a model containing the new version ID.
func (client VersionsClient) Clone(ctx context.Context, appID uuid.UUID, versionID string, versionCloneObject TaskUpdateObject) (result String, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VersionsClient.Clone")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ClonePreparer(ctx, appID, versionID, versionCloneObject)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.VersionsClient", "Clone", nil, "Failure preparing request")
		return
	}

	resp, err := client.CloneSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authoring.VersionsClient", "Clone", resp, "Failure sending request")
		return
	}

	result, err = client.CloneResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.VersionsClient", "Clone", resp, "Failure responding to request")
		return
	}

	return
}

// ClonePreparer prepares the Clone request.
func (client VersionsClient) ClonePreparer(ctx context.Context, appID uuid.UUID, versionID string, versionCloneObject TaskUpdateObject) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"appId":     autorest.Encode("path", appID),
		"versionId": autorest.Encode("path", versionID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/luis/api/v2.0", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/versions/{versionId}/clone", pathParameters),
		autorest.WithJSON(versionCloneObject))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CloneSender sends the Clone request. The method will close the
// http.Response Body if it receives an error.
func (client VersionsClient) CloneSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CloneResponder handles the response to the Clone request. The method always
// closes the http.Response Body.
func (client VersionsClient) CloneResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an application version.
// Parameters:
// appID - the application ID.
// versionID - the version ID.
func (client VersionsClient) Delete(ctx context.Context, appID uuid.UUID, versionID string) (result OperationStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VersionsClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, appID, versionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.VersionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authoring.VersionsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.VersionsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VersionsClient) DeletePreparer(ctx context.Context, appID uuid.UUID, versionID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"appId":     autorest.Encode("path", appID),
		"versionId": autorest.Encode("path", versionID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{Endpoint}/luis/api/v2.0", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/versions/{versionId}/", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VersionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VersionsClient) DeleteResponder(resp *http.Response) (result OperationStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteUnlabelledUtterance deleted an unlabelled utterance in a version of the application.
// Parameters:
// appID - the application ID.
// versionID - the version ID.
// utterance - the utterance text to delete.
func (client VersionsClient) DeleteUnlabelledUtterance(ctx context.Context, appID uuid.UUID, versionID string, utterance string) (result OperationStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VersionsClient.DeleteUnlabelledUtterance")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteUnlabelledUtterancePreparer(ctx, appID, versionID, utterance)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.VersionsClient", "DeleteUnlabelledUtterance", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteUnlabelledUtteranceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authoring.VersionsClient", "DeleteUnlabelledUtterance", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteUnlabelledUtteranceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.VersionsClient", "DeleteUnlabelledUtterance", resp, "Failure responding to request")
		return
	}

	return
}

// DeleteUnlabelledUtterancePreparer prepares the DeleteUnlabelledUtterance request.
func (client VersionsClient) DeleteUnlabelledUtterancePreparer(ctx context.Context, appID uuid.UUID, versionID string, utterance string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"appId":     autorest.Encode("path", appID),
		"versionId": autorest.Encode("path", versionID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{Endpoint}/luis/api/v2.0", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/versions/{versionId}/suggest", pathParameters),
		autorest.WithJSON(utterance))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteUnlabelledUtteranceSender sends the DeleteUnlabelledUtterance request. The method will close the
// http.Response Body if it receives an error.
func (client VersionsClient) DeleteUnlabelledUtteranceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteUnlabelledUtteranceResponder handles the response to the DeleteUnlabelledUtterance request. The method always
// closes the http.Response Body.
func (client VersionsClient) DeleteUnlabelledUtteranceResponder(resp *http.Response) (result OperationStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Export exports a LUIS application to JSON format.
// Parameters:
// appID - the application ID.
// versionID - the version ID.
func (client VersionsClient) Export(ctx context.Context, appID uuid.UUID, versionID string) (result LuisApp, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VersionsClient.Export")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ExportPreparer(ctx, appID, versionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.VersionsClient", "Export", nil, "Failure preparing request")
		return
	}

	resp, err := client.ExportSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authoring.VersionsClient", "Export", resp, "Failure sending request")
		return
	}

	result, err = client.ExportResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.VersionsClient", "Export", resp, "Failure responding to request")
		return
	}

	return
}

// ExportPreparer prepares the Export request.
func (client VersionsClient) ExportPreparer(ctx context.Context, appID uuid.UUID, versionID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"appId":     autorest.Encode("path", appID),
		"versionId": autorest.Encode("path", versionID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/luis/api/v2.0", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/versions/{versionId}/export", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ExportSender sends the Export request. The method will close the
// http.Response Body if it receives an error.
func (client VersionsClient) ExportSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ExportResponder handles the response to the Export request. The method always
// closes the http.Response Body.
func (client VersionsClient) ExportResponder(resp *http.Response) (result LuisApp, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets the version information such as date created, last modified date, endpoint URL, count of intents and
// entities, training and publishing status.
// Parameters:
// appID - the application ID.
// versionID - the version ID.
func (client VersionsClient) Get(ctx context.Context, appID uuid.UUID, versionID string) (result VersionInfo, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VersionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, appID, versionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.VersionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authoring.VersionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.VersionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client VersionsClient) GetPreparer(ctx context.Context, appID uuid.UUID, versionID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"appId":     autorest.Encode("path", appID),
		"versionId": autorest.Encode("path", versionID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/luis/api/v2.0", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/versions/{versionId}/", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VersionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VersionsClient) GetResponder(resp *http.Response) (result VersionInfo, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Import imports a new version into a LUIS application.
// Parameters:
// appID - the application ID.
// luisApp - a LUIS application structure.
// versionID - the new versionId to import. If not specified, the versionId will be read from the imported
// object.
func (client VersionsClient) Import(ctx context.Context, appID uuid.UUID, luisApp LuisApp, versionID string) (result String, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VersionsClient.Import")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ImportPreparer(ctx, appID, luisApp, versionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.VersionsClient", "Import", nil, "Failure preparing request")
		return
	}

	resp, err := client.ImportSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authoring.VersionsClient", "Import", resp, "Failure sending request")
		return
	}

	result, err = client.ImportResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.VersionsClient", "Import", resp, "Failure responding to request")
		return
	}

	return
}

// ImportPreparer prepares the Import request.
func (client VersionsClient) ImportPreparer(ctx context.Context, appID uuid.UUID, luisApp LuisApp, versionID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"appId": autorest.Encode("path", appID),
	}

	queryParameters := map[string]interface{}{}
	if len(versionID) > 0 {
		queryParameters["versionId"] = autorest.Encode("query", versionID)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/luis/api/v2.0", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/versions/import", pathParameters),
		autorest.WithJSON(luisApp),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ImportSender sends the Import request. The method will close the
// http.Response Body if it receives an error.
func (client VersionsClient) ImportSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ImportResponder handles the response to the Import request. The method always
// closes the http.Response Body.
func (client VersionsClient) ImportResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of versions for this application ID.
// Parameters:
// appID - the application ID.
// skip - the number of entries to skip. Default value is 0.
// take - the number of entries to return. Maximum page size is 500. Default is 100.
func (client VersionsClient) List(ctx context.Context, appID uuid.UUID, skip *int32, take *int32) (result ListVersionInfo, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VersionsClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil}}}}},
		{TargetValue: take,
			Constraints: []validation.Constraint{{Target: "take", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "take", Name: validation.InclusiveMaximum, Rule: int64(500), Chain: nil},
					{Target: "take", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("authoring.VersionsClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, appID, skip, take)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.VersionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authoring.VersionsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.VersionsClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client VersionsClient) ListPreparer(ctx context.Context, appID uuid.UUID, skip *int32, take *int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"appId": autorest.Encode("path", appID),
	}

	queryParameters := map[string]interface{}{}
	if skip != nil {
		queryParameters["skip"] = autorest.Encode("query", *skip)
	} else {
		queryParameters["skip"] = autorest.Encode("query", 0)
	}
	if take != nil {
		queryParameters["take"] = autorest.Encode("query", *take)
	} else {
		queryParameters["take"] = autorest.Encode("query", 100)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/luis/api/v2.0", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/versions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VersionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VersionsClient) ListResponder(resp *http.Response) (result ListVersionInfo, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates the name or description of the application version.
// Parameters:
// appID - the application ID.
// versionID - the version ID.
// versionUpdateObject - a model containing Name and Description of the application.
func (client VersionsClient) Update(ctx context.Context, appID uuid.UUID, versionID string, versionUpdateObject TaskUpdateObject) (result OperationStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VersionsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, appID, versionID, versionUpdateObject)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.VersionsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authoring.VersionsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authoring.VersionsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client VersionsClient) UpdatePreparer(ctx context.Context, appID uuid.UUID, versionID string, versionUpdateObject TaskUpdateObject) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"appId":     autorest.Encode("path", appID),
		"versionId": autorest.Encode("path", versionID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("{Endpoint}/luis/api/v2.0", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/versions/{versionId}/", pathParameters),
		autorest.WithJSON(versionUpdateObject))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client VersionsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client VersionsClient) UpdateResponder(resp *http.Response) (result OperationStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
