// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Retrieves the configuration for asynchronous invocation for a function, version,
// or alias. To configure options for asynchronous invocation, use
// PutFunctionEventInvokeConfig.
func (c *Client) GetFunctionEventInvokeConfig(ctx context.Context, params *GetFunctionEventInvokeConfigInput, optFns ...func(*Options)) (*GetFunctionEventInvokeConfigOutput, error) {
	if params == nil {
		params = &GetFunctionEventInvokeConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFunctionEventInvokeConfig", params, optFns, addOperationGetFunctionEventInvokeConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFunctionEventInvokeConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFunctionEventInvokeConfigInput struct {

	// The name of the Lambda function, version, or alias. Name formats
	//
	// * Function
	// name - my-function (name-only), my-function:v1 (with alias).
	//
	// * Function ARN -
	// arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	// * Partial ARN -
	// 123456789012:function:my-function.
	//
	// You can append a version number or alias to
	// any of the formats. The length constraint applies only to the full ARN. If you
	// specify only the function name, it is limited to 64 characters in length.
	//
	// This member is required.
	FunctionName *string

	// A version number or alias name.
	Qualifier *string
}

type GetFunctionEventInvokeConfigOutput struct {

	// A destination for events after they have been sent to a function for processing.
	// Destinations
	//
	// * Function - The Amazon Resource Name (ARN) of a Lambda
	// function.
	//
	// * Queue - The ARN of an SQS queue.
	//
	// * Topic - The ARN of an SNS
	// topic.
	//
	// * Event Bus - The ARN of an Amazon EventBridge event bus.
	DestinationConfig *types.DestinationConfig

	// The Amazon Resource Name (ARN) of the function.
	FunctionArn *string

	// The date and time that the configuration was last updated.
	LastModified *time.Time

	// The maximum age of a request that Lambda sends to a function for processing.
	MaximumEventAgeInSeconds *int32

	// The maximum number of times to retry when the function returns an error.
	MaximumRetryAttempts *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetFunctionEventInvokeConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetFunctionEventInvokeConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFunctionEventInvokeConfig{}, middleware.After)
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
	addOpGetFunctionEventInvokeConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetFunctionEventInvokeConfig(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetFunctionEventInvokeConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "GetFunctionEventInvokeConfig",
	}
}
