// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchain

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchain/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a proposal for a change to the network that other members of the network
// can vote on, for example, a proposal to add a new member to the network. Any
// member can create a proposal. Applies only to Hyperledger Fabric.
func (c *Client) CreateProposal(ctx context.Context, params *CreateProposalInput, optFns ...func(*Options)) (*CreateProposalOutput, error) {
	if params == nil {
		params = &CreateProposalInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProposal", params, optFns, addOperationCreateProposalMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProposalOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProposalInput struct {

	// The type of actions proposed, such as inviting a member or removing a member.
	// The types of Actions in a proposal are mutually exclusive. For example, a
	// proposal with Invitations actions cannot also contain Removals actions.
	//
	// This member is required.
	Actions *types.ProposalActions

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the operation. An idempotent operation completes no more than one time. This
	// identifier is required only if you make a service request directly using an HTTP
	// client. It is generated automatically if you use an AWS SDK or the AWS CLI.
	//
	// This member is required.
	ClientRequestToken *string

	// The unique identifier of the member that is creating the proposal. This
	// identifier is especially useful for identifying the member making the proposal
	// when multiple members exist in a single AWS account.
	//
	// This member is required.
	MemberId *string

	// The unique identifier of the network for which the proposal is made.
	//
	// This member is required.
	NetworkId *string

	// A description for the proposal that is visible to voting members, for example,
	// "Proposal to add Example Corp. as member."
	Description *string
}

type CreateProposalOutput struct {

	// The unique identifier of the proposal.
	ProposalId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateProposalMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateProposal{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateProposal{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateProposalMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateProposalValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProposal(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateProposal struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateProposal) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateProposal) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateProposalInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateProposalInput ")
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
func addIdempotencyToken_opCreateProposalMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateProposal{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateProposal(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "managedblockchain",
		OperationName: "CreateProposal",
	}
}
