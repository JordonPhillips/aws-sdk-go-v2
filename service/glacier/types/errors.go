// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Returned if there is insufficient capacity to process this expedited request.
// This error only applies to expedited retrievals and not to standard or bulk
// retrievals.
type InsufficientCapacityException struct {
	Message *string

	Type *string
	Code *string
}

func (e *InsufficientCapacityException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InsufficientCapacityException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InsufficientCapacityException) ErrorCode() string             { return "InsufficientCapacityException" }
func (e *InsufficientCapacityException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if a parameter of the request is incorrectly specified.
type InvalidParameterValueException struct {
	Message *string

	Type *string
	Code *string
}

func (e *InvalidParameterValueException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterValueException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterValueException) ErrorCode() string             { return "InvalidParameterValueException" }
func (e *InvalidParameterValueException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if the request results in a vault or account limit being exceeded.
type LimitExceededException struct {
	Message *string

	Type *string
	Code *string
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if a required header or parameter is missing from the request.
type MissingParameterValueException struct {
	Message *string

	Type *string
	Code *string
}

func (e *MissingParameterValueException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MissingParameterValueException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MissingParameterValueException) ErrorCode() string             { return "MissingParameterValueException" }
func (e *MissingParameterValueException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if a retrieval job would exceed the current data policy's retrieval
// rate limit. For more information about data retrieval policies,
type PolicyEnforcedException struct {
	Message *string

	Type *string
	Code *string
}

func (e *PolicyEnforcedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PolicyEnforcedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PolicyEnforcedException) ErrorCode() string             { return "PolicyEnforcedException" }
func (e *PolicyEnforcedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if, when uploading an archive, Amazon S3 Glacier times out while
// receiving the upload.
type RequestTimeoutException struct {
	Message *string

	Type *string
	Code *string
}

func (e *RequestTimeoutException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RequestTimeoutException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RequestTimeoutException) ErrorCode() string             { return "RequestTimeoutException" }
func (e *RequestTimeoutException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if the specified resource (such as a vault, upload ID, or job ID)
// doesn't exist.
type ResourceNotFoundException struct {
	Message *string

	Type *string
	Code *string
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Returned if the service cannot complete the request.
type ServiceUnavailableException struct {
	Message *string

	Type *string
	Code *string
}

func (e *ServiceUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceUnavailableException) ErrorCode() string             { return "ServiceUnavailableException" }
func (e *ServiceUnavailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }
