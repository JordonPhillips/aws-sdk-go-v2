// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This API is in preview release for Amazon Connect and is subject to change.
// Updates the name and description of a quick connect. The request accepts the
// following data in JSON format. At least Name or Description must be provided.
func (c *Client) UpdateQuickConnectName(ctx context.Context, params *UpdateQuickConnectNameInput, optFns ...func(*Options)) (*UpdateQuickConnectNameOutput, error) {
	if params == nil {
		params = &UpdateQuickConnectNameInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateQuickConnectName", params, optFns, addOperationUpdateQuickConnectNameMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateQuickConnectNameOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateQuickConnectNameInput struct {

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// The identifier for the quick connect.
	//
	// This member is required.
	QuickConnectId *string

	// The description of the quick connect.
	Description *string

	// The name of the quick connect.
	Name *string
}

type UpdateQuickConnectNameOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateQuickConnectNameMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateQuickConnectName{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateQuickConnectName{}, middleware.After)
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
	if err = addOpUpdateQuickConnectNameValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateQuickConnectName(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateQuickConnectName(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "UpdateQuickConnectName",
	}
}
