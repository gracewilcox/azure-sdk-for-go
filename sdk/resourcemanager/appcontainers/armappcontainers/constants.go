// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcontainers

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers"
	moduleVersion = "v3.1.0"
)

// AccessMode - Access mode for storage
type AccessMode string

const (
	AccessModeReadOnly  AccessMode = "ReadOnly"
	AccessModeReadWrite AccessMode = "ReadWrite"
)

// PossibleAccessModeValues returns the possible values for the AccessMode const type.
func PossibleAccessModeValues() []AccessMode {
	return []AccessMode{
		AccessModeReadOnly,
		AccessModeReadWrite,
	}
}

// Action - Allow or Deny rules to determine for incoming IP. Note: Rules can only consist of ALL Allow or ALL Deny
type Action string

const (
	ActionAllow Action = "Allow"
	ActionDeny  Action = "Deny"
)

// PossibleActionValues returns the possible values for the Action const type.
func PossibleActionValues() []Action {
	return []Action{
		ActionAllow,
		ActionDeny,
	}
}

// ActiveRevisionsMode - ActiveRevisionsMode controls how active revisions are handled for the Container app:Multiple: multiple
// revisions can be active.Single: Only one revision can be active at a time. Revision weights can
// not be used in this mode. If no value if provided, this is the default.
type ActiveRevisionsMode string

const (
	ActiveRevisionsModeMultiple ActiveRevisionsMode = "Multiple"
	ActiveRevisionsModeSingle   ActiveRevisionsMode = "Single"
)

// PossibleActiveRevisionsModeValues returns the possible values for the ActiveRevisionsMode const type.
func PossibleActiveRevisionsModeValues() []ActiveRevisionsMode {
	return []ActiveRevisionsMode{
		ActiveRevisionsModeMultiple,
		ActiveRevisionsModeSingle,
	}
}

// Affinity - Sticky Session Affinity
type Affinity string

const (
	AffinityNone   Affinity = "none"
	AffinitySticky Affinity = "sticky"
)

// PossibleAffinityValues returns the possible values for the Affinity const type.
func PossibleAffinityValues() []Affinity {
	return []Affinity{
		AffinityNone,
		AffinitySticky,
	}
}

// AppProtocol - Tells Dapr which protocol your application is using. Valid options are http and grpc. Default is http
type AppProtocol string

const (
	AppProtocolGrpc AppProtocol = "grpc"
	AppProtocolHTTP AppProtocol = "http"
)

// PossibleAppProtocolValues returns the possible values for the AppProtocol const type.
func PossibleAppProtocolValues() []AppProtocol {
	return []AppProtocol{
		AppProtocolGrpc,
		AppProtocolHTTP,
	}
}

// Applicability - indicates whether the profile is default for the location.
type Applicability string

const (
	ApplicabilityCustom          Applicability = "Custom"
	ApplicabilityLocationDefault Applicability = "LocationDefault"
)

// PossibleApplicabilityValues returns the possible values for the Applicability const type.
func PossibleApplicabilityValues() []Applicability {
	return []Applicability{
		ApplicabilityCustom,
		ApplicabilityLocationDefault,
	}
}

// BindingType - Custom Domain binding type.
type BindingType string

const (
	BindingTypeDisabled   BindingType = "Disabled"
	BindingTypeSniEnabled BindingType = "SniEnabled"
)

// PossibleBindingTypeValues returns the possible values for the BindingType const type.
func PossibleBindingTypeValues() []BindingType {
	return []BindingType{
		BindingTypeDisabled,
		BindingTypeSniEnabled,
	}
}

// CertificateProvisioningState - Provisioning state of the certificate.
type CertificateProvisioningState string

const (
	CertificateProvisioningStateCanceled     CertificateProvisioningState = "Canceled"
	CertificateProvisioningStateDeleteFailed CertificateProvisioningState = "DeleteFailed"
	CertificateProvisioningStateFailed       CertificateProvisioningState = "Failed"
	CertificateProvisioningStatePending      CertificateProvisioningState = "Pending"
	CertificateProvisioningStateSucceeded    CertificateProvisioningState = "Succeeded"
)

// PossibleCertificateProvisioningStateValues returns the possible values for the CertificateProvisioningState const type.
func PossibleCertificateProvisioningStateValues() []CertificateProvisioningState {
	return []CertificateProvisioningState{
		CertificateProvisioningStateCanceled,
		CertificateProvisioningStateDeleteFailed,
		CertificateProvisioningStateFailed,
		CertificateProvisioningStatePending,
		CertificateProvisioningStateSucceeded,
	}
}

