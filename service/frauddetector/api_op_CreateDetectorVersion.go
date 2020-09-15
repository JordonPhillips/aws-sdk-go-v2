// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a detector version. The detector version starts in a DRAFT status.
func (c *Client) CreateDetectorVersion(ctx context.Context, params *CreateDetectorVersionInput, optFns ...func(*Options)) (*CreateDetectorVersionOutput, error) {
	stack := middleware.NewStack("CreateDetectorVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateDetectorVersionMiddlewares(stack)
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
	addOpCreateDetectorVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDetectorVersion(options.Region), middleware.Before)

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
			OperationName: "CreateDetectorVersion",
			Err:           err,
		}
	}
	out := result.(*CreateDetectorVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDetectorVersionInput struct {
	// The description of the detector version.
	Description *string
	// A collection of key and value pairs.
	Tags []*types.Tag
	// The Amazon Sagemaker model endpoints to include in the detector version.
	ExternalModelEndpoints []*string
	// The model versions to include in the detector version.
	ModelVersions []*types.ModelVersion
	// The ID of the detector under which you want to create a new version.
	DetectorId *string
	// The rule execution mode for the rules included in the detector version. You can
	// define and edit the rule mode at the detector version level, when it is in draft
	// status. If you specify FIRST_MATCHED, Amazon Fraud Detector evaluates rules
	// sequentially, first to last, stopping at the first matched rule. Amazon Fraud
	// dectector then provides the outcomes for that single rule. If you specifiy
	// ALL_MATCHED, Amazon Fraud Detector evaluates all rules and returns the outcomes
	// for all matched rules. The default behavior is FIRST_MATCHED.
	RuleExecutionMode types.RuleExecutionMode
	// The rules to include in the detector version.
	Rules []*types.Rule
}

type CreateDetectorVersionOutput struct {
	// The status of the detector version.
	Status types.DetectorVersionStatus
	// The ID for the created version's parent detector.
	DetectorId *string
	// The ID for the created detector.
	DetectorVersionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateDetectorVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDetectorVersion{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDetectorVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDetectorVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "CreateDetectorVersion",
	}
}