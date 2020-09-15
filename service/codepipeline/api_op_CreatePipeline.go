// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a pipeline. In the pipeline structure, you must include either
// artifactStore or artifactStores in your pipeline, but you cannot use both. If
// you create a cross-region action in your pipeline, you must use artifactStores.
func (c *Client) CreatePipeline(ctx context.Context, params *CreatePipelineInput, optFns ...func(*Options)) (*CreatePipelineOutput, error) {
	stack := middleware.NewStack("CreatePipeline", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreatePipelineMiddlewares(stack)
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
	addOpCreatePipelineValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePipeline(options.Region), middleware.Before)

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
			OperationName: "CreatePipeline",
			Err:           err,
		}
	}
	out := result.(*CreatePipelineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a CreatePipeline action.
type CreatePipelineInput struct {
	// Represents the structure of actions and stages to be performed in the pipeline.
	Pipeline *types.PipelineDeclaration
	// The tags for the pipeline.
	Tags []*types.Tag
}

// Represents the output of a CreatePipeline action.
type CreatePipelineOutput struct {
	// Specifies the tags applied to the pipeline.
	Tags []*types.Tag
	// Represents the structure of actions and stages to be performed in the pipeline.
	Pipeline *types.PipelineDeclaration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreatePipelineMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePipeline{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePipeline{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreatePipeline(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "CreatePipeline",
	}
}