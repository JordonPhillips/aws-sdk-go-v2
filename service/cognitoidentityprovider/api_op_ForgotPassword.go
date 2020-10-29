// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Calling this API causes a message to be sent to the end user with a confirmation
// code that is required to change the user's password. For the Username parameter,
// you can use the username or user alias. The method used to send the confirmation
// code is sent according to the specified AccountRecoverySetting. For more
// information, see Recovering User Accounts
// (https://docs.aws.amazon.com/cognito/latest/developerguide/how-to-recover-a-user-account.html)
// in the Amazon Cognito Developer Guide. If neither a verified phone number nor a
// verified email exists, an InvalidParameterException is thrown. To use the
// confirmation code for resetting the password, call ConfirmForgotPassword
// (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_ConfirmForgotPassword.html).
func (c *Client) ForgotPassword(ctx context.Context, params *ForgotPasswordInput, optFns ...func(*Options)) (*ForgotPasswordOutput, error) {
	if params == nil {
		params = &ForgotPasswordInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ForgotPassword", params, optFns, addOperationForgotPasswordMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ForgotPasswordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to reset a user's password.
type ForgotPasswordInput struct {

	// The ID of the client associated with the user pool.
	//
	// This member is required.
	ClientId *string

	// The user name of the user for whom you want to enter a code to reset a forgotten
	// password.
	//
	// This member is required.
	Username *string

	// The Amazon Pinpoint analytics metadata for collecting metrics for ForgotPassword
	// calls.
	AnalyticsMetadata *types.AnalyticsMetadataType

	// A map of custom key-value pairs that you can provide as input for any custom
	// workflows that this action triggers. You create custom workflows by assigning
	// AWS Lambda functions to user pool triggers. When you use the ForgotPassword API
	// action, Amazon Cognito invokes any functions that are assigned to the following
	// triggers: pre sign-up, custom message, and user migration. When Amazon Cognito
	// invokes any of these functions, it passes a JSON payload, which the function
	// receives as input. This payload contains a clientMetadata attribute, which
	// provides the data that you assigned to the ClientMetadata parameter in your
	// ForgotPassword request. In your function code in AWS Lambda, you can process the
	// clientMetadata value to enhance your workflow for your specific needs. For more
	// information, see Customizing User Pool Workflows with Lambda Triggers
	// (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html)
	// in the Amazon Cognito Developer Guide. Take the following limitations into
	// consideration when you use the ClientMetadata parameter:
	//
	// * Amazon Cognito does
	// not store the ClientMetadata value. This data is available only to AWS Lambda
	// triggers that are assigned to a user pool to support custom workflows. If your
	// user pool configuration does not include triggers, the ClientMetadata parameter
	// serves no purpose.
	//
	// * Amazon Cognito does not validate the ClientMetadata
	// value.
	//
	// * Amazon Cognito does not encrypt the the ClientMetadata value, so don't
	// use it to provide sensitive information.
	ClientMetadata map[string]*string

	// A keyed-hash message authentication code (HMAC) calculated using the secret key
	// of a user pool client and username plus the client ID in the message.
	SecretHash *string

	// Contextual data such as the user's device fingerprint, IP address, or location
	// used for evaluating the risk of an unexpected event by Amazon Cognito advanced
	// security.
	UserContextData *types.UserContextDataType
}

// Respresents the response from the server regarding the request to reset a
// password.
type ForgotPasswordOutput struct {

	// The code delivery details returned by the server in response to the request to
	// reset a password.
	CodeDeliveryDetails *types.CodeDeliveryDetailsType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationForgotPasswordMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpForgotPassword{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpForgotPassword{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	addRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpForgotPasswordValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opForgotPassword(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opForgotPassword(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ForgotPassword",
	}
}
