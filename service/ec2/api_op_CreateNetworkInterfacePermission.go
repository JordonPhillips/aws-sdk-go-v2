// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Grants an AWS-authorized account permission to attach the specified network
// interface to an instance in their account. You can grant permission to a single
// AWS account only, and only one account at a time.
func (c *Client) CreateNetworkInterfacePermission(ctx context.Context, params *CreateNetworkInterfacePermissionInput, optFns ...func(*Options)) (*CreateNetworkInterfacePermissionOutput, error) {
	stack := middleware.NewStack("CreateNetworkInterfacePermission", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpCreateNetworkInterfacePermissionMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpCreateNetworkInterfacePermissionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNetworkInterfacePermission(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "CreateNetworkInterfacePermission",
			Err:           err,
		}
	}
	out := result.(*CreateNetworkInterfacePermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CreateNetworkInterfacePermission.
type CreateNetworkInterfacePermissionInput struct {
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The ID of the network interface.
	NetworkInterfaceId *string
	// The AWS account ID.
	AwsAccountId *string
	// The AWS service. Currently not supported.
	AwsService *string
	// The type of permission to grant.
	Permission types.InterfacePermissionType
}

// Contains the output of CreateNetworkInterfacePermission.
type CreateNetworkInterfacePermissionOutput struct {
	// Information about the permission for the network interface.
	InterfacePermission *types.NetworkInterfacePermission

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpCreateNetworkInterfacePermissionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpCreateNetworkInterfacePermission{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpCreateNetworkInterfacePermission{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateNetworkInterfacePermission(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateNetworkInterfacePermission",
	}
}