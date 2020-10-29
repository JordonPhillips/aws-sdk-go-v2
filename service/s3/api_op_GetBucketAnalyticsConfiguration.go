// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This implementation of the GET operation returns an analytics configuration
// (identified by the analytics configuration ID) from the bucket. To use this
// operation, you must have permissions to perform the s3:GetAnalyticsConfiguration
// action. The bucket owner has this permission by default. The bucket owner can
// grant this permission to others. For more information about permissions, see
// Permissions Related to Bucket Subresource Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html) in the
// Amazon Simple Storage Service Developer Guide. For information about Amazon S3
// analytics feature, see Amazon S3 Analytics – Storage Class Analysis
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/analytics-storage-class.html)
// in the Amazon Simple Storage Service Developer Guide. Related Resources
//
// *
// DeleteBucketAnalyticsConfiguration
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketAnalyticsConfiguration.html)
//
// *
// ListBucketAnalyticsConfigurations
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBucketAnalyticsConfigurations.html)
//
// *
// PutBucketAnalyticsConfiguration
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketAnalyticsConfiguration.html)
func (c *Client) GetBucketAnalyticsConfiguration(ctx context.Context, params *GetBucketAnalyticsConfigurationInput, optFns ...func(*Options)) (*GetBucketAnalyticsConfigurationOutput, error) {
	if params == nil {
		params = &GetBucketAnalyticsConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketAnalyticsConfiguration", params, optFns, addOperationGetBucketAnalyticsConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketAnalyticsConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketAnalyticsConfigurationInput struct {

	// The name of the bucket from which an analytics configuration is retrieved.
	//
	// This member is required.
	Bucket *string

	// The ID that identifies the analytics configuration.
	//
	// This member is required.
	Id *string

	// The account id of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string
}

type GetBucketAnalyticsConfigurationOutput struct {

	// The configuration and any analyses for the analytics filter.
	AnalyticsConfiguration *types.AnalyticsConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetBucketAnalyticsConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetBucketAnalyticsConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketAnalyticsConfiguration{}, middleware.After)
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
	addOpGetBucketAnalyticsConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketAnalyticsConfiguration(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetBucketAnalyticsConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketAnalyticsConfiguration",
	}
}
