// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Lightsail throws this exception when the user cannot be authenticated or uses
// invalid credentials to access a resource.
type AccessDeniedException struct {
	Message *string

	Code *string
	Docs *string
	Tip  *string
}

func (e *AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessDeniedException) ErrorCode() string             { return "AccessDeniedException" }
func (e *AccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Lightsail throws this exception when an account is still in the setup in
// progress state.
type AccountSetupInProgressException struct {
	Message *string

	Code *string
	Docs *string
	Tip  *string
}

func (e *AccountSetupInProgressException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccountSetupInProgressException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccountSetupInProgressException) ErrorCode() string {
	return "AccountSetupInProgressException"
}
func (e *AccountSetupInProgressException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Lightsail throws this exception when user input does not conform to the
// validation rules of an input field. Domain-related APIs are only available in
// the N. Virginia (us-east-1) Region. Please set your AWS Region configuration to
// us-east-1 to create, view, or edit these resources.
type InvalidInputException struct {
	Message *string

	Code *string
	Docs *string
	Tip  *string
}

func (e *InvalidInputException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidInputException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidInputException) ErrorCode() string             { return "InvalidInputException" }
func (e *InvalidInputException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Lightsail throws this exception when it cannot find a resource.
type NotFoundException struct {
	Message *string

	Code *string
	Docs *string
	Tip  *string
}

func (e *NotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotFoundException) ErrorCode() string             { return "NotFoundException" }
func (e *NotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Lightsail throws this exception when an operation fails to execute.
type OperationFailureException struct {
	Message *string

	Code *string
	Docs *string
	Tip  *string
}

func (e *OperationFailureException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OperationFailureException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OperationFailureException) ErrorCode() string             { return "OperationFailureException" }
func (e *OperationFailureException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A general service exception.
type ServiceException struct {
	Message *string

	Code *string
	Docs *string
	Tip  *string
}

func (e *ServiceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceException) ErrorCode() string             { return "ServiceException" }
func (e *ServiceException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Lightsail throws this exception when the user has not been authenticated.
type UnauthenticatedException struct {
	Message *string

	Code *string
	Docs *string
	Tip  *string
}

func (e *UnauthenticatedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnauthenticatedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnauthenticatedException) ErrorCode() string             { return "UnauthenticatedException" }
func (e *UnauthenticatedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
