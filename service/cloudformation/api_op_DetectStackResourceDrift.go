// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about whether a resource's actual configuration differs, or
// has drifted, from it's expected configuration, as defined in the stack template
// and any values specified as template parameters. This information includes
// actual and expected property values for resources in which AWS CloudFormation
// detects drift. Only resource properties explicitly defined in the stack template
// are checked for drift. For more information about stack and resource drift, see
// Detecting Unregulated Configuration Changes to Stacks and Resources
// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-drift.html).
// Use DetectStackResourceDrift to detect drift on individual resources, or
// DetectStackDrift () to detect drift on all resources in a given stack that
// support drift detection. Resources that do not currently support drift detection
// cannot be checked. For a list of resources that support drift detection, see
// Resources that Support Drift Detection
// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-drift-resource-list.html).
func (c *Client) DetectStackResourceDrift(ctx context.Context, params *DetectStackResourceDriftInput, optFns ...func(*Options)) (*DetectStackResourceDriftOutput, error) {
	stack := middleware.NewStack("DetectStackResourceDrift", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDetectStackResourceDriftMiddlewares(stack)
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
	addOpDetectStackResourceDriftValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetectStackResourceDrift(options.Region), middleware.Before)

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
			OperationName: "DetectStackResourceDrift",
			Err:           err,
		}
	}
	out := result.(*DetectStackResourceDriftOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectStackResourceDriftInput struct {
	// The name of the stack to which the resource belongs.
	StackName *string
	// The logical name of the resource for which to return drift information.
	LogicalResourceId *string
}

type DetectStackResourceDriftOutput struct {
	// Information about whether the resource's actual configuration has drifted from
	// its expected template configuration, including actual and expected property
	// values and any differences detected.
	StackResourceDrift *types.StackResourceDrift

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDetectStackResourceDriftMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDetectStackResourceDrift{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDetectStackResourceDrift{}, middleware.After)
}

func newServiceMetadataMiddleware_opDetectStackResourceDrift(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DetectStackResourceDrift",
	}
}