// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Lists and describes the models in an Amazon Rekognition Custom Labels project.
// You can specify up to 10 model versions in ProjectVersionArns. If you don't
// specify a value, descriptions for all models are returned. This operation
// requires permissions to perform the rekognition:DescribeProjectVersions action.
func (c *Client) DescribeProjectVersions(ctx context.Context, params *DescribeProjectVersionsInput, optFns ...func(*Options)) (*DescribeProjectVersionsOutput, error) {
	if params == nil {
		params = &DescribeProjectVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeProjectVersions", params, optFns, addOperationDescribeProjectVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeProjectVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeProjectVersionsInput struct {

	// The Amazon Resource Name (ARN) of the project that contains the models you want
	// to describe.
	//
	// This member is required.
	ProjectArn *string

	// The maximum number of results to return per paginated call. The largest value
	// you can specify is 100. If you specify a value greater than 100, a
	// ValidationException error occurs. The default value is 100.
	MaxResults *int32

	// If the previous response was incomplete (because there is more results to
	// retrieve), Amazon Rekognition Custom Labels returns a pagination token in the
	// response. You can use this pagination token to retrieve the next set of results.
	NextToken *string

	// A list of model version names that you want to describe. You can add up to 10
	// model version names to the list. If you don't specify a value, all model
	// descriptions are returned. A version name is part of a model (ProjectVersion)
	// ARN. For example, my-model.2020-01-21T09.10.15 is the version name in the
	// following ARN.
	// arn:aws:rekognition:us-east-1:123456789012:project/getting-started/version/my-model.2020-01-21T09.10.15/1234567890123.
	VersionNames []string
}

type DescribeProjectVersionsOutput struct {

	// If the previous response was incomplete (because there is more results to
	// retrieve), Amazon Rekognition Custom Labels returns a pagination token in the
	// response. You can use this pagination token to retrieve the next set of results.
	NextToken *string

	// A list of model descriptions. The list is sorted by the creation date and time
	// of the model versions, latest to earliest.
	ProjectVersionDescriptions []types.ProjectVersionDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeProjectVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeProjectVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeProjectVersions{}, middleware.After)
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
	if err = addOpDescribeProjectVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeProjectVersions(options.Region), middleware.Before); err != nil {
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

// DescribeProjectVersionsAPIClient is a client that implements the
// DescribeProjectVersions operation.
type DescribeProjectVersionsAPIClient interface {
	DescribeProjectVersions(context.Context, *DescribeProjectVersionsInput, ...func(*Options)) (*DescribeProjectVersionsOutput, error)
}

var _ DescribeProjectVersionsAPIClient = (*Client)(nil)

// DescribeProjectVersionsPaginatorOptions is the paginator options for
// DescribeProjectVersions
type DescribeProjectVersionsPaginatorOptions struct {
	// The maximum number of results to return per paginated call. The largest value
	// you can specify is 100. If you specify a value greater than 100, a
	// ValidationException error occurs. The default value is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeProjectVersionsPaginator is a paginator for DescribeProjectVersions
type DescribeProjectVersionsPaginator struct {
	options   DescribeProjectVersionsPaginatorOptions
	client    DescribeProjectVersionsAPIClient
	params    *DescribeProjectVersionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeProjectVersionsPaginator returns a new
// DescribeProjectVersionsPaginator
func NewDescribeProjectVersionsPaginator(client DescribeProjectVersionsAPIClient, params *DescribeProjectVersionsInput, optFns ...func(*DescribeProjectVersionsPaginatorOptions)) *DescribeProjectVersionsPaginator {
	options := DescribeProjectVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeProjectVersionsInput{}
	}

	return &DescribeProjectVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeProjectVersionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeProjectVersions page.
func (p *DescribeProjectVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeProjectVersionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.DescribeProjectVersions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ProjectVersionRunningWaiterOptions are waiter options for
// ProjectVersionRunningWaiter
type ProjectVersionRunningWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ProjectVersionRunningWaiter will use default minimum delay of 30 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, ProjectVersionRunningWaiter will use default max delay of 120 seconds.
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
	Retryable func(context.Context, *DescribeProjectVersionsInput, *DescribeProjectVersionsOutput, error) (bool, error)
}

// ProjectVersionRunningWaiter defines the waiters for ProjectVersionRunning
type ProjectVersionRunningWaiter struct {
	client DescribeProjectVersionsAPIClient

	options ProjectVersionRunningWaiterOptions
}

// NewProjectVersionRunningWaiter constructs a ProjectVersionRunningWaiter.
func NewProjectVersionRunningWaiter(client DescribeProjectVersionsAPIClient, optFns ...func(*ProjectVersionRunningWaiterOptions)) *ProjectVersionRunningWaiter {
	options := ProjectVersionRunningWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = projectVersionRunningStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ProjectVersionRunningWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ProjectVersionRunning waiter. The maxWaitDur
// is the maximum wait duration the waiter will wait. The maxWaitDur is required
// and must be greater than zero.
func (w *ProjectVersionRunningWaiter) Wait(ctx context.Context, params *DescribeProjectVersionsInput, maxWaitDur time.Duration, optFns ...func(*ProjectVersionRunningWaiterOptions)) error {
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

		out, err := w.client.DescribeProjectVersions(ctx, params, func(o *Options) {
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
	return fmt.Errorf("exceeded max wait time for ProjectVersionRunning waiter")
}

func projectVersionRunningStateRetryable(ctx context.Context, input *DescribeProjectVersionsInput, output *DescribeProjectVersionsOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("ProjectVersionDescriptions[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "RUNNING"
		var match = true
		listOfValues, ok := pathValue.([]string)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected []string value got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			if v != expectedValue {
				match = false
			}
		}

		if match {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("ProjectVersionDescriptions[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "FAILED"
		listOfValues, ok := pathValue.([]string)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected []string value got %T", pathValue)
		}

		for _, v := range listOfValues {
			if v == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	return true, nil
}

// ProjectVersionTrainingCompletedWaiterOptions are waiter options for
// ProjectVersionTrainingCompletedWaiter
type ProjectVersionTrainingCompletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ProjectVersionTrainingCompletedWaiter will use default minimum delay of 120
	// seconds. Note that MinDelay must resolve to a value lesser than or equal to the
	// MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, ProjectVersionTrainingCompletedWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
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
	Retryable func(context.Context, *DescribeProjectVersionsInput, *DescribeProjectVersionsOutput, error) (bool, error)
}

// ProjectVersionTrainingCompletedWaiter defines the waiters for
// ProjectVersionTrainingCompleted
type ProjectVersionTrainingCompletedWaiter struct {
	client DescribeProjectVersionsAPIClient

	options ProjectVersionTrainingCompletedWaiterOptions
}

// NewProjectVersionTrainingCompletedWaiter constructs a
// ProjectVersionTrainingCompletedWaiter.
func NewProjectVersionTrainingCompletedWaiter(client DescribeProjectVersionsAPIClient, optFns ...func(*ProjectVersionTrainingCompletedWaiterOptions)) *ProjectVersionTrainingCompletedWaiter {
	options := ProjectVersionTrainingCompletedWaiterOptions{}
	options.MinDelay = 120 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = projectVersionTrainingCompletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ProjectVersionTrainingCompletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ProjectVersionTrainingCompleted waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *ProjectVersionTrainingCompletedWaiter) Wait(ctx context.Context, params *DescribeProjectVersionsInput, maxWaitDur time.Duration, optFns ...func(*ProjectVersionTrainingCompletedWaiterOptions)) error {
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

		out, err := w.client.DescribeProjectVersions(ctx, params, func(o *Options) {
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
	return fmt.Errorf("exceeded max wait time for ProjectVersionTrainingCompleted waiter")
}

func projectVersionTrainingCompletedStateRetryable(ctx context.Context, input *DescribeProjectVersionsInput, output *DescribeProjectVersionsOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("ProjectVersionDescriptions[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "TRAINING_COMPLETED"
		var match = true
		listOfValues, ok := pathValue.([]string)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected []string value got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			if v != expectedValue {
				match = false
			}
		}

		if match {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("ProjectVersionDescriptions[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "TRAINING_FAILED"
		listOfValues, ok := pathValue.([]string)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected []string value got %T", pathValue)
		}

		for _, v := range listOfValues {
			if v == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeProjectVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "DescribeProjectVersions",
	}
}
