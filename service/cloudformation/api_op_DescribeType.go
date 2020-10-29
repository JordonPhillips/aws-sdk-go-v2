// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns detailed information about a type that has been registered. If you
// specify a VersionId, DescribeType returns information about that specific type
// version. Otherwise, it returns information about the default type version.
func (c *Client) DescribeType(ctx context.Context, params *DescribeTypeInput, optFns ...func(*Options)) (*DescribeTypeOutput, error) {
	if params == nil {
		params = &DescribeTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeType", params, optFns, addOperationDescribeTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTypeInput struct {

	// The Amazon Resource Name (ARN) of the type. Conditional: You must specify either
	// TypeName and Type, or Arn.
	Arn *string

	// The kind of type. Currently the only valid value is RESOURCE. Conditional: You
	// must specify either TypeName and Type, or Arn.
	Type types.RegistryType

	// The name of the type. Conditional: You must specify either TypeName and Type, or
	// Arn.
	TypeName *string

	// The ID of a specific version of the type. The version ID is the value at the end
	// of the Amazon Resource Name (ARN) assigned to the type version when it is
	// registered. If you specify a VersionId, DescribeType returns information about
	// that specific type version. Otherwise, it returns information about the default
	// type version.
	VersionId *string
}

type DescribeTypeOutput struct {

	// The Amazon Resource Name (ARN) of the type.
	Arn *string

	// The ID of the default version of the type. The default version is used when the
	// type version is not specified. To set the default version of a type, use
	// SetTypeDefaultVersion.
	DefaultVersionId *string

	// The deprecation status of the type. Valid values include:
	//
	// * LIVE: The type is
	// registered and can be used in CloudFormation operations, dependent on its
	// provisioning behavior and visibility scope.
	//
	// * DEPRECATED: The type has been
	// deregistered and can no longer be used in CloudFormation operations.
	DeprecatedStatus types.DeprecatedStatus

	// The description of the registered type.
	Description *string

	// The URL of a page providing detailed documentation for this type.
	DocumentationUrl *string

	// The Amazon Resource Name (ARN) of the IAM execution role used to register the
	// type. If your resource type calls AWS APIs in any of its handlers, you must
	// create an IAM execution role
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html) that includes
	// the necessary permissions to call those AWS APIs, and provision that execution
	// role in your account. CloudFormation then assumes that execution role to provide
	// your resource type with the appropriate credentials.
	ExecutionRoleArn *string

	// Whether the specified type version is set as the default version.
	IsDefaultVersion *bool

	// When the specified type version was registered.
	LastUpdated *time.Time

	// Contains logging configuration information for a type.
	LoggingConfig *types.LoggingConfig

	// The provisioning behavior of the type. AWS CloudFormation determines the
	// provisioning type during registration, based on the types of handlers in the
	// schema handler package submitted. Valid values include:
	//
	// * FULLY_MUTABLE: The
	// type includes an update handler to process updates to the type during stack
	// update operations.
	//
	// * IMMUTABLE: The type does not include an update handler, so
	// the type cannot be updated and must instead be replaced during stack update
	// operations.
	//
	// * NON_PROVISIONABLE: The type does not include all of the following
	// handlers, and therefore cannot actually be provisioned.
	//
	// * create
	//
	// * read
	//
	// *
	// delete
	ProvisioningType types.ProvisioningType

	// The schema that defines the type. For more information on type schemas, see
	// Resource Provider Schema
	// (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-schema.html)
	// in the CloudFormation CLI User Guide.
	Schema *string

	// The URL of the source code for the type.
	SourceUrl *string

	// When the specified type version was registered.
	TimeCreated *time.Time

	// The kind of type. Currently the only valid value is RESOURCE.
	Type types.RegistryType

	// The name of the registered type.
	TypeName *string

	// The scope at which the type is visible and usable in CloudFormation operations.
	// Valid values include:
	//
	// * PRIVATE: The type is only visible and usable within the
	// account in which it is registered. Currently, AWS CloudFormation marks any types
	// you register as PRIVATE.
	//
	// * PUBLIC: The type is publically visible and usable
	// within any Amazon account.
	Visibility types.Visibility

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeType{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeType(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeType(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DescribeType",
	}
}
