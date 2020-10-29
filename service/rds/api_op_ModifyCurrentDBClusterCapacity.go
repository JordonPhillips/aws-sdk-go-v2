// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Set the capacity of an Aurora Serverless DB cluster to a specific value. Aurora
// Serverless scales seamlessly based on the workload on the DB cluster. In some
// cases, the capacity might not scale fast enough to meet a sudden change in
// workload, such as a large number of new transactions. Call
// ModifyCurrentDBClusterCapacity to set the capacity explicitly. After this call
// sets the DB cluster capacity, Aurora Serverless can automatically scale the DB
// cluster based on the cooldown period for scaling up and the cooldown period for
// scaling down. For more information about Aurora Serverless, see Using Amazon
// Aurora Serverless
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.html)
// in the Amazon Aurora User Guide. If you call ModifyCurrentDBClusterCapacity with
// the default TimeoutAction, connections that prevent Aurora Serverless from
// finding a scaling point might be dropped. For more information about scaling
// points, see  Autoscaling for Aurora Serverless
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.how-it-works.html#aurora-serverless.how-it-works.auto-scaling)
// in the Amazon Aurora User Guide. This action only applies to Aurora DB clusters.
func (c *Client) ModifyCurrentDBClusterCapacity(ctx context.Context, params *ModifyCurrentDBClusterCapacityInput, optFns ...func(*Options)) (*ModifyCurrentDBClusterCapacityOutput, error) {
	if params == nil {
		params = &ModifyCurrentDBClusterCapacityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyCurrentDBClusterCapacity", params, optFns, addOperationModifyCurrentDBClusterCapacityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyCurrentDBClusterCapacityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyCurrentDBClusterCapacityInput struct {

	// The DB cluster identifier for the cluster being modified. This parameter isn't
	// case-sensitive. Constraints:
	//
	// * Must match the identifier of an existing DB
	// cluster.
	//
	// This member is required.
	DBClusterIdentifier *string

	// The DB cluster capacity. When you change the capacity of a paused Aurora
	// Serverless DB cluster, it automatically resumes. Constraints:
	//
	// * For Aurora
	// MySQL, valid capacity values are 1, 2, 4, 8, 16, 32, 64, 128, and 256.
	//
	// * For
	// Aurora PostgreSQL, valid capacity values are 2, 4, 8, 16, 32, 64, 192, and 384.
	Capacity *int32

	// The amount of time, in seconds, that Aurora Serverless tries to find a scaling
	// point to perform seamless scaling before enforcing the timeout action. The
	// default is 300.
	//
	// * Value must be from 10 through 600.
	SecondsBeforeTimeout *int32

	// The action to take when the timeout is reached, either ForceApplyCapacityChange
	// or RollbackCapacityChange. ForceApplyCapacityChange, the default, sets the
	// capacity to the specified value as soon as possible. RollbackCapacityChange
	// ignores the capacity change if a scaling point isn't found in the timeout
	// period.
	TimeoutAction *string
}

type ModifyCurrentDBClusterCapacityOutput struct {

	// The current capacity of the DB cluster.
	CurrentCapacity *int32

	// A user-supplied DB cluster identifier. This identifier is the unique key that
	// identifies a DB cluster.
	DBClusterIdentifier *string

	// A value that specifies the capacity that the DB cluster scales to next.
	PendingCapacity *int32

	// The number of seconds before a call to ModifyCurrentDBClusterCapacity times out.
	SecondsBeforeTimeout *int32

	// The timeout action of a call to ModifyCurrentDBClusterCapacity, either
	// ForceApplyCapacityChange or RollbackCapacityChange.
	TimeoutAction *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyCurrentDBClusterCapacityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyCurrentDBClusterCapacity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyCurrentDBClusterCapacity{}, middleware.After)
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
	addOpModifyCurrentDBClusterCapacityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyCurrentDBClusterCapacity(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opModifyCurrentDBClusterCapacity(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyCurrentDBClusterCapacity",
	}
}
