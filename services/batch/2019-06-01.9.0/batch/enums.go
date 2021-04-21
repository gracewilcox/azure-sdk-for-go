package batch

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AccessScope enumerates the values for access scope.
type AccessScope string

const (
	// Job Grants access to perform all operations on the Job containing the Task.
	Job AccessScope = "job"
)

// PossibleAccessScopeValues returns an array of possible values for the AccessScope const type.
func PossibleAccessScopeValues() []AccessScope {
	return []AccessScope{Job}
}

// AllocationState enumerates the values for allocation state.
type AllocationState string

const (
	// Resizing The Pool is resizing; that is, Compute Nodes are being added to or removed from the Pool.
	Resizing AllocationState = "resizing"
	// Steady The Pool is not resizing. There are no changes to the number of Compute Nodes in the Pool in
	// progress. A Pool enters this state when it is created and when no operations are being performed on the
	// Pool to change the number of Compute Nodes.
	Steady AllocationState = "steady"
	// Stopping The Pool was resizing, but the user has requested that the resize be stopped, but the stop
	// request has not yet been completed.
	Stopping AllocationState = "stopping"
)

// PossibleAllocationStateValues returns an array of possible values for the AllocationState const type.
func PossibleAllocationStateValues() []AllocationState {
	return []AllocationState{Resizing, Steady, Stopping}
}

// AutoUserScope enumerates the values for auto user scope.
type AutoUserScope string

const (
	// Pool Specifies that the Task runs as the common auto user Account which is created on every Compute Node
	// in a Pool.
	Pool AutoUserScope = "pool"
	// Task Specifies that the service should create a new user for the Task.
	Task AutoUserScope = "task"
)

// PossibleAutoUserScopeValues returns an array of possible values for the AutoUserScope const type.
func PossibleAutoUserScopeValues() []AutoUserScope {
	return []AutoUserScope{Pool, Task}
}

// CachingType enumerates the values for caching type.
type CachingType string

const (
	// None The caching mode for the disk is not enabled.
	None CachingType = "none"
	// ReadOnly The caching mode for the disk is read only.
	ReadOnly CachingType = "readonly"
	// ReadWrite The caching mode for the disk is read and write.
	ReadWrite CachingType = "readwrite"
)

// PossibleCachingTypeValues returns an array of possible values for the CachingType const type.
func PossibleCachingTypeValues() []CachingType {
	return []CachingType{None, ReadOnly, ReadWrite}
}

// CertificateFormat enumerates the values for certificate format.
type CertificateFormat string

const (
	// Cer The Certificate is a base64-encoded X.509 Certificate.
	Cer CertificateFormat = "cer"
	// Pfx The Certificate is a PFX (PKCS#12) formatted Certificate or Certificate chain.
	Pfx CertificateFormat = "pfx"
)

// PossibleCertificateFormatValues returns an array of possible values for the CertificateFormat const type.
func PossibleCertificateFormatValues() []CertificateFormat {
	return []CertificateFormat{Cer, Pfx}
}

// CertificateState enumerates the values for certificate state.
type CertificateState string

const (
	// Active The Certificate is available for use in Pools.
	Active CertificateState = "active"
	// DeleteFailed The user requested that the Certificate be deleted, but there are Pools that still have
	// references to the Certificate, or it is still installed on one or more Nodes. (The latter can occur if
	// the Certificate has been removed from the Pool, but the Compute Node has not yet restarted. Compute
	// Nodes refresh their Certificates only when they restart.) You may use the cancel Certificate delete
	// operation to cancel the delete, or the delete Certificate operation to retry the delete.
	DeleteFailed CertificateState = "deletefailed"
	// Deleting The user has requested that the Certificate be deleted, but the delete operation has not yet
	// completed. You may not reference the Certificate when creating or updating Pools.
	Deleting CertificateState = "deleting"
)

// PossibleCertificateStateValues returns an array of possible values for the CertificateState const type.
func PossibleCertificateStateValues() []CertificateState {
	return []CertificateState{Active, DeleteFailed, Deleting}
}

// CertificateStoreLocation enumerates the values for certificate store location.
type CertificateStoreLocation string

const (
	// CurrentUser Certificates should be installed to the CurrentUser Certificate store.
	CurrentUser CertificateStoreLocation = "currentuser"
	// LocalMachine Certificates should be installed to the LocalMachine Certificate store.
	LocalMachine CertificateStoreLocation = "localmachine"
)

// PossibleCertificateStoreLocationValues returns an array of possible values for the CertificateStoreLocation const type.
func PossibleCertificateStoreLocationValues() []CertificateStoreLocation {
	return []CertificateStoreLocation{CurrentUser, LocalMachine}
}

// CertificateVisibility enumerates the values for certificate visibility.
type CertificateVisibility string

