// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the firmware version and other information about a wireless gateway.
func (c *Client) GetWirelessGatewayFirmwareInformation(ctx context.Context, params *GetWirelessGatewayFirmwareInformationInput, optFns ...func(*Options)) (*GetWirelessGatewayFirmwareInformationOutput, error) {
	if params == nil {
		params = &GetWirelessGatewayFirmwareInformationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetWirelessGatewayFirmwareInformation", params, optFns, addOperationGetWirelessGatewayFirmwareInformationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetWirelessGatewayFirmwareInformationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWirelessGatewayFirmwareInformationInput struct {

	// The ID of the resource to get.
	//
	// This member is required.
	Id *string
}

type GetWirelessGatewayFirmwareInformationOutput struct {

	// Information about the wireless gateway's firmware.
	LoRaWAN *types.LoRaWANGatewayCurrentVersion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetWirelessGatewayFirmwareInformationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetWirelessGatewayFirmwareInformation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetWirelessGatewayFirmwareInformation{}, middleware.After)
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
	if err = addOpGetWirelessGatewayFirmwareInformationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetWirelessGatewayFirmwareInformation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetWirelessGatewayFirmwareInformation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "GetWirelessGatewayFirmwareInformation",
	}
}
