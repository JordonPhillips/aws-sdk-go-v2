// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the contents of the effective policy for specified policy type and
// account. The effective policy is the aggregation of any policies of the
// specified type that the account inherits, plus any policy of that type that is
// directly attached to the account. This operation applies only to policy types
// other than service control policies (SCPs). For more information about policy
// inheritance, see How Policy Inheritance Works
// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies-inheritance.html)
// in the AWS Organizations User Guide. This operation can be called only from the
// organization's management account or by a member account that is a delegated
// administrator for an AWS service.
func (c *Client) DescribeEffectivePolicy(ctx context.Context, params *DescribeEffectivePolicyInput, optFns ...func(*Options)) (*DescribeEffectivePolicyOutput, error) {
	if params == nil {
		params = &DescribeEffectivePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEffectivePolicy", params, optFns, addOperationDescribeEffectivePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEffectivePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEffectivePolicyInput struct {

	// The type of policy that you want information about. You can specify one of the
	// following values:
	//
	// * AISERVICES_OPT_OUT_POLICY
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_ai-opt-out.html)
	//
	// *
	// BACKUP_POLICY
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_backup.html)
	//
	// *
	// TAG_POLICY
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_tag-policies.html)
	//
	// This member is required.
	PolicyType types.EffectivePolicyType

	// When you're signed in as the management account, specify the ID of the account
	// that you want details about. Specifying an organization root or organizational
	// unit (OU) as the target is not supported.
	TargetId *string
}

type DescribeEffectivePolicyOutput struct {

	// The contents of the effective policy.
	EffectivePolicy *types.EffectivePolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEffectivePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEffectivePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEffectivePolicy{}, middleware.After)
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
	addOpDescribeEffectivePolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEffectivePolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeEffectivePolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "DescribeEffectivePolicy",
	}
}
