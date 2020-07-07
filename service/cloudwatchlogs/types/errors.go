// Code generated by smithy-go-codegen DO NOT EDIT.
package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// The event was already logged.
type DataAlreadyAcceptedException struct {
	Message *string

	ExpectedSequenceToken *string
}

func (e *DataAlreadyAcceptedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DataAlreadyAcceptedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DataAlreadyAcceptedException) ErrorCode() string             { return "DataAlreadyAcceptedException" }
func (e *DataAlreadyAcceptedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *DataAlreadyAcceptedException) GetExpectedSequenceToken() string {
	return ptr.ToString(e.ExpectedSequenceToken)
}
func (e *DataAlreadyAcceptedException) HasExpectedSequenceToken() bool {
	return e.ExpectedSequenceToken != nil
}
func (e *DataAlreadyAcceptedException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *DataAlreadyAcceptedException) HasMessage() bool {
	return e.Message != nil
}

// The operation is not valid on the specified resource.
type InvalidOperationException struct {
	Message *string
}

func (e *InvalidOperationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidOperationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidOperationException) ErrorCode() string             { return "InvalidOperationException" }
func (e *InvalidOperationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidOperationException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidOperationException) HasMessage() bool {
	return e.Message != nil
}

// A parameter is specified incorrectly.
type InvalidParameterException struct {
	Message *string
}

func (e *InvalidParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterException) ErrorCode() string             { return "InvalidParameterException" }
func (e *InvalidParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidParameterException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidParameterException) HasMessage() bool {
	return e.Message != nil
}

// The sequence token is not valid. You can get the correct sequence token in the
// expectedSequenceToken field in the InvalidSequenceTokenException message.
type InvalidSequenceTokenException struct {
	Message *string

	ExpectedSequenceToken *string
}

func (e *InvalidSequenceTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSequenceTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSequenceTokenException) ErrorCode() string             { return "InvalidSequenceTokenException" }
func (e *InvalidSequenceTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidSequenceTokenException) GetExpectedSequenceToken() string {
	return ptr.ToString(e.ExpectedSequenceToken)
}
func (e *InvalidSequenceTokenException) HasExpectedSequenceToken() bool {
	return e.ExpectedSequenceToken != nil
}
func (e *InvalidSequenceTokenException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidSequenceTokenException) HasMessage() bool {
	return e.Message != nil
}

// You have reached the maximum number of resources that can be created.
type LimitExceededException struct {
	Message *string
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
func (e *LimitExceededException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *LimitExceededException) HasMessage() bool {
	return e.Message != nil
}

// The query string is not valid. Details about this error are displayed in a
// QueryCompileError object. For more information, see . For more information about
// valid query syntax, see CloudWatch Logs Insights Query Syntax
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html).
type MalformedQueryException struct {
	Message *string

	QueryCompileError *QueryCompileError
}

func (e *MalformedQueryException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MalformedQueryException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MalformedQueryException) ErrorCode() string             { return "MalformedQueryException" }
func (e *MalformedQueryException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *MalformedQueryException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *MalformedQueryException) HasMessage() bool {
	return e.Message != nil
}
func (e *MalformedQueryException) GetQueryCompileError() *QueryCompileError {
	return e.QueryCompileError
}
func (e *MalformedQueryException) HasQueryCompileError() bool {
	return e.QueryCompileError != nil
}

// Multiple requests to update the same resource were in conflict.
type OperationAbortedException struct {
	Message *string
}

func (e *OperationAbortedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OperationAbortedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OperationAbortedException) ErrorCode() string             { return "OperationAbortedException" }
func (e *OperationAbortedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *OperationAbortedException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *OperationAbortedException) HasMessage() bool {
	return e.Message != nil
}

// The specified resource already exists.
type ResourceAlreadyExistsException struct {
	Message *string
}

func (e *ResourceAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceAlreadyExistsException) ErrorCode() string             { return "ResourceAlreadyExistsException" }
func (e *ResourceAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceAlreadyExistsException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceAlreadyExistsException) HasMessage() bool {
	return e.Message != nil
}

// The specified resource does not exist.
type ResourceNotFoundException struct {
	Message *string
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
func (e *ResourceNotFoundException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceNotFoundException) HasMessage() bool {
	return e.Message != nil
}

// The service cannot complete the request.
type ServiceUnavailableException struct {
	Message *string
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
func (e *ServiceUnavailableException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ServiceUnavailableException) HasMessage() bool {
	return e.Message != nil
}

// The most likely cause is an invalid AWS access key ID or secret key.
type UnrecognizedClientException struct {
	Message *string
}

func (e *UnrecognizedClientException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnrecognizedClientException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnrecognizedClientException) ErrorCode() string             { return "UnrecognizedClientException" }
func (e *UnrecognizedClientException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *UnrecognizedClientException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *UnrecognizedClientException) HasMessage() bool {
	return e.Message != nil
}
