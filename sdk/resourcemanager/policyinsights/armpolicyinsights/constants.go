//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armpolicyinsights

const (
	moduleName    = "armpolicyinsights"
	moduleVersion = "v0.7.0"
)

// ComplianceState - The compliance state that should be set on the resource.
type ComplianceState string

const (
	// ComplianceStateCompliant - The resource is in compliance with the policy.
	ComplianceStateCompliant ComplianceState = "Compliant"
	// ComplianceStateNonCompliant - The resource is not in compliance with the policy.
	ComplianceStateNonCompliant ComplianceState = "NonCompliant"
	// ComplianceStateUnknown - The compliance state of the resource is not known.
	ComplianceStateUnknown ComplianceState = "Unknown"
)

// PossibleComplianceStateValues returns the possible values for the ComplianceState const type.
func PossibleComplianceStateValues() []ComplianceState {
	return []ComplianceState{
		ComplianceStateCompliant,
		ComplianceStateNonCompliant,
		ComplianceStateUnknown,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// FieldRestrictionResult - The type of restriction that is imposed on the field.
type FieldRestrictionResult string

const (
	// FieldRestrictionResultDeny - The field and/or values will be denied by policy.
	FieldRestrictionResultDeny FieldRestrictionResult = "Deny"
	// FieldRestrictionResultRemoved - The field will be removed by policy.
	FieldRestrictionResultRemoved FieldRestrictionResult = "Removed"
	// FieldRestrictionResultRequired - The field and/or values are required by policy.
	FieldRestrictionResultRequired FieldRestrictionResult = "Required"
)

// PossibleFieldRestrictionResultValues returns the possible values for the FieldRestrictionResult const type.
func PossibleFieldRestrictionResultValues() []FieldRestrictionResult {
	return []FieldRestrictionResult{
		FieldRestrictionResultDeny,
		FieldRestrictionResultRemoved,
		FieldRestrictionResultRequired,
	}
}

type PolicyEventsResourceType string

const (
	PolicyEventsResourceTypeDefault PolicyEventsResourceType = "default"
)

// PossiblePolicyEventsResourceTypeValues returns the possible values for the PolicyEventsResourceType const type.
func PossiblePolicyEventsResourceTypeValues() []PolicyEventsResourceType {
	return []PolicyEventsResourceType{
		PolicyEventsResourceTypeDefault,
	}
}

type PolicyStatesResource string

const (
	PolicyStatesResourceDefault PolicyStatesResource = "default"
	PolicyStatesResourceLatest  PolicyStatesResource = "latest"
)

// PossiblePolicyStatesResourceValues returns the possible values for the PolicyStatesResource const type.
func PossiblePolicyStatesResourceValues() []PolicyStatesResource {
	return []PolicyStatesResource{
		PolicyStatesResourceDefault,
		PolicyStatesResourceLatest,
	}
}

type PolicyStatesSummaryResourceType string

const (
	PolicyStatesSummaryResourceTypeLatest PolicyStatesSummaryResourceType = "latest"
)

// PossiblePolicyStatesSummaryResourceTypeValues returns the possible values for the PolicyStatesSummaryResourceType const type.
func PossiblePolicyStatesSummaryResourceTypeValues() []PolicyStatesSummaryResourceType {
	return []PolicyStatesSummaryResourceType{
		PolicyStatesSummaryResourceTypeLatest,
	}
}

type PolicyTrackedResourcesResourceType string

const (
	PolicyTrackedResourcesResourceTypeDefault PolicyTrackedResourcesResourceType = "default"
)

// PossiblePolicyTrackedResourcesResourceTypeValues returns the possible values for the PolicyTrackedResourcesResourceType const type.
func PossiblePolicyTrackedResourcesResourceTypeValues() []PolicyTrackedResourcesResourceType {
	return []PolicyTrackedResourcesResourceType{
		PolicyTrackedResourcesResourceTypeDefault,
	}
}

// ResourceDiscoveryMode - The way resources to remediate are discovered. Defaults to ExistingNonCompliant if not specified.
type ResourceDiscoveryMode string

const (
	// ResourceDiscoveryModeExistingNonCompliant - Remediate resources that are already known to be non-compliant.
	ResourceDiscoveryModeExistingNonCompliant ResourceDiscoveryMode = "ExistingNonCompliant"
	// ResourceDiscoveryModeReEvaluateCompliance - Re-evaluate the compliance state of resources and then remediate the resources
	// found to be non-compliant.
	ResourceDiscoveryModeReEvaluateCompliance ResourceDiscoveryMode = "ReEvaluateCompliance"
)

// PossibleResourceDiscoveryModeValues returns the possible values for the ResourceDiscoveryMode const type.
func PossibleResourceDiscoveryModeValues() []ResourceDiscoveryMode {
	return []ResourceDiscoveryMode{
		ResourceDiscoveryModeExistingNonCompliant,
		ResourceDiscoveryModeReEvaluateCompliance,
	}
}
