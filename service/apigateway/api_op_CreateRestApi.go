// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a new RestApi resource.
func (c *Client) CreateRestApi(ctx context.Context, params *CreateRestApiInput, optFns ...func(*Options)) (*CreateRestApiOutput, error) {
	if params == nil {
		params = &CreateRestApiInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRestApi", params, optFns, addOperationCreateRestApiMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRestApiOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The POST Request to add a new RestApi resource to your collection.
type CreateRestApiInput struct {

	// [Required] The name of the RestApi.
	//
	// This member is required.
	Name *string

	// The source of the API key for metering requests according to a usage plan. Valid
	// values are:
	//
	// * HEADER to read the API key from the X-API-Key header of a
	// request.
	//
	// * AUTHORIZER to read the API key from the UsageIdentifierKey from a
	// custom authorizer.
	ApiKeySource types.ApiKeySourceType

	// The list of binary media types supported by the RestApi. By default, the RestApi
	// supports only UTF-8-encoded text payloads.
	BinaryMediaTypes []*string

	// The ID of the RestApi that you want to clone from.
	CloneFrom *string

	// The description of the RestApi.
	Description *string

	// The endpoint configuration of this RestApi showing the endpoint types of the
	// API.
	EndpointConfiguration *types.EndpointConfiguration

	// A nullable integer that is used to enable compression (with non-negative between
	// 0 and 10485760 (10M) bytes, inclusive) or disable compression (with a null
	// value) on an API. When compression is enabled, compression or decompression is
	// not applied on the payload if the payload size is smaller than this value.
	// Setting it to zero allows compression for any payload size.
	MinimumCompressionSize *int32

	// A stringified JSON policy document that applies to this RestApi regardless of
	// the caller and Method configuration.
	Policy *string

	// The key-value map of strings. The valid character set is [a-zA-Z+-=._:/]. The
	// tag key can be up to 128 characters and must not start with aws:. The tag value
	// can be up to 256 characters.
	Tags map[string]*string

	Template *bool

	TemplateSkipList []*string

	Title *string

	// A version identifier for the API.
	Version *string
}

// Represents a REST API. Create an API
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-create-api.html)
type CreateRestApiOutput struct {

	// The source of the API key for metering requests according to a usage plan. Valid
	// values are:
	//
	// * HEADER to read the API key from the X-API-Key header of a
	// request.
	//
	// * AUTHORIZER to read the API key from the UsageIdentifierKey from a
	// custom authorizer.
	ApiKeySource types.ApiKeySourceType

	// The list of binary media types supported by the RestApi. By default, the RestApi
	// supports only UTF-8-encoded text payloads.
	BinaryMediaTypes []*string

	// The timestamp when the API was created.
	CreatedDate *time.Time

	// The API's description.
	Description *string

	// The endpoint configuration of this RestApi showing the endpoint types of the
	// API.
	EndpointConfiguration *types.EndpointConfiguration

	// The API's identifier. This identifier is unique across all of your APIs in API
	// Gateway.
	Id *string

	// A nullable integer that is used to enable compression (with non-negative between
	// 0 and 10485760 (10M) bytes, inclusive) or disable compression (with a null
	// value) on an API. When compression is enabled, compression or decompression is
	// not applied on the payload if the payload size is smaller than this value.
	// Setting it to zero allows compression for any payload size.
	MinimumCompressionSize *int32

	// The API's name.
	Name *string

	// A stringified JSON policy document that applies to this RestApi regardless of
	// the caller and Method configuration.
	Policy *string

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]*string

	// A version identifier for the API.
	Version *string

	// The warning messages reported when failonwarnings is turned on during API
	// import.
	Warnings []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateRestApiMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateRestApi{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRestApi{}, middleware.After)
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
	addOpCreateRestApiValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRestApi(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateRestApi(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "CreateRestApi",
	}
}
