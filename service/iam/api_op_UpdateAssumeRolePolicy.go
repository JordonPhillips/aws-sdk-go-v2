// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the policy that grants an IAM entity permission to assume a role. This
// is typically referred to as the "role trust policy". For more information about
// roles, go to Using Roles to Delegate Permissions and Federate Identities
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/roles-toplevel.html).
func (c *Client) UpdateAssumeRolePolicy(ctx context.Context, params *UpdateAssumeRolePolicyInput, optFns ...func(*Options)) (*UpdateAssumeRolePolicyOutput, error) {
	if params == nil {
		params = &UpdateAssumeRolePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAssumeRolePolicy", params, optFns, addOperationUpdateAssumeRolePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAssumeRolePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAssumeRolePolicyInput struct {

	// The policy that grants an entity permission to assume the role. You must provide
	// policies in JSON format in IAM. However, for AWS CloudFormation templates
	// formatted in YAML, you can provide the policy in JSON or YAML format. AWS
	// CloudFormation always converts a YAML policy to JSON format before submitting it
	// to IAM. The regex pattern (http://wikipedia.org/wiki/regex) used to validate
	// this parameter is a string of characters consisting of the following:
	//
	// * Any
	// printable ASCII character ranging from the space character (\u0020) through the
	// end of the ASCII character range
	//
	// * The printable characters in the Basic Latin
	// and Latin-1 Supplement character set (through \u00FF)
	//
	// * The special characters
	// tab (\u0009), line feed (\u000A), and carriage return (\u000D)
	//
	// This member is required.
	PolicyDocument *string

	// The name of the role to update with the new policy. This parameter allows
	// (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters consisting of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: _+=,.@-
	//
	// This member is required.
	RoleName *string
}

type UpdateAssumeRolePolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateAssumeRolePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUpdateAssumeRolePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateAssumeRolePolicy{}, middleware.After)
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
	addOpUpdateAssumeRolePolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAssumeRolePolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateAssumeRolePolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "UpdateAssumeRolePolicy",
	}
}
