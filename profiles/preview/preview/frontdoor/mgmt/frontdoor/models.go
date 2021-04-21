// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package frontdoor

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/frontdoor/mgmt/2018-08-01-preview/frontdoor"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type Action = original.Action

const (
	Allow Action = original.Allow
	Block Action = original.Block
	Log   Action = original.Log
)

type Availability = original.Availability

const (
	Available   Availability = original.Available
	Unavailable Availability = original.Unavailable
)

type CertificateSource = original.CertificateSource

const (
	CertificateSourceAzureKeyVault CertificateSource = original.CertificateSourceAzureKeyVault
	CertificateSourceFrontDoor     CertificateSource = original.CertificateSourceFrontDoor
)

type CertificateType = original.CertificateType

const (
	Dedicated CertificateType = original.Dedicated
)

type CustomHTTPSProvisioningState = original.CustomHTTPSProvisioningState

const (
	Disabled  CustomHTTPSProvisioningState = original.Disabled
	Disabling CustomHTTPSProvisioningState = original.Disabling
	Enabled   CustomHTTPSProvisioningState = original.Enabled
	Enabling  CustomHTTPSProvisioningState = original.Enabling
	Failed    CustomHTTPSProvisioningState = original.Failed
)

type CustomHTTPSProvisioningSubstate = original.CustomHTTPSProvisioningSubstate

const (
	CertificateDeleted                            CustomHTTPSProvisioningSubstate = original.CertificateDeleted
	CertificateDeployed                           CustomHTTPSProvisioningSubstate = original.CertificateDeployed
	DeletingCertificate                           CustomHTTPSProvisioningSubstate = original.DeletingCertificate
	DeployingCertificate                          CustomHTTPSProvisioningSubstate = original.DeployingCertificate
	DomainControlValidationRequestApproved        CustomHTTPSProvisioningSubstate = original.DomainControlValidationRequestApproved
	DomainControlValidationRequestRejected        CustomHTTPSProvisioningSubstate = original.DomainControlValidationRequestRejected
	DomainControlValidationRequestTimedOut        CustomHTTPSProvisioningSubstate = original.DomainControlValidationRequestTimedOut
	IssuingCertificate                            CustomHTTPSProvisioningSubstate = original.IssuingCertificate
	PendingDomainControlValidationREquestApproval CustomHTTPSProvisioningSubstate = original.PendingDomainControlValidationREquestApproval
	SubmittingDomainControlValidationRequest      CustomHTTPSProvisioningSubstate = original.SubmittingDomainControlValidationRequest
)

type DynamicCompressionEnabled = original.DynamicCompressionEnabled

const (
	DynamicCompressionEnabledDisabled DynamicCompressionEnabled = original.DynamicCompressionEnabledDisabled
	DynamicCompressionEnabledEnabled  DynamicCompressionEnabled = original.DynamicCompressionEnabledEnabled
)

type EnabledState = original.EnabledState

const (
	EnabledStateDisabled EnabledState = original.EnabledStateDisabled
	EnabledStateEnabled  EnabledState = original.EnabledStateEnabled
)

type EnabledStateEnum = original.EnabledStateEnum

const (
	EnabledStateEnumDisabled EnabledStateEnum = original.EnabledStateEnumDisabled
	EnabledStateEnumEnabled  EnabledStateEnum = original.EnabledStateEnumEnabled
)

type ForwardingProtocol = original.ForwardingProtocol

const (
	HTTPOnly     ForwardingProtocol = original.HTTPOnly
	HTTPSOnly    ForwardingProtocol = original.HTTPSOnly
	MatchRequest ForwardingProtocol = original.MatchRequest
)

type MatchCondition = original.MatchCondition

const (
	PostArgs      MatchCondition = original.PostArgs
	QueryString   MatchCondition = original.QueryString
	RemoteAddr    MatchCondition = original.RemoteAddr
	RequestBody   MatchCondition = original.RequestBody
	RequestHeader MatchCondition = original.RequestHeader
	RequestMethod MatchCondition = original.RequestMethod
	RequestURI    MatchCondition = original.RequestURI
)

type Mode = original.Mode

const (
	Detection  Mode = original.Detection
	Prevention Mode = original.Prevention
)

type NetworkOperationStatus = original.NetworkOperationStatus

const (
	NetworkOperationStatusFailed     NetworkOperationStatus = original.NetworkOperationStatusFailed
	NetworkOperationStatusInProgress NetworkOperationStatus = original.NetworkOperationStatusInProgress
	NetworkOperationStatusSucceeded  NetworkOperationStatus = original.NetworkOperationStatusSucceeded
)

type Operator = original.Operator

