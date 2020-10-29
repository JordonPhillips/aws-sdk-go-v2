// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Suspends the specified automatic scaling processes, or all processes, for the
// specified Auto Scaling group. If you suspend either the Launch or Terminate
// process types, it can prevent other process types from functioning properly. For
// more information, see Suspending and Resuming Scaling Processes
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-suspend-resume-processes.html)
// in the Amazon EC2 Auto Scaling User Guide. To resume processes that have been
// suspended, call the ResumeProcesses API.
func (c *Client) SuspendProcesses(ctx context.Context, params *SuspendProcessesInput, optFns ...func(*Options)) (*SuspendProcessesOutput, error) {
	if params == nil {
		params = &SuspendProcessesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SuspendProcesses", params, optFns, addOperationSuspendProcessesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SuspendProcessesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SuspendProcessesInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// One or more of the following processes:
	//
	// * Launch
	//
	// * Terminate
	//
	// *
	// AddToLoadBalancer
	//
	// * AlarmNotification
	//
	// * AZRebalance
	//
	// * HealthCheck
	//
	// *
	// InstanceRefresh
	//
	// * ReplaceUnhealthy
	//
	// * ScheduledActions
	//
	// If you omit this
	// parameter, all processes are specified.
	ScalingProcesses []*string
}

type SuspendProcessesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSuspendProcessesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSuspendProcesses{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSuspendProcesses{}, middleware.After)
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
	addOpSuspendProcessesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSuspendProcesses(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSuspendProcesses(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "SuspendProcesses",
	}
}
