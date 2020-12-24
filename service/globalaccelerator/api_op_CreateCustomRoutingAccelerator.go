// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a custom routing accelerator. A custom routing accelerator directs
// traffic to one of possibly thousands of Amazon EC2 instance destinations running
// in a single or multiple virtual private clouds (VPC) subnet endpoints. Be aware
// that, by default, all destination EC2 instances in a VPC subnet endpoint cannot
// receive traffic. To enable all destinations to receive traffic, or to specify
// individual port mappings that can receive traffic, see the
// AllowCustomRoutingTraffic
// (https://docs.aws.amazon.com/global-accelerator/latest/api/API_AllowCustomRoutingTraffic.html)
// operation.
func (c *Client) CreateCustomRoutingAccelerator(ctx context.Context, params *CreateCustomRoutingAcceleratorInput, optFns ...func(*Options)) (*CreateCustomRoutingAcceleratorOutput, error) {
	if params == nil {
		params = &CreateCustomRoutingAcceleratorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCustomRoutingAccelerator", params, optFns, addOperationCreateCustomRoutingAcceleratorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCustomRoutingAcceleratorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCustomRoutingAcceleratorInput struct {

	// A unique, case-sensitive identifier that you provide to ensure the
	// idempotency—that is, the uniqueness—of the request.
	//
	// This member is required.
	IdempotencyToken *string

	// The name of a custom routing accelerator. The name can have a maximum of 64
	// characters, must contain only alphanumeric characters or hyphens (-), and must
	// not begin or end with a hyphen.
	//
	// This member is required.
	Name *string

	// Indicates whether an accelerator is enabled. The value is true or false. The
	// default value is true. If the value is set to true, an accelerator cannot be
	// deleted. If set to false, the accelerator can be deleted.
	Enabled *bool

	// The value for the address type must be IPv4.
	IpAddressType types.IpAddressType

	// Create tags for an accelerator. For more information, see Tagging in AWS Global
	// Accelerator
	// (https://docs.aws.amazon.com/global-accelerator/latest/dg/tagging-in-global-accelerator.html)
	// in the AWS Global Accelerator Developer Guide.
	Tags []types.Tag
}

type CreateCustomRoutingAcceleratorOutput struct {

	// The accelerator that is created.
	Accelerator *types.CustomRoutingAccelerator

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateCustomRoutingAcceleratorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCustomRoutingAccelerator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCustomRoutingAccelerator{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateCustomRoutingAcceleratorMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateCustomRoutingAcceleratorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCustomRoutingAccelerator(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateCustomRoutingAccelerator struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateCustomRoutingAccelerator) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateCustomRoutingAccelerator) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateCustomRoutingAcceleratorInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateCustomRoutingAcceleratorInput ")
	}

	if input.IdempotencyToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.IdempotencyToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateCustomRoutingAcceleratorMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateCustomRoutingAccelerator{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateCustomRoutingAccelerator(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "CreateCustomRoutingAccelerator",
	}
}
