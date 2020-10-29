// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the estimated number of decision tasks in the specified task list. The
// count returned is an approximation and isn't guaranteed to be exact. If you
// specify a task list that no decision task was ever scheduled in then 0 is
// returned. Access Control You can use IAM policies to control this action's
// access to Amazon SWF resources as follows:
//
// * Use a Resource element with the
// domain name to limit the action to only specified domains.
//
// * Use an Action
// element to allow or deny permission to call this action.
//
// * Constrain the
// taskList.name parameter by using a Condition element with the swf:taskList.name
// key to allow the action to access only certain task lists.
//
// If the caller
// doesn't have sufficient permissions to invoke the action, or the parameter
// values fall outside the specified constraints, the action fails. The associated
// event attribute's cause parameter is set to OPERATION_NOT_PERMITTED. For details
// and example IAM policies, see Using IAM to Manage Access to Amazon SWF Workflows
// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) CountPendingDecisionTasks(ctx context.Context, params *CountPendingDecisionTasksInput, optFns ...func(*Options)) (*CountPendingDecisionTasksOutput, error) {
	if params == nil {
		params = &CountPendingDecisionTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CountPendingDecisionTasks", params, optFns, addOperationCountPendingDecisionTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CountPendingDecisionTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CountPendingDecisionTasksInput struct {

	// The name of the domain that contains the task list.
	//
	// This member is required.
	Domain *string

	// The name of the task list.
	//
	// This member is required.
	TaskList *types.TaskList
}

// Contains the count of tasks in a task list.
type CountPendingDecisionTasksOutput struct {

	// The number of tasks in the task list.
	//
	// This member is required.
	Count *int32

	// If set to true, indicates that the actual count was more than the maximum
	// supported by this API and the count returned is the truncated value.
	Truncated *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCountPendingDecisionTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCountPendingDecisionTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCountPendingDecisionTasks{}, middleware.After)
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
	addOpCountPendingDecisionTasksValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCountPendingDecisionTasks(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCountPendingDecisionTasks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "CountPendingDecisionTasks",
	}
}
