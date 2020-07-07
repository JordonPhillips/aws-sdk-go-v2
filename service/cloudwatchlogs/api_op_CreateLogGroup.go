// Code generated by smithy-go-codegen DO NOT EDIT.
package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a log group with the specified name. You can create up to 20,000 log
// groups per account. You must use the following guidelines when naming a log
// group:
//
//     * Log group names must be unique within a region for an AWS
// account.
//
//     * Log group names can be between 1 and 512 characters long.
//
//     *
// Log group names consist of the following characters: a-z, A-Z, 0-9, '_'
// (underscore), '-' (hyphen), '/' (forward slash), '.' (period), and '#' (number
// sign)
//
// If you associate a AWS Key Management Service (AWS KMS) customer master
// key (CMK) with the log group, ingested data is encrypted using the CMK. This
// association is stored as long as the data encrypted with the CMK is still within
// Amazon CloudWatch Logs. This enables Amazon CloudWatch Logs to decrypt this data
// whenever it is requested. If you attempt to associate a CMK with the log group
// but the CMK does not exist or the CMK is disabled, you will receive an
// InvalidParameterException error. Important: CloudWatch Logs supports only
// symmetric CMKs. Do not associate an asymmetric CMK with your log group. For more
// information, see Using Symmetric and Asymmetric Keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html).
func (c *Client) CreateLogGroup(ctx context.Context, params *CreateLogGroupInput, optFns ...func(*Options)) (*CreateLogGroupOutput, error) {
	stack := middleware.NewStack("CreateLogGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	awsmiddleware.AddResolveServiceEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	v4.AddHTTPSignerMiddleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addOpCreateLogGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLogGroup(options.Region), middleware.Before)
	addawsAwsjson11_serdeOpCreateLogGroupMiddlewares(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "CreateLogGroup",
			Err:           err,
		}
	}
	out := result.(*CreateLogGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLogGroupInput struct {
	// The Amazon Resource Name (ARN) of the CMK to use when encrypting log data. For
	// more information, see Amazon Resource Names - AWS Key Management Service (AWS
	// KMS)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-kms).
	KmsKeyId *string
	// The name of the log group.
	LogGroupName *string
	// The key-value pairs to use for the tags.
	Tags map[string]*string
}

type CreateLogGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateLogGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLogGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLogGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateLogGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "CloudWatch Logs",
		ServiceID:      "cloudwatchlogs",
		EndpointPrefix: "cloudwatchlogs",
		SigningName:    "logs",
		OperationName:  "CreateLogGroup",
	}
}
