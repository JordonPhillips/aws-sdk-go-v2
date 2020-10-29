// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list ARNs for the report groups in the current AWS account.
func (c *Client) ListReportGroups(ctx context.Context, params *ListReportGroupsInput, optFns ...func(*Options)) (*ListReportGroupsOutput, error) {
	if params == nil {
		params = &ListReportGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListReportGroups", params, optFns, addOperationListReportGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListReportGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListReportGroupsInput struct {

	// The maximum number of paginated report groups returned per response. Use
	// nextToken to iterate pages in the list of returned ReportGroup objects. The
	// default value is 100.
	MaxResults *int32

	// During a previous call, the maximum number of items that can be returned is the
	// value specified in maxResults. If there more items in the list, then a unique
	// string called a nextToken is returned. To get the next batch of items in the
	// list, call this operation again, adding the next token to the call. To get all
	// of the items in the list, keep calling this operation with each subsequent next
	// token that is returned, until no more next tokens are returned.
	NextToken *string

	// The criterion to be used to list build report groups. Valid values include:
	//
	// *
	// CREATED_TIME: List based on when each report group was created.
	//
	// *
	// LAST_MODIFIED_TIME: List based on when each report group was last changed.
	//
	// *
	// NAME: List based on each report group's name.
	SortBy types.ReportGroupSortByType

	// Used to specify the order to sort the list of returned report groups. Valid
	// values are ASCENDING and DESCENDING.
	SortOrder types.SortOrderType
}

type ListReportGroupsOutput struct {

	// During a previous call, the maximum number of items that can be returned is the
	// value specified in maxResults. If there more items in the list, then a unique
	// string called a nextToken is returned. To get the next batch of items in the
	// list, call this operation again, adding the next token to the call. To get all
	// of the items in the list, keep calling this operation with each subsequent next
	// token that is returned, until no more next tokens are returned.
	NextToken *string

	// The list of ARNs for the report groups in the current AWS account.
	ReportGroups []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListReportGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListReportGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListReportGroups{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListReportGroups(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListReportGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "ListReportGroups",
	}
}
