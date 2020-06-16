// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mq

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type RebootBrokerInput struct {
	_ struct{} `type:"structure"`

	// BrokerId is a required field
	BrokerId *string `location:"uri" locationName:"broker-id" type:"string" required:"true"`
}

// String returns the string representation
func (s RebootBrokerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RebootBrokerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RebootBrokerInput"}

	if s.BrokerId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BrokerId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RebootBrokerInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.BrokerId != nil {
		v := *s.BrokerId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "broker-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type RebootBrokerOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RebootBrokerOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RebootBrokerOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opRebootBroker = "RebootBroker"

// RebootBrokerRequest returns a request value for making API operation for
// AmazonMQ.
//
// Reboots a broker. Note: This API is asynchronous.
//
//    // Example sending a request using RebootBrokerRequest.
//    req := client.RebootBrokerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/RebootBroker
func (c *Client) RebootBrokerRequest(input *RebootBrokerInput) RebootBrokerRequest {
	op := &aws.Operation{
		Name:       opRebootBroker,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/brokers/{broker-id}/reboot",
	}

	if input == nil {
		input = &RebootBrokerInput{}
	}

	req := c.newRequest(op, input, &RebootBrokerOutput{})

	return RebootBrokerRequest{Request: req, Input: input, Copy: c.RebootBrokerRequest}
}

// RebootBrokerRequest is the request type for the
// RebootBroker API operation.
type RebootBrokerRequest struct {
	*aws.Request
	Input *RebootBrokerInput
	Copy  func(*RebootBrokerInput) RebootBrokerRequest
}

// Send marshals and sends the RebootBroker API request.
func (r RebootBrokerRequest) Send(ctx context.Context) (*RebootBrokerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RebootBrokerResponse{
		RebootBrokerOutput: r.Request.Data.(*RebootBrokerOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RebootBrokerResponse is the response type for the
// RebootBroker API operation.
type RebootBrokerResponse struct {
	*RebootBrokerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RebootBroker request.
func (r *RebootBrokerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}