// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ActiveDirectoryErrorType string

// Enum values for ActiveDirectoryErrorType
const (
	ActiveDirectoryErrorTypeDomainNotFound         ActiveDirectoryErrorType = "DOMAIN_NOT_FOUND"
	ActiveDirectoryErrorTypeIncompatibleDomainMode ActiveDirectoryErrorType = "INCOMPATIBLE_DOMAIN_MODE"
	ActiveDirectoryErrorTypeWrongVpc               ActiveDirectoryErrorType = "WRONG_VPC"
	ActiveDirectoryErrorTypeInvalidDomainStage     ActiveDirectoryErrorType = "INVALID_DOMAIN_STAGE"
)

// Values returns all known values for ActiveDirectoryErrorType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ActiveDirectoryErrorType) Values() []ActiveDirectoryErrorType {
	return []ActiveDirectoryErrorType{
		"DOMAIN_NOT_FOUND",
		"INCOMPATIBLE_DOMAIN_MODE",
		"WRONG_VPC",
		"INVALID_DOMAIN_STAGE",
	}
}

type AdministrativeActionType string

// Enum values for AdministrativeActionType
const (
	AdministrativeActionTypeFileSystemUpdate              AdministrativeActionType = "FILE_SYSTEM_UPDATE"
	AdministrativeActionTypeStorageOptimization           AdministrativeActionType = "STORAGE_OPTIMIZATION"
	AdministrativeActionTypeFileSystemAliasAssociation    AdministrativeActionType = "FILE_SYSTEM_ALIAS_ASSOCIATION"
	AdministrativeActionTypeFileSystemAliasDisassociation AdministrativeActionType = "FILE_SYSTEM_ALIAS_DISASSOCIATION"
)

// Values returns all known values for AdministrativeActionType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AdministrativeActionType) Values() []AdministrativeActionType {
	return []AdministrativeActionType{
		"FILE_SYSTEM_UPDATE",
		"STORAGE_OPTIMIZATION",
		"FILE_SYSTEM_ALIAS_ASSOCIATION",
		"FILE_SYSTEM_ALIAS_DISASSOCIATION",
	}
}

type AliasLifecycle string

// Enum values for AliasLifecycle
const (
	AliasLifecycleAvailable    AliasLifecycle = "AVAILABLE"
	AliasLifecycleCreating     AliasLifecycle = "CREATING"
	AliasLifecycleDeleting     AliasLifecycle = "DELETING"
	AliasLifecycleCreateFailed AliasLifecycle = "CREATE_FAILED"
	AliasLifecycleDeleteFailed AliasLifecycle = "DELETE_FAILED"
)

// Values returns all known values for AliasLifecycle. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AliasLifecycle) Values() []AliasLifecycle {
	return []AliasLifecycle{
		"AVAILABLE",
		"CREATING",
		"DELETING",
		"CREATE_FAILED",
		"DELETE_FAILED",
	}
}

type AutoImportPolicyType string

// Enum values for AutoImportPolicyType
const (
	AutoImportPolicyTypeNone       AutoImportPolicyType = "NONE"
	AutoImportPolicyTypeNew        AutoImportPolicyType = "NEW"
	AutoImportPolicyTypeNewChanged AutoImportPolicyType = "NEW_CHANGED"
)

// Values returns all known values for AutoImportPolicyType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AutoImportPolicyType) Values() []AutoImportPolicyType {
	return []AutoImportPolicyType{
		"NONE",
		"NEW",
		"NEW_CHANGED",
	}
}

type BackupLifecycle string

// Enum values for BackupLifecycle
const (
	BackupLifecycleAvailable    BackupLifecycle = "AVAILABLE"
	BackupLifecycleCreating     BackupLifecycle = "CREATING"
	BackupLifecycleTransferring BackupLifecycle = "TRANSFERRING"
	BackupLifecycleDeleted      BackupLifecycle = "DELETED"
	BackupLifecycleFailed       BackupLifecycle = "FAILED"
	BackupLifecyclePending      BackupLifecycle = "PENDING"
)

// Values returns all known values for BackupLifecycle. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BackupLifecycle) Values() []BackupLifecycle {
	return []BackupLifecycle{
		"AVAILABLE",
		"CREATING",
		"TRANSFERRING",
		"DELETED",
		"FAILED",
		"PENDING",
	}
}

type BackupType string

// Enum values for BackupType
const (
	BackupTypeAutomatic     BackupType = "AUTOMATIC"
	BackupTypeUserInitiated BackupType = "USER_INITIATED"
	BackupTypeAwsBackup     BackupType = "AWS_BACKUP"
)

// Values returns all known values for BackupType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (BackupType) Values() []BackupType {
	return []BackupType{
		"AUTOMATIC",
		"USER_INITIATED",
		"AWS_BACKUP",
	}
}