const (
	// CertificateVisibilityRemoteUser The Certificate should be visible to the user Accounts under which users
	// remotely access the Compute Node.
	CertificateVisibilityRemoteUser CertificateVisibility = "remoteuser"
	// CertificateVisibilityStartTask The Certificate should be visible to the user Account under which the
	// start Task is run.
	CertificateVisibilityStartTask CertificateVisibility = "starttask"
	// CertificateVisibilityTask The Certificate should be visible to the user Accounts under which Job Tasks
	// are run.
	CertificateVisibilityTask CertificateVisibility = "task"
)

// PossibleCertificateVisibilityValues returns an array of possible values for the CertificateVisibility const type.
func PossibleCertificateVisibilityValues() []CertificateVisibility {
	return []CertificateVisibility{CertificateVisibilityRemoteUser, CertificateVisibilityStartTask, CertificateVisibilityTask}
}

// ComputeNodeDeallocationOption enumerates the values for compute node deallocation option.
type ComputeNodeDeallocationOption string

const (
	// Requeue Terminate running Task processes and requeue the Tasks. The Tasks will run again when a Compute
	// Node is available. Remove Compute Nodes as soon as Tasks have been terminated.
	Requeue ComputeNodeDeallocationOption = "requeue"
	// RetainedData Allow currently running Tasks to complete, then wait for all Task data retention periods to
	// expire. Schedule no new Tasks while waiting. Remove Compute Nodes when all Task retention periods have
	// expired.
	RetainedData ComputeNodeDeallocationOption = "retaineddata"
	// TaskCompletion Allow currently running Tasks to complete. Schedule no new Tasks while waiting. Remove
	// Compute Nodes when all Tasks have completed.
	TaskCompletion ComputeNodeDeallocationOption = "taskcompletion"
	// Terminate Terminate running Tasks. The Tasks will be completed with failureInfo indicating that they
	// were terminated, and will not run again. Remove Compute Nodes as soon as Tasks have been terminated.
	Terminate ComputeNodeDeallocationOption = "terminate"
)

// PossibleComputeNodeDeallocationOptionValues returns an array of possible values for the ComputeNodeDeallocationOption const type.
func PossibleComputeNodeDeallocationOptionValues() []ComputeNodeDeallocationOption {
	return []ComputeNodeDeallocationOption{Requeue, RetainedData, TaskCompletion, Terminate}
}

// ComputeNodeFillType enumerates the values for compute node fill type.
type ComputeNodeFillType string

const (
	// Pack As many Tasks as possible (maxTasksPerNode) should be assigned to each Compute Node in the Pool
	// before any Tasks are assigned to the next Compute Node in the Pool.
	Pack ComputeNodeFillType = "pack"
	// Spread Tasks should be assigned evenly across all Compute Nodes in the Pool.
	Spread ComputeNodeFillType = "spread"
)

// PossibleComputeNodeFillTypeValues returns an array of possible values for the ComputeNodeFillType const type.
func PossibleComputeNodeFillTypeValues() []ComputeNodeFillType {
	return []ComputeNodeFillType{Pack, Spread}
}

// ComputeNodeRebootOption enumerates the values for compute node reboot option.
type ComputeNodeRebootOption string

const (
	// ComputeNodeRebootOptionRequeue Terminate running Task processes and requeue the Tasks. The Tasks will
	// run again when a Compute Node is available. Restart the Compute Node as soon as Tasks have been
	// terminated.
	ComputeNodeRebootOptionRequeue ComputeNodeRebootOption = "requeue"
	// ComputeNodeRebootOptionRetainedData Allow currently running Tasks to complete, then wait for all Task
	// data retention periods to expire. Schedule no new Tasks while waiting. Restart the Compute Node when all
	// Task retention periods have expired.
	ComputeNodeRebootOptionRetainedData ComputeNodeRebootOption = "retaineddata"
	// ComputeNodeRebootOptionTaskCompletion Allow currently running Tasks to complete. Schedule no new Tasks
	// while waiting. Restart the Compute Node when all Tasks have completed.
	ComputeNodeRebootOptionTaskCompletion ComputeNodeRebootOption = "taskcompletion"
	// ComputeNodeRebootOptionTerminate Terminate running Tasks. The Tasks will be completed with failureInfo
	// indicating that they were terminated, and will not run again. Restart the Compute Node as soon as Tasks
	// have been terminated.
	ComputeNodeRebootOptionTerminate ComputeNodeRebootOption = "terminate"
)

// PossibleComputeNodeRebootOptionValues returns an array of possible values for the ComputeNodeRebootOption const type.
func PossibleComputeNodeRebootOptionValues() []ComputeNodeRebootOption {
	return []ComputeNodeRebootOption{ComputeNodeRebootOptionRequeue, ComputeNodeRebootOptionRetainedData, ComputeNodeRebootOptionTaskCompletion, ComputeNodeRebootOptionTerminate}
}

// ComputeNodeReimageOption enumerates the values for compute node reimage option.
type ComputeNodeReimageOption string