// CheckNameAvailabilityReason - The reason why the given name is not available.
type CheckNameAvailabilityReason string

const (
	CheckNameAvailabilityReasonAlreadyExists CheckNameAvailabilityReason = "AlreadyExists"
	CheckNameAvailabilityReasonInvalid       CheckNameAvailabilityReason = "Invalid"
)

// PossibleCheckNameAvailabilityReasonValues returns the possible values for the CheckNameAvailabilityReason const type.
func PossibleCheckNameAvailabilityReasonValues() []CheckNameAvailabilityReason {
	return []CheckNameAvailabilityReason{
		CheckNameAvailabilityReasonAlreadyExists,
		CheckNameAvailabilityReasonInvalid,
	}
}

// ConnectedEnvironmentProvisioningState - Provisioning state of the Kubernetes Environment.
type ConnectedEnvironmentProvisioningState string

const (
	ConnectedEnvironmentProvisioningStateCanceled                      ConnectedEnvironmentProvisioningState = "Canceled"
	ConnectedEnvironmentProvisioningStateFailed                        ConnectedEnvironmentProvisioningState = "Failed"
	ConnectedEnvironmentProvisioningStateInfrastructureSetupComplete   ConnectedEnvironmentProvisioningState = "InfrastructureSetupComplete"
	ConnectedEnvironmentProvisioningStateInfrastructureSetupInProgress ConnectedEnvironmentProvisioningState = "InfrastructureSetupInProgress"
	ConnectedEnvironmentProvisioningStateInitializationInProgress      ConnectedEnvironmentProvisioningState = "InitializationInProgress"
	ConnectedEnvironmentProvisioningStateScheduledForDelete            ConnectedEnvironmentProvisioningState = "ScheduledForDelete"
	ConnectedEnvironmentProvisioningStateSucceeded                     ConnectedEnvironmentProvisioningState = "Succeeded"
	ConnectedEnvironmentProvisioningStateWaiting                       ConnectedEnvironmentProvisioningState = "Waiting"
)

// PossibleConnectedEnvironmentProvisioningStateValues returns the possible values for the ConnectedEnvironmentProvisioningState const type.
func PossibleConnectedEnvironmentProvisioningStateValues() []ConnectedEnvironmentProvisioningState {
	return []ConnectedEnvironmentProvisioningState{
		ConnectedEnvironmentProvisioningStateCanceled,
		ConnectedEnvironmentProvisioningStateFailed,
		ConnectedEnvironmentProvisioningStateInfrastructureSetupComplete,
		ConnectedEnvironmentProvisioningStateInfrastructureSetupInProgress,
		ConnectedEnvironmentProvisioningStateInitializationInProgress,
		ConnectedEnvironmentProvisioningStateScheduledForDelete,
		ConnectedEnvironmentProvisioningStateSucceeded,
		ConnectedEnvironmentProvisioningStateWaiting,
	}
}

// ContainerAppContainerRunningState - Current running state of the container
type ContainerAppContainerRunningState string

const (
	ContainerAppContainerRunningStateRunning    ContainerAppContainerRunningState = "Running"
	ContainerAppContainerRunningStateTerminated ContainerAppContainerRunningState = "Terminated"
	ContainerAppContainerRunningStateWaiting    ContainerAppContainerRunningState = "Waiting"
)

// PossibleContainerAppContainerRunningStateValues returns the possible values for the ContainerAppContainerRunningState const type.
func PossibleContainerAppContainerRunningStateValues() []ContainerAppContainerRunningState {
	return []ContainerAppContainerRunningState{
		ContainerAppContainerRunningStateRunning,
		ContainerAppContainerRunningStateTerminated,
		ContainerAppContainerRunningStateWaiting,
	}
}

// ContainerAppProvisioningState - Provisioning state of the Container App.
type ContainerAppProvisioningState string

const (
	ContainerAppProvisioningStateCanceled   ContainerAppProvisioningState = "Canceled"
	ContainerAppProvisioningStateDeleting   ContainerAppProvisioningState = "Deleting"
	ContainerAppProvisioningStateFailed     ContainerAppProvisioningState = "Failed"
	ContainerAppProvisioningStateInProgress ContainerAppProvisioningState = "InProgress"
	ContainerAppProvisioningStateSucceeded  ContainerAppProvisioningState = "Succeeded"
)

