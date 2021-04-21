// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package devices

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/iothub/mgmt/2020-03-01/devices"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AccessRights = original.AccessRights

const (
	DeviceConnect                                        AccessRights = original.DeviceConnect
	RegistryRead                                         AccessRights = original.RegistryRead
	RegistryReadDeviceConnect                            AccessRights = original.RegistryReadDeviceConnect
	RegistryReadRegistryWrite                            AccessRights = original.RegistryReadRegistryWrite
	RegistryReadRegistryWriteDeviceConnect               AccessRights = original.RegistryReadRegistryWriteDeviceConnect
	RegistryReadRegistryWriteServiceConnect              AccessRights = original.RegistryReadRegistryWriteServiceConnect
	RegistryReadRegistryWriteServiceConnectDeviceConnect AccessRights = original.RegistryReadRegistryWriteServiceConnectDeviceConnect
	RegistryReadServiceConnect                           AccessRights = original.RegistryReadServiceConnect
	RegistryReadServiceConnectDeviceConnect              AccessRights = original.RegistryReadServiceConnectDeviceConnect
	RegistryWrite                                        AccessRights = original.RegistryWrite
	RegistryWriteDeviceConnect                           AccessRights = original.RegistryWriteDeviceConnect
	RegistryWriteServiceConnect                          AccessRights = original.RegistryWriteServiceConnect
	RegistryWriteServiceConnectDeviceConnect             AccessRights = original.RegistryWriteServiceConnectDeviceConnect
	ServiceConnect                                       AccessRights = original.ServiceConnect
	ServiceConnectDeviceConnect                          AccessRights = original.ServiceConnectDeviceConnect
)

type AuthenticationType = original.AuthenticationType

const (
	IdentityBased AuthenticationType = original.IdentityBased
	KeyBased      AuthenticationType = original.KeyBased
)

type Capabilities = original.Capabilities

const (
	DeviceManagement Capabilities = original.DeviceManagement
	None             Capabilities = original.None
)

type Encoding = original.Encoding

const (
	Avro        Encoding = original.Avro
	AvroDeflate Encoding = original.AvroDeflate
	JSON        Encoding = original.JSON
)

type EndpointHealthStatus = original.EndpointHealthStatus

const (
	Dead      EndpointHealthStatus = original.Dead
	Healthy   EndpointHealthStatus = original.Healthy
	Unhealthy EndpointHealthStatus = original.Unhealthy
	Unknown   EndpointHealthStatus = original.Unknown
)

type IPFilterActionType = original.IPFilterActionType

const (
	Accept IPFilterActionType = original.Accept
	Reject IPFilterActionType = original.Reject
)

type IotHubNameUnavailabilityReason = original.IotHubNameUnavailabilityReason

const (
	AlreadyExists IotHubNameUnavailabilityReason = original.AlreadyExists
	Invalid       IotHubNameUnavailabilityReason = original.Invalid
)

type IotHubReplicaRoleType = original.IotHubReplicaRoleType

const (
	Primary   IotHubReplicaRoleType = original.Primary
	Secondary IotHubReplicaRoleType = original.Secondary
)

type IotHubScaleType = original.IotHubScaleType

const (
	IotHubScaleTypeAutomatic IotHubScaleType = original.IotHubScaleTypeAutomatic
	IotHubScaleTypeManual    IotHubScaleType = original.IotHubScaleTypeManual
	IotHubScaleTypeNone      IotHubScaleType = original.IotHubScaleTypeNone
)

type IotHubSku = original.IotHubSku

const (
	B1 IotHubSku = original.B1
	B2 IotHubSku = original.B2
	B3 IotHubSku = original.B3
	F1 IotHubSku = original.F1
	S1 IotHubSku = original.S1
	S2 IotHubSku = original.S2
	S3 IotHubSku = original.S3
)

type IotHubSkuTier = original.IotHubSkuTier

const (
	Basic    IotHubSkuTier = original.Basic
	Free     IotHubSkuTier = original.Free
	Standard IotHubSkuTier = original.Standard
)

type JobStatus = original.JobStatus

const (
	JobStatusCancelled JobStatus = original.JobStatusCancelled
	JobStatusCompleted JobStatus = original.JobStatusCompleted
	JobStatusEnqueued  JobStatus = original.JobStatusEnqueued
	JobStatusFailed    JobStatus = original.JobStatusFailed
	JobStatusRunning   JobStatus = original.JobStatusRunning
	JobStatusUnknown   JobStatus = original.JobStatusUnknown
)

