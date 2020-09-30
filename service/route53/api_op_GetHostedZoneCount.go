// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	route53cust "github.com/aws/aws-sdk-go-v2/service/route53/internal/customizations"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the number of hosted zones that are associated with the current AWS
// account.
func (c *Client) GetHostedZoneCount(ctx context.Context, params *GetHostedZoneCountInput, optFns ...func(*Options)) (*GetHostedZoneCountOutput, error) {
	stack := middleware.NewStack("GetHostedZoneCount", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetHostedZoneCountMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	route53cust.AddSanitizeURLMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetHostedZoneCount(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "GetHostedZoneCount",
			Err:           err,
		}
	}
	out := result.(*GetHostedZoneCountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to retrieve a count of all the hosted zones that are associated with
// the current AWS account.
type GetHostedZoneCountInput struct {
}

// A complex type that contains the response to a GetHostedZoneCount request.
type GetHostedZoneCountOutput struct {

	// The total number of public and private hosted zones that are associated with the
	// current AWS account.
	//
	// This member is required.
	HostedZoneCount *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetHostedZoneCountMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetHostedZoneCount{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetHostedZoneCount{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetHostedZoneCount(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "GetHostedZoneCount",
	}
}
