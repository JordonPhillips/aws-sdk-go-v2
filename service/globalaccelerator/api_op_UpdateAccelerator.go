// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update an accelerator. Global Accelerator is a global service that supports
// endpoints in multiple AWS Regions but you must specify the US West (Oregon)
// Region to create or update accelerators.
func (c *Client) UpdateAccelerator(ctx context.Context, params *UpdateAcceleratorInput, optFns ...func(*Options)) (*UpdateAcceleratorOutput, error) {
	if params == nil {
		params = &UpdateAcceleratorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAccelerator", params, optFns, addOperationUpdateAcceleratorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAcceleratorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAcceleratorInput struct {

	// The Amazon Resource Name (ARN) of the accelerator to update.
	//
	// This member is required.
	AcceleratorArn *string

	// Indicates whether an accelerator is enabled. The value is true or false. The
	// default value is true. If the value is set to true, the accelerator cannot be
	// deleted. If set to false, the accelerator can be deleted.
	Enabled *bool

	// The IP address type, which must be IPv4.
	IpAddressType types.IpAddressType

	// The name of the accelerator. The name can have a maximum of 32 characters, must
	// contain only alphanumeric characters or hyphens (-), and must not begin or end
	// with a hyphen.
	Name *string
}

type UpdateAcceleratorOutput struct {

	// Information about the updated accelerator.
	Accelerator *types.Accelerator

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateAcceleratorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateAccelerator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateAccelerator{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateAcceleratorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAccelerator(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateAccelerator(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "UpdateAccelerator",
	}
}