// PossibleContainerAppProvisioningStateValues returns the possible values for the ContainerAppProvisioningState const type.
func PossibleContainerAppProvisioningStateValues() []ContainerAppProvisioningState {
	return []ContainerAppProvisioningState{
		ContainerAppProvisioningStateCanceled,
		ContainerAppProvisioningStateDeleting,
		ContainerAppProvisioningStateFailed,
		ContainerAppProvisioningStateInProgress,
		ContainerAppProvisioningStateSucceeded,
	}
}

// ContainerAppReplicaRunningState - Current running state of the replica
type ContainerAppReplicaRunningState string

const (
	ContainerAppReplicaRunningStateNotRunning ContainerAppReplicaRunningState = "NotRunning"
	ContainerAppReplicaRunningStateRunning    ContainerAppReplicaRunningState = "Running"
	ContainerAppReplicaRunningStateUnknown    ContainerAppReplicaRunningState = "Unknown"
)

// PossibleContainerAppReplicaRunningStateValues returns the possible values for the ContainerAppReplicaRunningState const type.
func PossibleContainerAppReplicaRunningStateValues() []ContainerAppReplicaRunningState {
	return []ContainerAppReplicaRunningState{
		ContainerAppReplicaRunningStateNotRunning,
		ContainerAppReplicaRunningStateRunning,
		ContainerAppReplicaRunningStateUnknown,
	}
}

// ContainerAppRunningStatus - Running status of the Container App.
type ContainerAppRunningStatus string

const (
	// ContainerAppRunningStatusProgressing - Container App is transitioning between Stopped and Running states.
	ContainerAppRunningStatusProgressing ContainerAppRunningStatus = "Progressing"
	// ContainerAppRunningStatusReady - Container App Job is in Ready state.
	ContainerAppRunningStatusReady ContainerAppRunningStatus = "Ready"
	// ContainerAppRunningStatusRunning - Container App is in Running state.
	ContainerAppRunningStatusRunning ContainerAppRunningStatus = "Running"
	// ContainerAppRunningStatusStopped - Container App is in Stopped state.
	ContainerAppRunningStatusStopped ContainerAppRunningStatus = "Stopped"
	// ContainerAppRunningStatusSuspended - Container App Job is in Suspended state.
	ContainerAppRunningStatusSuspended ContainerAppRunningStatus = "Suspended"
)

// PossibleContainerAppRunningStatusValues returns the possible values for the ContainerAppRunningStatus const type.
func PossibleContainerAppRunningStatusValues() []ContainerAppRunningStatus {
	return []ContainerAppRunningStatus{
		ContainerAppRunningStatusProgressing,
		ContainerAppRunningStatusReady,
		ContainerAppRunningStatusRunning,
		ContainerAppRunningStatusStopped,
		ContainerAppRunningStatusSuspended,
	}
}

// ContainerType - The container type of the sessions.
type ContainerType string

const (
	ContainerTypeCustomContainer ContainerType = "CustomContainer"
	ContainerTypePythonLTS       ContainerType = "PythonLTS"
)

// PossibleContainerTypeValues returns the possible values for the ContainerType const type.
func PossibleContainerTypeValues() []ContainerType {
	return []ContainerType{
		ContainerTypeCustomContainer,
		ContainerTypePythonLTS,
	}
}

// CookieExpirationConvention - The convention used when determining the session cookie's expiration.
type CookieExpirationConvention string

const (
	CookieExpirationConventionFixedTime               CookieExpirationConvention = "FixedTime"
	CookieExpirationConventionIdentityProviderDerived CookieExpirationConvention = "IdentityProviderDerived"
)