const (
	// ComputeNodeReimageOptionRequeue Terminate running Task processes and requeue the Tasks. The Tasks will
	// run again when a Compute Node is available. Reimage the Compute Node as soon as Tasks have been
	// terminated.
	ComputeNodeReimageOptionRequeue ComputeNodeReimageOption = "requeue"
	// ComputeNodeReimageOptionRetainedData Allow currently running Tasks to complete, then wait for all Task
	// data retention periods to expire. Schedule no new Tasks while waiting. Reimage the Compute Node when all
	// Task retention periods have expired.
	ComputeNodeReimageOptionRetainedData ComputeNodeReimageOption = "retaineddata"
	// ComputeNodeReimageOptionTaskCompletion Allow currently running Tasks to complete. Schedule no new Tasks
	// while waiting. Reimage the Compute Node when all Tasks have completed.
	ComputeNodeReimageOptionTaskCompletion ComputeNodeReimageOption = "taskcompletion"
	// ComputeNodeReimageOptionTerminate Terminate running Tasks. The Tasks will be completed with failureInfo
	// indicating that they were terminated, and will not run again. Reimage the Compute Node as soon as Tasks
	// have been terminated.
	ComputeNodeReimageOptionTerminate ComputeNodeReimageOption = "terminate"
)

// PossibleComputeNodeReimageOptionValues returns an array of possible values for the ComputeNodeReimageOption const type.
func PossibleComputeNodeReimageOptionValues() []ComputeNodeReimageOption {
	return []ComputeNodeReimageOption{ComputeNodeReimageOptionRequeue, ComputeNodeReimageOptionRetainedData, ComputeNodeReimageOptionTaskCompletion, ComputeNodeReimageOptionTerminate}
}

// ComputeNodeState enumerates the values for compute node state.
type ComputeNodeState string

const (
	// Creating The Batch service has obtained the underlying virtual machine from Azure Compute, but it has
	// not yet started to join the Pool.
	Creating ComputeNodeState = "creating"
	// Idle The Compute Node is not currently running a Task.
	Idle ComputeNodeState = "idle"
	// LeavingPool The Compute Node is leaving the Pool, either because the user explicitly removed it or
	// because the Pool is resizing or autoscaling down.
	LeavingPool ComputeNodeState = "leavingpool"
	// Offline The Compute Node is not currently running a Task, and scheduling of new Tasks to the Compute
	// Node is disabled.
	Offline ComputeNodeState = "offline"
	// Preempted The low-priority Compute Node has been preempted. Tasks which were running on the Compute Node
	// when it was preempted will be rescheduled when another Compute Node becomes available.
	Preempted ComputeNodeState = "preempted"
	// Rebooting The Compute Node is rebooting.
	Rebooting ComputeNodeState = "rebooting"
	// Reimaging The Compute Node is reimaging.
	Reimaging ComputeNodeState = "reimaging"
	// Running The Compute Node is running one or more Tasks (other than a start task).
	Running ComputeNodeState = "running"
	// Starting The Batch service is starting on the underlying virtual machine.
	Starting ComputeNodeState = "starting"
	// StartTaskFailed The start Task has failed on the Compute Node (and exhausted all retries), and
	// waitForSuccess is set. The Compute Node is not usable for running Tasks.
	StartTaskFailed ComputeNodeState = "starttaskfailed"
	// Unknown The Batch service has lost contact with the Compute Node, and does not know its true state.
	Unknown ComputeNodeState = "unknown"
	// Unusable The Compute Node cannot be used for Task execution due to errors.
	Unusable ComputeNodeState = "unusable"
	// WaitingForStartTask The start Task has started running on the Compute Node, but waitForSuccess is set
	// and the start Task has not yet completed.
	WaitingForStartTask ComputeNodeState = "waitingforstarttask"
)

// PossibleComputeNodeStateValues returns an array of possible values for the ComputeNodeState const type.
func PossibleComputeNodeStateValues() []ComputeNodeState {
	return []ComputeNodeState{Creating, Idle, LeavingPool, Offline, Preempted, Rebooting, Reimaging, Running, Starting, StartTaskFailed, Unknown, Unusable, WaitingForStartTask}
}

// ContainerWorkingDirectory enumerates the values for container working directory.
type ContainerWorkingDirectory string

const (
	// ContainerImageDefault Use the working directory defined in the container Image. Beware that this
	// directory will not contain the Resource Files downloaded by Batch.
	ContainerImageDefault ContainerWorkingDirectory = "containerImageDefault"
	// TaskWorkingDirectory Use the standard Batch service Task working directory, which will contain the Task
	// Resource Files populated by Batch.
	TaskWorkingDirectory ContainerWorkingDirectory = "taskWorkingDirectory"
)

// PossibleContainerWorkingDirectoryValues returns an array of possible values for the ContainerWorkingDirectory const type.
func PossibleContainerWorkingDirectoryValues() []ContainerWorkingDirectory {
	return []ContainerWorkingDirectory{ContainerImageDefault, TaskWorkingDirectory}
}

