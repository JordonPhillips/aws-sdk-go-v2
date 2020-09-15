// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes one or more of your linked EC2-Classic instances. This request only
// returns information about EC2-Classic instances linked to a VPC through
// ClassicLink. You cannot use this request to return information about other
// instances.
func (c *Client) DescribeClassicLinkInstances(ctx context.Context, params *DescribeClassicLinkInstancesInput, optFns ...func(*Options)) (*DescribeClassicLinkInstancesOutput, error) {
	stack := middleware.NewStack("DescribeClassicLinkInstances", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeClassicLinkInstancesMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClassicLinkInstances(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "DescribeClassicLinkInstances",
			Err:           err,
		}
	}
	out := result.(*DescribeClassicLinkInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeClassicLinkInstancesInput struct {
	// One or more instance IDs. Must be instances linked to a VPC through ClassicLink.
	InstanceIds []*string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	// Constraint: If the value is greater than 1000, we return only 1000 items.
	MaxResults *int32
	// The token for the next page of results.
	NextToken *string
	// One or more filters.
	//
	//     * group-id - The ID of a VPC security group that's
	// associated with the instance.
	//
	//     * instance-id - The ID of the instance.
	//
	//
	// * tag: - The key/value combination of a tag assigned to the resource. Use the
	// tag key in the filter name and the tag value as the filter value. For example,
	// to find all resources that have a tag with the key Owner and the value TeamA,
	// specify tag:Owner for the filter name and TeamA for the filter value.
	//
	//     *
	// tag-key - The key of a tag assigned to the resource. Use this filter to find all
	// resources assigned a tag with a specific key, regardless of the tag value.
	//
	//
	// * <p> <code>vpc-id</code> - The ID of the VPC to which the instance is
	// linked.</p> <p> <code>vpc-id</code> - The ID of the VPC that the instance is
	// linked to.</p> </li> </ul>
	Filters []*types.Filter
}

type DescribeClassicLinkInstancesOutput struct {
	// Information about one or more linked EC2-Classic instances.
	Instances []*types.ClassicLinkInstance
	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeClassicLinkInstancesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeClassicLinkInstances{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeClassicLinkInstances{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeClassicLinkInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeClassicLinkInstances",
	}
}