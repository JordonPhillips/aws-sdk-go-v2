// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of delegations from an audit owner to a delegate.
func (c *Client) GetDelegations(ctx context.Context, params *GetDelegationsInput, optFns ...func(*Options)) (*GetDelegationsOutput, error) {
	if params == nil {
		params = &GetDelegationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDelegations", params, optFns, addOperationGetDelegationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDelegationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDelegationsInput struct {

	// Represents the maximum number of results per page, or per API request call.
	MaxResults *int32

	// The pagination token used to fetch the next set of results.
	NextToken *string
}

type GetDelegationsOutput struct {

	// The list of delegations returned by the GetDelegations API.
	Delegations []types.DelegationMetadata

	// The pagination token used to fetch the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDelegationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDelegations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDelegations{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDelegations(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// GetDelegationsAPIClient is a client that implements the GetDelegations
// operation.
type GetDelegationsAPIClient interface {
	GetDelegations(context.Context, *GetDelegationsInput, ...func(*Options)) (*GetDelegationsOutput, error)
}

var _ GetDelegationsAPIClient = (*Client)(nil)

// GetDelegationsPaginatorOptions is the paginator options for GetDelegations
type GetDelegationsPaginatorOptions struct {
	// Represents the maximum number of results per page, or per API request call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetDelegationsPaginator is a paginator for GetDelegations
type GetDelegationsPaginator struct {
	options   GetDelegationsPaginatorOptions
	client    GetDelegationsAPIClient
	params    *GetDelegationsInput
	nextToken *string
	firstPage bool
}

// NewGetDelegationsPaginator returns a new GetDelegationsPaginator
func NewGetDelegationsPaginator(client GetDelegationsAPIClient, params *GetDelegationsInput, optFns ...func(*GetDelegationsPaginatorOptions)) *GetDelegationsPaginator {
	options := GetDelegationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &GetDelegationsInput{}
	}

	return &GetDelegationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetDelegationsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetDelegations page.
func (p *GetDelegationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetDelegationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.GetDelegations(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetDelegations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "auditmanager",
		OperationName: "GetDelegations",
	}
}