type JobType = original.JobType

const (
	JobTypeBackup                    JobType = original.JobTypeBackup
	JobTypeExport                    JobType = original.JobTypeExport
	JobTypeFactoryResetDevice        JobType = original.JobTypeFactoryResetDevice
	JobTypeFirmwareUpdate            JobType = original.JobTypeFirmwareUpdate
	JobTypeImport                    JobType = original.JobTypeImport
	JobTypeReadDeviceProperties      JobType = original.JobTypeReadDeviceProperties
	JobTypeRebootDevice              JobType = original.JobTypeRebootDevice
	JobTypeUnknown                   JobType = original.JobTypeUnknown
	JobTypeUpdateDeviceConfiguration JobType = original.JobTypeUpdateDeviceConfiguration
	JobTypeWriteDeviceProperties     JobType = original.JobTypeWriteDeviceProperties
)

type PrivateLinkServiceConnectionStatus = original.PrivateLinkServiceConnectionStatus

const (
	Approved     PrivateLinkServiceConnectionStatus = original.Approved
	Disconnected PrivateLinkServiceConnectionStatus = original.Disconnected
	Pending      PrivateLinkServiceConnectionStatus = original.Pending
	Rejected     PrivateLinkServiceConnectionStatus = original.Rejected
)

type PublicNetworkAccess = original.PublicNetworkAccess

const (
	Disabled PublicNetworkAccess = original.Disabled
	Enabled  PublicNetworkAccess = original.Enabled
)

type RouteErrorSeverity = original.RouteErrorSeverity

const (
	Error   RouteErrorSeverity = original.Error
	Warning RouteErrorSeverity = original.Warning
)

type RoutingSource = original.RoutingSource

const (
	RoutingSourceDeviceJobLifecycleEvents RoutingSource = original.RoutingSourceDeviceJobLifecycleEvents
	RoutingSourceDeviceLifecycleEvents    RoutingSource = original.RoutingSourceDeviceLifecycleEvents
	RoutingSourceDeviceMessages           RoutingSource = original.RoutingSourceDeviceMessages
	RoutingSourceInvalid                  RoutingSource = original.RoutingSourceInvalid
	RoutingSourceTwinChangeEvents         RoutingSource = original.RoutingSourceTwinChangeEvents
)

type TestResultStatus = original.TestResultStatus

const (
	False     TestResultStatus = original.False
	True      TestResultStatus = original.True
	Undefined TestResultStatus = original.Undefined
)

