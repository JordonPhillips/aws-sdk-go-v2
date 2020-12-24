// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of PackageVersionSummary
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageVersionSummary.html)
// objects for package versions in a repository that match the request parameters.
func (c *Client) ListPackageVersions(ctx context.Context, params *ListPackageVersionsInput, optFns ...func(*Options)) (*ListPackageVersionsOutput, error) {
	if params == nil {
		params = &ListPackageVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPackageVersions", params, optFns, addOperationListPackageVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPackageVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPackageVersionsInput struct {

	// The name of the domain that contains the repository that contains the returned
	// package versions.
	//
	// This member is required.
	Domain *string

	// The format of the returned packages. The valid package types are:
	//
	// * npm: A Node
	// Package Manager (npm) package.
	//
	// * pypi: A Python Package Index (PyPI)
	// package.
	//
	// * maven: A Maven package that contains compiled code in a
	// distributable format, such as a JAR file.
	//
	// * nuget: A NuGet package.
	//
	// This member is required.
	Format types.PackageFormat

	// The name of the package for which you want to return a list of package versions.
	//
	// This member is required.
	Package *string

	// The name of the repository that contains the package.
	//
	// This member is required.
	Repository *string

	// The 12-digit account number of the AWS account that owns the domain. It does not
	// include dashes or spaces.
	DomainOwner *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The namespace of the package. The package component that specifies its namespace
	// depends on its type. For example:
	//
	// * The namespace of a Maven package is its
	// groupId.
	//
	// * The namespace of an npm package is its scope.
	//
	// * A Python package
	// does not contain a corresponding component, so Python packages do not have a
	// namespace.
	//
	// * A NuGet package does not contain a corresponding component, so
	// NuGet packages do not have a namespace.
	Namespace *string

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// How to sort the returned list of package versions.
	SortBy types.PackageVersionSortType

	// A string that specifies the status of the package versions to include in the
	// returned list. It can be one of the following:
	//
	// * Published
	//
	// * Unfinished
	//
	// *
	// Unlisted
	//
	// * Archived
	//
	// * Disposed
	Status types.PackageVersionStatus
}

type ListPackageVersionsOutput struct {

	// The default package version to display. This depends on the package format:
	//
	// *
	// For Maven and PyPI packages, it's the most recently published package
	// version.
	//
	// * For npm packages, it's the version referenced by the latest tag. If
	// the latest tag is not set, it's the most recently published package version.
	DefaultDisplayVersion *string

	// A format of the package. Valid package format values are:
	//
	// * npm
	//
	// * pypi
	//
	// *
	// maven
	//
	// * nuget
	Format types.PackageFormat

	// The namespace of the package. The package component that specifies its namespace
	// depends on its type. For example:
	//
	// * The namespace of a Maven package is its
	// groupId.
	//
	// * The namespace of an npm package is its scope.
	//
	// * A Python package
	// does not contain a corresponding component, so Python packages do not have a
	// namespace.
	//
	// * A NuGet package does not contain a corresponding component, so
	// NuGet packages do not have a namespace.
	Namespace *string

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// The name of the package.
	Package *string

	// The returned list of PackageVersionSummary
	// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageVersionSummary.html)
	// objects.
	Versions []types.PackageVersionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListPackageVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPackageVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPackageVersions{}, middleware.After)
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
	if err = addOpListPackageVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPackageVersions(options.Region), middleware.Before); err != nil {
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

// ListPackageVersionsAPIClient is a client that implements the ListPackageVersions
// operation.
type ListPackageVersionsAPIClient interface {
	ListPackageVersions(context.Context, *ListPackageVersionsInput, ...func(*Options)) (*ListPackageVersionsOutput, error)
}

var _ ListPackageVersionsAPIClient = (*Client)(nil)

// ListPackageVersionsPaginatorOptions is the paginator options for
// ListPackageVersions
type ListPackageVersionsPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPackageVersionsPaginator is a paginator for ListPackageVersions
type ListPackageVersionsPaginator struct {
	options   ListPackageVersionsPaginatorOptions
	client    ListPackageVersionsAPIClient
	params    *ListPackageVersionsInput
	nextToken *string
	firstPage bool
}

// NewListPackageVersionsPaginator returns a new ListPackageVersionsPaginator
func NewListPackageVersionsPaginator(client ListPackageVersionsAPIClient, params *ListPackageVersionsInput, optFns ...func(*ListPackageVersionsPaginatorOptions)) *ListPackageVersionsPaginator {
	options := ListPackageVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListPackageVersionsInput{}
	}

	return &ListPackageVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPackageVersionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListPackageVersions page.
func (p *ListPackageVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPackageVersionsOutput, error) {
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

	result, err := p.client.ListPackageVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPackageVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "ListPackageVersions",
	}
}