type DataRepositoryLifecycle string

// Enum values for DataRepositoryLifecycle
const (
	DataRepositoryLifecycleCreating      DataRepositoryLifecycle = "CREATING"
	DataRepositoryLifecycleAvailable     DataRepositoryLifecycle = "AVAILABLE"
	DataRepositoryLifecycleMisconfigured DataRepositoryLifecycle = "MISCONFIGURED"
	DataRepositoryLifecycleUpdating      DataRepositoryLifecycle = "UPDATING"
	DataRepositoryLifecycleDeleting      DataRepositoryLifecycle = "DELETING"
)

// Values returns all known values for DataRepositoryLifecycle. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DataRepositoryLifecycle) Values() []DataRepositoryLifecycle {
	return []DataRepositoryLifecycle{
		"CREATING",
		"AVAILABLE",
		"MISCONFIGURED",
		"UPDATING",
		"DELETING",
	}
}

type DataRepositoryTaskFilterName string

// Enum values for DataRepositoryTaskFilterName
const (
	DataRepositoryTaskFilterNameFileSystemId  DataRepositoryTaskFilterName = "file-system-id"
	DataRepositoryTaskFilterNameTaskLifecycle DataRepositoryTaskFilterName = "task-lifecycle"
)

// Values returns all known values for DataRepositoryTaskFilterName. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (DataRepositoryTaskFilterName) Values() []DataRepositoryTaskFilterName {
	return []DataRepositoryTaskFilterName{
		"file-system-id",
		"task-lifecycle",
	}
}

type DataRepositoryTaskLifecycle string

// Enum values for DataRepositoryTaskLifecycle
const (
	DataRepositoryTaskLifecyclePending   DataRepositoryTaskLifecycle = "PENDING"
	DataRepositoryTaskLifecycleExecuting DataRepositoryTaskLifecycle = "EXECUTING"
	DataRepositoryTaskLifecycleFailed    DataRepositoryTaskLifecycle = "FAILED"
	DataRepositoryTaskLifecycleSucceeded DataRepositoryTaskLifecycle = "SUCCEEDED"
	DataRepositoryTaskLifecycleCanceled  DataRepositoryTaskLifecycle = "CANCELED"
	DataRepositoryTaskLifecycleCanceling DataRepositoryTaskLifecycle = "CANCELING"
)

// Values returns all known values for DataRepositoryTaskLifecycle. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (DataRepositoryTaskLifecycle) Values() []DataRepositoryTaskLifecycle {
	return []DataRepositoryTaskLifecycle{
		"PENDING",
		"EXECUTING",
		"FAILED",
		"SUCCEEDED",
		"CANCELED",
		"CANCELING",
	}
}

type DataRepositoryTaskType string

// Enum values for DataRepositoryTaskType
const (
	DataRepositoryTaskTypeExport DataRepositoryTaskType = "EXPORT_TO_REPOSITORY"
)

// Values returns all known values for DataRepositoryTaskType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DataRepositoryTaskType) Values() []DataRepositoryTaskType {
	return []DataRepositoryTaskType{
		"EXPORT_TO_REPOSITORY",
	}
}

type DriveCacheType string

// Enum values for DriveCacheType
const (
	DriveCacheTypeNone DriveCacheType = "NONE"
	DriveCacheTypeRead DriveCacheType = "READ"
)

// Values returns all known values for DriveCacheType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DriveCacheType) Values() []DriveCacheType {
	return []DriveCacheType{
		"NONE",
		"READ",
	}
}

type FileSystemLifecycle string

// Enum values for FileSystemLifecycle
const (
	FileSystemLifecycleAvailable     FileSystemLifecycle = "AVAILABLE"
	FileSystemLifecycleCreating      FileSystemLifecycle = "CREATING"
	FileSystemLifecycleFailed        FileSystemLifecycle = "FAILED"
	FileSystemLifecycleDeleting      FileSystemLifecycle = "DELETING"
	FileSystemLifecycleMisconfigured FileSystemLifecycle = "MISCONFIGURED"
	FileSystemLifecycleUpdating      FileSystemLifecycle = "UPDATING"
)

// Values returns all known values for FileSystemLifecycle. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FileSystemLifecycle) Values() []FileSystemLifecycle {
	return []FileSystemLifecycle{
		"AVAILABLE",
		"CREATING",
		"FAILED",
		"DELETING",
		"MISCONFIGURED",
		"UPDATING",
	}
}

type FileSystemMaintenanceOperation string

// Enum values for FileSystemMaintenanceOperation
const (
	FileSystemMaintenanceOperationPatching  FileSystemMaintenanceOperation = "PATCHING"
	FileSystemMaintenanceOperationBackingUp FileSystemMaintenanceOperation = "BACKING_UP"
)

