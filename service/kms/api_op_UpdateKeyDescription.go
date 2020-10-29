// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the description of a customer master key (CMK). To see the description
// of a CMK, use DescribeKey. You cannot perform this operation on a CMK in a
// different AWS account. The CMK that you use for this operation must be in a
// compatible key state. For details, see How Key State Affects Use of a Customer
// Master Key
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// AWS Key Management Service Developer Guide.
func (c *Client) UpdateKeyDescription(ctx context.Context, params *UpdateKeyDescriptionInput, optFns ...func(*Options)) (*UpdateKeyDescriptionOutput, error) {
	if params == nil {
		params = &UpdateKeyDescriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateKeyDescription", params, optFns, addOperationUpdateKeyDescriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateKeyDescriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateKeyDescriptionInput struct {

	// New description for the CMK.
	//
	// This member is required.
	Description *string

	// A unique identifier for the customer master key (CMK). Specify the key ID or the
	// Amazon Resource Name (ARN) of the CMK. For example:
	//
	// * Key ID:
	// 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To
	// get the key ID and key ARN for a CMK, use ListKeys or DescribeKey.
	//
	// This member is required.
	KeyId *string
}

type UpdateKeyDescriptionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateKeyDescriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateKeyDescription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateKeyDescription{}, middleware.After)
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
	addOpUpdateKeyDescriptionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateKeyDescription(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateKeyDescription(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "UpdateKeyDescription",
	}
}
