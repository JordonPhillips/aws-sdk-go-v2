// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of AWS Batch jobs. You must specify only one of the following:
//
// *
// a job queue ID to return a list of jobs in that job queue
//
// * a multi-node
// parallel job ID to return a list of that job's nodes
//
// * an array job ID to
// return a list of that job's children
//
// You can filter the results by job status
// with the jobStatus parameter. If you do not specify a status, only RUNNING jobs
// are returned.
func (c *Client) ListJobs(ctx context.Context, params *ListJobsInput, optFns ...func(*Options)) (*ListJobsOutput, error) {
	if params == nil {
		params = &ListJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListJobs", params, optFns, addOperationListJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListJobsInput struct {

	// The job ID for an array job. Specifying an array job ID with this parameter
	// lists all child jobs from within the specified array.
	ArrayJobId *string

	// The name or full Amazon Resource Name (ARN) of the job queue with which to list
	// jobs.
	JobQueue *string

	// The job status with which to filter jobs in the specified queue. If you do not
	// specify a status, only RUNNING jobs are returned.
	JobStatus types.JobStatus

	// The maximum number of results returned by ListJobs in paginated output. When
	// this parameter is used, ListJobs only returns maxResults results in a single
	// page along with a nextToken response element. The remaining results of the
	// initial request can be seen by sending another ListJobs request with the
	// returned nextToken value. This value can be between 1 and 100. If this parameter
	// is not used, then ListJobs returns up to 100 results and a nextToken value if
	// applicable.
	MaxResults *int32

	// The job ID for a multi-node parallel job. Specifying a multi-node parallel job
	// ID with this parameter lists all nodes that are associated with the specified
	// job.
	MultiNodeJobId *string

	// The nextToken value returned from a previous paginated ListJobs request where
	// maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value. This value is null when there are no more results to return.
	// This token should be treated as an opaque identifier that is only used to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string
}

type ListJobsOutput struct {

	// A list of job summaries that match the request.
	//
	// This member is required.
	JobSummaryList []*types.JobSummary

	// The nextToken value to include in a future ListJobs request. When the results of
	// a ListJobs request exceed maxResults, this value can be used to retrieve the
	// next page of results. This value is null when there are no more results to
	// return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListJobs{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListJobs(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "batch",
		OperationName: "ListJobs",
	}
}
