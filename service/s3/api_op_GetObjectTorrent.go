// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/service/s3/internal/arn"
)

type GetObjectTorrentInput struct {
	_ struct{} `type:"structure"`

	// The name of the bucket containing the object for which to get the torrent
	// files.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// The object key for which to get the information.
	//
	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`
}

// String returns the string representation
func (s GetObjectTorrentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetObjectTorrentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetObjectTorrentInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetObjectTorrentInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetObjectTorrentInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.RequestPayer) > 0 {
		v := s.RequestPayer

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-payer", v, metadata)
	}
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.Key != nil {
		v := *s.Key

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Key", protocol.StringValue(v), metadata)
	}
	return nil
}

func (s *GetObjectTorrentInput) getEndpointARN() (arn.Resource, error) {
	if s.Bucket == nil {
		return nil, fmt.Errorf("member Bucket is nil")
	}
	return parseEndpointARN(*s.Bucket)
}

func (s *GetObjectTorrentInput) hasEndpointARN() bool {
	if s.Bucket == nil {
		return false
	}
	return arn.IsARN(*s.Bucket)
}

type GetObjectTorrentOutput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// A Bencoded dictionary as defined by the BitTorrent specification
	Body io.ReadCloser `type:"blob"`

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged RequestCharged `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"true"`
}

// String returns the string representation
func (s GetObjectTorrentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetObjectTorrentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.RequestCharged) > 0 {
		v := s.RequestCharged

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-charged", v, metadata)
	}
	// Skipping Body Output type's body not valid.
	return nil
}

const opGetObjectTorrent = "GetObjectTorrent"

// GetObjectTorrentRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Return torrent files from a bucket. BitTorrent can save you bandwidth when
// you're distributing large files. For more information about BitTorrent, see
// Amazon S3 Torrent (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3Torrent.html).
//
// You can get torrent only for objects that are less than 5 GB in size and
// that are not encrypted using server-side encryption with customer-provided
// encryption key.
//
// To use GET, you must have READ access to the object.
//
// The following operation is related to GetObjectTorrent:
//
//    * GetObject
//
//    // Example sending a request using GetObjectTorrentRequest.
//    req := client.GetObjectTorrentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetObjectTorrent
func (c *Client) GetObjectTorrentRequest(input *GetObjectTorrentInput) GetObjectTorrentRequest {
	op := &aws.Operation{
		Name:       opGetObjectTorrent,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}/{Key+}?torrent",
	}

	if input == nil {
		input = &GetObjectTorrentInput{}
	}

	req := c.newRequest(op, input, &GetObjectTorrentOutput{})

	return GetObjectTorrentRequest{Request: req, Input: input, Copy: c.GetObjectTorrentRequest}
}

// GetObjectTorrentRequest is the request type for the
// GetObjectTorrent API operation.
type GetObjectTorrentRequest struct {
	*aws.Request
	Input *GetObjectTorrentInput
	Copy  func(*GetObjectTorrentInput) GetObjectTorrentRequest
}

// Send marshals and sends the GetObjectTorrent API request.
func (r GetObjectTorrentRequest) Send(ctx context.Context) (*GetObjectTorrentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetObjectTorrentResponse{
		GetObjectTorrentOutput: r.Request.Data.(*GetObjectTorrentOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetObjectTorrentResponse is the response type for the
// GetObjectTorrent API operation.
type GetObjectTorrentResponse struct {
	*GetObjectTorrentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetObjectTorrent request.
func (r *GetObjectTorrentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}