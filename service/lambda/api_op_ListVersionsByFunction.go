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

// Returns a list of versions
// (https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html), with the
// version-specific configuration of each. Lambda returns up to 50 versions per
// call.
func (c *Client) ListVersionsByFunction(ctx context.Context, params *ListVersionsByFunctionInput, optFns ...func(*Options)) (*ListVersionsByFunctionOutput, error) {
	if params == nil {
		params = &ListVersionsByFunctionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVersionsByFunction", params, optFns, addOperationListVersionsByFunctionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVersionsByFunctionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVersionsByFunctionInput struct {

	// The name of the Lambda function. Name formats
	//
	// * Function name - MyFunction.
	//
	// *
	// Function ARN - arn:aws:lambda:us-west-2:123456789012:function:MyFunction.
	//
	// *
	// Partial ARN - 123456789012:function:MyFunction.
	//
	// The length constraint applies
	// only to the full ARN. If you specify only the function name, it is limited to 64
	// characters in length.
	//
	// This member is required.
	FunctionName *string

	// Specify the pagination token that's returned by a previous request to retrieve
	// the next page of results.
	Marker *string

	// The maximum number of versions to return.
	MaxItems *int32
}

type ListVersionsByFunctionOutput struct {

	// The pagination token that's included if more results are available.
	NextMarker *string

	// A list of Lambda function versions.
	Versions []*types.FunctionConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListVersionsByFunctionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListVersionsByFunction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListVersionsByFunction{}, middleware.After)
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
	addOpListVersionsByFunctionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListVersionsByFunction(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListVersionsByFunction(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "ListVersionsByFunction",
	}
}
