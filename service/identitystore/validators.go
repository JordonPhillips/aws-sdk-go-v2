// Code generated by smithy-go-codegen DO NOT EDIT.

package identitystore

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpDescribeGroup struct {
}

func (*validateOpDescribeGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeUser struct {
}

func (*validateOpDescribeUser) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeUser) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeUserInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeUserInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListGroups struct {
}

func (*validateOpListGroups) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListGroups) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListGroupsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListGroupsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListUsers struct {
}

func (*validateOpListUsers) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListUsers) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListUsersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListUsersInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpDescribeGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeGroup{}, middleware.After)
}

func addOpDescribeUserValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeUser{}, middleware.After)
}

func addOpListGroupsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListGroups{}, middleware.After)
}

func addOpListUsersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListUsers{}, middleware.After)
}

func validateFilter(v *types.Filter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Filter"}
	if v.AttributePath == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AttributePath"))
	}
	if v.AttributeValue == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AttributeValue"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateFilters(v []types.Filter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Filters"}
	for i := range v {
		if err := validateFilter(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeGroupInput(v *DescribeGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeGroupInput"}
	if v.IdentityStoreId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityStoreId"))
	}
	if v.GroupId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GroupId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeUserInput(v *DescribeUserInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeUserInput"}
	if v.IdentityStoreId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityStoreId"))
	}
	if v.UserId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UserId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListGroupsInput(v *ListGroupsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListGroupsInput"}
	if v.IdentityStoreId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityStoreId"))
	}
	if v.Filters != nil {
		if err := validateFilters(v.Filters); err != nil {
			invalidParams.AddNested("Filters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListUsersInput(v *ListUsersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListUsersInput"}
	if v.IdentityStoreId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityStoreId"))
	}
	if v.Filters != nil {
		if err := validateFilters(v.Filters); err != nil {
			invalidParams.AddNested("Filters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
