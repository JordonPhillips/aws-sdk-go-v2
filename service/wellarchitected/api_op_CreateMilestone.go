// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a milestone for an existing workload.
func (c *Client) CreateMilestone(ctx context.Context, params *CreateMilestoneInput, optFns ...func(*Options)) (*CreateMilestoneOutput, error) {
	if params == nil {
		params = &CreateMilestoneInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMilestone", params, optFns, addOperationCreateMilestoneMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMilestoneOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for milestone creation.
type CreateMilestoneInput struct {

	// A unique case-sensitive string used to ensure that this request is idempotent
	// (executes only once). You should not reuse the same token for other requests. If
	// you retry a request with the same client request token and the same parameters
	// after it has completed successfully, the result of the original request is
	// returned. This token is listed as required, however, if you do not specify it,
	// the AWS SDKs automatically generate one for you. If you are not using the AWS
	// SDK or the AWS CLI, you must provide this token or the request will fail.
	//
	// This member is required.
	ClientRequestToken *string

	// The name of the milestone in a workload. Milestone names must be unique within a
	// workload.
	//
	// This member is required.
	MilestoneName *string

	// The ID assigned to the workload. This ID is unique within an AWS Region.
	//
	// This member is required.
	WorkloadId *string
}

// Output of a create milestone call.
type CreateMilestoneOutput struct {

	// The milestone number. A workload can have a maximum of 100 milestones.
	MilestoneNumber int32

	// The ID assigned to the workload. This ID is unique within an AWS Region.
	WorkloadId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateMilestoneMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMilestone{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMilestone{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateMilestoneMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateMilestoneValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMilestone(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateMilestone struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateMilestone) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateMilestone) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateMilestoneInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateMilestoneInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateMilestoneMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateMilestone{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateMilestone(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wellarchitected",
		OperationName: "CreateMilestone",
	}
}