// DependencyAction enumerates the values for dependency action.
type DependencyAction string

const (
	// Block Block the Task's dependencies.
	Block DependencyAction = "block"
	// Satisfy Satisfy the Task's dependencies.
	Satisfy DependencyAction = "satisfy"
)

// PossibleDependencyActionValues returns an array of possible values for the DependencyAction const type.
func PossibleDependencyActionValues() []DependencyAction {
	return []DependencyAction{Block, Satisfy}
}

// DisableComputeNodeSchedulingOption enumerates the values for disable compute node scheduling option.
type DisableComputeNodeSchedulingOption string

const (
	// DisableComputeNodeSchedulingOptionRequeue Terminate running Task processes and requeue the Tasks. The
	// Tasks may run again on other Compute Nodes, or when Task scheduling is re-enabled on this Compute Node.
	// Enter offline state as soon as Tasks have been terminated.
	DisableComputeNodeSchedulingOptionRequeue DisableComputeNodeSchedulingOption = "requeue"
	// DisableComputeNodeSchedulingOptionTaskCompletion Allow currently running Tasks to complete. Schedule no
	// new Tasks while waiting. Enter offline state when all Tasks have completed.
	DisableComputeNodeSchedulingOptionTaskCompletion DisableComputeNodeSchedulingOption = "taskcompletion"
	// DisableComputeNodeSchedulingOptionTerminate Terminate running Tasks. The Tasks will be completed with
	// failureInfo indicating that they were terminated, and will not run again. Enter offline state as soon as
	// Tasks have been terminated.
	DisableComputeNodeSchedulingOptionTerminate DisableComputeNodeSchedulingOption = "terminate"
)

// PossibleDisableComputeNodeSchedulingOptionValues returns an array of possible values for the DisableComputeNodeSchedulingOption const type.
func PossibleDisableComputeNodeSchedulingOptionValues() []DisableComputeNodeSchedulingOption {
	return []DisableComputeNodeSchedulingOption{DisableComputeNodeSchedulingOptionRequeue, DisableComputeNodeSchedulingOptionTaskCompletion, DisableComputeNodeSchedulingOptionTerminate}
}

// DisableJobOption enumerates the values for disable job option.
type DisableJobOption string

const (
	// DisableJobOptionRequeue Terminate running Tasks and requeue them. The Tasks will run again when the Job
	// is enabled.
	DisableJobOptionRequeue DisableJobOption = "requeue"
	// DisableJobOptionTerminate Terminate running Tasks. The Tasks will be completed with failureInfo
	// indicating that they were terminated, and will not run again.
	DisableJobOptionTerminate DisableJobOption = "terminate"
	// DisableJobOptionWait Allow currently running Tasks to complete.
	DisableJobOptionWait DisableJobOption = "wait"
)

// PossibleDisableJobOptionValues returns an array of possible values for the DisableJobOption const type.
func PossibleDisableJobOptionValues() []DisableJobOption {
	return []DisableJobOption{DisableJobOptionRequeue, DisableJobOptionTerminate, DisableJobOptionWait}
}

// DynamicVNetAssignmentScope enumerates the values for dynamic v net assignment scope.
type DynamicVNetAssignmentScope string

const (
	// DynamicVNetAssignmentScopeJob Dynamic VNet assignment is done per-job.
	DynamicVNetAssignmentScopeJob DynamicVNetAssignmentScope = "job"
	// DynamicVNetAssignmentScopeNone No dynamic VNet assignment is enabled.
	DynamicVNetAssignmentScopeNone DynamicVNetAssignmentScope = "none"
)

// PossibleDynamicVNetAssignmentScopeValues returns an array of possible values for the DynamicVNetAssignmentScope const type.
func PossibleDynamicVNetAssignmentScopeValues() []DynamicVNetAssignmentScope {
	return []DynamicVNetAssignmentScope{DynamicVNetAssignmentScopeJob, DynamicVNetAssignmentScopeNone}
}

// ElevationLevel enumerates the values for elevation level.
type ElevationLevel string

const (
	// Admin The user is a user with elevated access and operates with full Administrator permissions.
	Admin ElevationLevel = "admin"
	// NonAdmin The user is a standard user without elevated access.
	NonAdmin ElevationLevel = "nonadmin"
)

// PossibleElevationLevelValues returns an array of possible values for the ElevationLevel const type.
func PossibleElevationLevelValues() []ElevationLevel {
	return []ElevationLevel{Admin, NonAdmin}
}

// ErrorCategory enumerates the values for error category.
type ErrorCategory string

const (
	// ServerError The error is due to an internal server issue.
	ServerError ErrorCategory = "servererror"
	// UserError The error is due to a user issue, such as misconfiguration.
	UserError ErrorCategory = "usererror"
)

