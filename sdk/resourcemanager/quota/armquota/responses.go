//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armquota

// ClientCreateOrUpdateResponse contains the response from method Client.BeginCreateOrUpdate.
type ClientCreateOrUpdateResponse struct {
	// Quota limit.
	CurrentQuotaLimitBase
}

// ClientGetResponse contains the response from method Client.Get.
type ClientGetResponse struct {
	// Quota limit.
	CurrentQuotaLimitBase

	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// ClientListResponse contains the response from method Client.NewListPager.
type ClientListResponse struct {
	// Quota limits.
	Limits

	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// ClientUpdateResponse contains the response from method Client.BeginUpdate.
type ClientUpdateResponse struct {
	// Quota limit.
	CurrentQuotaLimitBase
}

// GroupQuotaLimitsClientGetResponse contains the response from method GroupQuotaLimitsClient.Get.
type GroupQuotaLimitsClientGetResponse struct {
	// Group Quota limit.
	GroupQuotaLimit
}

// GroupQuotaLimitsClientListResponse contains the response from method GroupQuotaLimitsClient.NewListPager.
type GroupQuotaLimitsClientListResponse struct {
	// List of Group Quota Limit details.
	GroupQuotaLimitList
}

// GroupQuotaLimitsRequestClientCreateOrUpdateResponse contains the response from method GroupQuotaLimitsRequestClient.BeginCreateOrUpdate.
type GroupQuotaLimitsRequestClientCreateOrUpdateResponse struct {
	// Status of a single GroupQuota request.
	SubmittedResourceRequestStatus
}

// GroupQuotaLimitsRequestClientGetResponse contains the response from method GroupQuotaLimitsRequestClient.Get.
type GroupQuotaLimitsRequestClientGetResponse struct {
	// Status of a single GroupQuota request.
	SubmittedResourceRequestStatus
}

// GroupQuotaLimitsRequestClientListResponse contains the response from method GroupQuotaLimitsRequestClient.NewListPager.
type GroupQuotaLimitsRequestClientListResponse struct {
	// Share Quota Entity list.
	SubmittedResourceRequestStatusList
}

// GroupQuotaLimitsRequestClientUpdateResponse contains the response from method GroupQuotaLimitsRequestClient.BeginUpdate.
type GroupQuotaLimitsRequestClientUpdateResponse struct {
	// Status of a single GroupQuota request.
	SubmittedResourceRequestStatus
}

// GroupQuotaLocationSettingsClientCreateOrUpdateResponse contains the response from method GroupQuotaLocationSettingsClient.BeginCreateOrUpdate.
type GroupQuotaLocationSettingsClientCreateOrUpdateResponse struct {
	// The GroupQuota Enforcement status for a Azure Location/Region.
	GroupQuotasEnforcementResponse
}

// GroupQuotaLocationSettingsClientGetResponse contains the response from method GroupQuotaLocationSettingsClient.Get.
type GroupQuotaLocationSettingsClientGetResponse struct {
	// The GroupQuota Enforcement status for a Azure Location/Region.
	GroupQuotasEnforcementResponse
}

// GroupQuotaLocationSettingsClientListResponse contains the response from method GroupQuotaLocationSettingsClient.NewListPager.
type GroupQuotaLocationSettingsClientListResponse struct {
	// List of Azure regions, where the group quotas is enabled for enforcement.
	GroupQuotasEnforcementListResponse
}

// GroupQuotaLocationSettingsClientUpdateResponse contains the response from method GroupQuotaLocationSettingsClient.BeginUpdate.
type GroupQuotaLocationSettingsClientUpdateResponse struct {
	// The GroupQuota Enforcement status for a Azure Location/Region.
	GroupQuotasEnforcementResponse
}

// GroupQuotaSubscriptionAllocationClientGetResponse contains the response from method GroupQuotaSubscriptionAllocationClient.Get.
type GroupQuotaSubscriptionAllocationClientGetResponse struct {
	// Quota allocated to a subscription for the specific Resource Provider, Location, ResourceName. This will include the GroupQuota
	// and total quota allocated to the subscription. Only the Group quota allocated to the subscription can be allocated back
	// to the MG Group Quota.
	SubscriptionQuotaAllocations
}

// GroupQuotaSubscriptionAllocationClientListResponse contains the response from method GroupQuotaSubscriptionAllocationClient.NewListPager.
type GroupQuotaSubscriptionAllocationClientListResponse struct {
	// Subscription quota list.
	SubscriptionQuotaAllocationsList
}

// GroupQuotaSubscriptionAllocationRequestClientCreateOrUpdateResponse contains the response from method GroupQuotaSubscriptionAllocationRequestClient.BeginCreateOrUpdate.
type GroupQuotaSubscriptionAllocationRequestClientCreateOrUpdateResponse struct {
	// The subscription quota allocation status.
	AllocationRequestStatus
}

// GroupQuotaSubscriptionAllocationRequestClientGetResponse contains the response from method GroupQuotaSubscriptionAllocationRequestClient.Get.
type GroupQuotaSubscriptionAllocationRequestClientGetResponse struct {
	// The subscription quota allocation status.
	AllocationRequestStatus
}

// GroupQuotaSubscriptionAllocationRequestClientListResponse contains the response from method GroupQuotaSubscriptionAllocationRequestClient.NewListPager.
type GroupQuotaSubscriptionAllocationRequestClientListResponse struct {
	// List of QuotaAllocation Request Status
	AllocationRequestStatusList
}

// GroupQuotaSubscriptionAllocationRequestClientUpdateResponse contains the response from method GroupQuotaSubscriptionAllocationRequestClient.BeginUpdate.
type GroupQuotaSubscriptionAllocationRequestClientUpdateResponse struct {
	// The subscription quota allocation status.
	AllocationRequestStatus
}

// GroupQuotaSubscriptionRequestsClientGetResponse contains the response from method GroupQuotaSubscriptionRequestsClient.Get.
type GroupQuotaSubscriptionRequestsClientGetResponse struct {
	// The new quota limit request status.
	GroupQuotaSubscriptionRequestStatus
}

// GroupQuotaSubscriptionRequestsClientListResponse contains the response from method GroupQuotaSubscriptionRequestsClient.NewListPager.
type GroupQuotaSubscriptionRequestsClientListResponse struct {
	// List of GroupQuotaSubscriptionRequests Status
	GroupQuotaSubscriptionRequestStatusList
}

// GroupQuotaSubscriptionsClientCreateOrUpdateResponse contains the response from method GroupQuotaSubscriptionsClient.BeginCreateOrUpdate.
type GroupQuotaSubscriptionsClientCreateOrUpdateResponse struct {
	// This represents a Azure subscriptionId that is associated with a GroupQuotasEntity.
	GroupQuotaSubscriptionID
}

// GroupQuotaSubscriptionsClientDeleteResponse contains the response from method GroupQuotaSubscriptionsClient.BeginDelete.
type GroupQuotaSubscriptionsClientDeleteResponse struct {
	// placeholder for future response values
}

// GroupQuotaSubscriptionsClientGetResponse contains the response from method GroupQuotaSubscriptionsClient.Get.
type GroupQuotaSubscriptionsClientGetResponse struct {
	// This represents a Azure subscriptionId that is associated with a GroupQuotasEntity.
	GroupQuotaSubscriptionID
}

// GroupQuotaSubscriptionsClientListResponse contains the response from method GroupQuotaSubscriptionsClient.NewListPager.
type GroupQuotaSubscriptionsClientListResponse struct {
	// List of GroupQuotaSubscriptionIds
	GroupQuotaSubscriptionIDList
}

// GroupQuotaSubscriptionsClientUpdateResponse contains the response from method GroupQuotaSubscriptionsClient.BeginUpdate.
type GroupQuotaSubscriptionsClientUpdateResponse struct {
	// This represents a Azure subscriptionId that is associated with a GroupQuotasEntity.
	GroupQuotaSubscriptionID
}

// GroupQuotaUsagesClientListResponse contains the response from method GroupQuotaUsagesClient.NewListPager.
type GroupQuotaUsagesClientListResponse struct {
	// List of resource usages and quotas for GroupQuota.
	ResourceUsageList
}

// GroupQuotasClientCreateOrUpdateResponse contains the response from method GroupQuotasClient.BeginCreateOrUpdate.
type GroupQuotasClientCreateOrUpdateResponse struct {
	// Properties and filters for ShareQuota. The request parameter is optional, if there are no filters specified.
	GroupQuotasEntity
}

// GroupQuotasClientDeleteResponse contains the response from method GroupQuotasClient.BeginDelete.
type GroupQuotasClientDeleteResponse struct {
	// placeholder for future response values
}

// GroupQuotasClientGetResponse contains the response from method GroupQuotasClient.Get.
type GroupQuotasClientGetResponse struct {
	// Properties and filters for ShareQuota. The request parameter is optional, if there are no filters specified.
	GroupQuotasEntity
}

// GroupQuotasClientListResponse contains the response from method GroupQuotasClient.NewListPager.
type GroupQuotasClientListResponse struct {
	// List of Group Quotas at MG level.
	GroupQuotaList
}

// GroupQuotasClientUpdateResponse contains the response from method GroupQuotasClient.BeginUpdate.
type GroupQuotasClientUpdateResponse struct {
	// Properties and filters for ShareQuota. The request parameter is optional, if there are no filters specified.
	GroupQuotasEntity
}

// OperationClientListResponse contains the response from method OperationClient.NewListPager.
type OperationClientListResponse struct {
	OperationList
}

// RequestStatusClientGetResponse contains the response from method RequestStatusClient.Get.
type RequestStatusClientGetResponse struct {
	// List of quota requests with details.
	RequestDetails
}

// RequestStatusClientListResponse contains the response from method RequestStatusClient.NewListPager.
type RequestStatusClientListResponse struct {
	// Quota request information.
	RequestDetailsList
}

// UsagesClientGetResponse contains the response from method UsagesClient.Get.
type UsagesClientGetResponse struct {
	// Resource usage.
	CurrentUsagesBase

	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// UsagesClientListResponse contains the response from method UsagesClient.NewListPager.
type UsagesClientListResponse struct {
	// Quota limits.
	UsagesLimits

	// ETag contains the information returned from the ETag header response.
	ETag *string
}
