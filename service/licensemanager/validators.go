// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanager

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpCreateLicenseConfiguration struct {
}

func (*validateOpCreateLicenseConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateLicenseConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateLicenseConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateLicenseConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteLicenseConfiguration struct {
}

func (*validateOpDeleteLicenseConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteLicenseConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteLicenseConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteLicenseConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetLicenseConfiguration struct {
}

func (*validateOpGetLicenseConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetLicenseConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetLicenseConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetLicenseConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListAssociationsForLicenseConfiguration struct {
}

func (*validateOpListAssociationsForLicenseConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListAssociationsForLicenseConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListAssociationsForLicenseConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListAssociationsForLicenseConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListFailuresForLicenseConfigurationOperations struct {
}

func (*validateOpListFailuresForLicenseConfigurationOperations) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListFailuresForLicenseConfigurationOperations) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListFailuresForLicenseConfigurationOperationsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListFailuresForLicenseConfigurationOperationsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListLicenseSpecificationsForResource struct {
}

func (*validateOpListLicenseSpecificationsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListLicenseSpecificationsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListLicenseSpecificationsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListLicenseSpecificationsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListResourceInventory struct {
}

func (*validateOpListResourceInventory) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListResourceInventory) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListResourceInventoryInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListResourceInventoryInput(input); err != nil {
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

type validateOpListUsageForLicenseConfiguration struct {
}

func (*validateOpListUsageForLicenseConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListUsageForLicenseConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListUsageForLicenseConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListUsageForLicenseConfigurationInput(input); err != nil {
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

type validateOpUpdateLicenseConfiguration struct {
}

func (*validateOpUpdateLicenseConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateLicenseConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateLicenseConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateLicenseConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateLicenseSpecificationsForResource struct {
}

func (*validateOpUpdateLicenseSpecificationsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateLicenseSpecificationsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateLicenseSpecificationsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateLicenseSpecificationsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateServiceSettings struct {
}

func (*validateOpUpdateServiceSettings) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateServiceSettings) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateServiceSettingsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateServiceSettingsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateLicenseConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateLicenseConfiguration{}, middleware.After)
}

func addOpDeleteLicenseConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteLicenseConfiguration{}, middleware.After)
}

func addOpGetLicenseConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetLicenseConfiguration{}, middleware.After)
}

func addOpListAssociationsForLicenseConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListAssociationsForLicenseConfiguration{}, middleware.After)
}

func addOpListFailuresForLicenseConfigurationOperationsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListFailuresForLicenseConfigurationOperations{}, middleware.After)
}

func addOpListLicenseSpecificationsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListLicenseSpecificationsForResource{}, middleware.After)
}

func addOpListResourceInventoryValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListResourceInventory{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpListUsageForLicenseConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListUsageForLicenseConfiguration{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateLicenseConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateLicenseConfiguration{}, middleware.After)
}

func addOpUpdateLicenseSpecificationsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateLicenseSpecificationsForResource{}, middleware.After)
}

func addOpUpdateServiceSettingsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateServiceSettings{}, middleware.After)
}

func validateInventoryFilter(v *types.InventoryFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InventoryFilter"}
	if len(v.Condition) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Condition"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateInventoryFilterList(v []types.InventoryFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InventoryFilterList"}
	for i := range v {
		if err := validateInventoryFilter(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateLicenseSpecification(v *types.LicenseSpecification) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "LicenseSpecification"}
	if v.LicenseConfigurationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LicenseConfigurationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateLicenseSpecifications(v []types.LicenseSpecification) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "LicenseSpecifications"}
	for i := range v {
		if err := validateLicenseSpecification(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOrganizationConfiguration(v *types.OrganizationConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "OrganizationConfiguration"}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateProductInformation(v *types.ProductInformation) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ProductInformation"}
	if v.ResourceType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceType"))
	}
	if v.ProductInformationFilterList == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProductInformationFilterList"))
	} else if v.ProductInformationFilterList != nil {
		if err := validateProductInformationFilterList(v.ProductInformationFilterList); err != nil {
			invalidParams.AddNested("ProductInformationFilterList", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateProductInformationFilter(v *types.ProductInformationFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ProductInformationFilter"}
	if v.ProductInformationFilterValue == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProductInformationFilterValue"))
	}
	if v.ProductInformationFilterComparator == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProductInformationFilterComparator"))
	}
	if v.ProductInformationFilterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProductInformationFilterName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateProductInformationFilterList(v []types.ProductInformationFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ProductInformationFilterList"}
	for i := range v {
		if err := validateProductInformationFilter(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateProductInformationList(v []types.ProductInformation) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ProductInformationList"}
	for i := range v {
		if err := validateProductInformation(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateLicenseConfigurationInput(v *CreateLicenseConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateLicenseConfigurationInput"}
	if v.ProductInformationList != nil {
		if err := validateProductInformationList(v.ProductInformationList); err != nil {
			invalidParams.AddNested("ProductInformationList", err.(smithy.InvalidParamsError))
		}
	}
	if len(v.LicenseCountingType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("LicenseCountingType"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteLicenseConfigurationInput(v *DeleteLicenseConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteLicenseConfigurationInput"}
	if v.LicenseConfigurationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LicenseConfigurationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetLicenseConfigurationInput(v *GetLicenseConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetLicenseConfigurationInput"}
	if v.LicenseConfigurationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LicenseConfigurationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListAssociationsForLicenseConfigurationInput(v *ListAssociationsForLicenseConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListAssociationsForLicenseConfigurationInput"}
	if v.LicenseConfigurationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LicenseConfigurationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListFailuresForLicenseConfigurationOperationsInput(v *ListFailuresForLicenseConfigurationOperationsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListFailuresForLicenseConfigurationOperationsInput"}
	if v.LicenseConfigurationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LicenseConfigurationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListLicenseSpecificationsForResourceInput(v *ListLicenseSpecificationsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListLicenseSpecificationsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListResourceInventoryInput(v *ListResourceInventoryInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListResourceInventoryInput"}
	if v.Filters != nil {
		if err := validateInventoryFilterList(v.Filters); err != nil {
			invalidParams.AddNested("Filters", err.(smithy.InvalidParamsError))
		}
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

func validateOpListUsageForLicenseConfigurationInput(v *ListUsageForLicenseConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListUsageForLicenseConfigurationInput"}
	if v.LicenseConfigurationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LicenseConfigurationArn"))
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
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateLicenseConfigurationInput(v *UpdateLicenseConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateLicenseConfigurationInput"}
	if v.ProductInformationList != nil {
		if err := validateProductInformationList(v.ProductInformationList); err != nil {
			invalidParams.AddNested("ProductInformationList", err.(smithy.InvalidParamsError))
		}
	}
	if v.LicenseConfigurationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LicenseConfigurationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateLicenseSpecificationsForResourceInput(v *UpdateLicenseSpecificationsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateLicenseSpecificationsForResourceInput"}
	if v.AddLicenseSpecifications != nil {
		if err := validateLicenseSpecifications(v.AddLicenseSpecifications); err != nil {
			invalidParams.AddNested("AddLicenseSpecifications", err.(smithy.InvalidParamsError))
		}
	}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.RemoveLicenseSpecifications != nil {
		if err := validateLicenseSpecifications(v.RemoveLicenseSpecifications); err != nil {
			invalidParams.AddNested("RemoveLicenseSpecifications", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateServiceSettingsInput(v *UpdateServiceSettingsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateServiceSettingsInput"}
	if v.OrganizationConfiguration != nil {
		if err := validateOrganizationConfiguration(v.OrganizationConfiguration); err != nil {
			invalidParams.AddNested("OrganizationConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