const (
	Any                Operator = original.Any
	BeginsWith         Operator = original.BeginsWith
	Contains           Operator = original.Contains
	EndsWith           Operator = original.EndsWith
	Equal              Operator = original.Equal
	GeoMatch           Operator = original.GeoMatch
	GreaterThan        Operator = original.GreaterThan
	GreaterThanOrEqual Operator = original.GreaterThanOrEqual
	IPMatch            Operator = original.IPMatch
	LessThan           Operator = original.LessThan
	LessThanOrEqual    Operator = original.LessThanOrEqual
)

type Protocol = original.Protocol

const (
	HTTP  Protocol = original.HTTP
	HTTPS Protocol = original.HTTPS
)

type Query = original.Query

const (
	StripAll  Query = original.StripAll
	StripNone Query = original.StripNone
)

type ResourceState = original.ResourceState

const (
	ResourceStateCreating  ResourceState = original.ResourceStateCreating
	ResourceStateDeleting  ResourceState = original.ResourceStateDeleting
	ResourceStateDisabled  ResourceState = original.ResourceStateDisabled
	ResourceStateDisabling ResourceState = original.ResourceStateDisabling
	ResourceStateEnabled   ResourceState = original.ResourceStateEnabled
	ResourceStateEnabling  ResourceState = original.ResourceStateEnabling
)

type ResourceType = original.ResourceType

const (
	MicrosoftNetworkfrontDoors                  ResourceType = original.MicrosoftNetworkfrontDoors
	MicrosoftNetworkfrontDoorsfrontendEndpoints ResourceType = original.MicrosoftNetworkfrontDoorsfrontendEndpoints
)

type RuleGroupOverride = original.RuleGroupOverride

const (
	SQLInjection RuleGroupOverride = original.SQLInjection
	XSS          RuleGroupOverride = original.XSS
)

type RuleSetType = original.RuleSetType

const (
	RuleSetTypeAzureManagedRuleSet RuleSetType = original.RuleSetTypeAzureManagedRuleSet
	RuleSetTypeUnknown             RuleSetType = original.RuleSetTypeUnknown
)

type RuleType = original.RuleType

const (
	MatchRule     RuleType = original.MatchRule
	RateLimitRule RuleType = original.RateLimitRule
)

type SessionAffinityEnabledState = original.SessionAffinityEnabledState

const (
	SessionAffinityEnabledStateDisabled SessionAffinityEnabledState = original.SessionAffinityEnabledStateDisabled
	SessionAffinityEnabledStateEnabled  SessionAffinityEnabledState = original.SessionAffinityEnabledStateEnabled
)

type TLSProtocolType = original.TLSProtocolType

const (
	ServerNameIndication TLSProtocolType = original.ServerNameIndication
)

type Transform = original.Transform

const (
	HTMLEntityDecode Transform = original.HTMLEntityDecode
	Lowercase        Transform = original.Lowercase
	RemoveNulls      Transform = original.RemoveNulls
	Trim             Transform = original.Trim
	Uppercase        Transform = original.Uppercase
	URLDecode        Transform = original.URLDecode
	URLEncode        Transform = original.URLEncode
)

type WebApplicationFirewallPolicy = original.WebApplicationFirewallPolicy

const (
	WebApplicationFirewallPolicyCreating  WebApplicationFirewallPolicy = original.WebApplicationFirewallPolicyCreating
	WebApplicationFirewallPolicyDeleting  WebApplicationFirewallPolicy = original.WebApplicationFirewallPolicyDeleting
	WebApplicationFirewallPolicyDisabled  WebApplicationFirewallPolicy = original.WebApplicationFirewallPolicyDisabled
	WebApplicationFirewallPolicyDisabling WebApplicationFirewallPolicy = original.WebApplicationFirewallPolicyDisabling
	WebApplicationFirewallPolicyEnabled   WebApplicationFirewallPolicy = original.WebApplicationFirewallPolicyEnabled
	WebApplicationFirewallPolicyEnabling  WebApplicationFirewallPolicy = original.WebApplicationFirewallPolicyEnabling
)

