// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationinsights

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/applicationinsights/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateApplication struct {
}

func (*validateOpCreateApplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateApplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateComponent struct {
}

func (*validateOpCreateComponent) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateComponent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateComponentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateComponentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateLogPattern struct {
}

func (*validateOpCreateLogPattern) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateLogPattern) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateLogPatternInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateLogPatternInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteApplication struct {
}

func (*validateOpDeleteApplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteApplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteComponent struct {
}

func (*validateOpDeleteComponent) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteComponent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteComponentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteComponentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteLogPattern struct {
}

func (*validateOpDeleteLogPattern) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteLogPattern) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteLogPatternInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteLogPatternInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeApplication struct {
}

func (*validateOpDescribeApplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeApplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeComponentConfiguration struct {
}

func (*validateOpDescribeComponentConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeComponentConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeComponentConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeComponentConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeComponentConfigurationRecommendation struct {
}

func (*validateOpDescribeComponentConfigurationRecommendation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeComponentConfigurationRecommendation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeComponentConfigurationRecommendationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeComponentConfigurationRecommendationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeComponent struct {
}

func (*validateOpDescribeComponent) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeComponent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeComponentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeComponentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeLogPattern struct {
}

func (*validateOpDescribeLogPattern) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeLogPattern) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeLogPatternInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeLogPatternInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeObservation struct {
}

func (*validateOpDescribeObservation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeObservation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeObservationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeObservationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeProblem struct {
}

func (*validateOpDescribeProblem) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeProblem) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeProblemInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeProblemInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeProblemObservations struct {
}

func (*validateOpDescribeProblemObservations) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeProblemObservations) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeProblemObservationsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeProblemObservationsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListComponents struct {
}

func (*validateOpListComponents) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListComponents) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListComponentsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListComponentsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListLogPatternSets struct {
}

func (*validateOpListLogPatternSets) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListLogPatternSets) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListLogPatternSetsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListLogPatternSetsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListLogPatterns struct {
}

func (*validateOpListLogPatterns) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListLogPatterns) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListLogPatternsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListLogPatternsInput(input); err != nil {
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

type validateOpUpdateApplication struct {
}

func (*validateOpUpdateApplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateApplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateComponentConfiguration struct {
}

func (*validateOpUpdateComponentConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateComponentConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateComponentConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateComponentConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateComponent struct {
}

func (*validateOpUpdateComponent) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateComponent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateComponentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateComponentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateLogPattern struct {
}

func (*validateOpUpdateLogPattern) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateLogPattern) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateLogPatternInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateLogPatternInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateApplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateApplication{}, middleware.After)
}

func addOpCreateComponentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateComponent{}, middleware.After)
}

func addOpCreateLogPatternValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateLogPattern{}, middleware.After)
}

func addOpDeleteApplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteApplication{}, middleware.After)
}

func addOpDeleteComponentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteComponent{}, middleware.After)
}

func addOpDeleteLogPatternValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteLogPattern{}, middleware.After)
}

func addOpDescribeApplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeApplication{}, middleware.After)
}

func addOpDescribeComponentConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeComponentConfiguration{}, middleware.After)
}

func addOpDescribeComponentConfigurationRecommendationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeComponentConfigurationRecommendation{}, middleware.After)
}

func addOpDescribeComponentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeComponent{}, middleware.After)
}

func addOpDescribeLogPatternValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeLogPattern{}, middleware.After)
}

func addOpDescribeObservationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeObservation{}, middleware.After)
}

func addOpDescribeProblemValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeProblem{}, middleware.After)
}

func addOpDescribeProblemObservationsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeProblemObservations{}, middleware.After)
}

func addOpListComponentsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListComponents{}, middleware.After)
}

func addOpListLogPatternSetsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListLogPatternSets{}, middleware.After)
}

func addOpListLogPatternsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListLogPatterns{}, middleware.After)
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

func addOpUpdateApplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateApplication{}, middleware.After)
}

func addOpUpdateComponentConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateComponentConfiguration{}, middleware.After)
}

func addOpUpdateComponentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateComponent{}, middleware.After)
}

func addOpUpdateLogPatternValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateLogPattern{}, middleware.After)
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

func validateTagList(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagList"}
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

func validateOpCreateApplicationInput(v *CreateApplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateApplicationInput"}
	if v.ResourceGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceGroupName"))
	}
	if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateComponentInput(v *CreateComponentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateComponentInput"}
	if v.ResourceGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceGroupName"))
	}
	if v.ComponentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ComponentName"))
	}
	if v.ResourceList == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceList"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateLogPatternInput(v *CreateLogPatternInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateLogPatternInput"}
	if v.ResourceGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceGroupName"))
	}
	if v.PatternSetName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PatternSetName"))
	}
	if v.PatternName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PatternName"))
	}
	if v.Pattern == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Pattern"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteApplicationInput(v *DeleteApplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteApplicationInput"}
	if v.ResourceGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceGroupName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteComponentInput(v *DeleteComponentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteComponentInput"}
	if v.ResourceGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceGroupName"))
	}
	if v.ComponentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ComponentName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteLogPatternInput(v *DeleteLogPatternInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteLogPatternInput"}
	if v.ResourceGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceGroupName"))
	}
	if v.PatternSetName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PatternSetName"))
	}
	if v.PatternName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PatternName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeApplicationInput(v *DescribeApplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeApplicationInput"}
	if v.ResourceGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceGroupName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeComponentConfigurationInput(v *DescribeComponentConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeComponentConfigurationInput"}
	if v.ResourceGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceGroupName"))
	}
	if v.ComponentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ComponentName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeComponentConfigurationRecommendationInput(v *DescribeComponentConfigurationRecommendationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeComponentConfigurationRecommendationInput"}
	if v.ResourceGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceGroupName"))
	}
	if v.ComponentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ComponentName"))
	}
	if len(v.Tier) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Tier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeComponentInput(v *DescribeComponentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeComponentInput"}
	if v.ResourceGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceGroupName"))
	}
	if v.ComponentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ComponentName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeLogPatternInput(v *DescribeLogPatternInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeLogPatternInput"}
	if v.ResourceGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceGroupName"))
	}
	if v.PatternSetName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PatternSetName"))
	}
	if v.PatternName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PatternName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeObservationInput(v *DescribeObservationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeObservationInput"}
	if v.ObservationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ObservationId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeProblemInput(v *DescribeProblemInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeProblemInput"}
	if v.ProblemId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProblemId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeProblemObservationsInput(v *DescribeProblemObservationsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeProblemObservationsInput"}
	if v.ProblemId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProblemId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListComponentsInput(v *ListComponentsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListComponentsInput"}
	if v.ResourceGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceGroupName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListLogPatternSetsInput(v *ListLogPatternSetsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListLogPatternSetsInput"}
	if v.ResourceGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceGroupName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListLogPatternsInput(v *ListLogPatternsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListLogPatternsInput"}
	if v.ResourceGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceGroupName"))
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
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
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
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	} else if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
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
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
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

func validateOpUpdateApplicationInput(v *UpdateApplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateApplicationInput"}
	if v.ResourceGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceGroupName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateComponentConfigurationInput(v *UpdateComponentConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateComponentConfigurationInput"}
	if v.ResourceGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceGroupName"))
	}
	if v.ComponentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ComponentName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateComponentInput(v *UpdateComponentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateComponentInput"}
	if v.ResourceGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceGroupName"))
	}
	if v.ComponentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ComponentName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateLogPatternInput(v *UpdateLogPatternInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateLogPatternInput"}
	if v.ResourceGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceGroupName"))
	}
	if v.PatternSetName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PatternSetName"))
	}
	if v.PatternName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PatternName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
