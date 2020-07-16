// Code generated by smithy-go-codegen DO NOT EDIT.
package dynamodbgov2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbgov2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The UpdateTimeToLive method enables or disables Time to Live (TTL) for the
// specified table. A successful UpdateTimeToLive call returns the current
// TimeToLiveSpecification. It can take up to one hour for the change to fully
// process. Any additional UpdateTimeToLive calls for the same table during this
// one hour duration result in a ValidationException. TTL compares the current time
// in epoch time format to the time stored in the TTL attribute of an item. If the
// epoch time value stored in the attribute is less than the current time, the item
// is marked as expired and subsequently deleted. The epoch time format is the
// number of seconds elapsed since 12:00:00 AM January 1, 1970 UTC. DynamoDB
// deletes expired items on a best-effort basis to ensure availability of
// throughput for other data operations. DynamoDB typically deletes expired items
// within two days of expiration. The exact duration within which an item gets
// deleted after expiration is specific to the nature of the workload. Items that
// have expired and not been deleted will still show up in reads, queries, and
// scans. As items are deleted, they are removed from any local secondary index and
// global secondary index immediately in the same eventually consistent way as a
// standard delete operation. For more information, see Time To Live
// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/TTL.html) in
// the Amazon DynamoDB Developer Guide.
func (c *Client) UpdateTimeToLive(ctx context.Context, params *UpdateTimeToLiveInput, optFns ...func(*Options)) (*UpdateTimeToLiveOutput, error) {
	stack := middleware.NewStack("UpdateTimeToLive", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	awsmiddleware.AddResolveServiceEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	v4.AddHTTPSignerMiddleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addOpUpdateTimeToLiveValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTimeToLive(options.Region), middleware.Before)
	addawsAwsjson10_serdeOpUpdateTimeToLiveMiddlewares(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "UpdateTimeToLive",
			Err:           err,
		}
	}
	out := result.(*UpdateTimeToLiveOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of an UpdateTimeToLive operation.
type UpdateTimeToLiveInput struct {
	// The name of the table to be configured.
	TableName *string
	// Represents the settings used to enable or disable Time to Live for the specified
	// table.
	TimeToLiveSpecification *types.TimeToLiveSpecification
}

type UpdateTimeToLiveOutput struct {
	// Represents the output of an UpdateTimeToLive operation.
	TimeToLiveSpecification *types.TimeToLiveSpecification

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpUpdateTimeToLiveMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateTimeToLive{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateTimeToLive{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateTimeToLive(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "DynamoDB GoV2",
		ServiceID:      "dynamodbgov2",
		EndpointPrefix: "dynamodbgov2",
		SigningName:    "dynamodb",
		OperationName:  "UpdateTimeToLive",
	}
}
