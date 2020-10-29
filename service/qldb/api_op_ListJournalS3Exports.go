// Code generated by smithy-go-codegen DO NOT EDIT.

package qldb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns an array of journal export job descriptions for all ledgers that are
// associated with the current AWS account and Region. This action returns a
// maximum of MaxResults items, and is paginated so that you can retrieve all the
// items by calling ListJournalS3Exports multiple times. This action does not
// return any expired export jobs. For more information, see Export Job Expiration
// (https://docs.aws.amazon.com/qldb/latest/developerguide/export-journal.request.html#export-journal.request.expiration)
// in the Amazon QLDB Developer Guide.
func (c *Client) ListJournalS3Exports(ctx context.Context, params *ListJournalS3ExportsInput, optFns ...func(*Options)) (*ListJournalS3ExportsOutput, error) {
	if params == nil {
		params = &ListJournalS3ExportsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListJournalS3Exports", params, optFns, addOperationListJournalS3ExportsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListJournalS3ExportsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListJournalS3ExportsInput struct {

	// The maximum number of results to return in a single ListJournalS3Exports
	// request. (The actual number of results returned might be fewer.)
	MaxResults *int32

	// A pagination token, indicating that you want to retrieve the next page of
	// results. If you received a value for NextToken in the response from a previous
	// ListJournalS3Exports call, then you should use that value as input here.
	NextToken *string
}

type ListJournalS3ExportsOutput struct {

	// The array of journal export job descriptions for all ledgers that are associated
	// with the current AWS account and Region.
	JournalS3Exports []*types.JournalS3ExportDescription

	// * If NextToken is empty, then the last page of results has been processed and
	// there are no more results to be retrieved.
	//
	// * If NextToken is not empty, then
	// there are more results available. To retrieve the next page of results, use the
	// value of NextToken in a subsequent ListJournalS3Exports call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListJournalS3ExportsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListJournalS3Exports{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListJournalS3Exports{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListJournalS3Exports(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListJournalS3Exports(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "ListJournalS3Exports",
	}
}
