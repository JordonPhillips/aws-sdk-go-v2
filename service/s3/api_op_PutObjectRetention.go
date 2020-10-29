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

// Places an Object Retention configuration on an object. This action is not
// supported by Amazon S3 on Outposts. Related Resources
//
// * Locking Objects
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html)
func (c *Client) PutObjectRetention(ctx context.Context, params *PutObjectRetentionInput, optFns ...func(*Options)) (*PutObjectRetentionOutput, error) {
	if params == nil {
		params = &PutObjectRetentionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutObjectRetention", params, optFns, addOperationPutObjectRetentionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutObjectRetentionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutObjectRetentionInput struct {

	// The bucket name that contains the object you want to apply this Object Retention
	// configuration to. When using this API with an access point, you must direct
	// requests to the access point hostname. The access point hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// operation with an access point through the AWS SDKs, you provide the access
	// point ARN in place of the bucket name. For more information about access point
	// ARNs, see Using Access Points
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) in
	// the Amazon Simple Storage Service Developer Guide.
	//
	// This member is required.
	Bucket *string

	// The key name for the object that you want to apply this Object Retention
	// configuration to.
	//
	// This member is required.
	Key *string

	// Indicates whether this operation should bypass Governance-mode restrictions.
	BypassGovernanceRetention *bool

	// The MD5 hash for the request body.
	ContentMD5 *string

	// The account id of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer types.RequestPayer

	// The container element for the Object Retention configuration.
	Retention *types.ObjectLockRetention

	// The version ID for the object that you want to apply this Object Retention
	// configuration to.
	VersionId *string
}

type PutObjectRetentionOutput struct {

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged types.RequestCharged

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutObjectRetentionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutObjectRetention{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutObjectRetention{}, middleware.After)
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
	addOpPutObjectRetentionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutObjectRetention(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)
	smithyhttp.AddChecksumMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutObjectRetention(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutObjectRetention",
	}
}
