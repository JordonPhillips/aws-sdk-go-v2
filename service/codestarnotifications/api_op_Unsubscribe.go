// Code generated by smithy-go-codegen DO NOT EDIT.

package codestarnotifications

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes an association between a notification rule and an Amazon SNS topic so
// that subscribers to that topic stop receiving notifications when the events
// described in the rule are triggered.
func (c *Client) Unsubscribe(ctx context.Context, params *UnsubscribeInput, optFns ...func(*Options)) (*UnsubscribeOutput, error) {
	stack := middleware.NewStack("Unsubscribe", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUnsubscribeMiddlewares(stack)
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
	addOpUnsubscribeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUnsubscribe(options.Region), middleware.Before)

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
			OperationName: "Unsubscribe",
			Err:           err,
		}
	}
	out := result.(*UnsubscribeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UnsubscribeInput struct {
	// The Amazon Resource Name (ARN) of the notification rule.
	Arn *string
	// The ARN of the SNS topic to unsubscribe from the notification rule.
	TargetAddress *string
}

type UnsubscribeOutput struct {
	// The Amazon Resource Name (ARN) of the the notification rule from which you have
	// removed a subscription.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUnsubscribeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUnsubscribe{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUnsubscribe{}, middleware.After)
}

func newServiceMetadataMiddleware_opUnsubscribe(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codestar-notifications",
		OperationName: "Unsubscribe",
	}
}