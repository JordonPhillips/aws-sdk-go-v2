// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes all versions of the intent, including the $LATEST version. To delete a
// specific version of the intent, use the DeleteIntentVersion () operation. You
// can delete a version of an intent only if it is not referenced. To delete an
// intent that is referred to in one or more bots (see how-it-works ()), you must
// remove those references first. If you get the ResourceInUseException exception,
// it provides an example reference that shows where the intent is referenced. To
// remove the reference to the intent, either update the bot or delete it. If you
// get the same exception when you attempt to delete the intent again, repeat until
// the intent has no references and the call to DeleteIntent is successful.  <p>
// This operation requires permission for the <code>lex:DeleteIntent</code> action.
// </p>
func (c *Client) DeleteIntent(ctx context.Context, params *DeleteIntentInput, optFns ...func(*Options)) (*DeleteIntentOutput, error) {
	stack := middleware.NewStack("DeleteIntent", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteIntentMiddlewares(stack)
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
	addOpDeleteIntentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteIntent(options.Region), middleware.Before)

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
			OperationName: "DeleteIntent",
			Err:           err,
		}
	}
	out := result.(*DeleteIntentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteIntentInput struct {
	// The name of the intent. The name is case sensitive.
	Name *string
}

type DeleteIntentOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteIntentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteIntent{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteIntent{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteIntent(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "DeleteIntent",
	}
}