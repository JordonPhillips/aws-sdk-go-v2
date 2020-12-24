// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Describes the state of the specified instances with respect to the specified
// load balancer. If no instances are specified, the call describes the state of
// all instances that are currently registered with the load balancer. If instances
// are specified, their state is returned even if they are no longer registered
// with the load balancer. The state of terminated instances is not returned.
func (c *Client) DescribeInstanceHealth(ctx context.Context, params *DescribeInstanceHealthInput, optFns ...func(*Options)) (*DescribeInstanceHealthOutput, error) {
	if params == nil {
		params = &DescribeInstanceHealthInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInstanceHealth", params, optFns, addOperationDescribeInstanceHealthMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInstanceHealthOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeInstanceHealth.
type DescribeInstanceHealthInput struct {

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string

	// The IDs of the instances.
	Instances []types.Instance
}

// Contains the output for DescribeInstanceHealth.
type DescribeInstanceHealthOutput struct {

	// Information about the health of the instances.
	InstanceStates []types.InstanceState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeInstanceHealthMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeInstanceHealth{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeInstanceHealth{}, middleware.After)
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
	if err = addOpDescribeInstanceHealthValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstanceHealth(options.Region), middleware.Before); err != nil {
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

// DescribeInstanceHealthAPIClient is a client that implements the
// DescribeInstanceHealth operation.
type DescribeInstanceHealthAPIClient interface {
	DescribeInstanceHealth(context.Context, *DescribeInstanceHealthInput, ...func(*Options)) (*DescribeInstanceHealthOutput, error)
}

var _ DescribeInstanceHealthAPIClient = (*Client)(nil)

// AnyInstanceInServiceWaiterOptions are waiter options for
// AnyInstanceInServiceWaiter
type AnyInstanceInServiceWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// AnyInstanceInServiceWaiter will use default minimum delay of 15 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, AnyInstanceInServiceWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
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
	Retryable func(context.Context, *DescribeInstanceHealthInput, *DescribeInstanceHealthOutput, error) (bool, error)
}

// AnyInstanceInServiceWaiter defines the waiters for AnyInstanceInService
type AnyInstanceInServiceWaiter struct {
	client DescribeInstanceHealthAPIClient

	options AnyInstanceInServiceWaiterOptions
}

// NewAnyInstanceInServiceWaiter constructs a AnyInstanceInServiceWaiter.
func NewAnyInstanceInServiceWaiter(client DescribeInstanceHealthAPIClient, optFns ...func(*AnyInstanceInServiceWaiterOptions)) *AnyInstanceInServiceWaiter {
	options := AnyInstanceInServiceWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = anyInstanceInServiceStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &AnyInstanceInServiceWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for AnyInstanceInService waiter. The maxWaitDur
// is the maximum wait duration the waiter will wait. The maxWaitDur is required
// and must be greater than zero.
func (w *AnyInstanceInServiceWaiter) Wait(ctx context.Context, params *DescribeInstanceHealthInput, maxWaitDur time.Duration, optFns ...func(*AnyInstanceInServiceWaiterOptions)) error {
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

		out, err := w.client.DescribeInstanceHealth(ctx, params, func(o *Options) {
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
	return fmt.Errorf("exceeded max wait time for AnyInstanceInService waiter")
}

func anyInstanceInServiceStateRetryable(ctx context.Context, input *DescribeInstanceHealthInput, output *DescribeInstanceHealthOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("InstanceStates[].State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "InService"
		listOfValues, ok := pathValue.([]string)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected []string value got %T", pathValue)
		}

		for _, v := range listOfValues {
			if v == expectedValue {
				return false, nil
			}
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeInstanceHealth(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "DescribeInstanceHealth",
	}
}
