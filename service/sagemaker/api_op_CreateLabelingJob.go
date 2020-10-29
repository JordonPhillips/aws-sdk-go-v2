// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a job that uses workers to label the data objects in your input dataset.
// You can use the labeled data to train machine learning models. You can select
// your workforce from one of three providers:
//
// * A private workforce that you
// create. It can include employees, contractors, and outside experts. Use a
// private workforce when want the data to stay within your organization or when a
// specific set of skills is required.
//
// * One or more vendors that you select from
// the AWS Marketplace. Vendors provide expertise in specific areas.
//
// * The Amazon
// Mechanical Turk workforce. This is the largest workforce, but it should only be
// used for public data or data that has been stripped of any personally
// identifiable information.
//
// You can also use automated data labeling to reduce
// the number of data objects that need to be labeled by a human. Automated data
// labeling uses active learning to determine if a data object can be labeled by
// machine or if it needs to be sent to a human worker. For more information, see
// Using Automated Data Labeling
// (https://docs.aws.amazon.com/sagemaker/latest/dg/sms-automated-labeling.html).
// The data objects to be labeled are contained in an Amazon S3 bucket. You create
// a manifest file that describes the location of each object. For more
// information, see Using Input and Output Data
// (https://docs.aws.amazon.com/sagemaker/latest/dg/sms-data.html). The output can
// be used as the manifest file for another labeling job or as training data for
// your machine learning models.
func (c *Client) CreateLabelingJob(ctx context.Context, params *CreateLabelingJobInput, optFns ...func(*Options)) (*CreateLabelingJobOutput, error) {
	if params == nil {
		params = &CreateLabelingJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLabelingJob", params, optFns, addOperationCreateLabelingJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLabelingJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLabelingJobInput struct {

	// Configures the labeling task and how it is presented to workers; including, but
	// not limited to price, keywords, and batch size (task count).
	//
	// This member is required.
	HumanTaskConfig *types.HumanTaskConfig

	// Input data for the labeling job, such as the Amazon S3 location of the data
	// objects and the location of the manifest file that describes the data objects.
	//
	// This member is required.
	InputConfig *types.LabelingJobInputConfig

	// The attribute name to use for the label in the output manifest file. This is the
	// key for the key/value pair formed with the label that a worker assigns to the
	// object. The name can't end with "-metadata". If you are running a semantic
	// segmentation labeling job, the attribute name must end with "-ref". If you are
	// running any other kind of labeling job, the attribute name must not end with
	// "-ref".
	//
	// This member is required.
	LabelAttributeName *string

	// The name of the labeling job. This name is used to identify the job in a list of
	// labeling jobs.
	//
	// This member is required.
	LabelingJobName *string

	// The location of the output data and the AWS Key Management Service key ID for
	// the key used to encrypt the output data, if any.
	//
	// This member is required.
	OutputConfig *types.LabelingJobOutputConfig

	// The Amazon Resource Number (ARN) that Amazon SageMaker assumes to perform tasks
	// on your behalf during data labeling. You must grant this role the necessary
	// permissions so that Amazon SageMaker can successfully complete data labeling.
	//
	// This member is required.
	RoleArn *string

	// The S3 URL of the file that defines the categories used to label the data
	// objects. For 3D point cloud task types, see Create a Labeling Category
	// Configuration File for 3D Point Cloud Labeling Jobs
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/sms-point-cloud-label-category-config.html).
	// For all other built-in task types
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/sms-task-types.html) and custom
	// tasks
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/sms-custom-templates.html),
	// your label category configuration file must be a JSON file in the following
	// format. Identify the labels you want to use by replacing label_1,
	// label_2,...,label_n with your label categories. {
	//     "document-version":
	// "2018-11-28"
	//
	//     "labels": [
	//
	//     {
	//
	//     "label": "label_1"
	//
	//     },
	//
	//     {
	//
	//
	// "label": "label_2"
	//
	//     },
	//
	//     ...
	//
	//     {
	//
	//     "label": "label_n"
	//
	//     }
	//
	//
	// ]
	//
	//     }
	LabelCategoryConfigS3Uri *string

	// Configures the information required to perform automated data labeling.
	LabelingJobAlgorithmsConfig *types.LabelingJobAlgorithmsConfig

	// A set of conditions for stopping the labeling job. If any of the conditions are
	// met, the job is automatically stopped. You can use these conditions to control
	// the cost of data labeling.
	StoppingConditions *types.LabelingJobStoppingConditions

	// An array of key/value pairs. For more information, see Using Cost Allocation
	// Tags
	// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what)
	// in the AWS Billing and Cost Management User Guide.
	Tags []*types.Tag
}

type CreateLabelingJobOutput struct {

	// The Amazon Resource Name (ARN) of the labeling job. You use this ARN to identify
	// the labeling job.
	//
	// This member is required.
	LabelingJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateLabelingJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLabelingJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLabelingJob{}, middleware.After)
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
	addOpCreateLabelingJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLabelingJob(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateLabelingJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateLabelingJob",
	}
}