// PossibleErrorCategoryValues returns an array of possible values for the ErrorCategory const type.
func PossibleErrorCategoryValues() []ErrorCategory {
	return []ErrorCategory{ServerError, UserError}
}

// InboundEndpointProtocol enumerates the values for inbound endpoint protocol.
type InboundEndpointProtocol string

const (
	// TCP Use TCP for the endpoint.
	TCP InboundEndpointProtocol = "tcp"
	// UDP Use UDP for the endpoint.
	UDP InboundEndpointProtocol = "udp"
)

// PossibleInboundEndpointProtocolValues returns an array of possible values for the InboundEndpointProtocol const type.
func PossibleInboundEndpointProtocolValues() []InboundEndpointProtocol {
	return []InboundEndpointProtocol{TCP, UDP}
}

// JobAction enumerates the values for job action.
type JobAction string

const (
	// JobActionDisable Disable the Job. This is equivalent to calling the disable Job API, with a disableTasks
	// value of requeue.
	JobActionDisable JobAction = "disable"
	// JobActionNone Take no action.
	JobActionNone JobAction = "none"
	// JobActionTerminate Terminate the Job. The terminateReason in the Job's executionInfo is set to
	// "TaskFailed".
	JobActionTerminate JobAction = "terminate"
)

// PossibleJobActionValues returns an array of possible values for the JobAction const type.
func PossibleJobActionValues() []JobAction {
	return []JobAction{JobActionDisable, JobActionNone, JobActionTerminate}
}

// JobPreparationTaskState enumerates the values for job preparation task state.
type JobPreparationTaskState string

const (
	// JobPreparationTaskStateCompleted The Task has exited with exit code 0, or the Task has exhausted its
	// retry limit, or the Batch service was unable to start the Task due to Task preparation errors (such as
	// resource file download failures).
	JobPreparationTaskStateCompleted JobPreparationTaskState = "completed"
	// JobPreparationTaskStateRunning The Task is currently running (including retrying).
	JobPreparationTaskStateRunning JobPreparationTaskState = "running"
)

// PossibleJobPreparationTaskStateValues returns an array of possible values for the JobPreparationTaskState const type.
func PossibleJobPreparationTaskStateValues() []JobPreparationTaskState {
	return []JobPreparationTaskState{JobPreparationTaskStateCompleted, JobPreparationTaskStateRunning}
}

// JobReleaseTaskState enumerates the values for job release task state.
type JobReleaseTaskState string

const (
	// JobReleaseTaskStateCompleted The Task has exited with exit code 0, or the Task has exhausted its retry
	// limit, or the Batch service was unable to start the Task due to Task preparation errors (such as
	// resource file download failures).
	JobReleaseTaskStateCompleted JobReleaseTaskState = "completed"
	// JobReleaseTaskStateRunning The Task is currently running (including retrying).
	JobReleaseTaskStateRunning JobReleaseTaskState = "running"
)

// PossibleJobReleaseTaskStateValues returns an array of possible values for the JobReleaseTaskState const type.
func PossibleJobReleaseTaskStateValues() []JobReleaseTaskState {
	return []JobReleaseTaskState{JobReleaseTaskStateCompleted, JobReleaseTaskStateRunning}
}

// JobScheduleState enumerates the values for job schedule state.
type JobScheduleState string

const (
	// JobScheduleStateActive The Job Schedule is active and will create Jobs as per its schedule.
	JobScheduleStateActive JobScheduleState = "active"
	// JobScheduleStateCompleted The Job Schedule has terminated, either by reaching its end time or by the
	// user terminating it explicitly.
	JobScheduleStateCompleted JobScheduleState = "completed"
	// JobScheduleStateDeleting The user has requested that the Job Schedule be deleted, but the delete
	// operation is still in progress. The scheduler will not initiate any new Jobs for this Job Schedule, and
	// will delete any existing Jobs and Tasks under the Job Schedule, including any active Job. The Job
	// Schedule will be deleted when all Jobs and Tasks under the Job Schedule have been deleted.
	JobScheduleStateDeleting JobScheduleState = "deleting"
	// JobScheduleStateDisabled The user has disabled the Job Schedule. The scheduler will not initiate any new
	// Jobs will on this schedule, but any existing active Job will continue to run.
	JobScheduleStateDisabled JobScheduleState = "disabled"
	// JobScheduleStateTerminating The Job Schedule has no more work to do, or has been explicitly terminated
	// by the user, but the termination operation is still in progress. The scheduler will not initiate any new
	// Jobs for this Job Schedule, nor is any existing Job active.
	JobScheduleStateTerminating JobScheduleState = "terminating"
)

// PossibleJobScheduleStateValues returns an array of possible values for the JobScheduleState const type.
func PossibleJobScheduleStateValues() []JobScheduleState {
	return []JobScheduleState{JobScheduleStateActive, JobScheduleStateCompleted, JobScheduleStateDeleting, JobScheduleStateDisabled, JobScheduleStateTerminating}
}

