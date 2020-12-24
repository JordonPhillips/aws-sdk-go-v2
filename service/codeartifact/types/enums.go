// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type DomainStatus string

// Enum values for DomainStatus
const (
	DomainStatusActive  DomainStatus = "Active"
	DomainStatusDeleted DomainStatus = "Deleted"
)

// Values returns all known values for DomainStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (DomainStatus) Values() []DomainStatus {
	return []DomainStatus{
		"Active",
		"Deleted",
	}
}

type ExternalConnectionStatus string

// Enum values for ExternalConnectionStatus
const (
	ExternalConnectionStatusAvailable ExternalConnectionStatus = "Available"
)

// Values returns all known values for ExternalConnectionStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExternalConnectionStatus) Values() []ExternalConnectionStatus {
	return []ExternalConnectionStatus{
		"Available",
	}
}

type HashAlgorithm string

// Enum values for HashAlgorithm
const (
	HashAlgorithmMd5    HashAlgorithm = "MD5"
	HashAlgorithmSha1   HashAlgorithm = "SHA-1"
	HashAlgorithmSha256 HashAlgorithm = "SHA-256"
	HashAlgorithmSha512 HashAlgorithm = "SHA-512"
)

// Values returns all known values for HashAlgorithm. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (HashAlgorithm) Values() []HashAlgorithm {
	return []HashAlgorithm{
		"MD5",
		"SHA-1",
		"SHA-256",
		"SHA-512",
	}
}

type PackageFormat string

// Enum values for PackageFormat
const (
	PackageFormatNpm   PackageFormat = "npm"
	PackageFormatPypi  PackageFormat = "pypi"
	PackageFormatMaven PackageFormat = "maven"
	PackageFormatNuget PackageFormat = "nuget"
)

// Values returns all known values for PackageFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PackageFormat) Values() []PackageFormat {
	return []PackageFormat{
		"npm",
		"pypi",
		"maven",
		"nuget",
	}
}

type PackageVersionErrorCode string

// Enum values for PackageVersionErrorCode
const (
	PackageVersionErrorCodeAlreadyExists      PackageVersionErrorCode = "ALREADY_EXISTS"
	PackageVersionErrorCodeMismatchedRevision PackageVersionErrorCode = "MISMATCHED_REVISION"
	PackageVersionErrorCodeMismatchedStatus   PackageVersionErrorCode = "MISMATCHED_STATUS"
	PackageVersionErrorCodeNotAllowed         PackageVersionErrorCode = "NOT_ALLOWED"
	PackageVersionErrorCodeNotFound           PackageVersionErrorCode = "NOT_FOUND"
	PackageVersionErrorCodeSkipped            PackageVersionErrorCode = "SKIPPED"
)

// Values returns all known values for PackageVersionErrorCode. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PackageVersionErrorCode) Values() []PackageVersionErrorCode {
	return []PackageVersionErrorCode{
		"ALREADY_EXISTS",
		"MISMATCHED_REVISION",
		"MISMATCHED_STATUS",
		"NOT_ALLOWED",
		"NOT_FOUND",
		"SKIPPED",
	}
}

type PackageVersionSortType string

// Enum values for PackageVersionSortType
const (
	PackageVersionSortTypePublishedTime PackageVersionSortType = "PUBLISHED_TIME"
)

// Values returns all known values for PackageVersionSortType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PackageVersionSortType) Values() []PackageVersionSortType {
	return []PackageVersionSortType{
		"PUBLISHED_TIME",
	}
}

type PackageVersionStatus string

// Enum values for PackageVersionStatus
const (
	PackageVersionStatusPublished  PackageVersionStatus = "Published"
	PackageVersionStatusUnfinished PackageVersionStatus = "Unfinished"
	PackageVersionStatusUnlisted   PackageVersionStatus = "Unlisted"
	PackageVersionStatusArchived   PackageVersionStatus = "Archived"
	PackageVersionStatusDisposed   PackageVersionStatus = "Disposed"
	PackageVersionStatusDeleted    PackageVersionStatus = "Deleted"
)

// Values returns all known values for PackageVersionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PackageVersionStatus) Values() []PackageVersionStatus {
	return []PackageVersionStatus{
		"Published",
		"Unfinished",
		"Unlisted",
		"Archived",
		"Disposed",
		"Deleted",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeDomain         ResourceType = "domain"
	ResourceTypeRepository     ResourceType = "repository"
	ResourceTypePackage        ResourceType = "package"
	ResourceTypePackageVersion ResourceType = "package-version"
	ResourceTypeAsset          ResourceType = "asset"
)

// Values returns all known values for ResourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"domain",
		"repository",
		"package",
		"package-version",
		"asset",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonCannotParse           ValidationExceptionReason = "CANNOT_PARSE"
	ValidationExceptionReasonEncryptionKeyError    ValidationExceptionReason = "ENCRYPTION_KEY_ERROR"
	ValidationExceptionReasonFieldValidationFailed ValidationExceptionReason = "FIELD_VALIDATION_FAILED"
	ValidationExceptionReasonUnknownOperation      ValidationExceptionReason = "UNKNOWN_OPERATION"
	ValidationExceptionReasonOther                 ValidationExceptionReason = "OTHER"
)

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"CANNOT_PARSE",
		"ENCRYPTION_KEY_ERROR",
		"FIELD_VALIDATION_FAILED",
		"UNKNOWN_OPERATION",
		"OTHER",
	}
}
