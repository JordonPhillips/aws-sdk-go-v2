// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
// 2019. For information, including how to migrate your AWS WAF resources from the
// prior release, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
// Deletes the specified RegexPatternSet.
func (c *Client) DeleteRegexPatternSet(ctx context.Context, params *DeleteRegexPatternSetInput, optFns ...func(*Options)) (*DeleteRegexPatternSetOutput, error) {
	if params == nil {
		params = &DeleteRegexPatternSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteRegexPatternSet", params, optFns, addOperationDeleteRegexPatternSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteRegexPatternSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRegexPatternSetInput struct {

	// A unique identifier for the set. This ID is returned in the responses to create
	// and list commands. You provide it to operations like update and delete.
	//
	// This member is required.
	Id *string

	// A token used for optimistic locking. AWS WAF returns a token to your get and
	// list requests, to mark the state of the entity at the time of the request. To
	// make changes to the entity associated with the token, you provide the token to
	// operations like update and delete. AWS WAF uses the token to ensure that no
	// changes have been made to the entity since you last retrieved it. If a change
	// has been made, the update fails with a WAFOptimisticLockException. If this
	// happens, perform another get, and use the new token returned by that operation.
	//
	// This member is required.
	LockToken *string

	// The name of the set. You cannot change the name after you create the set.
	//
	// This member is required.
	Name *string

	// Specifies whether this is for an AWS CloudFront distribution or for a regional
	// application. A regional application can be an Application Load Balancer (ALB),
	// an API Gateway REST API, or an AppSync GraphQL API. To work with CloudFront, you
	// must also specify the Region US East (N. Virginia) as follows:
	//
	// * CLI - Specify
	// the Region when you use the CloudFront scope: --scope=CLOUDFRONT
	// --region=us-east-1.
	//
	// * API and SDKs - For all calls, use the Region endpoint
	// us-east-1.
	//
	// This member is required.
	Scope types.Scope
}

type DeleteRegexPatternSetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteRegexPatternSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteRegexPatternSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteRegexPatternSet{}, middleware.After)
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
	addOpDeleteRegexPatternSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRegexPatternSet(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteRegexPatternSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wafv2",
		OperationName: "DeleteRegexPatternSet",
	}
}
