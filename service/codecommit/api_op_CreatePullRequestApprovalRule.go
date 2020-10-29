// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an approval rule for a pull request.
func (c *Client) CreatePullRequestApprovalRule(ctx context.Context, params *CreatePullRequestApprovalRuleInput, optFns ...func(*Options)) (*CreatePullRequestApprovalRuleOutput, error) {
	if params == nil {
		params = &CreatePullRequestApprovalRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePullRequestApprovalRule", params, optFns, addOperationCreatePullRequestApprovalRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePullRequestApprovalRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePullRequestApprovalRuleInput struct {

	// The content of the approval rule, including the number of approvals needed and
	// the structure of an approval pool defined for approvals, if any. For more
	// information about approval pools, see the AWS CodeCommit User Guide. When you
	// create the content of the approval rule, you can specify approvers in an
	// approval pool in one of two ways:
	//
	// * CodeCommitApprovers: This option only
	// requires an AWS account and a resource. It can be used for both IAM users and
	// federated access users whose name matches the provided resource name. This is a
	// very powerful option that offers a great deal of flexibility. For example, if
	// you specify the AWS account 123456789012 and Mary_Major, all of the following
	// would be counted as approvals coming from that user:
	//
	// * An IAM user in the
	// account (arn:aws:iam::123456789012:user/Mary_Major)
	//
	// * A federated user
	// identified in IAM as Mary_Major
	// (arn:aws:sts::123456789012:federated-user/Mary_Major)
	//
	// This option does not
	// recognize an active session of someone assuming the role of CodeCommitReview
	// with a role session name of Mary_Major
	// (arn:aws:sts::123456789012:assumed-role/CodeCommitReview/Mary_Major) unless you
	// include a wildcard (*Mary_Major).
	//
	// * Fully qualified ARN: This option allows you
	// to specify the fully qualified Amazon Resource Name (ARN) of the IAM user or
	// role.
	//
	// For more information about IAM ARNs, wildcards, and formats, see IAM
	// Identifiers
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html) in
	// the IAM User Guide.
	//
	// This member is required.
	ApprovalRuleContent *string

	// The name for the approval rule.
	//
	// This member is required.
	ApprovalRuleName *string

	// The system-generated ID of the pull request for which you want to create the
	// approval rule.
	//
	// This member is required.
	PullRequestId *string
}

type CreatePullRequestApprovalRuleOutput struct {

	// Information about the created approval rule.
	//
	// This member is required.
	ApprovalRule *types.ApprovalRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreatePullRequestApprovalRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePullRequestApprovalRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePullRequestApprovalRule{}, middleware.After)
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
	addOpCreatePullRequestApprovalRuleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePullRequestApprovalRule(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreatePullRequestApprovalRule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "CreatePullRequestApprovalRule",
	}
}