// JobState enumerates the values for job state.
type JobState string

const (
	// JobStateActive The Job is available to have Tasks scheduled.
	JobStateActive JobState = "active"
	// JobStateCompleted All Tasks have terminated, and the system will not accept any more Tasks or any
	// further changes to the Job.
	JobStateCompleted JobState = "completed"
	// JobStateDeleting A user has requested that the Job be deleted, but the delete operation is still in
	// progress (for example, because the system is still terminating running Tasks).
	JobStateDeleting JobState = "deleting"
	// JobStateDisabled A user has disabled the Job. No Tasks are running, and no new Tasks will be scheduled.
	JobStateDisabled JobState = "disabled"
	// JobStateDisabling A user has requested that the Job be disabled, but the disable operation is still in
	// progress (for example, waiting for Tasks to terminate).
	JobStateDisabling JobState = "disabling"
	// JobStateEnabling A user has requested that the Job be enabled, but the enable operation is still in
	// progress.
	JobStateEnabling JobState = "enabling"
	// JobStateTerminating The Job is about to complete, either because a Job Manager Task has completed or
	// because the user has terminated the Job, but the terminate operation is still in progress (for example,
	// because Job Release Tasks are running).
	JobStateTerminating JobState = "terminating"
)

// PossibleJobStateValues returns an array of possible values for the JobState const type.
func PossibleJobStateValues() []JobState {
	return []JobState{JobStateActive, JobStateCompleted, JobStateDeleting, JobStateDisabled, JobStateDisabling, JobStateEnabling, JobStateTerminating}
}

// LoginMode enumerates the values for login mode.
type LoginMode string

const (
	// Batch The LOGON32_LOGON_BATCH Win32 login mode. The batch login mode is recommended for long running
	// parallel processes.
	Batch LoginMode = "batch"
	// Interactive The LOGON32_LOGON_INTERACTIVE Win32 login mode. UAC is enabled on Windows
	// VirtualMachineConfiguration Pools. If this option is used with an elevated user identity in a Windows
	// VirtualMachineConfiguration Pool, the user session will not be elevated unless the application executed
	// by the Task command line is configured to always require administrative privilege or to always require
	// maximum privilege.
	Interactive LoginMode = "interactive"
)

// PossibleLoginModeValues returns an array of possible values for the LoginMode const type.
func PossibleLoginModeValues() []LoginMode {
	return []LoginMode{Batch, Interactive}
}

// NetworkSecurityGroupRuleAccess enumerates the values for network security group rule access.
type NetworkSecurityGroupRuleAccess string

const (
	// Allow Allow access.
	Allow NetworkSecurityGroupRuleAccess = "allow"
	// Deny Deny access.
	Deny NetworkSecurityGroupRuleAccess = "deny"
)

// PossibleNetworkSecurityGroupRuleAccessValues returns an array of possible values for the NetworkSecurityGroupRuleAccess const type.
func PossibleNetworkSecurityGroupRuleAccessValues() []NetworkSecurityGroupRuleAccess {
	return []NetworkSecurityGroupRuleAccess{Allow, Deny}
}

// OnAllTasksComplete enumerates the values for on all tasks complete.
type OnAllTasksComplete string

const (
	// NoAction Do nothing. The Job remains active unless terminated or disabled by some other means.
	NoAction OnAllTasksComplete = "noaction"
	// TerminateJob Terminate the Job. The Job's terminateReason is set to 'AllTasksComplete'.
	TerminateJob OnAllTasksComplete = "terminatejob"
)

// PossibleOnAllTasksCompleteValues returns an array of possible values for the OnAllTasksComplete const type.
func PossibleOnAllTasksCompleteValues() []OnAllTasksComplete {
	return []OnAllTasksComplete{NoAction, TerminateJob}
}

// OnTaskFailure enumerates the values for on task failure.
type OnTaskFailure string

const (
	// OnTaskFailureNoAction Do nothing. The Job remains active unless terminated or disabled by some other
	// means.
	OnTaskFailureNoAction OnTaskFailure = "noaction"
	// OnTaskFailurePerformExitOptionsJobAction Take the action associated with the Task exit condition in the
	// Task's exitConditions collection. (This may still result in no action being taken, if that is what the
	// Task specifies.)
	OnTaskFailurePerformExitOptionsJobAction OnTaskFailure = "performexitoptionsjobaction"
)

// PossibleOnTaskFailureValues returns an array of possible values for the OnTaskFailure const type.
func PossibleOnTaskFailureValues() []OnTaskFailure {
	return []OnTaskFailure{OnTaskFailureNoAction, OnTaskFailurePerformExitOptionsJobAction}
}

// OSType enumerates the values for os type.
type OSType string

const (
	// Linux The Linux operating system.
	Linux OSType = "linux"
	// Windows The Windows operating system.
	Windows OSType = "windows"
)