// PossibleCookieExpirationConventionValues returns the possible values for the CookieExpirationConvention const type.
func PossibleCookieExpirationConventionValues() []CookieExpirationConvention {
	return []CookieExpirationConvention{
		CookieExpirationConventionFixedTime,
		CookieExpirationConventionIdentityProviderDerived,
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

// DNSVerificationTestResult - DNS verification test result.
type DNSVerificationTestResult string

const (
	DNSVerificationTestResultFailed  DNSVerificationTestResult = "Failed"
	DNSVerificationTestResultPassed  DNSVerificationTestResult = "Passed"
	DNSVerificationTestResultSkipped DNSVerificationTestResult = "Skipped"
)

// PossibleDNSVerificationTestResultValues returns the possible values for the DNSVerificationTestResult const type.
func PossibleDNSVerificationTestResultValues() []DNSVerificationTestResult {
	return []DNSVerificationTestResult{
		DNSVerificationTestResultFailed,
		DNSVerificationTestResultPassed,
		DNSVerificationTestResultSkipped,
	}
}

// EnvironmentProvisioningState - Provisioning state of the Environment.
type EnvironmentProvisioningState string

const (
	EnvironmentProvisioningStateCanceled                      EnvironmentProvisioningState = "Canceled"
	EnvironmentProvisioningStateFailed                        EnvironmentProvisioningState = "Failed"
	EnvironmentProvisioningStateInfrastructureSetupComplete   EnvironmentProvisioningState = "InfrastructureSetupComplete"
	EnvironmentProvisioningStateInfrastructureSetupInProgress EnvironmentProvisioningState = "InfrastructureSetupInProgress"
	EnvironmentProvisioningStateInitializationInProgress      EnvironmentProvisioningState = "InitializationInProgress"
	EnvironmentProvisioningStateScheduledForDelete            EnvironmentProvisioningState = "ScheduledForDelete"
	EnvironmentProvisioningStateSucceeded                     EnvironmentProvisioningState = "Succeeded"
	EnvironmentProvisioningStateUpgradeFailed                 EnvironmentProvisioningState = "UpgradeFailed"
	EnvironmentProvisioningStateUpgradeRequested              EnvironmentProvisioningState = "UpgradeRequested"
	EnvironmentProvisioningStateWaiting                       EnvironmentProvisioningState = "Waiting"
)

// PossibleEnvironmentProvisioningStateValues returns the possible values for the EnvironmentProvisioningState const type.
func PossibleEnvironmentProvisioningStateValues() []EnvironmentProvisioningState {
	return []EnvironmentProvisioningState{
		EnvironmentProvisioningStateCanceled,
		EnvironmentProvisioningStateFailed,
		EnvironmentProvisioningStateInfrastructureSetupComplete,
		EnvironmentProvisioningStateInfrastructureSetupInProgress,
		EnvironmentProvisioningStateInitializationInProgress,
		EnvironmentProvisioningStateScheduledForDelete,
		EnvironmentProvisioningStateSucceeded,
		EnvironmentProvisioningStateUpgradeFailed,
		EnvironmentProvisioningStateUpgradeRequested,
		EnvironmentProvisioningStateWaiting,
	}
}

// ExtendedLocationTypes - The type of extendedLocation.
type ExtendedLocationTypes string

const (
	ExtendedLocationTypesCustomLocation ExtendedLocationTypes = "CustomLocation"
)

// PossibleExtendedLocationTypesValues returns the possible values for the ExtendedLocationTypes const type.
func PossibleExtendedLocationTypesValues() []ExtendedLocationTypes {
	return []ExtendedLocationTypes{
		ExtendedLocationTypesCustomLocation,
	}
}

// ForwardProxyConvention - The convention used to determine the url of the request made.
type ForwardProxyConvention string

const (
	ForwardProxyConventionCustom   ForwardProxyConvention = "Custom"
	ForwardProxyConventionNoProxy  ForwardProxyConvention = "NoProxy"
	ForwardProxyConventionStandard ForwardProxyConvention = "Standard"
)

// PossibleForwardProxyConventionValues returns the possible values for the ForwardProxyConvention const type.
func PossibleForwardProxyConventionValues() []ForwardProxyConvention {
	return []ForwardProxyConvention{
		ForwardProxyConventionCustom,
		ForwardProxyConventionNoProxy,
		ForwardProxyConventionStandard,
	}
}

// IdentitySettingsLifeCycle - Use to select the lifecycle stages of a Container App during which the Managed Identity should
// be available.
type IdentitySettingsLifeCycle string

const (
	IdentitySettingsLifeCycleAll  IdentitySettingsLifeCycle = "All"
	IdentitySettingsLifeCycleInit IdentitySettingsLifeCycle = "Init"
	IdentitySettingsLifeCycleMain IdentitySettingsLifeCycle = "Main"
	IdentitySettingsLifeCycleNone IdentitySettingsLifeCycle = "None"
)

// PossibleIdentitySettingsLifeCycleValues returns the possible values for the IdentitySettingsLifeCycle const type.
func PossibleIdentitySettingsLifeCycleValues() []IdentitySettingsLifeCycle {
	return []IdentitySettingsLifeCycle{
		IdentitySettingsLifeCycleAll,
		IdentitySettingsLifeCycleInit,
		IdentitySettingsLifeCycleMain,
		IdentitySettingsLifeCycleNone,
	}
}

// IngressClientCertificateMode - Client certificate mode for mTLS authentication. Ignore indicates server drops client certificate
// on forwarding. Accept indicates server forwards client certificate but does not require a client
// certificate. Require indicates server requires a client certificate.
type IngressClientCertificateMode string

const (
	IngressClientCertificateModeAccept  IngressClientCertificateMode = "accept"
	IngressClientCertificateModeIgnore  IngressClientCertificateMode = "ignore"
	IngressClientCertificateModeRequire IngressClientCertificateMode = "require"
)

// PossibleIngressClientCertificateModeValues returns the possible values for the IngressClientCertificateMode const type.
func PossibleIngressClientCertificateModeValues() []IngressClientCertificateMode {
	return []IngressClientCertificateMode{
		IngressClientCertificateModeAccept,
		IngressClientCertificateModeIgnore,
		IngressClientCertificateModeRequire,
	}
}

// IngressTransportMethod - Ingress transport protocol
type IngressTransportMethod string

const (
	IngressTransportMethodAuto  IngressTransportMethod = "auto"
	IngressTransportMethodHTTP  IngressTransportMethod = "http"
	IngressTransportMethodHTTP2 IngressTransportMethod = "http2"
	IngressTransportMethodTCP   IngressTransportMethod = "tcp"
)

// PossibleIngressTransportMethodValues returns the possible values for the IngressTransportMethod const type.
func PossibleIngressTransportMethodValues() []IngressTransportMethod {
	return []IngressTransportMethod{
		IngressTransportMethodAuto,
		IngressTransportMethodHTTP,
		IngressTransportMethodHTTP2,
		IngressTransportMethodTCP,
	}
}

// JavaComponentProvisioningState - Provisioning state of the Java Component.
type JavaComponentProvisioningState string

const (
	JavaComponentProvisioningStateCanceled   JavaComponentProvisioningState = "Canceled"
	JavaComponentProvisioningStateDeleting   JavaComponentProvisioningState = "Deleting"
	JavaComponentProvisioningStateFailed     JavaComponentProvisioningState = "Failed"
	JavaComponentProvisioningStateInProgress JavaComponentProvisioningState = "InProgress"
	JavaComponentProvisioningStateSucceeded  JavaComponentProvisioningState = "Succeeded"
)

// PossibleJavaComponentProvisioningStateValues returns the possible values for the JavaComponentProvisioningState const type.
func PossibleJavaComponentProvisioningStateValues() []JavaComponentProvisioningState {
	return []JavaComponentProvisioningState{
		JavaComponentProvisioningStateCanceled,
		JavaComponentProvisioningStateDeleting,
		JavaComponentProvisioningStateFailed,
		JavaComponentProvisioningStateInProgress,
		JavaComponentProvisioningStateSucceeded,
	}
}

// JavaComponentType - Type of the Java Component.
type JavaComponentType string

const (
	JavaComponentTypeSpringBootAdmin   JavaComponentType = "SpringBootAdmin"
	JavaComponentTypeSpringCloudConfig JavaComponentType = "SpringCloudConfig"
	JavaComponentTypeSpringCloudEureka JavaComponentType = "SpringCloudEureka"
)

// PossibleJavaComponentTypeValues returns the possible values for the JavaComponentType const type.
func PossibleJavaComponentTypeValues() []JavaComponentType {
	return []JavaComponentType{
		JavaComponentTypeSpringBootAdmin,
		JavaComponentTypeSpringCloudConfig,
		JavaComponentTypeSpringCloudEureka,
	}
}

// JobExecutionRunningState - Current running State of the job
type JobExecutionRunningState string

const (
	JobExecutionRunningStateDegraded   JobExecutionRunningState = "Degraded"
	JobExecutionRunningStateFailed     JobExecutionRunningState = "Failed"
	JobExecutionRunningStateProcessing JobExecutionRunningState = "Processing"
	JobExecutionRunningStateRunning    JobExecutionRunningState = "Running"
	JobExecutionRunningStateStopped    JobExecutionRunningState = "Stopped"
	JobExecutionRunningStateSucceeded  JobExecutionRunningState = "Succeeded"
	JobExecutionRunningStateUnknown    JobExecutionRunningState = "Unknown"
)

// PossibleJobExecutionRunningStateValues returns the possible values for the JobExecutionRunningState const type.
func PossibleJobExecutionRunningStateValues() []JobExecutionRunningState {
	return []JobExecutionRunningState{
		JobExecutionRunningStateDegraded,
		JobExecutionRunningStateFailed,
		JobExecutionRunningStateProcessing,
		JobExecutionRunningStateRunning,
		JobExecutionRunningStateStopped,
		JobExecutionRunningStateSucceeded,
		JobExecutionRunningStateUnknown,
	}
}

// JobProvisioningState - Provisioning state of the Container Apps Job.
type JobProvisioningState string

const (
	JobProvisioningStateCanceled   JobProvisioningState = "Canceled"
	JobProvisioningStateDeleting   JobProvisioningState = "Deleting"
	JobProvisioningStateFailed     JobProvisioningState = "Failed"
	JobProvisioningStateInProgress JobProvisioningState = "InProgress"
	JobProvisioningStateSucceeded  JobProvisioningState = "Succeeded"
)

// PossibleJobProvisioningStateValues returns the possible values for the JobProvisioningState const type.
func PossibleJobProvisioningStateValues() []JobProvisioningState {
	return []JobProvisioningState{
		JobProvisioningStateCanceled,
		JobProvisioningStateDeleting,
		JobProvisioningStateFailed,
		JobProvisioningStateInProgress,
		JobProvisioningStateSucceeded,
	}
}

// LifecycleType - The lifecycle type of the session pool.
type LifecycleType string

const (
	LifecycleTypeOnContainerExit LifecycleType = "OnContainerExit"
	LifecycleTypeTimed           LifecycleType = "Timed"
)

// PossibleLifecycleTypeValues returns the possible values for the LifecycleType const type.
func PossibleLifecycleTypeValues() []LifecycleType {
	return []LifecycleType{
		LifecycleTypeOnContainerExit,
		LifecycleTypeTimed,
	}
}

// LogLevel - Sets the log level for the Dapr sidecar. Allowed values are debug, info, warn, error. Default is info.
type LogLevel string

const (
	LogLevelDebug LogLevel = "debug"
	LogLevelError LogLevel = "error"
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
)

// PossibleLogLevelValues returns the possible values for the LogLevel const type.
func PossibleLogLevelValues() []LogLevel {
	return []LogLevel{
		LogLevelDebug,
		LogLevelError,
		LogLevelInfo,
		LogLevelWarn,
	}
}

// ManagedCertificateDomainControlValidation - Selected type of domain control validation for managed certificates.
type ManagedCertificateDomainControlValidation string

const (
	ManagedCertificateDomainControlValidationCNAME ManagedCertificateDomainControlValidation = "CNAME"
	ManagedCertificateDomainControlValidationHTTP  ManagedCertificateDomainControlValidation = "HTTP"
	ManagedCertificateDomainControlValidationTXT   ManagedCertificateDomainControlValidation = "TXT"
)

// PossibleManagedCertificateDomainControlValidationValues returns the possible values for the ManagedCertificateDomainControlValidation const type.
func PossibleManagedCertificateDomainControlValidationValues() []ManagedCertificateDomainControlValidation {
	return []ManagedCertificateDomainControlValidation{
		ManagedCertificateDomainControlValidationCNAME,
		ManagedCertificateDomainControlValidationHTTP,
		ManagedCertificateDomainControlValidationTXT,
	}
}

// ManagedServiceIdentityType - Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                       ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeSystemAssigned             ManagedServiceIdentityType = "SystemAssigned"
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned,UserAssigned"
	ManagedServiceIdentityTypeUserAssigned               ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeNone,
		ManagedServiceIdentityTypeSystemAssigned,
		ManagedServiceIdentityTypeSystemAssignedUserAssigned,
		ManagedServiceIdentityTypeUserAssigned,
	}
}