type BaseClient = original.BaseClient
type CertificateBodyDescription = original.CertificateBodyDescription
type CertificateDescription = original.CertificateDescription
type CertificateListDescription = original.CertificateListDescription
type CertificateProperties = original.CertificateProperties
type CertificatePropertiesWithNonce = original.CertificatePropertiesWithNonce
type CertificateVerificationDescription = original.CertificateVerificationDescription
type CertificateWithNonceDescription = original.CertificateWithNonceDescription
type CertificatesClient = original.CertificatesClient
type CloudToDeviceProperties = original.CloudToDeviceProperties
type EndpointHealthData = original.EndpointHealthData
type EndpointHealthDataListResult = original.EndpointHealthDataListResult
type EndpointHealthDataListResultIterator = original.EndpointHealthDataListResultIterator
type EndpointHealthDataListResultPage = original.EndpointHealthDataListResultPage
type EnrichmentProperties = original.EnrichmentProperties
type ErrorDetails = original.ErrorDetails
type EventHubConsumerGroupInfo = original.EventHubConsumerGroupInfo
type EventHubConsumerGroupsListResult = original.EventHubConsumerGroupsListResult
type EventHubConsumerGroupsListResultIterator = original.EventHubConsumerGroupsListResultIterator
type EventHubConsumerGroupsListResultPage = original.EventHubConsumerGroupsListResultPage
type EventHubProperties = original.EventHubProperties
type ExportDevicesRequest = original.ExportDevicesRequest
type FailoverInput = original.FailoverInput
type FallbackRouteProperties = original.FallbackRouteProperties
type FeedbackProperties = original.FeedbackProperties
type GroupIDInformation = original.GroupIDInformation
type GroupIDInformationProperties = original.GroupIDInformationProperties
type IPFilterRule = original.IPFilterRule
type ImportDevicesRequest = original.ImportDevicesRequest
type IotHubCapacity = original.IotHubCapacity
type IotHubClient = original.IotHubClient
type IotHubDescription = original.IotHubDescription
type IotHubDescriptionListResult = original.IotHubDescriptionListResult
type IotHubDescriptionListResultIterator = original.IotHubDescriptionListResultIterator
type IotHubDescriptionListResultPage = original.IotHubDescriptionListResultPage
type IotHubLocationDescription = original.IotHubLocationDescription
type IotHubManualFailoverFuture = original.IotHubManualFailoverFuture
type IotHubNameAvailabilityInfo = original.IotHubNameAvailabilityInfo
type IotHubProperties = original.IotHubProperties
type IotHubQuotaMetricInfo = original.IotHubQuotaMetricInfo
type IotHubQuotaMetricInfoListResult = original.IotHubQuotaMetricInfoListResult
type IotHubQuotaMetricInfoListResultIterator = original.IotHubQuotaMetricInfoListResultIterator
type IotHubQuotaMetricInfoListResultPage = original.IotHubQuotaMetricInfoListResultPage
type IotHubResourceClient = original.IotHubResourceClient
type IotHubResourceCreateOrUpdateFuture = original.IotHubResourceCreateOrUpdateFuture
type IotHubResourceDeleteFuture = original.IotHubResourceDeleteFuture
type IotHubResourceUpdateFuture = original.IotHubResourceUpdateFuture
type IotHubSkuDescription = original.IotHubSkuDescription
type IotHubSkuDescriptionListResult = original.IotHubSkuDescriptionListResult
type IotHubSkuDescriptionListResultIterator = original.IotHubSkuDescriptionListResultIterator
type IotHubSkuDescriptionListResultPage = original.IotHubSkuDescriptionListResultPage
type IotHubSkuInfo = original.IotHubSkuInfo
type JobResponse = original.JobResponse
type JobResponseListResult = original.JobResponseListResult
type JobResponseListResultIterator = original.JobResponseListResultIterator
type JobResponseListResultPage = original.JobResponseListResultPage
type ListPrivateEndpointConnection = original.ListPrivateEndpointConnection
type MatchedRoute = original.MatchedRoute
type MessagingEndpointProperties = original.MessagingEndpointProperties
type Name = original.Name
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationInputs = original.OperationInputs
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateEndpointConnectionsDeleteFuture = original.PrivateEndpointConnectionsDeleteFuture
type PrivateEndpointConnectionsUpdateFuture = original.PrivateEndpointConnectionsUpdateFuture
type PrivateLinkResources = original.PrivateLinkResources
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type PrivateLinkServiceConnectionState = original.PrivateLinkServiceConnectionState
type RegistryStatistics = original.RegistryStatistics
type Resource = original.Resource
type ResourceProviderCommonClient = original.ResourceProviderCommonClient
type RouteCompilationError = original.RouteCompilationError
type RouteErrorPosition = original.RouteErrorPosition
type RouteErrorRange = original.RouteErrorRange
type RouteProperties = original.RouteProperties
type RoutingEndpoints = original.RoutingEndpoints
type RoutingEventHubProperties = original.RoutingEventHubProperties
type RoutingMessage = original.RoutingMessage
type RoutingProperties = original.RoutingProperties
type RoutingServiceBusQueueEndpointProperties = original.RoutingServiceBusQueueEndpointProperties
type RoutingServiceBusTopicEndpointProperties = original.RoutingServiceBusTopicEndpointProperties
type RoutingStorageContainerProperties = original.RoutingStorageContainerProperties
type RoutingTwin = original.RoutingTwin
type RoutingTwinProperties = original.RoutingTwinProperties
type SetObject = original.SetObject
type SharedAccessSignatureAuthorizationRule = original.SharedAccessSignatureAuthorizationRule
type SharedAccessSignatureAuthorizationRuleListResult = original.SharedAccessSignatureAuthorizationRuleListResult
type SharedAccessSignatureAuthorizationRuleListResultIterator = original.SharedAccessSignatureAuthorizationRuleListResultIterator
type SharedAccessSignatureAuthorizationRuleListResultPage = original.SharedAccessSignatureAuthorizationRuleListResultPage
type StorageEndpointProperties = original.StorageEndpointProperties
type TagsResource = original.TagsResource
type TestAllRoutesInput = original.TestAllRoutesInput
type TestAllRoutesResult = original.TestAllRoutesResult
type TestRouteInput = original.TestRouteInput
type TestRouteResult = original.TestRouteResult
type TestRouteResultDetails = original.TestRouteResultDetails
type UserSubscriptionQuota = original.UserSubscriptionQuota
type UserSubscriptionQuotaListResult = original.UserSubscriptionQuotaListResult

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewCertificatesClient(subscriptionID string) CertificatesClient {
	return original.NewCertificatesClient(subscriptionID)
}
func NewCertificatesClientWithBaseURI(baseURI string, subscriptionID string) CertificatesClient {
	return original.NewCertificatesClientWithBaseURI(baseURI, subscriptionID)
}
func NewEndpointHealthDataListResultIterator(page EndpointHealthDataListResultPage) EndpointHealthDataListResultIterator {
	return original.NewEndpointHealthDataListResultIterator(page)
}
func NewEndpointHealthDataListResultPage(cur EndpointHealthDataListResult, getNextPage func(context.Context, EndpointHealthDataListResult) (EndpointHealthDataListResult, error)) EndpointHealthDataListResultPage {
	return original.NewEndpointHealthDataListResultPage(cur, getNextPage)
}
func NewEventHubConsumerGroupsListResultIterator(page EventHubConsumerGroupsListResultPage) EventHubConsumerGroupsListResultIterator {
	return original.NewEventHubConsumerGroupsListResultIterator(page)
}
func NewEventHubConsumerGroupsListResultPage(cur EventHubConsumerGroupsListResult, getNextPage func(context.Context, EventHubConsumerGroupsListResult) (EventHubConsumerGroupsListResult, error)) EventHubConsumerGroupsListResultPage {
	return original.NewEventHubConsumerGroupsListResultPage(cur, getNextPage)
}
func NewIotHubClient(subscriptionID string) IotHubClient {
	return original.NewIotHubClient(subscriptionID)
}
func NewIotHubClientWithBaseURI(baseURI string, subscriptionID string) IotHubClient {
	return original.NewIotHubClientWithBaseURI(baseURI, subscriptionID)
}
func NewIotHubClient(subscriptionID string) IotHubClient {
	return original.NewIotHubClient(subscriptionID)
}
func NewIotHubClientWithBaseURI(baseURI string, subscriptionID string) IotHubClient {
	return original.NewIotHubClientWithBaseURI(baseURI, subscriptionID)
}
func NewIotHubDescriptionListResultIterator(page IotHubDescriptionListResultPage) IotHubDescriptionListResultIterator {
	return original.NewIotHubDescriptionListResultIterator(page)
}
func NewIotHubDescriptionListResultPage(cur IotHubDescriptionListResult, getNextPage func(context.Context, IotHubDescriptionListResult) (IotHubDescriptionListResult, error)) IotHubDescriptionListResultPage {
	return original.NewIotHubDescriptionListResultPage(cur, getNextPage)
}
func NewIotHubQuotaMetricInfoListResultIterator(page IotHubQuotaMetricInfoListResultPage) IotHubQuotaMetricInfoListResultIterator {
	return original.NewIotHubQuotaMetricInfoListResultIterator(page)
}
func NewIotHubQuotaMetricInfoListResultPage(cur IotHubQuotaMetricInfoListResult, getNextPage func(context.Context, IotHubQuotaMetricInfoListResult) (IotHubQuotaMetricInfoListResult, error)) IotHubQuotaMetricInfoListResultPage {
	return original.NewIotHubQuotaMetricInfoListResultPage(cur, getNextPage)
}
func NewIotHubResourceClient(subscriptionID string) IotHubResourceClient {
	return original.NewIotHubResourceClient(subscriptionID)
}
func NewIotHubResourceClientWithBaseURI(baseURI string, subscriptionID string) IotHubResourceClient {
	return original.NewIotHubResourceClientWithBaseURI(baseURI, subscriptionID)
}
func NewIotHubSkuDescriptionListResultIterator(page IotHubSkuDescriptionListResultPage) IotHubSkuDescriptionListResultIterator {
	return original.NewIotHubSkuDescriptionListResultIterator(page)
}
func NewIotHubSkuDescriptionListResultPage(cur IotHubSkuDescriptionListResult, getNextPage func(context.Context, IotHubSkuDescriptionListResult) (IotHubSkuDescriptionListResult, error)) IotHubSkuDescriptionListResultPage {
	return original.NewIotHubSkuDescriptionListResultPage(cur, getNextPage)
}
func NewJobResponseListResultIterator(page JobResponseListResultPage) JobResponseListResultIterator {
	return original.NewJobResponseListResultIterator(page)
}
func NewJobResponseListResultPage(cur JobResponseListResult, getNextPage func(context.Context, JobResponseListResult) (JobResponseListResult, error)) JobResponseListResultPage {
	return original.NewJobResponseListResultPage(cur, getNextPage)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClient(subscriptionID)
}
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClient(subscriptionID)
}
func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceProviderCommonClient(subscriptionID string) ResourceProviderCommonClient {
	return original.NewResourceProviderCommonClient(subscriptionID)
}
func NewResourceProviderCommonClientWithBaseURI(baseURI string, subscriptionID string) ResourceProviderCommonClient {
	return original.NewResourceProviderCommonClientWithBaseURI(baseURI, subscriptionID)
}
func NewSharedAccessSignatureAuthorizationRuleListResultIterator(page SharedAccessSignatureAuthorizationRuleListResultPage) SharedAccessSignatureAuthorizationRuleListResultIterator {
	return original.NewSharedAccessSignatureAuthorizationRuleListResultIterator(page)
}
func NewSharedAccessSignatureAuthorizationRuleListResultPage(cur SharedAccessSignatureAuthorizationRuleListResult, getNextPage func(context.Context, SharedAccessSignatureAuthorizationRuleListResult) (SharedAccessSignatureAuthorizationRuleListResult, error)) SharedAccessSignatureAuthorizationRuleListResultPage {
	return original.NewSharedAccessSignatureAuthorizationRuleListResultPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccessRightsValues() []AccessRights {
	return original.PossibleAccessRightsValues()
}
func PossibleAuthenticationTypeValues() []AuthenticationType {
	return original.PossibleAuthenticationTypeValues()
}
func PossibleCapabilitiesValues() []Capabilities {
	return original.PossibleCapabilitiesValues()
}
func PossibleEncodingValues() []Encoding {
	return original.PossibleEncodingValues()
}
func PossibleEndpointHealthStatusValues() []EndpointHealthStatus {
	return original.PossibleEndpointHealthStatusValues()
}
func PossibleIPFilterActionTypeValues() []IPFilterActionType {
	return original.PossibleIPFilterActionTypeValues()
}
func PossibleIotHubNameUnavailabilityReasonValues() []IotHubNameUnavailabilityReason {
	return original.PossibleIotHubNameUnavailabilityReasonValues()
}
func PossibleIotHubReplicaRoleTypeValues() []IotHubReplicaRoleType {
	return original.PossibleIotHubReplicaRoleTypeValues()
}
func PossibleIotHubScaleTypeValues() []IotHubScaleType {
	return original.PossibleIotHubScaleTypeValues()
}
func PossibleIotHubSkuTierValues() []IotHubSkuTier {
	return original.PossibleIotHubSkuTierValues()
}
func PossibleIotHubSkuValues() []IotHubSku {
	return original.PossibleIotHubSkuValues()
}
func PossibleJobStatusValues() []JobStatus {
	return original.PossibleJobStatusValues()
}
func PossibleJobTypeValues() []JobType {
	return original.PossibleJobTypeValues()
}
func PossiblePrivateLinkServiceConnectionStatusValues() []PrivateLinkServiceConnectionStatus {
	return original.PossiblePrivateLinkServiceConnectionStatusValues()
}
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return original.PossiblePublicNetworkAccessValues()
}
func PossibleRouteErrorSeverityValues() []RouteErrorSeverity {
	return original.PossibleRouteErrorSeverityValues()
}
func PossibleRoutingSourceValues() []RoutingSource {
	return original.PossibleRoutingSourceValues()
}
func PossibleTestResultStatusValues() []TestResultStatus {
	return original.PossibleTestResultStatusValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
