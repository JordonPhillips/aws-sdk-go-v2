// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Increase the number of node groups in the Global Datastore
func (c *Client) IncreaseNodeGroupsInGlobalReplicationGroup(ctx context.Context, params *IncreaseNodeGroupsInGlobalReplicationGroupInput, optFns ...func(*Options)) (*IncreaseNodeGroupsInGlobalReplicationGroupOutput, error) {
	if params == nil {
		params = &IncreaseNodeGroupsInGlobalReplicationGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "IncreaseNodeGroupsInGlobalReplicationGroup", params, optFns, addOperationIncreaseNodeGroupsInGlobalReplicationGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*IncreaseNodeGroupsInGlobalReplicationGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type IncreaseNodeGroupsInGlobalReplicationGroupInput struct {

	// Indicates that the process begins immediately. At present, the only permitted
	// value for this parameter is true.
	//
	// This member is required.
	ApplyImmediately *bool

	// The name of the Global Datastore
	//
	// This member is required.
	GlobalReplicationGroupId *string

	// The number of node groups you wish to add
	//
	// This member is required.
	NodeGroupCount *int32

	// Describes the replication group IDs, the AWS regions where they are stored and
	// the shard configuration for each that comprise the Global Datastore
	RegionalConfigurations []*types.RegionalConfiguration
}

type IncreaseNodeGroupsInGlobalReplicationGroupOutput struct {

	// Consists of a primary cluster that accepts writes and an associated secondary
	// cluster that resides in a different AWS region. The secondary cluster accepts
	// only reads. The primary cluster automatically replicates updates to the
	// secondary cluster.
	//
	// * The GlobalReplicationGroupIdSuffix represents the name of
	// the Global Datastore, which is what you use to associate a secondary cluster.
	GlobalReplicationGroup *types.GlobalReplicationGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationIncreaseNodeGroupsInGlobalReplicationGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpIncreaseNodeGroupsInGlobalReplicationGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpIncreaseNodeGroupsInGlobalReplicationGroup{}, middleware.After)
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
	addOpIncreaseNodeGroupsInGlobalReplicationGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opIncreaseNodeGroupsInGlobalReplicationGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opIncreaseNodeGroupsInGlobalReplicationGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "IncreaseNodeGroupsInGlobalReplicationGroup",
	}
}
