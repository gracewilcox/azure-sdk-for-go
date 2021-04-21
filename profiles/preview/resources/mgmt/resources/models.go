// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package resources

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-10-01/resources"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AliasPathAttributes = original.AliasPathAttributes

const (
	Modifiable AliasPathAttributes = original.Modifiable
	None       AliasPathAttributes = original.None
)

type AliasPathTokenType = original.AliasPathTokenType

const (
	Any          AliasPathTokenType = original.Any
	Array        AliasPathTokenType = original.Array
	Boolean      AliasPathTokenType = original.Boolean
	Integer      AliasPathTokenType = original.Integer
	NotSpecified AliasPathTokenType = original.NotSpecified
	Number       AliasPathTokenType = original.Number
	Object       AliasPathTokenType = original.Object
	String       AliasPathTokenType = original.String
)

type AliasPatternType = original.AliasPatternType

const (
	AliasPatternTypeExtract      AliasPatternType = original.AliasPatternTypeExtract
	AliasPatternTypeNotSpecified AliasPatternType = original.AliasPatternTypeNotSpecified
)

type AliasType = original.AliasType

const (
	AliasTypeMask         AliasType = original.AliasTypeMask
	AliasTypeNotSpecified AliasType = original.AliasTypeNotSpecified
	AliasTypePlainText    AliasType = original.AliasTypePlainText
)

type ChangeType = original.ChangeType

const (
	Create   ChangeType = original.Create
	Delete   ChangeType = original.Delete
	Deploy   ChangeType = original.Deploy
	Ignore   ChangeType = original.Ignore
	Modify   ChangeType = original.Modify
	NoChange ChangeType = original.NoChange
)

type DeploymentMode = original.DeploymentMode

const (
	Complete    DeploymentMode = original.Complete
	Incremental DeploymentMode = original.Incremental
)

type ExpressionEvaluationOptionsScopeType = original.ExpressionEvaluationOptionsScopeType

const (
	ExpressionEvaluationOptionsScopeTypeInner        ExpressionEvaluationOptionsScopeType = original.ExpressionEvaluationOptionsScopeTypeInner
	ExpressionEvaluationOptionsScopeTypeNotSpecified ExpressionEvaluationOptionsScopeType = original.ExpressionEvaluationOptionsScopeTypeNotSpecified
	ExpressionEvaluationOptionsScopeTypeOuter        ExpressionEvaluationOptionsScopeType = original.ExpressionEvaluationOptionsScopeTypeOuter
)

type OnErrorDeploymentType = original.OnErrorDeploymentType

const (
	LastSuccessful     OnErrorDeploymentType = original.LastSuccessful
	SpecificDeployment OnErrorDeploymentType = original.SpecificDeployment
)

type PropertyChangeType = original.PropertyChangeType

const (
	PropertyChangeTypeArray  PropertyChangeType = original.PropertyChangeTypeArray
	PropertyChangeTypeCreate PropertyChangeType = original.PropertyChangeTypeCreate
	PropertyChangeTypeDelete PropertyChangeType = original.PropertyChangeTypeDelete
	PropertyChangeTypeModify PropertyChangeType = original.PropertyChangeTypeModify
)

type ProvisioningOperation = original.ProvisioningOperation

