// Code generated by smithy-go-codegen DO NOT EDIT.

package timestreamquery

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/timestreamquery/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// DescribeEndpoints returns a list of available endpoints to make Timestream API
// calls against. This API is available through both Write and Query. Because
// Timestream’s SDKs are designed to transparently work with the service’s
// architecture, including the management and mapping of the service endpoints, it
// is not recommended that you use this API unless:
//
// * Your application uses a
// programming language that does not yet have SDK support
//
// * You require better
// control over the client-side implementation
//
// For detailed information on how to
// use DescribeEndpoints, see The Endpoint Discovery Pattern and REST APIs
// (https://docs.aws.amazon.com/timestream/latest/developerguide/Using-API.endpoint-discovery.html).
func (c *Client) DescribeEndpoints(ctx context.Context, params *DescribeEndpointsInput, optFns ...func(*Options)) (*DescribeEndpointsOutput, error) {
	if params == nil {
		params = &DescribeEndpointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEndpoints", params, optFns, addOperationDescribeEndpointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEndpointsInput struct {
}

type DescribeEndpointsOutput struct {

	// An Endpoints object is returned when a DescribeEndpoints request is made.
	//
	// This member is required.
	Endpoints []*types.Endpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEndpointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeEndpoints{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEndpoints(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeEndpoints(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "timestream",
		OperationName: "DescribeEndpoints",
	}
}
