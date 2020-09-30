// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	route53cust "github.com/aws/aws-sdk-go-v2/service/route53/internal/customizations"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a configuration for DNS query logging. If you delete a configuration,
// Amazon Route 53 stops sending query logs to CloudWatch Logs. Route 53 doesn't
// delete any logs that are already in CloudWatch Logs.  <p>For more information
// about DNS query logs, see <a
// href="https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateQueryLoggingConfig.html">CreateQueryLoggingConfig</a>.</p>
func (c *Client) DeleteQueryLoggingConfig(ctx context.Context, params *DeleteQueryLoggingConfigInput, optFns ...func(*Options)) (*DeleteQueryLoggingConfigOutput, error) {
	stack := middleware.NewStack("DeleteQueryLoggingConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpDeleteQueryLoggingConfigMiddlewares(stack)
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
	route53cust.AddSanitizeURLMiddleware(stack)
	addOpDeleteQueryLoggingConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteQueryLoggingConfig(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

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
			OperationName: "DeleteQueryLoggingConfig",
			Err:           err,
		}
	}
	out := result.(*DeleteQueryLoggingConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteQueryLoggingConfigInput struct {

	// The ID of the configuration that you want to delete.
	//
	// This member is required.
	Id *string
}

type DeleteQueryLoggingConfigOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpDeleteQueryLoggingConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpDeleteQueryLoggingConfig{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteQueryLoggingConfig{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteQueryLoggingConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "DeleteQueryLoggingConfig",
	}
}
