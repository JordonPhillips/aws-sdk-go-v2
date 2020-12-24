// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an asset model from specified property and hierarchy definitions. You
// create assets from asset models. With asset models, you can easily create assets
// of the same type that have standardized definitions. Each asset created from a
// model inherits the asset model's property and hierarchy definitions. For more
// information, see Defining asset models
// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/define-models.html)
// in the AWS IoT SiteWise User Guide.
func (c *Client) CreateAssetModel(ctx context.Context, params *CreateAssetModelInput, optFns ...func(*Options)) (*CreateAssetModelOutput, error) {
	if params == nil {
		params = &CreateAssetModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAssetModel", params, optFns, addOperationCreateAssetModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAssetModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAssetModelInput struct {

	// A unique, friendly name for the asset model.
	//
	// This member is required.
	AssetModelName *string

	// The composite asset models that are part of this asset model. Composite asset
	// models are asset models that contain specific properties. Each composite model
	// has a type that defines the properties that the composite model supports. Use
	// composite asset models to define alarms on this asset model.
	AssetModelCompositeModels []types.AssetModelCompositeModelDefinition

	// A description for the asset model.
	AssetModelDescription *string

	// The hierarchy definitions of the asset model. Each hierarchy specifies an asset
	// model whose assets can be children of any other assets created from this asset
	// model. For more information, see Asset hierarchies
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-hierarchies.html)
	// in the AWS IoT SiteWise User Guide. You can specify up to 10 hierarchies per
	// asset model. For more information, see Quotas
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/quotas.html) in the
	// AWS IoT SiteWise User Guide.
	AssetModelHierarchies []types.AssetModelHierarchyDefinition

	// The property definitions of the asset model. For more information, see Asset
	// properties
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-properties.html)
	// in the AWS IoT SiteWise User Guide. You can specify up to 200 properties per
	// asset model. For more information, see Quotas
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/quotas.html) in the
	// AWS IoT SiteWise User Guide.
	AssetModelProperties []types.AssetModelPropertyDefinition

	// A unique case-sensitive identifier that you can provide to ensure the
	// idempotency of the request. Don't reuse this client token if a new idempotent
	// request is required.
	ClientToken *string

	// A list of key-value pairs that contain metadata for the asset model. For more
	// information, see Tagging your AWS IoT SiteWise resources
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html)
	// in the AWS IoT SiteWise User Guide.
	Tags map[string]string
}

type CreateAssetModelOutput struct {

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the asset model, which has the following format.
	// arn:${Partition}:iotsitewise:${Region}:${Account}:asset-model/${AssetModelId}
	//
	// This member is required.
	AssetModelArn *string

	// The ID of the asset model. You can use this ID when you call other AWS IoT
	// SiteWise APIs.
	//
	// This member is required.
	AssetModelId *string

	// The status of the asset model, which contains a state (CREATING after
	// successfully calling this operation) and any error message.
	//
	// This member is required.
	AssetModelStatus *types.AssetModelStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateAssetModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAssetModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAssetModel{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opCreateAssetModelMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreateAssetModelMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateAssetModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAssetModel(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opCreateAssetModelMiddleware struct {
}

func (*endpointPrefix_opCreateAssetModelMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCreateAssetModelMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "model." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opCreateAssetModelMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opCreateAssetModelMiddleware{}, `OperationSerializer`, middleware.After)
}

type idempotencyToken_initializeOpCreateAssetModel struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateAssetModel) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateAssetModel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateAssetModelInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateAssetModelInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateAssetModelMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateAssetModel{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateAssetModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "CreateAssetModel",
	}
}
