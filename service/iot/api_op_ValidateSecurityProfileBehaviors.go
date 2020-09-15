// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Validates a Device Defender security profile behaviors specification.
func (c *Client) ValidateSecurityProfileBehaviors(ctx context.Context, params *ValidateSecurityProfileBehaviorsInput, optFns ...func(*Options)) (*ValidateSecurityProfileBehaviorsOutput, error) {
	stack := middleware.NewStack("ValidateSecurityProfileBehaviors", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpValidateSecurityProfileBehaviorsMiddlewares(stack)
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
	addOpValidateSecurityProfileBehaviorsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opValidateSecurityProfileBehaviors(options.Region), middleware.Before)

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
			OperationName: "ValidateSecurityProfileBehaviors",
			Err:           err,
		}
	}
	out := result.(*ValidateSecurityProfileBehaviorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ValidateSecurityProfileBehaviorsInput struct {
	// Specifies the behaviors that, when violated by a device (thing), cause an alert.
	Behaviors []*types.Behavior
}

type ValidateSecurityProfileBehaviorsOutput struct {
	// True if the behaviors were valid.
	Valid *bool
	// The list of any errors found in the behaviors.
	ValidationErrors []*types.ValidationError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpValidateSecurityProfileBehaviorsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpValidateSecurityProfileBehaviors{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpValidateSecurityProfileBehaviors{}, middleware.After)
}

func newServiceMetadataMiddleware_opValidateSecurityProfileBehaviors(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ValidateSecurityProfileBehaviors",
	}
}