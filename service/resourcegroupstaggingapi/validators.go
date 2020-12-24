// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroupstaggingapi

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpGetTagValues struct {
}

func (*validateOpGetTagValues) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetTagValues) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetTagValuesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetTagValuesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartReportCreation struct {
}

func (*validateOpStartReportCreation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartReportCreation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartReportCreationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartReportCreationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResources struct {
}

func (*validateOpTagResources) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResources) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourcesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourcesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResources struct {
}

func (*validateOpUntagResources) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResources) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourcesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourcesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpGetTagValuesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetTagValues{}, middleware.After)
}

func addOpStartReportCreationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartReportCreation{}, middleware.After)
}

func addOpTagResourcesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResources{}, middleware.After)
}

func addOpUntagResourcesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResources{}, middleware.After)
}

func validateOpGetTagValuesInput(v *GetTagValuesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetTagValuesInput"}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartReportCreationInput(v *StartReportCreationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartReportCreationInput"}
	if v.S3Bucket == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Bucket"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourcesInput(v *TagResourcesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourcesInput"}
	if v.ResourceARNList == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARNList"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourcesInput(v *UntagResourcesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourcesInput"}
	if v.ResourceARNList == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARNList"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
