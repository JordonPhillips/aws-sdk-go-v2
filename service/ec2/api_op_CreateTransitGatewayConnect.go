// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Connect attachment from a specified transit gateway attachment. A
// Connect attachment is a GRE-based tunnel attachment that you can use to
// establish a connection between a transit gateway and an appliance. A Connect
// attachment uses an existing VPC or AWS Direct Connect attachment as the
// underlying transport mechanism.
func (c *Client) CreateTransitGatewayConnect(ctx context.Context, params *CreateTransitGatewayConnectInput, optFns ...func(*Options)) (*CreateTransitGatewayConnectOutput, error) {
	if params == nil {
		params = &CreateTransitGatewayConnectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTransitGatewayConnect", params, optFns, addOperationCreateTransitGatewayConnectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTransitGatewayConnectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTransitGatewayConnectInput struct {

	// The Connect attachment options.
	//
	// This member is required.
	Options *types.CreateTransitGatewayConnectRequestOptions

	// The ID of the transit gateway attachment. You can specify a VPC attachment or a
	// AWS Direct Connect attachment.
	//
	// This member is required.
	TransportTransitGatewayAttachmentId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// The tags to apply to the Connect attachment.
	TagSpecifications []types.TagSpecification
}

type CreateTransitGatewayConnectOutput struct {

	// Information about the Connect attachment.
	TransitGatewayConnect *types.TransitGatewayConnect

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateTransitGatewayConnectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateTransitGatewayConnect{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateTransitGatewayConnect{}, middleware.After)
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
	if err = addOpCreateTransitGatewayConnectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTransitGatewayConnect(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTransitGatewayConnect(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateTransitGatewayConnect",
	}
}
