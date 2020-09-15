// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deregisters an Amazon ECS container instance from the specified cluster. This
// instance is no longer available to run tasks. If you intend to use the container
// instance for some other purpose after deregistration, you should stop all of the
// tasks running on the container instance before deregistration. That prevents any
// orphaned tasks from consuming resources. Deregistering a container instance
// removes the instance from a cluster, but it does not terminate the EC2 instance.
// If you are finished using the instance, be sure to terminate it in the Amazon
// EC2 console to stop billing. If you terminate a running container instance,
// Amazon ECS automatically deregisters the instance from your cluster (stopped
// container instances or instances with disconnected agents are not automatically
// deregistered when terminated).
func (c *Client) DeregisterContainerInstance(ctx context.Context, params *DeregisterContainerInstanceInput, optFns ...func(*Options)) (*DeregisterContainerInstanceOutput, error) {
	stack := middleware.NewStack("DeregisterContainerInstance", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeregisterContainerInstanceMiddlewares(stack)
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
	addOpDeregisterContainerInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterContainerInstance(options.Region), middleware.Before)

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
			OperationName: "DeregisterContainerInstance",
			Err:           err,
		}
	}
	out := result.(*DeregisterContainerInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterContainerInstanceInput struct {
	// The container instance ID or full ARN of the container instance to deregister.
	// The ARN contains the arn:aws:ecs namespace, followed by the Region of the
	// container instance, the AWS account ID of the container instance owner, the
	// container-instance namespace, and then the container instance ID. For example,
	// arn:aws:ecs:region:aws_account_id:container-instance/container_instance_ID.
	ContainerInstance *string
	// Forces the deregistration of the container instance. If you have tasks running
	// on the container instance when you deregister it with the force option, these
	// tasks remain running until you terminate the instance or the tasks stop through
	// some other means, but they are orphaned (no longer monitored or accounted for by
	// Amazon ECS). If an orphaned task on your container instance is part of an Amazon
	// ECS service, then the service scheduler starts another copy of that task, on a
	// different container instance if possible. Any containers in orphaned service
	// tasks that are registered with a Classic Load Balancer or an Application Load
	// Balancer target group are deregistered. They begin connection draining according
	// to the settings on the load balancer or target group.
	Force *bool
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the
	// container instance to deregister. If you do not specify a cluster, the default
	// cluster is assumed.
	Cluster *string
}

type DeregisterContainerInstanceOutput struct {
	// The container instance that was deregistered.
	ContainerInstance *types.ContainerInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeregisterContainerInstanceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeregisterContainerInstance{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeregisterContainerInstance{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeregisterContainerInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "DeregisterContainerInstance",
	}
}