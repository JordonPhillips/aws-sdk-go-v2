// Code generated by smithy-go-codegen DO NOT EDIT.

package dlm

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dlm/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateLifecyclePolicy struct {
}

func (*validateOpCreateLifecyclePolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateLifecyclePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateLifecyclePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateLifecyclePolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteLifecyclePolicy struct {
}

func (*validateOpDeleteLifecyclePolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteLifecyclePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteLifecyclePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteLifecyclePolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetLifecyclePolicy struct {
}

func (*validateOpGetLifecyclePolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetLifecyclePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetLifecyclePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetLifecyclePolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateLifecyclePolicy struct {
}

func (*validateOpUpdateLifecyclePolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateLifecyclePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateLifecyclePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateLifecyclePolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateLifecyclePolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateLifecyclePolicy{}, middleware.After)
}

func addOpDeleteLifecyclePolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteLifecyclePolicy{}, middleware.After)
}

func addOpGetLifecyclePolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetLifecyclePolicy{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateLifecyclePolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateLifecyclePolicy{}, middleware.After)
}

func validateAction(v *types.Action) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Action"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.CrossRegionCopy == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CrossRegionCopy"))
	} else if v.CrossRegionCopy != nil {
		if err := validateCrossRegionCopyActionList(v.CrossRegionCopy); err != nil {
			invalidParams.AddNested("CrossRegionCopy", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateActionList(v []types.Action) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ActionList"}
	for i := range v {
		if err := validateAction(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCrossRegionCopyAction(v *types.CrossRegionCopyAction) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CrossRegionCopyAction"}
	if v.Target == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Target"))
	}
	if v.EncryptionConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EncryptionConfiguration"))
	} else if v.EncryptionConfiguration != nil {
		if err := validateEncryptionConfiguration(v.EncryptionConfiguration); err != nil {
			invalidParams.AddNested("EncryptionConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCrossRegionCopyActionList(v []types.CrossRegionCopyAction) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CrossRegionCopyActionList"}
	for i := range v {
		if err := validateCrossRegionCopyAction(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCrossRegionCopyRule(v *types.CrossRegionCopyRule) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CrossRegionCopyRule"}
	if v.TargetRegion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TargetRegion"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCrossRegionCopyRules(v []types.CrossRegionCopyRule) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CrossRegionCopyRules"}
	for i := range v {
		if err := validateCrossRegionCopyRule(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEncryptionConfiguration(v *types.EncryptionConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EncryptionConfiguration"}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEventParameters(v *types.EventParameters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EventParameters"}
	if len(v.EventType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("EventType"))
	}
	if v.SnapshotOwner == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SnapshotOwner"))
	}
	if v.DescriptionRegex == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DescriptionRegex"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEventSource(v *types.EventSource) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EventSource"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.Parameters != nil {
		if err := validateEventParameters(v.Parameters); err != nil {
			invalidParams.AddNested("Parameters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateFastRestoreRule(v *types.FastRestoreRule) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "FastRestoreRule"}
	if v.AvailabilityZones == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AvailabilityZones"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePolicyDetails(v *types.PolicyDetails) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PolicyDetails"}
	if v.TargetTags != nil {
		if err := validateTargetTagList(v.TargetTags); err != nil {
			invalidParams.AddNested("TargetTags", err.(smithy.InvalidParamsError))
		}
	}
	if v.Schedules != nil {
		if err := validateScheduleList(v.Schedules); err != nil {
			invalidParams.AddNested("Schedules", err.(smithy.InvalidParamsError))
		}
	}
	if v.EventSource != nil {
		if err := validateEventSource(v.EventSource); err != nil {
			invalidParams.AddNested("EventSource", err.(smithy.InvalidParamsError))
		}
	}
	if v.Actions != nil {
		if err := validateActionList(v.Actions); err != nil {
			invalidParams.AddNested("Actions", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSchedule(v *types.Schedule) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Schedule"}
	if v.TagsToAdd != nil {
		if err := validateTagsToAddList(v.TagsToAdd); err != nil {
			invalidParams.AddNested("TagsToAdd", err.(smithy.InvalidParamsError))
		}
	}
	if v.VariableTags != nil {
		if err := validateVariableTagsList(v.VariableTags); err != nil {
			invalidParams.AddNested("VariableTags", err.(smithy.InvalidParamsError))
		}
	}
	if v.FastRestoreRule != nil {
		if err := validateFastRestoreRule(v.FastRestoreRule); err != nil {
			invalidParams.AddNested("FastRestoreRule", err.(smithy.InvalidParamsError))
		}
	}
	if v.CrossRegionCopyRules != nil {
		if err := validateCrossRegionCopyRules(v.CrossRegionCopyRules); err != nil {
			invalidParams.AddNested("CrossRegionCopyRules", err.(smithy.InvalidParamsError))
		}
	}
	if v.ShareRules != nil {
		if err := validateShareRules(v.ShareRules); err != nil {
			invalidParams.AddNested("ShareRules", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateScheduleList(v []types.Schedule) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ScheduleList"}
	for i := range v {
		if err := validateSchedule(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateShareRule(v *types.ShareRule) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ShareRule"}
	if v.TargetAccounts == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TargetAccounts"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateShareRules(v []types.ShareRule) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ShareRules"}
	for i := range v {
		if err := validateShareRule(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTag(v *types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Tag"}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTagsToAddList(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagsToAddList"}
	for i := range v {
		if err := validateTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTargetTagList(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TargetTagList"}
	for i := range v {
		if err := validateTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateVariableTagsList(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "VariableTagsList"}
	for i := range v {
		if err := validateTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateLifecyclePolicyInput(v *CreateLifecyclePolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateLifecyclePolicyInput"}
	if v.ExecutionRoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExecutionRoleArn"))
	}
	if v.Description == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Description"))
	}
	if len(v.State) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("State"))
	}
	if v.PolicyDetails == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyDetails"))
	} else if v.PolicyDetails != nil {
		if err := validatePolicyDetails(v.PolicyDetails); err != nil {
			invalidParams.AddNested("PolicyDetails", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteLifecyclePolicyInput(v *DeleteLifecyclePolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteLifecyclePolicyInput"}
	if v.PolicyId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetLifecyclePolicyInput(v *GetLifecyclePolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetLifecyclePolicyInput"}
	if v.PolicyId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
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

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
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

func validateOpUpdateLifecyclePolicyInput(v *UpdateLifecyclePolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateLifecyclePolicyInput"}
	if v.PolicyId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyId"))
	}
	if v.PolicyDetails != nil {
		if err := validatePolicyDetails(v.PolicyDetails); err != nil {
			invalidParams.AddNested("PolicyDetails", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
