// Code generated by smithy-go-codegen DO NOT EDIT.

package qldb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Exports journal contents within a date and time range from a ledger into a
// specified Amazon Simple Storage Service (Amazon S3) bucket. The data is written
// as files in Amazon Ion format. If the ledger with the given Name doesn't exist,
// then throws ResourceNotFoundException. If the ledger with the given Name is in
// CREATING status, then throws ResourcePreconditionNotMetException. You can
// initiate up to two concurrent journal export requests for each ledger. Beyond
// this limit, journal export requests throw LimitExceededException.
func (c *Client) ExportJournalToS3(ctx context.Context, params *ExportJournalToS3Input, optFns ...func(*Options)) (*ExportJournalToS3Output, error) {
	if params == nil {
		params = &ExportJournalToS3Input{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExportJournalToS3", params, optFns, addOperationExportJournalToS3Middlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExportJournalToS3Output)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportJournalToS3Input struct {

	// The exclusive end date and time for the range of journal contents that you want
	// to export. The ExclusiveEndTime must be in ISO 8601 date and time format and in
	// Universal Coordinated Time (UTC). For example: 2019-06-13T21:36:34Z The
	// ExclusiveEndTime must be less than or equal to the current UTC date and time.
	//
	// This member is required.
	ExclusiveEndTime *time.Time

	// The inclusive start date and time for the range of journal contents that you
	// want to export. The InclusiveStartTime must be in ISO 8601 date and time format
	// and in Universal Coordinated Time (UTC). For example: 2019-06-13T21:36:34Z The
	// InclusiveStartTime must be before ExclusiveEndTime. If you provide an
	// InclusiveStartTime that is before the ledger's CreationDateTime, Amazon QLDB
	// defaults it to the ledger's CreationDateTime.
	//
	// This member is required.
	InclusiveStartTime *time.Time

	// The name of the ledger.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of the IAM role that grants QLDB permissions for
	// a journal export job to do the following:
	//
	// * Write objects into your Amazon
	// Simple Storage Service (Amazon S3) bucket.
	//
	// * (Optional) Use your customer
	// master key (CMK) in AWS Key Management Service (AWS KMS) for server-side
	// encryption of your exported data.
	//
	// This member is required.
	RoleArn *string

	// The configuration settings of the Amazon S3 bucket destination for your export
	// request.
	//
	// This member is required.
	S3ExportConfiguration *types.S3ExportConfiguration
}

type ExportJournalToS3Output struct {

	// The unique ID that QLDB assigns to each journal export job. To describe your
	// export request and check the status of the job, you can use ExportId to call
	// DescribeJournalS3Export.
	//
	// This member is required.
	ExportId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationExportJournalToS3Middlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpExportJournalToS3{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpExportJournalToS3{}, middleware.After)
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
	addOpExportJournalToS3ValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opExportJournalToS3(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opExportJournalToS3(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "ExportJournalToS3",
	}
}
