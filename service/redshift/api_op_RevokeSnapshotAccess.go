// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the ability of the specified AWS customer account to restore the
// specified snapshot. If the account is currently restoring the snapshot, the
// restore will run to completion. For more information about working with
// snapshots, go to Amazon Redshift Snapshots
// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-snapshots.html)
// in the Amazon Redshift Cluster Management Guide.
func (c *Client) RevokeSnapshotAccess(ctx context.Context, params *RevokeSnapshotAccessInput, optFns ...func(*Options)) (*RevokeSnapshotAccessOutput, error) {
	stack := middleware.NewStack("RevokeSnapshotAccess", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpRevokeSnapshotAccessMiddlewares(stack)
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
	addOpRevokeSnapshotAccessValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRevokeSnapshotAccess(options.Region), middleware.Before)

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
			OperationName: "RevokeSnapshotAccess",
			Err:           err,
		}
	}
	out := result.(*RevokeSnapshotAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type RevokeSnapshotAccessInput struct {
	// The identifier of the snapshot that the account can no longer access.
	SnapshotIdentifier *string
	// The identifier of the AWS customer account that can no longer restore the
	// specified snapshot.
	AccountWithRestoreAccess *string
	// The identifier of the cluster the snapshot was created from. This parameter is
	// required if your IAM user has a policy containing a snapshot resource element
	// that specifies anything other than * for the cluster name.
	SnapshotClusterIdentifier *string
}

type RevokeSnapshotAccessOutput struct {
	// Describes a snapshot.
	Snapshot *types.Snapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpRevokeSnapshotAccessMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpRevokeSnapshotAccess{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpRevokeSnapshotAccess{}, middleware.After)
}

func newServiceMetadataMiddleware_opRevokeSnapshotAccess(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "RevokeSnapshotAccess",
	}
}