// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The requested bucket name is not available. The bucket namespace is shared by
// all users of the system. Select a different name and try again.
type BucketAlreadyExists struct {
	Message *string
}

func (e *BucketAlreadyExists) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BucketAlreadyExists) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BucketAlreadyExists) ErrorCode() string             { return "BucketAlreadyExists" }
func (e *BucketAlreadyExists) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The bucket you tried to create already exists, and you own it. Amazon S3 returns
// this error in all AWS Regions except in the North Virginia Region. For legacy
// compatibility, if you re-create an existing bucket that you already own in the
// North Virginia Region, Amazon S3 returns 200 OK and resets the bucket access
// control lists (ACLs).
type BucketAlreadyOwnedByYou struct {
	Message *string
}

func (e *BucketAlreadyOwnedByYou) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BucketAlreadyOwnedByYou) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BucketAlreadyOwnedByYou) ErrorCode() string             { return "BucketAlreadyOwnedByYou" }
func (e *BucketAlreadyOwnedByYou) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Object is archived and inaccessible until restored.
type InvalidObjectState struct {
	Message *string

	StorageClass StorageClass
	AccessTier   IntelligentTieringAccessTier
}

func (e *InvalidObjectState) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidObjectState) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidObjectState) ErrorCode() string             { return "InvalidObjectState" }
func (e *InvalidObjectState) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified bucket does not exist.
type NoSuchBucket struct {
	Message *string
}

func (e *NoSuchBucket) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NoSuchBucket) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NoSuchBucket) ErrorCode() string             { return "NoSuchBucket" }
func (e *NoSuchBucket) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified key does not exist.
type NoSuchKey struct {
	Message *string
}

func (e *NoSuchKey) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NoSuchKey) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NoSuchKey) ErrorCode() string             { return "NoSuchKey" }
func (e *NoSuchKey) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified multipart upload does not exist.
type NoSuchUpload struct {
	Message *string
}

func (e *NoSuchUpload) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NoSuchUpload) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NoSuchUpload) ErrorCode() string             { return "NoSuchUpload" }
func (e *NoSuchUpload) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This operation is not allowed against this storage tier.
type ObjectAlreadyInActiveTierError struct {
	Message *string
}

func (e *ObjectAlreadyInActiveTierError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ObjectAlreadyInActiveTierError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ObjectAlreadyInActiveTierError) ErrorCode() string             { return "ObjectAlreadyInActiveTierError" }
func (e *ObjectAlreadyInActiveTierError) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The source object of the COPY operation is not in the active tier and is only
// stored in Amazon S3 Glacier.
type ObjectNotInActiveTierError struct {
	Message *string
}

func (e *ObjectNotInActiveTierError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ObjectNotInActiveTierError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ObjectNotInActiveTierError) ErrorCode() string             { return "ObjectNotInActiveTierError" }
func (e *ObjectNotInActiveTierError) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