// PoolManagementType - The pool management type of the session pool.
type PoolManagementType string

const (
	PoolManagementTypeDynamic PoolManagementType = "Dynamic"
	PoolManagementTypeManual  PoolManagementType = "Manual"
)

// PossiblePoolManagementTypeValues returns the possible values for the PoolManagementType const type.
func PossiblePoolManagementTypeValues() []PoolManagementType {
	return []PoolManagementType{
		PoolManagementTypeDynamic,
		PoolManagementTypeManual,
	}
}

// RevisionHealthState - Current health State of the revision
type RevisionHealthState string

const (
	RevisionHealthStateHealthy   RevisionHealthState = "Healthy"
	RevisionHealthStateNone      RevisionHealthState = "None"
	RevisionHealthStateUnhealthy RevisionHealthState = "Unhealthy"
)

// PossibleRevisionHealthStateValues returns the possible values for the RevisionHealthState const type.
func PossibleRevisionHealthStateValues() []RevisionHealthState {
	return []RevisionHealthState{
		RevisionHealthStateHealthy,
		RevisionHealthStateNone,
		RevisionHealthStateUnhealthy,
	}
}

// RevisionProvisioningState - Current provisioning State of the revision
type RevisionProvisioningState string

const (
	RevisionProvisioningStateDeprovisioned  RevisionProvisioningState = "Deprovisioned"
	RevisionProvisioningStateDeprovisioning RevisionProvisioningState = "Deprovisioning"
	RevisionProvisioningStateFailed         RevisionProvisioningState = "Failed"
	RevisionProvisioningStateProvisioned    RevisionProvisioningState = "Provisioned"
	RevisionProvisioningStateProvisioning   RevisionProvisioningState = "Provisioning"
)

