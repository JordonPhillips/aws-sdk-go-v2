// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Retrieves information about an asset model.
func (c *Client) DescribeAssetModel(ctx context.Context, params *DescribeAssetModelInput, optFns ...func(*Options)) (*DescribeAssetModelOutput, error) {
	if params == nil {
		params = &DescribeAssetModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAssetModel", params, optFns, addOperationDescribeAssetModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAssetModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAssetModelInput struct {

	// The ID of the asset model.
	//
	// This member is required.
	AssetModelId *string
}

type DescribeAssetModelOutput struct {

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the asset model, which has the following format.
	// arn:${Partition}:iotsitewise:${Region}:${Account}:asset-model/${AssetModelId}
	//
	// This member is required.
	AssetModelArn *string

	// The date the asset model was created, in Unix epoch time.
	//
	// This member is required.
	AssetModelCreationDate *time.Time

	// The asset model's description.
	//
	// This member is required.
	AssetModelDescription *string

	// A list of asset model hierarchies that each contain a childAssetModelId and a
	// hierarchyId (named id). A hierarchy specifies allowed parent/child asset
	// relationships for an asset model.
	//
	// This member is required.
	AssetModelHierarchies []types.AssetModelHierarchy

	// The ID of the asset model.
	//
	// This member is required.
	AssetModelId *string

	// The date the asset model was last updated, in Unix epoch time.
	//
	// This member is required.
	AssetModelLastUpdateDate *time.Time

	// The name of the asset model.
	//
	// This member is required.
	AssetModelName *string

	// The list of asset properties for the asset model. This object doesn't include
	// properties that you define in composite models. You can find composite model
	// properties in the assetModelCompositeModels object.
	//
	// This member is required.
	AssetModelProperties []types.AssetModelProperty

	// The current status of the asset model, which contains a state and any error
	// message.
	//
	// This member is required.
	AssetModelStatus *types.AssetModelStatus

	// The list of composite asset models for the asset model.
	AssetModelCompositeModels []types.AssetModelCompositeModel

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAssetModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAssetModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAssetModel{}, middleware.After)
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
	if err = addEndpointPrefix_opDescribeAssetModelMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeAssetModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAssetModel(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opDescribeAssetModelMiddleware struct {
}

func (*endpointPrefix_opDescribeAssetModelMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDescribeAssetModelMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "model." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opDescribeAssetModelMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opDescribeAssetModelMiddleware{}, `OperationSerializer`, middleware.After)
}

// DescribeAssetModelAPIClient is a client that implements the DescribeAssetModel
// operation.
type DescribeAssetModelAPIClient interface {
	DescribeAssetModel(context.Context, *DescribeAssetModelInput, ...func(*Options)) (*DescribeAssetModelOutput, error)
}

var _ DescribeAssetModelAPIClient = (*Client)(nil)

// AssetModelActiveWaiterOptions are waiter options for AssetModelActiveWaiter
type AssetModelActiveWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// AssetModelActiveWaiter will use default minimum delay of 3 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, AssetModelActiveWaiter will use default max delay of 120 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
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
	Retryable func(context.Context, *DescribeAssetModelInput, *DescribeAssetModelOutput, error) (bool, error)
}

// AssetModelActiveWaiter defines the waiters for AssetModelActive
type AssetModelActiveWaiter struct {
	client DescribeAssetModelAPIClient

	options AssetModelActiveWaiterOptions
}

// NewAssetModelActiveWaiter constructs a AssetModelActiveWaiter.
func NewAssetModelActiveWaiter(client DescribeAssetModelAPIClient, optFns ...func(*AssetModelActiveWaiterOptions)) *AssetModelActiveWaiter {
	options := AssetModelActiveWaiterOptions{}
	options.MinDelay = 3 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = assetModelActiveStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &AssetModelActiveWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for AssetModelActive waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *AssetModelActiveWaiter) Wait(ctx context.Context, params *DescribeAssetModelInput, maxWaitDur time.Duration, optFns ...func(*AssetModelActiveWaiterOptions)) error {
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

		out, err := w.client.DescribeAssetModel(ctx, params, func(o *Options) {
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
	return fmt.Errorf("exceeded max wait time for AssetModelActive waiter")
}

func assetModelActiveStateRetryable(ctx context.Context, input *DescribeAssetModelInput, output *DescribeAssetModelOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("assetModelStatus.state", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "ACTIVE"
		value, ok := pathValue.(string)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected string value got %T", pathValue)
		}

		if value == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("assetModelStatus.state", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "FAILED"
		value, ok := pathValue.(string)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected string value got %T", pathValue)
		}

		if value == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

// AssetModelNotExistsWaiterOptions are waiter options for
// AssetModelNotExistsWaiter
type AssetModelNotExistsWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// AssetModelNotExistsWaiter will use default minimum delay of 3 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, AssetModelNotExistsWaiter will use default max delay of 120 seconds.
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
	Retryable func(context.Context, *DescribeAssetModelInput, *DescribeAssetModelOutput, error) (bool, error)
}

// AssetModelNotExistsWaiter defines the waiters for AssetModelNotExists
type AssetModelNotExistsWaiter struct {
	client DescribeAssetModelAPIClient

	options AssetModelNotExistsWaiterOptions
}

// NewAssetModelNotExistsWaiter constructs a AssetModelNotExistsWaiter.
func NewAssetModelNotExistsWaiter(client DescribeAssetModelAPIClient, optFns ...func(*AssetModelNotExistsWaiterOptions)) *AssetModelNotExistsWaiter {
	options := AssetModelNotExistsWaiterOptions{}
	options.MinDelay = 3 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = assetModelNotExistsStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &AssetModelNotExistsWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for AssetModelNotExists waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *AssetModelNotExistsWaiter) Wait(ctx context.Context, params *DescribeAssetModelInput, maxWaitDur time.Duration, optFns ...func(*AssetModelNotExistsWaiterOptions)) error {
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

		out, err := w.client.DescribeAssetModel(ctx, params, func(o *Options) {
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
	return fmt.Errorf("exceeded max wait time for AssetModelNotExists waiter")
}

func assetModelNotExistsStateRetryable(ctx context.Context, input *DescribeAssetModelInput, output *DescribeAssetModelOutput, err error) (bool, error) {

	if err != nil {
		var errorType *types.ResourceNotFoundException
		if errors.As(err, &errorType) {
			return false, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeAssetModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "DescribeAssetModel",
	}
}
