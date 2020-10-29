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

// This operation lists in-progress multipart uploads. An in-progress multipart
// upload is a multipart upload that has been initiated using the Initiate
// Multipart Upload request, but has not yet been completed or aborted. This
// operation returns at most 1,000 multipart uploads in the response. 1,000
// multipart uploads is the maximum number of uploads a response can include, which
// is also the default value. You can further limit the number of uploads in a
// response by specifying the max-uploads parameter in the response. If additional
// multipart uploads satisfy the list criteria, the response will contain an
// IsTruncated element with the value true. To list the additional multipart
// uploads, use the key-marker and upload-id-marker request parameters. In the
// response, the uploads are sorted by key. If your application has initiated more
// than one multipart upload using the same object key, then uploads in the
// response are first sorted by key. Additionally, uploads are sorted in ascending
// order within each key by the upload initiation time. For more information on
// multipart uploads, see Uploading Objects Using Multipart Upload
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/uploadobjusingmpu.html). For
// information on permissions required to use the multipart upload API, see
// Multipart Upload API and Permissions
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuAndPermissions.html). The
// following operations are related to ListMultipartUploads:
//
// *
// CreateMultipartUpload
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html)
//
// *
// UploadPart
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html)
//
// *
// CompleteMultipartUpload
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompleteMultipartUpload.html)
//
// *
// ListParts
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html)
//
// *
// AbortMultipartUpload
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_AbortMultipartUpload.html)
func (c *Client) ListMultipartUploads(ctx context.Context, params *ListMultipartUploadsInput, optFns ...func(*Options)) (*ListMultipartUploadsOutput, error) {
	if params == nil {
		params = &ListMultipartUploadsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMultipartUploads", params, optFns, addOperationListMultipartUploadsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMultipartUploadsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMultipartUploadsInput struct {

	// The name of the bucket to which the multipart upload was initiated. When using
	// this API with an access point, you must direct requests to the access point
	// hostname. The access point hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// operation with an access point through the AWS SDKs, you provide the access
	// point ARN in place of the bucket name. For more information about access point
	// ARNs, see Using Access Points
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) in
	// the Amazon Simple Storage Service Developer Guide. When using this API with
	// Amazon S3 on Outposts, you must direct requests to the S3 on Outposts hostname.
	// The S3 on Outposts hostname takes the form
	// AccessPointName-AccountId.outpostID.s3-outposts.Region.amazonaws.com. When using
	// this operation using S3 on Outposts through the AWS SDKs, you provide the
	// Outposts bucket ARN in place of the bucket name. For more information about S3
	// on Outposts ARNs, see Using S3 on Outposts
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html) in the
	// Amazon Simple Storage Service Developer Guide.
	//
	// This member is required.
	Bucket *string

	// Character you use to group keys. All keys that contain the same string between
	// the prefix, if specified, and the first occurrence of the delimiter after the
	// prefix are grouped under a single result element, CommonPrefixes. If you don't
	// specify the prefix parameter, then the substring starts at the beginning of the
	// key. The keys that are grouped under CommonPrefixes result element are not
	// returned elsewhere in the response.
	Delimiter *string

	// Requests Amazon S3 to encode the object keys in the response and specifies the
	// encoding method to use. An object key may contain any Unicode character;
	// however, XML 1.0 parser cannot parse some characters, such as characters with an
	// ASCII value from 0 to 10. For characters that are not supported in XML 1.0, you
	// can add this parameter to request that Amazon S3 encode the keys in the
	// response.
	EncodingType types.EncodingType

	// The account id of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string

	// Together with upload-id-marker, this parameter specifies the multipart upload
	// after which listing should begin. If upload-id-marker is not specified, only the
	// keys lexicographically greater than the specified key-marker will be included in
	// the list. If upload-id-marker is specified, any multipart uploads for a key
	// equal to the key-marker might also be included, provided those multipart uploads
	// have upload IDs lexicographically greater than the specified upload-id-marker.
	KeyMarker *string

	// Sets the maximum number of multipart uploads, from 1 to 1,000, to return in the
	// response body. 1,000 is the maximum number of uploads that can be returned in a
	// response.
	MaxUploads *int32

	// Lists in-progress uploads only for those keys that begin with the specified
	// prefix. You can use prefixes to separate a bucket into different grouping of
	// keys. (You can think of using prefix to make groups in the same way you'd use a
	// folder in a file system.)
	Prefix *string

	// Together with key-marker, specifies the multipart upload after which listing
	// should begin. If key-marker is not specified, the upload-id-marker parameter is
	// ignored. Otherwise, any multipart uploads for a key equal to the key-marker
	// might be included in the list only if they have an upload ID lexicographically
	// greater than the specified upload-id-marker.
	UploadIdMarker *string
}

type ListMultipartUploadsOutput struct {

	// The name of the bucket to which the multipart upload was initiated.
	Bucket *string

	// If you specify a delimiter in the request, then the result returns each distinct
	// key prefix containing the delimiter in a CommonPrefixes element. The distinct
	// key prefixes are returned in the Prefix child element.
	CommonPrefixes []*types.CommonPrefix

	// Contains the delimiter you specified in the request. If you don't specify a
	// delimiter in your request, this element is absent from the response.
	Delimiter *string

	// Encoding type used by Amazon S3 to encode object keys in the response. If you
	// specify encoding-type request parameter, Amazon S3 includes this element in the
	// response, and returns encoded key name values in the following response
	// elements: Delimiter, KeyMarker, Prefix, NextKeyMarker, Key.
	EncodingType types.EncodingType

	// Indicates whether the returned list of multipart uploads is truncated. A value
	// of true indicates that the list was truncated. The list can be truncated if the
	// number of multipart uploads exceeds the limit allowed or specified by max
	// uploads.
	IsTruncated *bool

	// The key at or after which the listing began.
	KeyMarker *string

	// Maximum number of multipart uploads that could have been included in the
	// response.
	MaxUploads *int32

	// When a list is truncated, this element specifies the value that should be used
	// for the key-marker request parameter in a subsequent request.
	NextKeyMarker *string

	// When a list is truncated, this element specifies the value that should be used
	// for the upload-id-marker request parameter in a subsequent request.
	NextUploadIdMarker *string

	// When a prefix is provided in the request, this field contains the specified
	// prefix. The result contains only keys starting with the specified prefix.
	Prefix *string

	// Upload ID after which listing began.
	UploadIdMarker *string

	// Container for elements related to a particular multipart upload. A response can
	// contain zero or more Upload elements.
	Uploads []*types.MultipartUpload

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListMultipartUploadsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListMultipartUploads{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListMultipartUploads{}, middleware.After)
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
	addOpListMultipartUploadsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListMultipartUploads(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)
	return nil
}

func newServiceMetadataMiddleware_opListMultipartUploads(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "ListMultipartUploads",
	}
}