// PossibleRevisionProvisioningStateValues returns the possible values for the RevisionProvisioningState const type.
func PossibleRevisionProvisioningStateValues() []RevisionProvisioningState {
	return []RevisionProvisioningState{
		RevisionProvisioningStateDeprovisioned,
		RevisionProvisioningStateDeprovisioning,
		RevisionProvisioningStateFailed,
		RevisionProvisioningStateProvisioned,
		RevisionProvisioningStateProvisioning,
	}
}

// RevisionRunningState - Current running state of the revision
type RevisionRunningState string

const (
	RevisionRunningStateDegraded   RevisionRunningState = "Degraded"
	RevisionRunningStateFailed     RevisionRunningState = "Failed"
	RevisionRunningStateProcessing RevisionRunningState = "Processing"
	RevisionRunningStateRunning    RevisionRunningState = "Running"
	RevisionRunningStateStopped    RevisionRunningState = "Stopped"
	RevisionRunningStateUnknown    RevisionRunningState = "Unknown"
)

// PossibleRevisionRunningStateValues returns the possible values for the RevisionRunningState const type.
func PossibleRevisionRunningStateValues() []RevisionRunningState {
	return []RevisionRunningState{
		RevisionRunningStateDegraded,
		RevisionRunningStateFailed,
		RevisionRunningStateProcessing,
		RevisionRunningStateRunning,
		RevisionRunningStateStopped,
		RevisionRunningStateUnknown,
	}
}

