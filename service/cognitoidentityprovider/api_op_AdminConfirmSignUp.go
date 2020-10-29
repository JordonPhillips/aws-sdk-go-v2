// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Confirms user registration as an admin without using a confirmation code. Works
// on any user. Calling this action requires developer credentials.
func (c *Client) AdminConfirmSignUp(ctx context.Context, params *AdminConfirmSignUpInput, optFns ...func(*Options)) (*AdminConfirmSignUpOutput, error) {
	if params == nil {
		params = &AdminConfirmSignUpInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdminConfirmSignUp", params, optFns, addOperationAdminConfirmSignUpMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdminConfirmSignUpOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to confirm user registration.
type AdminConfirmSignUpInput struct {

	// The user pool ID for which you want to confirm user registration.
	//
	// This member is required.
	UserPoolId *string

	// The user name for which you want to confirm user registration.
	//
	// This member is required.
	Username *string

	// A map of custom key-value pairs that you can provide as input for any custom
	// workflows that this action triggers. If your user pool configuration includes
	// triggers, the AdminConfirmSignUp API action invokes the AWS Lambda function that
	// is specified for the post confirmation trigger. When Amazon Cognito invokes this
	// function, it passes a JSON payload, which the function receives as input. In
	// this payload, the clientMetadata attribute provides the data that you assigned
	// to the ClientMetadata parameter in your AdminConfirmSignUp request. In your
	// function code in AWS Lambda, you can process the ClientMetadata value to enhance
	// your workflow for your specific needs. For more information, see Customizing
	// User Pool Workflows with Lambda Triggers
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
}

// Represents the response from the server for the request to confirm registration.
type AdminConfirmSignUpOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAdminConfirmSignUpMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdminConfirmSignUp{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminConfirmSignUp{}, middleware.After)
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
	addOpAdminConfirmSignUpValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAdminConfirmSignUp(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAdminConfirmSignUp(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminConfirmSignUp",
	}
}
