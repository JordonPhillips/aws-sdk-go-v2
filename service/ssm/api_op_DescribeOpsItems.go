// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Query a set of OpsItems. You must have permission in AWS Identity and Access
// Management (IAM) to query a list of OpsItems. For more information, see Getting
// started with OpsCenter
// (https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter-getting-started.html)
// in the AWS Systems Manager User Guide. Operations engineers and IT professionals
// use OpsCenter to view, investigate, and remediate operational issues impacting
// the performance and health of their AWS resources. For more information, see AWS
// Systems Manager OpsCenter
// (https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter.html) in
// the AWS Systems Manager User Guide.
func (c *Client) DescribeOpsItems(ctx context.Context, params *DescribeOpsItemsInput, optFns ...func(*Options)) (*DescribeOpsItemsOutput, error) {
	if params == nil {
		params = &DescribeOpsItemsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOpsItems", params, optFns, addOperationDescribeOpsItemsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOpsItemsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOpsItemsInput struct {

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string

	// One or more filters to limit the response.
	//
	// * Key: CreatedTime Operations:
	// GreaterThan, LessThan
	//
	// * Key: LastModifiedBy Operations: Contains, Equals
	//
	// *
	// Key: LastModifiedTime Operations: GreaterThan, LessThan
	//
	// * Key: Priority
	// Operations: Equals
	//
	// * Key: Source Operations: Contains, Equals
	//
	// * Key: Status
	// Operations: Equals
	//
	// * Key: Title Operations: Contains
	//
	// * Key: OperationalData*
	// Operations: Equals
	//
	// * Key: OperationalDataKey Operations: Equals
	//
	// * Key:
	// OperationalDataValue Operations: Equals, Contains
	//
	// * Key: OpsItemId Operations:
	// Equals
	//
	// * Key: ResourceId Operations: Contains
	//
	// * Key: AutomationId Operations:
	// Equals
	//
	// *If you filter the response by using the OperationalData operator,
	// specify a key-value pair by using the following JSON format:
	// {"key":"key_name","value":"a_value"}
	OpsItemFilters []*types.OpsItemFilter
}

type DescribeOpsItemsOutput struct {

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// A list of OpsItems.
	OpsItemSummaries []*types.OpsItemSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeOpsItemsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeOpsItems{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeOpsItems{}, middleware.After)
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
	addOpDescribeOpsItemsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOpsItems(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeOpsItems(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeOpsItems",
	}
}
