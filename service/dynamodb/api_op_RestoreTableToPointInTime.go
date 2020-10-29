// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Restores the specified table to the specified point in time within
// EarliestRestorableDateTime and LatestRestorableDateTime. You can restore your
// table to any point in time during the last 35 days. Any number of users can
// execute up to 4 concurrent restores (any type of restore) in a given account.
// When you restore using point in time recovery, DynamoDB restores your table data
// to the state based on the selected date and time (day:hour:minute:second) to a
// new table. Along with data, the following are also included on the new restored
// table using point in time recovery:
//
// * Global secondary indexes (GSIs)
//
// * Local
// secondary indexes (LSIs)
//
// * Provisioned read and write capacity
//
// * Encryption
// settings All these settings come from the current settings of the source table
// at the time of restore.
//
// You must manually set up the following on the restored
// table:
//
// * Auto scaling policies
//
// * IAM policies
//
// * Amazon CloudWatch metrics and
// alarms
//
// * Tags
//
// * Stream settings
//
// * Time to Live (TTL) settings
//
// * Point in
// time recovery settings
func (c *Client) RestoreTableToPointInTime(ctx context.Context, params *RestoreTableToPointInTimeInput, optFns ...func(*Options)) (*RestoreTableToPointInTimeOutput, error) {
	if params == nil {
		params = &RestoreTableToPointInTimeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreTableToPointInTime", params, optFns, addOperationRestoreTableToPointInTimeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreTableToPointInTimeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreTableToPointInTimeInput struct {

	// The name of the new table to which it must be restored to.
	//
	// This member is required.
	TargetTableName *string

	// The billing mode of the restored table.
	BillingModeOverride types.BillingMode

	// List of global secondary indexes for the restored table. The indexes provided
	// should match existing secondary indexes. You can choose to exclude some or all
	// of the indexes at the time of restore.
	GlobalSecondaryIndexOverride []*types.GlobalSecondaryIndex

	// List of local secondary indexes for the restored table. The indexes provided
	// should match existing secondary indexes. You can choose to exclude some or all
	// of the indexes at the time of restore.
	LocalSecondaryIndexOverride []*types.LocalSecondaryIndex

	// Provisioned throughput settings for the restored table.
	ProvisionedThroughputOverride *types.ProvisionedThroughput

	// Time in the past to restore the table to.
	RestoreDateTime *time.Time

	// The new server-side encryption settings for the restored table.
	SSESpecificationOverride *types.SSESpecification

	// The DynamoDB table that will be restored. This value is an Amazon Resource Name
	// (ARN).
	SourceTableArn *string

	// Name of the source table that is being restored.
	SourceTableName *string

	// Restore the table to the latest possible time. LatestRestorableDateTime is
	// typically 5 minutes before the current time.
	UseLatestRestorableTime *bool
}

type RestoreTableToPointInTimeOutput struct {

	// Represents the properties of a table.
	TableDescription *types.TableDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRestoreTableToPointInTimeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpRestoreTableToPointInTime{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpRestoreTableToPointInTime{}, middleware.After)
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
	addOpRestoreTableToPointInTimeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreTableToPointInTime(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addValidateResponseChecksum(stack, options)
	addAcceptEncodingGzip(stack, options)
	return nil
}

func newServiceMetadataMiddleware_opRestoreTableToPointInTime(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "RestoreTableToPointInTime",
	}
}
