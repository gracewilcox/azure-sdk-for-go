package storagecache

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// FirmwareStatusType enumerates the values for firmware status type.
type FirmwareStatusType string

const (
	// Available ...
	Available FirmwareStatusType = "available"
	// Unavailable ...
	Unavailable FirmwareStatusType = "unavailable"
)

// PossibleFirmwareStatusTypeValues returns an array of possible values for the FirmwareStatusType const type.
func PossibleFirmwareStatusTypeValues() []FirmwareStatusType {
	return []FirmwareStatusType{Available, Unavailable}
}

// HealthStateType enumerates the values for health state type.
type HealthStateType string

const (
	// Degraded ...
	Degraded HealthStateType = "Degraded"
	// Down ...
	Down HealthStateType = "Down"
	// Flushing ...
	Flushing HealthStateType = "Flushing"
	// Healthy ...
	Healthy HealthStateType = "Healthy"
	// Stopped ...
	Stopped HealthStateType = "Stopped"
	// Stopping ...
	Stopping HealthStateType = "Stopping"
	// Transitioning ...
	Transitioning HealthStateType = "Transitioning"
	// Unknown ...
	Unknown HealthStateType = "Unknown"
	// Upgrading ...
	Upgrading HealthStateType = "Upgrading"
)

// PossibleHealthStateTypeValues returns an array of possible values for the HealthStateType const type.
func PossibleHealthStateTypeValues() []HealthStateType {
	return []HealthStateType{Degraded, Down, Flushing, Healthy, Stopped, Stopping, Transitioning, Unknown, Upgrading}
}

// ProvisioningStateType enumerates the values for provisioning state type.
type ProvisioningStateType string

const (
	// Cancelled ...
	Cancelled ProvisioningStateType = "Cancelled"
	// Creating ...
	Creating ProvisioningStateType = "Creating"
	// Deleting ...
	Deleting ProvisioningStateType = "Deleting"
	// Failed ...
	Failed ProvisioningStateType = "Failed"
	// Succeeded ...
	Succeeded ProvisioningStateType = "Succeeded"
	// Updating ...
	Updating ProvisioningStateType = "Updating"
)

// PossibleProvisioningStateTypeValues returns an array of possible values for the ProvisioningStateType const type.
func PossibleProvisioningStateTypeValues() []ProvisioningStateType {
	return []ProvisioningStateType{Cancelled, Creating, Deleting, Failed, Succeeded, Updating}
}

// ReasonCode enumerates the values for reason code.
type ReasonCode string

const (
	// NotAvailableForSubscription ...
	NotAvailableForSubscription ReasonCode = "NotAvailableForSubscription"
	// QuotaID ...
	QuotaID ReasonCode = "QuotaId"
)

// PossibleReasonCodeValues returns an array of possible values for the ReasonCode const type.
func PossibleReasonCodeValues() []ReasonCode {
	return []ReasonCode{NotAvailableForSubscription, QuotaID}
}

// StorageTargetType enumerates the values for storage target type.
type StorageTargetType string

const (
	// StorageTargetTypeClfs ...
	StorageTargetTypeClfs StorageTargetType = "clfs"
	// StorageTargetTypeNfs3 ...
	StorageTargetTypeNfs3 StorageTargetType = "nfs3"
	// StorageTargetTypeUnknown ...
	StorageTargetTypeUnknown StorageTargetType = "unknown"
)

// PossibleStorageTargetTypeValues returns an array of possible values for the StorageTargetType const type.
func PossibleStorageTargetTypeValues() []StorageTargetType {
	return []StorageTargetType{StorageTargetTypeClfs, StorageTargetTypeNfs3, StorageTargetTypeUnknown}
}
