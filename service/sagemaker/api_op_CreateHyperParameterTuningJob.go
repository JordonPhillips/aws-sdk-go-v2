// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a hyperparameter tuning job. A hyperparameter tuning job finds the best
// version of a model by running many training jobs on your dataset using the
// algorithm you choose and values for hyperparameters within ranges that you
// specify. It then chooses the hyperparameter values that result in a model that
// performs the best, as measured by an objective metric that you choose.
func (c *Client) CreateHyperParameterTuningJob(ctx context.Context, params *CreateHyperParameterTuningJobInput, optFns ...func(*Options)) (*CreateHyperParameterTuningJobOutput, error) {
	if params == nil {
		params = &CreateHyperParameterTuningJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateHyperParameterTuningJob", params, optFns, addOperationCreateHyperParameterTuningJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateHyperParameterTuningJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateHyperParameterTuningJobInput struct {

	// The HyperParameterTuningJobConfig object that describes the tuning job,
	// including the search strategy, the objective metric used to evaluate training
	// jobs, ranges of parameters to search, and resource limits for the tuning job.
	// For more information, see How Hyperparameter Tuning Works
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-how-it-works.html).
	//
	// This member is required.
	HyperParameterTuningJobConfig *types.HyperParameterTuningJobConfig

	// The name of the tuning job. This name is the prefix for the names of all
	// training jobs that this tuning job launches. The name must be unique within the
	// same AWS account and AWS Region. The name must have 1 to 32 characters. Valid
	// characters are a-z, A-Z, 0-9, and : + = @ _ % - (hyphen). The name is not case
	// sensitive.
	//
	// This member is required.
	HyperParameterTuningJobName *string

	// An array of key-value pairs. You can use tags to categorize your AWS resources
	// in different ways, for example, by purpose, owner, or environment. For more
	// information, see Tagging AWS Resources
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html). Tags that you
	// specify for the tuning job are also added to all training jobs that the tuning
	// job launches.
	Tags []types.Tag

	// The HyperParameterTrainingJobDefinition object that describes the training jobs
	// that this tuning job launches, including static hyperparameters, input data
	// configuration, output data configuration, resource configuration, and stopping
	// condition.
	TrainingJobDefinition *types.HyperParameterTrainingJobDefinition

	// A list of the HyperParameterTrainingJobDefinition objects launched for this
	// tuning job.
	TrainingJobDefinitions []types.HyperParameterTrainingJobDefinition

	// Specifies the configuration for starting the hyperparameter tuning job using one
	// or more previous tuning jobs as a starting point. The results of previous tuning
	// jobs are used to inform which combinations of hyperparameters to search over in
	// the new tuning job. All training jobs launched by the new hyperparameter tuning
	// job are evaluated by using the objective metric. If you specify
	// IDENTICAL_DATA_AND_ALGORITHM as the WarmStartType value for the warm start
	// configuration, the training job that performs the best in the new tuning job is
	// compared to the best training jobs from the parent tuning jobs. From these, the
	// training job that performs the best as measured by the objective metric is
	// returned as the overall best training job. All training jobs launched by parent
	// hyperparameter tuning jobs and the new hyperparameter tuning jobs count against
	// the limit of training jobs for the tuning job.
	WarmStartConfig *types.HyperParameterTuningJobWarmStartConfig
}

type CreateHyperParameterTuningJobOutput struct {

	// The Amazon Resource Name (ARN) of the tuning job. Amazon SageMaker assigns an
	// ARN to a hyperparameter tuning job when you create it.
	//
	// This member is required.
	HyperParameterTuningJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateHyperParameterTuningJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateHyperParameterTuningJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateHyperParameterTuningJob{}, middleware.After)
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
	if err = addOpCreateHyperParameterTuningJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHyperParameterTuningJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateHyperParameterTuningJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateHyperParameterTuningJob",
	}
}
