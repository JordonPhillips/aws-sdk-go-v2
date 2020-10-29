// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an AWS Lambda layer
// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) from a
// ZIP archive. Each time you call PublishLayerVersion with the same layer name, a
// new version is created. Add layers to your function with CreateFunction or
// UpdateFunctionConfiguration.
func (c *Client) PublishLayerVersion(ctx context.Context, params *PublishLayerVersionInput, optFns ...func(*Options)) (*PublishLayerVersionOutput, error) {
	if params == nil {
		params = &PublishLayerVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PublishLayerVersion", params, optFns, addOperationPublishLayerVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PublishLayerVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PublishLayerVersionInput struct {

	// The function layer archive.
	//
	// This member is required.
	Content *types.LayerVersionContentInput

	// The name or Amazon Resource Name (ARN) of the layer.
	//
	// This member is required.
	LayerName *string

	// A list of compatible function runtimes
	// (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html). Used for
	// filtering with ListLayers and ListLayerVersions.
	CompatibleRuntimes []types.Runtime

	// The description of the version.
	Description *string

	// The layer's software license. It can be any of the following:
	//
	// * An SPDX license
	// identifier (https://spdx.org/licenses/). For example, MIT.
	//
	// * The URL of a
	// license hosted on the internet. For example,
	// https://opensource.org/licenses/MIT.
	//
	// * The full text of the license.
	LicenseInfo *string
}

type PublishLayerVersionOutput struct {

	// The layer's compatible runtimes.
	CompatibleRuntimes []types.Runtime

	// Details about the layer version.
	Content *types.LayerVersionContentOutput

	// The date that the layer version was created, in ISO-8601 format
	// (https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).
	CreatedDate *string

	// The description of the version.
	Description *string

	// The ARN of the layer.
	LayerArn *string

	// The ARN of the layer version.
	LayerVersionArn *string

	// The layer's software license.
	LicenseInfo *string

	// The version number.
	Version *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPublishLayerVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPublishLayerVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPublishLayerVersion{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpPublishLayerVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPublishLayerVersion(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPublishLayerVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "PublishLayerVersion",
	}
}
