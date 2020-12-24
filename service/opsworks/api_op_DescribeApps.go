// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"time"
)

// Requests a description of a specified set of apps. This call accepts only one
// resource-identifying parameter. Required Permissions: To use this action, an IAM
// user must have a Show, Deploy, or Manage permissions level for the stack, or an
// attached policy that explicitly grants permissions. For more information about
// user permissions, see Managing User Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) DescribeApps(ctx context.Context, params *DescribeAppsInput, optFns ...func(*Options)) (*DescribeAppsOutput, error) {
	if params == nil {
		params = &DescribeAppsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeApps", params, optFns, addOperationDescribeAppsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAppsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAppsInput struct {

	// An array of app IDs for the apps to be described. If you use this parameter,
	// DescribeApps returns a description of the specified apps. Otherwise, it returns
	// a description of every app.
	AppIds []string

	// The app stack ID. If you use this parameter, DescribeApps returns a description
	// of the apps in the specified stack.
	StackId *string
}

// Contains the response to a DescribeApps request.
type DescribeAppsOutput struct {

	// An array of App objects that describe the specified apps.
	Apps []types.App

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAppsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeApps{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeApps{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeApps(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeAppsAPIClient is a client that implements the DescribeApps operation.
type DescribeAppsAPIClient interface {
	DescribeApps(context.Context, *DescribeAppsInput, ...func(*Options)) (*DescribeAppsOutput, error)
}

var _ DescribeAppsAPIClient = (*Client)(nil)

// AppExistsWaiterOptions are waiter options for AppExistsWaiter
type AppExistsWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// AppExistsWaiter will use default minimum delay of 1 seconds. Note that MinDelay
	// must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, AppExistsWaiter will use default max delay of 120 seconds. Note that
	// MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeAppsInput, *DescribeAppsOutput, error) (bool, error)
}

// AppExistsWaiter defines the waiters for AppExists
type AppExistsWaiter struct {
	client DescribeAppsAPIClient

	options AppExistsWaiterOptions
}

// NewAppExistsWaiter constructs a AppExistsWaiter.
func NewAppExistsWaiter(client DescribeAppsAPIClient, optFns ...func(*AppExistsWaiterOptions)) *AppExistsWaiter {
	options := AppExistsWaiterOptions{}
	options.MinDelay = 1 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = appExistsStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &AppExistsWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for AppExists waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *AppExistsWaiter) Wait(ctx context.Context, params *DescribeAppsInput, maxWaitDur time.Duration, optFns ...func(*AppExistsWaiterOptions)) error {
	if maxWaitDur <= 0 {
		return fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeApps(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return err
		}
		if !retryable {
			return nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return fmt.Errorf("exceeded max wait time for AppExists waiter")
}

func appExistsStateRetryable(ctx context.Context, input *DescribeAppsInput, output *DescribeAppsOutput, err error) (bool, error) {

	if err == nil {
		return false, nil
	}

	if err == nil {
		return false, fmt.Errorf("waiter state transitioned to Failure")
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeApps(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "DescribeApps",
	}
}
