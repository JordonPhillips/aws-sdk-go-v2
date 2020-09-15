// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an empty dataset and adds it to the specified dataset group. Use
// CreateDatasetImportJob () to import your training data to a dataset. There are
// three types of datasets:
//
//     * Interactions
//
//     * Items
//
//     * Users
//
// Each
// dataset type has an associated schema with required field types. Only the
// Interactions dataset is required in order to train a model (also referred to as
// creating a solution). A dataset can be in one of the following states:
//
//     *
// CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
//     * DELETE
// PENDING > DELETE IN_PROGRESS
//
// To get the status of the dataset, call
// DescribeDataset (). Related APIs
//
//     * CreateDatasetGroup ()
//
//     *
// ListDatasets ()
//
//     * DescribeDataset ()
//
//     * DeleteDataset ()
func (c *Client) CreateDataset(ctx context.Context, params *CreateDatasetInput, optFns ...func(*Options)) (*CreateDatasetOutput, error) {
	stack := middleware.NewStack("CreateDataset", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateDatasetMiddlewares(stack)
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
	addOpCreateDatasetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataset(options.Region), middleware.Before)

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
			OperationName: "CreateDataset",
			Err:           err,
		}
	}
	out := result.(*CreateDatasetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDatasetInput struct {
	// The Amazon Resource Name (ARN) of the dataset group to add the dataset to.
	DatasetGroupArn *string
	// The name for the dataset.
	Name *string
	// The ARN of the schema to associate with the dataset. The schema defines the
	// dataset fields.
	SchemaArn *string
	// The type of dataset. One of the following (case insensitive) values:
	//
	//     *
	// Interactions
	//
	//     * Items
	//
	//     * Users
	DatasetType *string
}

type CreateDatasetOutput struct {
	// The ARN of the dataset.
	DatasetArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateDatasetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDataset{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDataset{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDataset(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "CreateDataset",
	}
}