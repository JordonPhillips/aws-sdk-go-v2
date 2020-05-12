// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type BatchWriteInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that is associated with the Directory. For
	// more information, see arns.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// A list of operations that are part of the batch.
	//
	// Operations is a required field
	Operations []BatchWriteOperation `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchWriteInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchWriteInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchWriteInput"}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}

	if s.Operations == nil {
		invalidParams.Add(aws.NewErrParamRequired("Operations"))
	}
	if s.Operations != nil {
		for i, v := range s.Operations {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Operations", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchWriteInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Operations != nil {
		v := s.Operations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Operations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.DirectoryArn != nil {
		v := *s.DirectoryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type BatchWriteOutput struct {
	_ struct{} `type:"structure"`

	// A list of all the responses for each batch write.
	Responses []BatchWriteOperationResponse `type:"list"`
}

// String returns the string representation
func (s BatchWriteOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchWriteOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Responses != nil {
		v := s.Responses

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Responses", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opBatchWrite = "BatchWrite"

// BatchWriteRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Performs all the write operations in a batch. Either all the operations succeed
// or none.
//
//    // Example sending a request using BatchWriteRequest.
//    req := client.BatchWriteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/BatchWrite
func (c *Client) BatchWriteRequest(input *BatchWriteInput) BatchWriteRequest {
	op := &aws.Operation{
		Name:       opBatchWrite,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/batchwrite",
	}

	if input == nil {
		input = &BatchWriteInput{}
	}

	req := c.newRequest(op, input, &BatchWriteOutput{})

	return BatchWriteRequest{Request: req, Input: input, Copy: c.BatchWriteRequest}
}

// BatchWriteRequest is the request type for the
// BatchWrite API operation.
type BatchWriteRequest struct {
	*aws.Request
	Input *BatchWriteInput
	Copy  func(*BatchWriteInput) BatchWriteRequest
}

// Send marshals and sends the BatchWrite API request.
func (r BatchWriteRequest) Send(ctx context.Context) (*BatchWriteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchWriteResponse{
		BatchWriteOutput: r.Request.Data.(*BatchWriteOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchWriteResponse is the response type for the
// BatchWrite API operation.
type BatchWriteResponse struct {
	*BatchWriteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchWrite request.
func (r *BatchWriteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}