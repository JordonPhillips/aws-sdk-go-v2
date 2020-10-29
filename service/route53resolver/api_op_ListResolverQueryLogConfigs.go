// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists information about the specified query logging configurations. Each
// configuration defines where you want Resolver to save DNS query logs and
// specifies the VPCs that you want to log queries for.
func (c *Client) ListResolverQueryLogConfigs(ctx context.Context, params *ListResolverQueryLogConfigsInput, optFns ...func(*Options)) (*ListResolverQueryLogConfigsOutput, error) {
	if params == nil {
		params = &ListResolverQueryLogConfigsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResolverQueryLogConfigs", params, optFns, addOperationListResolverQueryLogConfigsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResolverQueryLogConfigsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResolverQueryLogConfigsInput struct {

	// An optional specification to return a subset of query logging configurations. If
	// you submit a second or subsequent ListResolverQueryLogConfigs request and
	// specify the NextToken parameter, you must use the same values for Filters, if
	// any, as in the previous request.
	Filters []*types.Filter

	// The maximum number of query logging configurations that you want to return in
	// the response to a ListResolverQueryLogConfigs request. If you don't specify a
	// value for MaxResults, Resolver returns up to 100 query logging configurations.
	MaxResults *int32

	// For the first ListResolverQueryLogConfigs request, omit this value. If there are
	// more than MaxResults query logging configurations that match the values that you
	// specify for Filters, you can submit another ListResolverQueryLogConfigs request
	// to get the next group of configurations. In the next request, specify the value
	// of NextToken from the previous response.
	NextToken *string

	// The element that you want Resolver to sort query logging configurations by. If
	// you submit a second or subsequent ListResolverQueryLogConfigs request and
	// specify the NextToken parameter, you must use the same value for SortBy, if any,
	// as in the previous request. Valid values include the following elements:
	//
	// * Arn:
	// The ARN of the query logging configuration
	//
	// * AssociationCount: The number of
	// VPCs that are associated with the specified configuration
	//
	// * CreationTime: The
	// date and time that Resolver returned when the configuration was created
	//
	// *
	// CreatorRequestId: The value that was specified for CreatorRequestId when the
	// configuration was created
	//
	// * DestinationArn: The location that logs are sent
	// to
	//
	// * Id: The ID of the configuration
	//
	// * Name: The name of the configuration
	//
	// *
	// OwnerId: The AWS account number of the account that created the configuration
	//
	// *
	// ShareStatus: Whether the configuration is shared with other AWS accounts or
	// shared with the current account by another AWS account. Sharing is configured
	// through AWS Resource Access Manager (AWS RAM).
	//
	// * Status: The current status of
	// the configuration. Valid values include the following:
	//
	// * CREATING: Resolver is
	// creating the query logging configuration.
	//
	// * CREATED: The query logging
	// configuration was successfully created. Resolver is logging queries that
	// originate in the specified VPC.
	//
	// * DELETING: Resolver is deleting this query
	// logging configuration.
	//
	// * FAILED: Resolver either couldn't create or couldn't
	// delete the query logging configuration. Here are two common causes:
	//
	// * The
	// specified destination (for example, an Amazon S3 bucket) was deleted.
	//
	// *
	// Permissions don't allow sending logs to the destination.
	SortBy *string

	// If you specified a value for SortBy, the order that you want query logging
	// configurations to be listed in, ASCENDING or DESCENDING. If you submit a second
	// or subsequent ListResolverQueryLogConfigs request and specify the NextToken
	// parameter, you must use the same value for SortOrder, if any, as in the previous
	// request.
	SortOrder types.SortOrder
}

type ListResolverQueryLogConfigsOutput struct {

	// If there are more than MaxResults query logging configurations, you can submit
	// another ListResolverQueryLogConfigs request to get the next group of
	// configurations. In the next request, specify the value of NextToken from the
	// previous response.
	NextToken *string

	// A list that contains one ResolverQueryLogConfig element for each query logging
	// configuration that matches the values that you specified for Filter.
	ResolverQueryLogConfigs []*types.ResolverQueryLogConfig

	// The total number of query logging configurations that were created by the
	// current account in the specified Region. This count can differ from the number
	// of query logging configurations that are returned in a
	// ListResolverQueryLogConfigs response, depending on the values that you specify
	// in the request.
	TotalCount *int32

	// The total number of query logging configurations that were created by the
	// current account in the specified Region and that match the filters that were
	// specified in the ListResolverQueryLogConfigs request. For the total number of
	// query logging configurations that were created by the current account in the
	// specified Region, see TotalCount.
	TotalFilteredCount *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListResolverQueryLogConfigsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListResolverQueryLogConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListResolverQueryLogConfigs{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListResolverQueryLogConfigs(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListResolverQueryLogConfigs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "ListResolverQueryLogConfigs",
	}
}
