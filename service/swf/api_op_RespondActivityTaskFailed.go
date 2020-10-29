// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Used by workers to tell the service that the ActivityTask identified by the
// taskToken has failed with reason (if specified). The reason and details appear
// in the ActivityTaskFailed event added to the workflow history. A task is
// considered open from the time that it is scheduled until it is closed. Therefore
// a task is reported as open while a worker is processing it. A task is closed
// after it has been specified in a call to RespondActivityTaskCompleted,
// RespondActivityTaskCanceled, RespondActivityTaskFailed, or the task has timed
// out
// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dg-basic.html#swf-dev-timeout-types).
// Access Control You can use IAM policies to control this action's access to
// Amazon SWF resources as follows:
//
// * Use a Resource element with the domain name
// to limit the action to only specified domains.
//
// * Use an Action element to allow
// or deny permission to call this action.
//
// * You cannot use an IAM policy to
// constrain this action's parameters.
//
// If the caller doesn't have sufficient
// permissions to invoke the action, or the parameter values fall outside the
// specified constraints, the action fails. The associated event attribute's cause
// parameter is set to OPERATION_NOT_PERMITTED. For details and example IAM
// policies, see Using IAM to Manage Access to Amazon SWF Workflows
// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) RespondActivityTaskFailed(ctx context.Context, params *RespondActivityTaskFailedInput, optFns ...func(*Options)) (*RespondActivityTaskFailedOutput, error) {
	if params == nil {
		params = &RespondActivityTaskFailedInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RespondActivityTaskFailed", params, optFns, addOperationRespondActivityTaskFailedMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RespondActivityTaskFailedOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RespondActivityTaskFailedInput struct {

	// The taskToken of the ActivityTask. taskToken is generated by the service and
	// should be treated as an opaque value. If the task is passed to another process,
	// its taskToken must also be passed. This enables it to provide its progress and
	// respond with results.
	//
	// This member is required.
	TaskToken *string

	// Detailed information about the failure.
	Details *string

	// Description of the error that may assist in diagnostics.
	Reason *string
}

type RespondActivityTaskFailedOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRespondActivityTaskFailedMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpRespondActivityTaskFailed{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpRespondActivityTaskFailed{}, middleware.After)
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
	addOpRespondActivityTaskFailedValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRespondActivityTaskFailed(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRespondActivityTaskFailed(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "RespondActivityTaskFailed",
	}
}