type AzureAsyncOperationResult = original.AzureAsyncOperationResult
type AzureManagedOverrideRuleGroup = original.AzureManagedOverrideRuleGroup
type AzureManagedRuleSet = original.AzureManagedRuleSet
type Backend = original.Backend
type BackendPool = original.BackendPool
type BackendPoolListResult = original.BackendPoolListResult
type BackendPoolProperties = original.BackendPoolProperties
type BackendPoolUpdateParameters = original.BackendPoolUpdateParameters
type BaseClient = original.BaseClient
type BasicManagedRuleSet = original.BasicManagedRuleSet
type CacheConfiguration = original.CacheConfiguration
type CertificateSourceParameters = original.CertificateSourceParameters
type CheckNameAvailabilityInput = original.CheckNameAvailabilityInput
type CheckNameAvailabilityOutput = original.CheckNameAvailabilityOutput
type CustomHTTPSConfiguration = original.CustomHTTPSConfiguration
type CustomRule = original.CustomRule
type CustomRules = original.CustomRules
type EndpointsClient = original.EndpointsClient
type EndpointsPurgeContentFuture = original.EndpointsPurgeContentFuture
type Error = original.Error
type ErrorDetails = original.ErrorDetails
type ErrorResponse = original.ErrorResponse
type FrontDoor = original.FrontDoor
type FrontDoorsClient = original.FrontDoorsClient
type FrontDoorsCreateOrUpdateFutureType = original.FrontDoorsCreateOrUpdateFutureType
type FrontDoorsDeleteFutureType = original.FrontDoorsDeleteFutureType
type FrontendEndpoint = original.FrontendEndpoint
type FrontendEndpointProperties = original.FrontendEndpointProperties
type FrontendEndpointUpdateParameters = original.FrontendEndpointUpdateParameters
type FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink = original.FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink
type FrontendEndpointsClient = original.FrontendEndpointsClient
type FrontendEndpointsDisableHTTPSFuture = original.FrontendEndpointsDisableHTTPSFuture
type FrontendEndpointsEnableHTTPSFuture = original.FrontendEndpointsEnableHTTPSFuture
type FrontendEndpointsListResult = original.FrontendEndpointsListResult
type FrontendEndpointsListResultIterator = original.FrontendEndpointsListResultIterator
type FrontendEndpointsListResultPage = original.FrontendEndpointsListResultPage
type HealthProbeSettingsListResult = original.HealthProbeSettingsListResult
type HealthProbeSettingsModel = original.HealthProbeSettingsModel
type HealthProbeSettingsProperties = original.HealthProbeSettingsProperties
type HealthProbeSettingsUpdateParameters = original.HealthProbeSettingsUpdateParameters
type KeyVaultCertificateSourceParameters = original.KeyVaultCertificateSourceParameters
type KeyVaultCertificateSourceParametersVault = original.KeyVaultCertificateSourceParametersVault
type ListResult = original.ListResult
type ListResultIterator = original.ListResultIterator
type ListResultPage = original.ListResultPage
type LoadBalancingSettingsListResult = original.LoadBalancingSettingsListResult
type LoadBalancingSettingsModel = original.LoadBalancingSettingsModel
type LoadBalancingSettingsProperties = original.LoadBalancingSettingsProperties
type LoadBalancingSettingsUpdateParameters = original.LoadBalancingSettingsUpdateParameters
type ManagedRuleSet = original.ManagedRuleSet
type ManagedRuleSets = original.ManagedRuleSets
type MatchCondition1 = original.MatchCondition1
type PoliciesClient = original.PoliciesClient
type PoliciesDeleteFuture = original.PoliciesDeleteFuture
type PolicySettings = original.PolicySettings
type Properties = original.Properties
type PurgeParameters = original.PurgeParameters
type Resource = original.Resource
type RoutingRule = original.RoutingRule
type RoutingRuleListResult = original.RoutingRuleListResult
type RoutingRuleProperties = original.RoutingRuleProperties
type RoutingRuleUpdateParameters = original.RoutingRuleUpdateParameters
type SubResource = original.SubResource
type TagsObject = original.TagsObject
type UpdateParameters = original.UpdateParameters
type ValidateCustomDomainInput = original.ValidateCustomDomainInput
type ValidateCustomDomainOutput = original.ValidateCustomDomainOutput
type WebApplicationFirewallPolicy1 = original.WebApplicationFirewallPolicy1
type WebApplicationFirewallPolicyListResult = original.WebApplicationFirewallPolicyListResult
type WebApplicationFirewallPolicyListResultIterator = original.WebApplicationFirewallPolicyListResultIterator
type WebApplicationFirewallPolicyListResultPage = original.WebApplicationFirewallPolicyListResultPage
type WebApplicationFirewallPolicyPropertiesFormat = original.WebApplicationFirewallPolicyPropertiesFormat

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewEndpointsClient(subscriptionID string) EndpointsClient {
	return original.NewEndpointsClient(subscriptionID)
}
func NewEndpointsClientWithBaseURI(baseURI string, subscriptionID string) EndpointsClient {
	return original.NewEndpointsClientWithBaseURI(baseURI, subscriptionID)
}
func NewFrontDoorsClient(subscriptionID string) FrontDoorsClient {
	return original.NewFrontDoorsClient(subscriptionID)
}
func NewFrontDoorsClientWithBaseURI(baseURI string, subscriptionID string) FrontDoorsClient {
	return original.NewFrontDoorsClientWithBaseURI(baseURI, subscriptionID)
}
func NewFrontendEndpointsClient(subscriptionID string) FrontendEndpointsClient {
	return original.NewFrontendEndpointsClient(subscriptionID)
}
func NewFrontendEndpointsClientWithBaseURI(baseURI string, subscriptionID string) FrontendEndpointsClient {
	return original.NewFrontendEndpointsClientWithBaseURI(baseURI, subscriptionID)
}
func NewFrontendEndpointsListResultIterator(page FrontendEndpointsListResultPage) FrontendEndpointsListResultIterator {
	return original.NewFrontendEndpointsListResultIterator(page)
}
func NewFrontendEndpointsListResultPage(cur FrontendEndpointsListResult, getNextPage func(context.Context, FrontendEndpointsListResult) (FrontendEndpointsListResult, error)) FrontendEndpointsListResultPage {
	return original.NewFrontendEndpointsListResultPage(cur, getNextPage)
}
func NewListResultIterator(page ListResultPage) ListResultIterator {
	return original.NewListResultIterator(page)
}
func NewListResultPage(cur ListResult, getNextPage func(context.Context, ListResult) (ListResult, error)) ListResultPage {
	return original.NewListResultPage(cur, getNextPage)
}
func NewPoliciesClient(subscriptionID string) PoliciesClient {
	return original.NewPoliciesClient(subscriptionID)
}
func NewPoliciesClientWithBaseURI(baseURI string, subscriptionID string) PoliciesClient {
	return original.NewPoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWebApplicationFirewallPolicyListResultIterator(page WebApplicationFirewallPolicyListResultPage) WebApplicationFirewallPolicyListResultIterator {
	return original.NewWebApplicationFirewallPolicyListResultIterator(page)
}
func NewWebApplicationFirewallPolicyListResultPage(cur WebApplicationFirewallPolicyListResult, getNextPage func(context.Context, WebApplicationFirewallPolicyListResult) (WebApplicationFirewallPolicyListResult, error)) WebApplicationFirewallPolicyListResultPage {
	return original.NewWebApplicationFirewallPolicyListResultPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleActionValues() []Action {
	return original.PossibleActionValues()
}
func PossibleAvailabilityValues() []Availability {
	return original.PossibleAvailabilityValues()
}
func PossibleCertificateSourceValues() []CertificateSource {
	return original.PossibleCertificateSourceValues()
}
func PossibleCertificateTypeValues() []CertificateType {
	return original.PossibleCertificateTypeValues()
}
func PossibleCustomHTTPSProvisioningStateValues() []CustomHTTPSProvisioningState {
	return original.PossibleCustomHTTPSProvisioningStateValues()
}
func PossibleCustomHTTPSProvisioningSubstateValues() []CustomHTTPSProvisioningSubstate {
	return original.PossibleCustomHTTPSProvisioningSubstateValues()
}
func PossibleDynamicCompressionEnabledValues() []DynamicCompressionEnabled {
	return original.PossibleDynamicCompressionEnabledValues()
}
func PossibleEnabledStateEnumValues() []EnabledStateEnum {
	return original.PossibleEnabledStateEnumValues()
}
func PossibleEnabledStateValues() []EnabledState {
	return original.PossibleEnabledStateValues()
}
func PossibleForwardingProtocolValues() []ForwardingProtocol {
	return original.PossibleForwardingProtocolValues()
}
func PossibleMatchConditionValues() []MatchCondition {
	return original.PossibleMatchConditionValues()
}
func PossibleModeValues() []Mode {
	return original.PossibleModeValues()
}
func PossibleNetworkOperationStatusValues() []NetworkOperationStatus {
	return original.PossibleNetworkOperationStatusValues()
}
func PossibleOperatorValues() []Operator {
	return original.PossibleOperatorValues()
}
func PossibleProtocolValues() []Protocol {
	return original.PossibleProtocolValues()
}
func PossibleQueryValues() []Query {
	return original.PossibleQueryValues()
}
func PossibleResourceStateValues() []ResourceState {
	return original.PossibleResourceStateValues()
}
func PossibleResourceTypeValues() []ResourceType {
	return original.PossibleResourceTypeValues()
}
func PossibleRuleGroupOverrideValues() []RuleGroupOverride {
	return original.PossibleRuleGroupOverrideValues()
}
func PossibleRuleSetTypeValues() []RuleSetType {
	return original.PossibleRuleSetTypeValues()
}
func PossibleRuleTypeValues() []RuleType {
	return original.PossibleRuleTypeValues()
}
func PossibleSessionAffinityEnabledStateValues() []SessionAffinityEnabledState {
	return original.PossibleSessionAffinityEnabledStateValues()
}
func PossibleTLSProtocolTypeValues() []TLSProtocolType {
	return original.PossibleTLSProtocolTypeValues()
}
func PossibleTransformValues() []Transform {
	return original.PossibleTransformValues()
}
func PossibleWebApplicationFirewallPolicyValues() []WebApplicationFirewallPolicy {
	return original.PossibleWebApplicationFirewallPolicyValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
