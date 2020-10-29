// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Gets a list of the Git repositories in your account.
func (c *Client) ListCodeRepositories(ctx context.Context, params *ListCodeRepositoriesInput, optFns ...func(*Options)) (*ListCodeRepositoriesOutput, error) {
	if params == nil {
		params = &ListCodeRepositoriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCodeRepositories", params, optFns, addOperationListCodeRepositoriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCodeRepositoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCodeRepositoriesInput struct {

	// A filter that returns only Git repositories that were created after the
	// specified time.
	CreationTimeAfter *time.Time

	// A filter that returns only Git repositories that were created before the
	// specified time.
	CreationTimeBefore *time.Time

	// A filter that returns only Git repositories that were last modified after the
	// specified time.
	LastModifiedTimeAfter *time.Time

	// A filter that returns only Git repositories that were last modified before the
	// specified time.
	LastModifiedTimeBefore *time.Time

	// The maximum number of Git repositories to return in the response.
	MaxResults *int32

	// A string in the Git repositories name. This filter returns only repositories
	// whose name contains the specified string.
	NameContains *string

	// If the result of a ListCodeRepositoriesOutput request was truncated, the
	// response includes a NextToken. To get the next set of Git repositories, use the
	// token in the next request.
	NextToken *string

	// The field to sort results by. The default is Name.
	SortBy types.CodeRepositorySortBy

	// The sort order for results. The default is Ascending.
	SortOrder types.CodeRepositorySortOrder
}

type ListCodeRepositoriesOutput struct {

	// Gets a list of summaries of the Git repositories. Each summary specifies the
	// following values for the repository:
	//
	// * Name
	//
	// * Amazon Resource Name (ARN)
	//
	// *
	// Creation time
	//
	// * Last modified time
	//
	// * Configuration information, including the
	// URL location of the repository and the ARN of the AWS Secrets Manager secret
	// that contains the credentials used to access the repository.
	//
	// This member is required.
	CodeRepositorySummaryList []*types.CodeRepositorySummary

	// If the result of a ListCodeRepositoriesOutput request was truncated, the
	// response includes a NextToken. To get the next set of Git repositories, use the
	// token in the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListCodeRepositoriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCodeRepositories{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCodeRepositories{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListCodeRepositories(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListCodeRepositories(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListCodeRepositories",
	}
}