// Scheme - Scheme to use for connecting to the host. Defaults to HTTP.
type Scheme string

const (
	SchemeHTTP  Scheme = "HTTP"
	SchemeHTTPS Scheme = "HTTPS"
)

// PossibleSchemeValues returns the possible values for the Scheme const type.
func PossibleSchemeValues() []Scheme {
	return []Scheme{
		SchemeHTTP,
		SchemeHTTPS,
	}
}

// SessionNetworkStatus - Network status for the sessions.
type SessionNetworkStatus string

const (
	SessionNetworkStatusEgressDisabled SessionNetworkStatus = "EgressDisabled"
	SessionNetworkStatusEgressEnabled  SessionNetworkStatus = "EgressEnabled"
)

// PossibleSessionNetworkStatusValues returns the possible values for the SessionNetworkStatus const type.
func PossibleSessionNetworkStatusValues() []SessionNetworkStatus {
	return []SessionNetworkStatus{
		SessionNetworkStatusEgressDisabled,
		SessionNetworkStatusEgressEnabled,
	}
}

// SessionPoolProvisioningState - Provisioning state of the session pool.
type SessionPoolProvisioningState string

const (
	SessionPoolProvisioningStateCanceled   SessionPoolProvisioningState = "Canceled"
	SessionPoolProvisioningStateDeleting   SessionPoolProvisioningState = "Deleting"
	SessionPoolProvisioningStateFailed     SessionPoolProvisioningState = "Failed"
	SessionPoolProvisioningStateInProgress SessionPoolProvisioningState = "InProgress"
	SessionPoolProvisioningStateSucceeded  SessionPoolProvisioningState = "Succeeded"
)