// PossibleOSTypeValues returns an array of possible values for the OSType const type.
func PossibleOSTypeValues() []OSType {
	return []OSType{Linux, Windows}
}

// OutputFileUploadCondition enumerates the values for output file upload condition.
type OutputFileUploadCondition string

const (
	// OutputFileUploadConditionTaskCompletion Upload the file(s) after the Task process exits, no matter what
	// the exit code was.
	OutputFileUploadConditionTaskCompletion OutputFileUploadCondition = "taskcompletion"
	// OutputFileUploadConditionTaskFailure Upload the file(s) only after the Task process exits with a nonzero
	// exit code.
	OutputFileUploadConditionTaskFailure OutputFileUploadCondition = "taskfailure"
	// OutputFileUploadConditionTaskSuccess Upload the file(s) only after the Task process exits with an exit
	// code of 0.
	OutputFileUploadConditionTaskSuccess OutputFileUploadCondition = "tasksuccess"
)

// PossibleOutputFileUploadConditionValues returns an array of possible values for the OutputFileUploadCondition const type.
func PossibleOutputFileUploadConditionValues() []OutputFileUploadCondition {
	return []OutputFileUploadCondition{OutputFileUploadConditionTaskCompletion, OutputFileUploadConditionTaskFailure, OutputFileUploadConditionTaskSuccess}
}

// PoolLifetimeOption enumerates the values for pool lifetime option.
type PoolLifetimeOption string

const (
	// PoolLifetimeOptionJob The Pool exists for the lifetime of the Job to which it is dedicated. The Batch
	// service creates the Pool when it creates the Job. If the 'job' option is applied to a Job Schedule, the
	// Batch service creates a new auto Pool for every Job created on the schedule.
	PoolLifetimeOptionJob PoolLifetimeOption = "job"
	// PoolLifetimeOptionJobSchedule The Pool exists for the lifetime of the Job Schedule. The Batch Service
	// creates the Pool when it creates the first Job on the schedule. You may apply this option only to Job
	// Schedules, not to Jobs.
	PoolLifetimeOptionJobSchedule PoolLifetimeOption = "jobschedule"
)

// PossiblePoolLifetimeOptionValues returns an array of possible values for the PoolLifetimeOption const type.
func PossiblePoolLifetimeOptionValues() []PoolLifetimeOption {
	return []PoolLifetimeOption{PoolLifetimeOptionJob, PoolLifetimeOptionJobSchedule}
}

// PoolState enumerates the values for pool state.
type PoolState string

const (
	// PoolStateActive The Pool is available to run Tasks subject to the availability of Compute Nodes.
	PoolStateActive PoolState = "active"
	// PoolStateDeleting The user has requested that the Pool be deleted, but the delete operation has not yet
	// completed.
	PoolStateDeleting PoolState = "deleting"
)

// PossiblePoolStateValues returns an array of possible values for the PoolState const type.
func PossiblePoolStateValues() []PoolState {
	return []PoolState{PoolStateActive, PoolStateDeleting}
}

// SchedulingState enumerates the values for scheduling state.
type SchedulingState string

const (
	// Disabled No new Tasks will be scheduled on the Compute Node. Tasks already running on the Compute Node
	// may still run to completion. All Compute Nodes start with scheduling enabled.
	Disabled SchedulingState = "disabled"
	// Enabled Tasks can be scheduled on the Compute Node.
	Enabled SchedulingState = "enabled"
)

// PossibleSchedulingStateValues returns an array of possible values for the SchedulingState const type.
func PossibleSchedulingStateValues() []SchedulingState {
	return []SchedulingState{Disabled, Enabled}
}

// StartTaskState enumerates the values for start task state.
type StartTaskState string

const (
	// StartTaskStateCompleted The start Task has exited with exit code 0, or the start Task has failed and the
	// retry limit has reached, or the start Task process did not run due to Task preparation errors (such as
	// resource file download failures).
	StartTaskStateCompleted StartTaskState = "completed"
	// StartTaskStateRunning The start Task is currently running.
	StartTaskStateRunning StartTaskState = "running"
)

// PossibleStartTaskStateValues returns an array of possible values for the StartTaskState const type.
func PossibleStartTaskStateValues() []StartTaskState {
	return []StartTaskState{StartTaskStateCompleted, StartTaskStateRunning}
}

// StorageAccountType enumerates the values for storage account type.
type StorageAccountType string

const (
	// PremiumLRS The data disk should use premium locally redundant storage.
	PremiumLRS StorageAccountType = "premium_lrs"
	// StandardLRS The data disk should use standard locally redundant storage.
	StandardLRS StorageAccountType = "standard_lrs"
)

// PossibleStorageAccountTypeValues returns an array of possible values for the StorageAccountType const type.
func PossibleStorageAccountTypeValues() []StorageAccountType {
	return []StorageAccountType{PremiumLRS, StandardLRS}
}