const (
	ProvisioningOperationAction                     ProvisioningOperation = original.ProvisioningOperationAction
	ProvisioningOperationAzureAsyncOperationWaiting ProvisioningOperation = original.ProvisioningOperationAzureAsyncOperationWaiting
	ProvisioningOperationCreate                     ProvisioningOperation = original.ProvisioningOperationCreate
	ProvisioningOperationDelete                     ProvisioningOperation = original.ProvisioningOperationDelete
	ProvisioningOperationDeploymentCleanup          ProvisioningOperation = original.ProvisioningOperationDeploymentCleanup
	ProvisioningOperationEvaluateDeploymentOutput   ProvisioningOperation = original.ProvisioningOperationEvaluateDeploymentOutput
	ProvisioningOperationNotSpecified               ProvisioningOperation = original.ProvisioningOperationNotSpecified
	ProvisioningOperationRead                       ProvisioningOperation = original.ProvisioningOperationRead
	ProvisioningOperationResourceCacheWaiting       ProvisioningOperation = original.ProvisioningOperationResourceCacheWaiting
	ProvisioningOperationWaiting                    ProvisioningOperation = original.ProvisioningOperationWaiting
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateAccepted     ProvisioningState = original.ProvisioningStateAccepted
	ProvisioningStateCanceled     ProvisioningState = original.ProvisioningStateCanceled
	ProvisioningStateCreated      ProvisioningState = original.ProvisioningStateCreated
	ProvisioningStateCreating     ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateDeleted      ProvisioningState = original.ProvisioningStateDeleted
	ProvisioningStateDeleting     ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed       ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateNotSpecified ProvisioningState = original.ProvisioningStateNotSpecified
	ProvisioningStateReady        ProvisioningState = original.ProvisioningStateReady
	ProvisioningStateRunning      ProvisioningState = original.ProvisioningStateRunning
	ProvisioningStateSucceeded    ProvisioningState = original.ProvisioningStateSucceeded
	ProvisioningStateUpdating     ProvisioningState = original.ProvisioningStateUpdating
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	ResourceIdentityTypeNone                       ResourceIdentityType = original.ResourceIdentityTypeNone
	ResourceIdentityTypeSystemAssigned             ResourceIdentityType = original.ResourceIdentityTypeSystemAssigned
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = original.ResourceIdentityTypeSystemAssignedUserAssigned
	ResourceIdentityTypeUserAssigned               ResourceIdentityType = original.ResourceIdentityTypeUserAssigned
)

type TagsPatchOperation = original.TagsPatchOperation

const (
	TagsPatchOperationDelete  TagsPatchOperation = original.TagsPatchOperationDelete
	TagsPatchOperationMerge   TagsPatchOperation = original.TagsPatchOperationMerge
	TagsPatchOperationReplace TagsPatchOperation = original.TagsPatchOperationReplace
)

type WhatIfResultFormat = original.WhatIfResultFormat

const (
	FullResourcePayloads WhatIfResultFormat = original.FullResourcePayloads
	ResourceIDOnly       WhatIfResultFormat = original.ResourceIDOnly
)