// PossibleSessionPoolProvisioningStateValues returns the possible values for the SessionPoolProvisioningState const type.
func PossibleSessionPoolProvisioningStateValues() []SessionPoolProvisioningState {
	return []SessionPoolProvisioningState{
		SessionPoolProvisioningStateCanceled,
		SessionPoolProvisioningStateDeleting,
		SessionPoolProvisioningStateFailed,
		SessionPoolProvisioningStateInProgress,
		SessionPoolProvisioningStateSucceeded,
	}
}

// SourceControlOperationState - Current provisioning State of the operation
type SourceControlOperationState string

const (
	SourceControlOperationStateCanceled   SourceControlOperationState = "Canceled"
	SourceControlOperationStateFailed     SourceControlOperationState = "Failed"
	SourceControlOperationStateInProgress SourceControlOperationState = "InProgress"
	SourceControlOperationStateSucceeded  SourceControlOperationState = "Succeeded"
)

// PossibleSourceControlOperationStateValues returns the possible values for the SourceControlOperationState const type.
func PossibleSourceControlOperationStateValues() []SourceControlOperationState {
	return []SourceControlOperationState{
		SourceControlOperationStateCanceled,
		SourceControlOperationStateFailed,
		SourceControlOperationStateInProgress,
		SourceControlOperationStateSucceeded,
	}
}

// StorageType - Storage type for the volume. If not provided, use EmptyDir.
type StorageType string

const (
	StorageTypeAzureFile    StorageType = "AzureFile"
	StorageTypeEmptyDir     StorageType = "EmptyDir"
	StorageTypeNfsAzureFile StorageType = "NfsAzureFile"
	StorageTypeSecret       StorageType = "Secret"
)

// PossibleStorageTypeValues returns the possible values for the StorageType const type.
func PossibleStorageTypeValues() []StorageType {
	return []StorageType{
		StorageTypeAzureFile,
		StorageTypeEmptyDir,
		StorageTypeNfsAzureFile,
		StorageTypeSecret,
	}
}

// TriggerType - Trigger type of the job
type TriggerType string

const (
	TriggerTypeEvent    TriggerType = "Event"
	TriggerTypeManual   TriggerType = "Manual"
	TriggerTypeSchedule TriggerType = "Schedule"
)

// PossibleTriggerTypeValues returns the possible values for the TriggerType const type.
func PossibleTriggerTypeValues() []TriggerType {
	return []TriggerType{
		TriggerTypeEvent,
		TriggerTypeManual,
		TriggerTypeSchedule,
	}
}

// Type - The type of probe.
type Type string

const (
	TypeLiveness  Type = "Liveness"
	TypeReadiness Type = "Readiness"
	TypeStartup   Type = "Startup"
)

// PossibleTypeValues returns the possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{
		TypeLiveness,
		TypeReadiness,
		TypeStartup,
	}
}

// UnauthenticatedClientActionV2 - The action to take when an unauthenticated client attempts to access the app.
type UnauthenticatedClientActionV2 string

const (
	UnauthenticatedClientActionV2AllowAnonymous      UnauthenticatedClientActionV2 = "AllowAnonymous"
	UnauthenticatedClientActionV2RedirectToLoginPage UnauthenticatedClientActionV2 = "RedirectToLoginPage"
	UnauthenticatedClientActionV2Return401           UnauthenticatedClientActionV2 = "Return401"
	UnauthenticatedClientActionV2Return403           UnauthenticatedClientActionV2 = "Return403"
)

// PossibleUnauthenticatedClientActionV2Values returns the possible values for the UnauthenticatedClientActionV2 const type.
func PossibleUnauthenticatedClientActionV2Values() []UnauthenticatedClientActionV2 {
	return []UnauthenticatedClientActionV2{
		UnauthenticatedClientActionV2AllowAnonymous,
		UnauthenticatedClientActionV2RedirectToLoginPage,
		UnauthenticatedClientActionV2Return401,
		UnauthenticatedClientActionV2Return403,
	}
}