// SubtaskState enumerates the values for subtask state.
type SubtaskState string

const (
	// SubtaskStateCompleted The Task is no longer eligible to run, usually because the Task has finished
	// successfully, or the Task has finished unsuccessfully and has exhausted its retry limit. A Task is also
	// marked as completed if an error occurred launching the Task, or when the Task has been terminated.
	SubtaskStateCompleted SubtaskState = "completed"
	// SubtaskStatePreparing The Task has been assigned to a Compute Node, but is waiting for a required Job
	// Preparation Task to complete on the Compute Node. If the Job Preparation Task succeeds, the Task will
	// move to running. If the Job Preparation Task fails, the Task will return to active and will be eligible
	// to be assigned to a different Compute Node.
	SubtaskStatePreparing SubtaskState = "preparing"
	// SubtaskStateRunning The Task is running on a Compute Node. This includes task-level preparation such as
	// downloading resource files or deploying Packages specified on the Task - it does not necessarily mean
	// that the Task command line has started executing.
	SubtaskStateRunning SubtaskState = "running"
)

// PossibleSubtaskStateValues returns an array of possible values for the SubtaskState const type.
func PossibleSubtaskStateValues() []SubtaskState {
	return []SubtaskState{SubtaskStateCompleted, SubtaskStatePreparing, SubtaskStateRunning}
}

// TaskAddStatus enumerates the values for task add status.
type TaskAddStatus string

const (
	// TaskAddStatusClientError The Task failed to add due to a client error and should not be retried without
	// modifying the request as appropriate.
	TaskAddStatusClientError TaskAddStatus = "clienterror"
	// TaskAddStatusServerError Task failed to add due to a server error and can be retried without
	// modification.
	TaskAddStatusServerError TaskAddStatus = "servererror"
	// TaskAddStatusSuccess The Task was added successfully.
	TaskAddStatusSuccess TaskAddStatus = "success"
)

// PossibleTaskAddStatusValues returns an array of possible values for the TaskAddStatus const type.
func PossibleTaskAddStatusValues() []TaskAddStatus {
	return []TaskAddStatus{TaskAddStatusClientError, TaskAddStatusServerError, TaskAddStatusSuccess}
}

// TaskExecutionResult enumerates the values for task execution result.
type TaskExecutionResult string

const (
	// Failure There was an error during processing of the Task. The failure may have occurred before the Task
	// process was launched, while the Task process was executing, or after the Task process exited.
	Failure TaskExecutionResult = "failure"
	// Success The Task ran successfully.
	Success TaskExecutionResult = "success"
)

// PossibleTaskExecutionResultValues returns an array of possible values for the TaskExecutionResult const type.
func PossibleTaskExecutionResultValues() []TaskExecutionResult {
	return []TaskExecutionResult{Failure, Success}
}

// TaskState enumerates the values for task state.
type TaskState string

const (
	// TaskStateActive The Task is queued and able to run, but is not currently assigned to a Compute Node. A
	// Task enters this state when it is created, when it is enabled after being disabled, or when it is
	// awaiting a retry after a failed run.
	TaskStateActive TaskState = "active"
	// TaskStateCompleted The Task is no longer eligible to run, usually because the Task has finished
	// successfully, or the Task has finished unsuccessfully and has exhausted its retry limit. A Task is also
	// marked as completed if an error occurred launching the Task, or when the Task has been terminated.
	TaskStateCompleted TaskState = "completed"
	// TaskStatePreparing The Task has been assigned to a Compute Node, but is waiting for a required Job
	// Preparation Task to complete on the Compute Node. If the Job Preparation Task succeeds, the Task will
	// move to running. If the Job Preparation Task fails, the Task will return to active and will be eligible
	// to be assigned to a different Compute Node.
	TaskStatePreparing TaskState = "preparing"
	// TaskStateRunning The Task is running on a Compute Node. This includes task-level preparation such as
	// downloading resource files or deploying Packages specified on the Task - it does not necessarily mean
	// that the Task command line has started executing.
	TaskStateRunning TaskState = "running"
)

// PossibleTaskStateValues returns an array of possible values for the TaskState const type.
func PossibleTaskStateValues() []TaskState {
	return []TaskState{TaskStateActive, TaskStateCompleted, TaskStatePreparing, TaskStateRunning}
}

// VerificationType enumerates the values for verification type.
type VerificationType string

const (
	// Unverified The associated Compute Node agent SKU should have binary compatibility with the Image, but
	// specific functionality has not been verified.
	Unverified VerificationType = "unverified"
	// Verified The Image is guaranteed to be compatible with the associated Compute Node agent SKU and all
	// Batch features have been confirmed to work as expected.
	Verified VerificationType = "verified"
)

// PossibleVerificationTypeValues returns an array of possible values for the VerificationType const type.
func PossibleVerificationTypeValues() []VerificationType {
	return []VerificationType{Unverified, Verified}
}
