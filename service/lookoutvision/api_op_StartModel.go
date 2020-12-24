// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutvision

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lookoutvision/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts the running of the version of an Amazon Lookout for Vision model.
// Starting a model takes a while to complete. To check the current state of the
// model, use DescribeModel. Once the model is running, you can detect custom
// labels in new images by calling DetectAnomalies. You are charged for the amount
// of time that the model is running. To stop a running model, call StopModel.
func (c *Client) StartModel(ctx context.Context, params *StartModelInput, optFns ...func(*Options)) (*StartModelOutput, error) {
	if params == nil {
		params = &StartModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartModel", params, optFns, addOperationStartModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartModelInput struct {

	// The minimum number of inference units to use. A single inference unit represents
	// 1 hour of processing and can support up to 5 Transaction Pers Second (TPS). Use
	// a higher number to increase the TPS throughput of your model. You are charged
	// for the number of inference units that you use.
	//
	// This member is required.
	MinInferenceUnits *int32

	// The version of the model that you want to start.
	//
	// This member is required.
	ModelVersion *string

	// The name of the project that contains the model that you want to start.
	//
	// This member is required.
	ProjectName *string

	// ClientToken is an idempotency token that ensures a call to StartModel completes
	// only once. You choose the value to pass. For example, An issue, such as an
	// network outage, might prevent you from getting a response from StartModel. In
	// this case, safely retry your call to StartModel by using the same ClientToken
	// parameter value. An error occurs if the other input parameters are not the same
	// as in the first request. Using a different value for ClientToken is considered a
	// new call to StartModel. An idempotency token is active for 8 hours.
	ClientToken *string
}

type StartModelOutput struct {

	// The current running status of the model.
	Status types.ModelHostingStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartModel{}, middleware.After)
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
	if err = addIdempotencyToken_opStartModelMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartModel(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartModel struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartModel) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartModel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartModelInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartModelInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartModelMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartModel{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutvision",
		OperationName: "StartModel",
	}
}
