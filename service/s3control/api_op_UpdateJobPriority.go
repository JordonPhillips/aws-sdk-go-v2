// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"strings"
)

// Updates an existing S3 Batch Operations job's priority. For more information,
// see S3 Batch Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-basics.html) in the
// Amazon Simple Storage Service Developer Guide. Related actions include:
//
// *
// CreateJob
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateJob.html)
//
// *
// ListJobs
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListJobs.html)
//
// *
// DescribeJob
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DescribeJob.html)
//
// *
// UpdateJobStatus
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_UpdateJobStatus.html)
func (c *Client) UpdateJobPriority(ctx context.Context, params *UpdateJobPriorityInput, optFns ...func(*Options)) (*UpdateJobPriorityOutput, error) {
	if params == nil {
		params = &UpdateJobPriorityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateJobPriority", params, optFns, addOperationUpdateJobPriorityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateJobPriorityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateJobPriorityInput struct {

	//
	//
	// This member is required.
	AccountId *string

	// The ID for the job whose priority you want to update.
	//
	// This member is required.
	JobId *string

	// The priority you want to assign to this job.
	//
	// This member is required.
	Priority *int32
}

type UpdateJobPriorityOutput struct {

	// The ID for the job whose priority Amazon S3 updated.
	//
	// This member is required.
	JobId *string

	// The new priority assigned to the specified job.
	//
	// This member is required.
	Priority *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateJobPriorityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpUpdateJobPriority{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpUpdateJobPriority{}, middleware.After)
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
	addEndpointPrefix_opUpdateJobPriorityMiddleware(stack)
	addOpUpdateJobPriorityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateJobPriority(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	return nil
}

type endpointPrefix_opUpdateJobPriorityMiddleware struct {
}

func (*endpointPrefix_opUpdateJobPriorityMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opUpdateJobPriorityMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*UpdateJobPriorityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", in.Parameters)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.HostPrefix = prefix.String()

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opUpdateJobPriorityMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opUpdateJobPriorityMiddleware{}, `OperationSerializer`, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateJobPriority(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "UpdateJobPriority",
	}
}
