// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of PackageSummary
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageSummary.html)
// objects for packages in a repository that match the request parameters.
func (c *Client) ListPackages(ctx context.Context, params *ListPackagesInput, optFns ...func(*Options)) (*ListPackagesOutput, error) {
	if params == nil {
		params = &ListPackagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPackages", params, optFns, addOperationListPackagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPackagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPackagesInput struct {

	// The domain that contains the repository that contains the requested list of
	// packages.
	//
	// This member is required.
	Domain *string

	// The name of the repository from which packages are to be listed.
	//
	// This member is required.
	Repository *string

	// The 12-digit account number of the AWS account that owns the domain. It does not
	// include dashes or spaces.
	DomainOwner *string

	// The format of the packages. The valid package types are:
	//
	// * npm: A Node Package
	// Manager (npm) package.
	//
	// * pypi: A Python Package Index (PyPI) package.
	//
	// * maven:
	// A Maven package that contains compiled code in a distributable format, such as a
	// JAR file.
	Format types.PackageFormat

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
	Namespace *string

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// A prefix used to filter returned repositories. Only repositories with names that
	// start with repositoryPrefix are returned.
	PackagePrefix *string
}

type ListPackagesOutput struct {

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// The list of returned PackageSummary
	// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageSummary.html)
	// objects.
	Packages []*types.PackageSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListPackagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPackages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPackages{}, middleware.After)
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
	addOpListPackagesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPackages(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListPackages(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "ListPackages",
	}
}