// Values returns all known values for FileSystemMaintenanceOperation. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (FileSystemMaintenanceOperation) Values() []FileSystemMaintenanceOperation {
	return []FileSystemMaintenanceOperation{
		"PATCHING",
		"BACKING_UP",
	}
}

type FileSystemType string

// Enum values for FileSystemType
const (
	FileSystemTypeWindows FileSystemType = "WINDOWS"
	FileSystemTypeLustre  FileSystemType = "LUSTRE"
)

// Values returns all known values for FileSystemType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FileSystemType) Values() []FileSystemType {
	return []FileSystemType{
		"WINDOWS",
		"LUSTRE",
	}
}

type FilterName string

// Enum values for FilterName
const (
	FilterNameFileSystemId   FilterName = "file-system-id"
	FilterNameBackupType     FilterName = "backup-type"
	FilterNameFileSystemType FilterName = "file-system-type"
)

// Values returns all known values for FilterName. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (FilterName) Values() []FilterName {
	return []FilterName{
		"file-system-id",
		"backup-type",
		"file-system-type",
	}
}

type LustreDeploymentType string

// Enum values for LustreDeploymentType
const (
	LustreDeploymentTypeScratch1    LustreDeploymentType = "SCRATCH_1"
	LustreDeploymentTypeScratch2    LustreDeploymentType = "SCRATCH_2"
	LustreDeploymentTypePersistent1 LustreDeploymentType = "PERSISTENT_1"
)

// Values returns all known values for LustreDeploymentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LustreDeploymentType) Values() []LustreDeploymentType {
	return []LustreDeploymentType{
		"SCRATCH_1",
		"SCRATCH_2",
		"PERSISTENT_1",
	}
}

type ReportFormat string

// Enum values for ReportFormat
const (
	ReportFormatReportCsv20191124 ReportFormat = "REPORT_CSV_20191124"
)

// Values returns all known values for ReportFormat. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ReportFormat) Values() []ReportFormat {
	return []ReportFormat{
		"REPORT_CSV_20191124",
	}
}

type ReportScope string

// Enum values for ReportScope
const (
	ReportScopeFailedFilesOnly ReportScope = "FAILED_FILES_ONLY"
)

// Values returns all known values for ReportScope. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ReportScope) Values() []ReportScope {
	return []ReportScope{
		"FAILED_FILES_ONLY",
	}
}

type ServiceLimit string

// Enum values for ServiceLimit
const (
	ServiceLimitFileSystemCount           ServiceLimit = "FILE_SYSTEM_COUNT"
	ServiceLimitTotalThroughputCapacity   ServiceLimit = "TOTAL_THROUGHPUT_CAPACITY"
	ServiceLimitTotalStorage              ServiceLimit = "TOTAL_STORAGE"
	ServiceLimitTotalUserInitiatedBackups ServiceLimit = "TOTAL_USER_INITIATED_BACKUPS"
)

// Values returns all known values for ServiceLimit. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ServiceLimit) Values() []ServiceLimit {
	return []ServiceLimit{
		"FILE_SYSTEM_COUNT",
		"TOTAL_THROUGHPUT_CAPACITY",
		"TOTAL_STORAGE",
		"TOTAL_USER_INITIATED_BACKUPS",
	}
}

type Status string

// Enum values for Status
const (
	StatusFailed            Status = "FAILED"
	StatusInProgress        Status = "IN_PROGRESS"
	StatusPending           Status = "PENDING"
	StatusCompleted         Status = "COMPLETED"
	StatusUpdatedOptimizing Status = "UPDATED_OPTIMIZING"
)

// Values returns all known values for Status. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Status) Values() []Status {
	return []Status{
		"FAILED",
		"IN_PROGRESS",
		"PENDING",
		"COMPLETED",
		"UPDATED_OPTIMIZING",
	}
}

type StorageType string

// Enum values for StorageType
const (
	StorageTypeSsd StorageType = "SSD"
	StorageTypeHdd StorageType = "HDD"
)

// Values returns all known values for StorageType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (StorageType) Values() []StorageType {
	return []StorageType{
		"SSD",
		"HDD",
	}
}

type WindowsDeploymentType string

// Enum values for WindowsDeploymentType
const (
	WindowsDeploymentTypeMultiAz1  WindowsDeploymentType = "MULTI_AZ_1"
	WindowsDeploymentTypeSingleAz1 WindowsDeploymentType = "SINGLE_AZ_1"
	WindowsDeploymentTypeSingleAz2 WindowsDeploymentType = "SINGLE_AZ_2"
)

// Values returns all known values for WindowsDeploymentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WindowsDeploymentType) Values() []WindowsDeploymentType {
	return []WindowsDeploymentType{
		"MULTI_AZ_1",
		"SINGLE_AZ_1",
		"SINGLE_AZ_2",
	}
}
