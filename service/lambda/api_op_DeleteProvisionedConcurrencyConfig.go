// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the provisioned concurrency configuration for a function.
func (c *Client) DeleteProvisionedConcurrencyConfig(ctx context.Context, params *DeleteProvisionedConcurrencyConfigInput, optFns ...func(*Options)) (*DeleteProvisionedConcurrencyConfigOutput, error) {
	if params == nil {
		params = &DeleteProvisionedConcurrencyConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteProvisionedConcurrencyConfig", params, optFns, addOperationDeleteProvisionedConcurrencyConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteProvisionedConcurrencyConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteProvisionedConcurrencyConfigInput struct {

	// The name of the Lambda function. Name formats
	//
	// * Function name - my-function.
	//
	// *
	// Function ARN - arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	// *
	// Partial ARN - 123456789012:function:my-function.
	//
	// The length constraint applies
	// only to the full ARN. If you specify only the function name, it is limited to 64
	// characters in length.
	//
	// This member is required.
	FunctionName *string

	// The version number or alias name.
	//
	// This member is required.
	Qualifier *string
}

type DeleteProvisionedConcurrencyConfigOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteProvisionedConcurrencyConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteProvisionedConcurrencyConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteProvisionedConcurrencyConfig{}, middleware.After)
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
	addOpDeleteProvisionedConcurrencyConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteProvisionedConcurrencyConfig(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteProvisionedConcurrencyConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "DeleteProvisionedConcurrencyConfig",
	}
}