type APIProfile = original.APIProfile
type Alias = original.Alias
type AliasPath = original.AliasPath
type AliasPathMetadata = original.AliasPathMetadata
type AliasPattern = original.AliasPattern
type BaseClient = original.BaseClient
type BasicDependency = original.BasicDependency
type Client = original.Client
type CloudError = original.CloudError
type CreateOrUpdateByIDFuture = original.CreateOrUpdateByIDFuture
type CreateOrUpdateFuture = original.CreateOrUpdateFuture
type DebugSetting = original.DebugSetting
type DeleteByIDFuture = original.DeleteByIDFuture
type DeleteFuture = original.DeleteFuture
type Dependency = original.Dependency
type Deployment = original.Deployment
type DeploymentExportResult = original.DeploymentExportResult
type DeploymentExtended = original.DeploymentExtended
type DeploymentExtendedFilter = original.DeploymentExtendedFilter
type DeploymentListResult = original.DeploymentListResult
type DeploymentListResultIterator = original.DeploymentListResultIterator
type DeploymentListResultPage = original.DeploymentListResultPage
type DeploymentOperation = original.DeploymentOperation
type DeploymentOperationProperties = original.DeploymentOperationProperties
type DeploymentOperationsClient = original.DeploymentOperationsClient
type DeploymentOperationsListResult = original.DeploymentOperationsListResult
type DeploymentOperationsListResultIterator = original.DeploymentOperationsListResultIterator
type DeploymentOperationsListResultPage = original.DeploymentOperationsListResultPage
type DeploymentProperties = original.DeploymentProperties
type DeploymentPropertiesExtended = original.DeploymentPropertiesExtended
type DeploymentValidateResult = original.DeploymentValidateResult
type DeploymentWhatIf = original.DeploymentWhatIf
type DeploymentWhatIfProperties = original.DeploymentWhatIfProperties
type DeploymentWhatIfSettings = original.DeploymentWhatIfSettings
type DeploymentsClient = original.DeploymentsClient
type DeploymentsCreateOrUpdateAtManagementGroupScopeFuture = original.DeploymentsCreateOrUpdateAtManagementGroupScopeFuture
type DeploymentsCreateOrUpdateAtScopeFuture = original.DeploymentsCreateOrUpdateAtScopeFuture
type DeploymentsCreateOrUpdateAtSubscriptionScopeFuture = original.DeploymentsCreateOrUpdateAtSubscriptionScopeFuture
type DeploymentsCreateOrUpdateAtTenantScopeFuture = original.DeploymentsCreateOrUpdateAtTenantScopeFuture
type DeploymentsCreateOrUpdateFuture = original.DeploymentsCreateOrUpdateFuture
type DeploymentsDeleteAtManagementGroupScopeFuture = original.DeploymentsDeleteAtManagementGroupScopeFuture
type DeploymentsDeleteAtScopeFuture = original.DeploymentsDeleteAtScopeFuture
type DeploymentsDeleteAtSubscriptionScopeFuture = original.DeploymentsDeleteAtSubscriptionScopeFuture
type DeploymentsDeleteAtTenantScopeFuture = original.DeploymentsDeleteAtTenantScopeFuture
type DeploymentsDeleteFuture = original.DeploymentsDeleteFuture
type DeploymentsValidateAtManagementGroupScopeFuture = original.DeploymentsValidateAtManagementGroupScopeFuture
type DeploymentsValidateAtScopeFuture = original.DeploymentsValidateAtScopeFuture
type DeploymentsValidateAtSubscriptionScopeFuture = original.DeploymentsValidateAtSubscriptionScopeFuture
type DeploymentsValidateAtTenantScopeFuture = original.DeploymentsValidateAtTenantScopeFuture
type DeploymentsValidateFuture = original.DeploymentsValidateFuture
type DeploymentsWhatIfAtManagementGroupScopeFuture = original.DeploymentsWhatIfAtManagementGroupScopeFuture
type DeploymentsWhatIfAtSubscriptionScopeFuture = original.DeploymentsWhatIfAtSubscriptionScopeFuture
type DeploymentsWhatIfAtTenantScopeFuture = original.DeploymentsWhatIfAtTenantScopeFuture
type DeploymentsWhatIfFuture = original.DeploymentsWhatIfFuture
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorResponse = original.ErrorResponse
type ExportTemplateRequest = original.ExportTemplateRequest
type ExpressionEvaluationOptions = original.ExpressionEvaluationOptions
type GenericResource = original.GenericResource
type GenericResourceExpanded = original.GenericResourceExpanded
type GenericResourceFilter = original.GenericResourceFilter
type Group = original.Group
type GroupExportResult = original.GroupExportResult
type GroupFilter = original.GroupFilter
type GroupListResult = original.GroupListResult
type GroupListResultIterator = original.GroupListResultIterator
type GroupListResultPage = original.GroupListResultPage
type GroupPatchable = original.GroupPatchable
type GroupProperties = original.GroupProperties
type GroupsClient = original.GroupsClient
type GroupsDeleteFuture = original.GroupsDeleteFuture
type GroupsExportTemplateFuture = original.GroupsExportTemplateFuture
type HTTPMessage = original.HTTPMessage
type Identity = original.Identity
type IdentityUserAssignedIdentitiesValue = original.IdentityUserAssignedIdentitiesValue
type ListResult = original.ListResult
type ListResultIterator = original.ListResultIterator
type ListResultPage = original.ListResultPage
type MoveInfo = original.MoveInfo
type MoveResourcesFuture = original.MoveResourcesFuture
type OnErrorDeployment = original.OnErrorDeployment
type OnErrorDeploymentExtended = original.OnErrorDeploymentExtended
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type ParametersLink = original.ParametersLink
type Plan = original.Plan
type Provider = original.Provider
type ProviderExtendedLocation = original.ProviderExtendedLocation
type ProviderListResult = original.ProviderListResult
type ProviderListResultIterator = original.ProviderListResultIterator
type ProviderListResultPage = original.ProviderListResultPage
type ProviderOperationDisplayProperties = original.ProviderOperationDisplayProperties
type ProviderResourceType = original.ProviderResourceType
type ProviderResourceTypeListResult = original.ProviderResourceTypeListResult
type ProviderResourceTypesClient = original.ProviderResourceTypesClient
type ProvidersClient = original.ProvidersClient
type Reference = original.Reference
type Resource = original.Resource
type ScopedDeployment = original.ScopedDeployment
type ScopedDeploymentWhatIf = original.ScopedDeploymentWhatIf
type Sku = original.Sku
type StatusMessage = original.StatusMessage
type SubResource = original.SubResource
type TagCount = original.TagCount
type TagDetails = original.TagDetails
type TagValue = original.TagValue
type Tags = original.Tags
type TagsClient = original.TagsClient
type TagsListResult = original.TagsListResult
type TagsListResultIterator = original.TagsListResultIterator
type TagsListResultPage = original.TagsListResultPage
type TagsPatchResource = original.TagsPatchResource
type TagsResource = original.TagsResource
type TargetResource = original.TargetResource
type TemplateHashResult = original.TemplateHashResult
type TemplateLink = original.TemplateLink
type UpdateByIDFuture = original.UpdateByIDFuture
type UpdateFuture = original.UpdateFuture
type ValidateMoveResourcesFuture = original.ValidateMoveResourcesFuture
type WhatIfChange = original.WhatIfChange
type WhatIfOperationProperties = original.WhatIfOperationProperties
type WhatIfOperationResult = original.WhatIfOperationResult
type WhatIfPropertyChange = original.WhatIfPropertyChange

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewClient(subscriptionID string) Client {
	return original.NewClient(subscriptionID)
}
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return original.NewClientWithBaseURI(baseURI, subscriptionID)
}
func NewDeploymentListResultIterator(page DeploymentListResultPage) DeploymentListResultIterator {
	return original.NewDeploymentListResultIterator(page)
}
func NewDeploymentListResultPage(cur DeploymentListResult, getNextPage func(context.Context, DeploymentListResult) (DeploymentListResult, error)) DeploymentListResultPage {
	return original.NewDeploymentListResultPage(cur, getNextPage)
}
func NewDeploymentOperationsClient(subscriptionID string) DeploymentOperationsClient {
	return original.NewDeploymentOperationsClient(subscriptionID)
}
func NewDeploymentOperationsClientWithBaseURI(baseURI string, subscriptionID string) DeploymentOperationsClient {
	return original.NewDeploymentOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDeploymentOperationsListResultIterator(page DeploymentOperationsListResultPage) DeploymentOperationsListResultIterator {
	return original.NewDeploymentOperationsListResultIterator(page)
}
func NewDeploymentOperationsListResultPage(cur DeploymentOperationsListResult, getNextPage func(context.Context, DeploymentOperationsListResult) (DeploymentOperationsListResult, error)) DeploymentOperationsListResultPage {
	return original.NewDeploymentOperationsListResultPage(cur, getNextPage)
}
func NewDeploymentsClient(subscriptionID string) DeploymentsClient {
	return original.NewDeploymentsClient(subscriptionID)
}
func NewDeploymentsClientWithBaseURI(baseURI string, subscriptionID string) DeploymentsClient {
	return original.NewDeploymentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewGroupListResultIterator(page GroupListResultPage) GroupListResultIterator {
	return original.NewGroupListResultIterator(page)
}
func NewGroupListResultPage(cur GroupListResult, getNextPage func(context.Context, GroupListResult) (GroupListResult, error)) GroupListResultPage {
	return original.NewGroupListResultPage(cur, getNextPage)
}
func NewGroupsClient(subscriptionID string) GroupsClient {
	return original.NewGroupsClient(subscriptionID)
}
func NewGroupsClientWithBaseURI(baseURI string, subscriptionID string) GroupsClient {
	return original.NewGroupsClientWithBaseURI(baseURI, subscriptionID)
}
func NewListResultIterator(page ListResultPage) ListResultIterator {
	return original.NewListResultIterator(page)
}
func NewListResultPage(cur ListResult, getNextPage func(context.Context, ListResult) (ListResult, error)) ListResultPage {
	return original.NewListResultPage(cur, getNextPage)
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
func NewProviderListResultIterator(page ProviderListResultPage) ProviderListResultIterator {
	return original.NewProviderListResultIterator(page)
}
func NewProviderListResultPage(cur ProviderListResult, getNextPage func(context.Context, ProviderListResult) (ProviderListResult, error)) ProviderListResultPage {
	return original.NewProviderListResultPage(cur, getNextPage)
}
func NewProviderResourceTypesClient(subscriptionID string) ProviderResourceTypesClient {
	return original.NewProviderResourceTypesClient(subscriptionID)
}
func NewProviderResourceTypesClientWithBaseURI(baseURI string, subscriptionID string) ProviderResourceTypesClient {
	return original.NewProviderResourceTypesClientWithBaseURI(baseURI, subscriptionID)
}
func NewProvidersClient(subscriptionID string) ProvidersClient {
	return original.NewProvidersClient(subscriptionID)
}
func NewProvidersClientWithBaseURI(baseURI string, subscriptionID string) ProvidersClient {
	return original.NewProvidersClientWithBaseURI(baseURI, subscriptionID)
}
func NewTagsClient(subscriptionID string) TagsClient {
	return original.NewTagsClient(subscriptionID)
}
func NewTagsClientWithBaseURI(baseURI string, subscriptionID string) TagsClient {
	return original.NewTagsClientWithBaseURI(baseURI, subscriptionID)
}
func NewTagsListResultIterator(page TagsListResultPage) TagsListResultIterator {
	return original.NewTagsListResultIterator(page)
}
func NewTagsListResultPage(cur TagsListResult, getNextPage func(context.Context, TagsListResult) (TagsListResult, error)) TagsListResultPage {
	return original.NewTagsListResultPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAliasPathAttributesValues() []AliasPathAttributes {
	return original.PossibleAliasPathAttributesValues()
}
func PossibleAliasPathTokenTypeValues() []AliasPathTokenType {
	return original.PossibleAliasPathTokenTypeValues()
}
func PossibleAliasPatternTypeValues() []AliasPatternType {
	return original.PossibleAliasPatternTypeValues()
}
func PossibleAliasTypeValues() []AliasType {
	return original.PossibleAliasTypeValues()
}
func PossibleChangeTypeValues() []ChangeType {
	return original.PossibleChangeTypeValues()
}
func PossibleDeploymentModeValues() []DeploymentMode {
	return original.PossibleDeploymentModeValues()
}
func PossibleExpressionEvaluationOptionsScopeTypeValues() []ExpressionEvaluationOptionsScopeType {
	return original.PossibleExpressionEvaluationOptionsScopeTypeValues()
}
func PossibleOnErrorDeploymentTypeValues() []OnErrorDeploymentType {
	return original.PossibleOnErrorDeploymentTypeValues()
}
func PossiblePropertyChangeTypeValues() []PropertyChangeType {
	return original.PossiblePropertyChangeTypeValues()
}
func PossibleProvisioningOperationValues() []ProvisioningOperation {
	return original.PossibleProvisioningOperationValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleTagsPatchOperationValues() []TagsPatchOperation {
	return original.PossibleTagsPatchOperationValues()
}
func PossibleWhatIfResultFormatValues() []WhatIfResultFormat {
	return original.PossibleWhatIfResultFormatValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
