// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a resouce policy to control access to a model group. For information about
// resoure policies, see Identity-based policies and resource-based policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_identity-vs-resource.html)
// in the AWS Identity and Access Management User Guide..
func (c *Client) PutModelPackageGroupPolicy(ctx context.Context, params *PutModelPackageGroupPolicyInput, optFns ...func(*Options)) (*PutModelPackageGroupPolicyOutput, error) {
	if params == nil {
		params = &PutModelPackageGroupPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutModelPackageGroupPolicy", params, optFns, addOperationPutModelPackageGroupPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutModelPackageGroupPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutModelPackageGroupPolicyInput struct {

	// The name of the model group to add a resource policy to.
	//
	// This member is required.
	ModelPackageGroupName *string

	// The resource policy for the model group.
	//
	// This member is required.
	ResourcePolicy *string
}

type PutModelPackageGroupPolicyOutput struct {

	// The Amazon Resource Name (ARN) of the model package group.
	//
	// This member is required.
	ModelPackageGroupArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutModelPackageGroupPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutModelPackageGroupPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutModelPackageGroupPolicy{}, middleware.After)
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
	if err = addOpPutModelPackageGroupPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutModelPackageGroupPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutModelPackageGroupPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "PutModelPackageGroupPolicy",
	}
}
