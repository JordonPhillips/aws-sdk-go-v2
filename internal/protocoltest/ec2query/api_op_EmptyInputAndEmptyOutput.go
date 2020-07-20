// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The example tests how requests and responses are serialized when there's no
// request or response members. While this should be rare, code generators must
// support this.
func (c *Client) EmptyInputAndEmptyOutput(ctx context.Context, params *EmptyInputAndEmptyOutputInput, optFns ...func(*Options)) (*EmptyInputAndEmptyOutputOutput, error) {
	stack := middleware.NewStack("EmptyInputAndEmptyOutput", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opEmptyInputAndEmptyOutput(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "EmptyInputAndEmptyOutput",
			Err:           err,
		}
	}
	out := result.(*EmptyInputAndEmptyOutputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EmptyInputAndEmptyOutputInput struct {
}

type EmptyInputAndEmptyOutputOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func newServiceMetadataMiddleware_opEmptyInputAndEmptyOutput(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "EC2 Protocol",
		ServiceID:      "ec2protocol",
		EndpointPrefix: "ec2protocol",
		OperationName:  "EmptyInputAndEmptyOutput",
	}
}
