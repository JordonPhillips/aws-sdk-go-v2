// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elastictranscoder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The ReadPresetRequest structure.
type ReadPresetInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the preset for which you want to get detailed information.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" type:"string" required:"true"`
}

// String returns the string representation
func (s ReadPresetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ReadPresetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ReadPresetInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ReadPresetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The ReadPresetResponse structure.
type ReadPresetOutput struct {
	_ struct{} `type:"structure"`

	// A section of the response body that provides information about the preset.
	Preset *Preset `type:"structure"`
}

// String returns the string representation
func (s ReadPresetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ReadPresetOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Preset != nil {
		v := s.Preset

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Preset", v, metadata)
	}
	return nil
}

const opReadPreset = "ReadPreset"

// ReadPresetRequest returns a request value for making API operation for
// Amazon Elastic Transcoder.
//
// The ReadPreset operation gets detailed information about a preset.
//
//    // Example sending a request using ReadPresetRequest.
//    req := client.ReadPresetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ReadPresetRequest(input *ReadPresetInput) ReadPresetRequest {
	op := &aws.Operation{
		Name:       opReadPreset,
		HTTPMethod: "GET",
		HTTPPath:   "/2012-09-25/presets/{Id}",
	}

	if input == nil {
		input = &ReadPresetInput{}
	}

	req := c.newRequest(op, input, &ReadPresetOutput{})

	return ReadPresetRequest{Request: req, Input: input, Copy: c.ReadPresetRequest}
}

// ReadPresetRequest is the request type for the
// ReadPreset API operation.
type ReadPresetRequest struct {
	*aws.Request
	Input *ReadPresetInput
	Copy  func(*ReadPresetInput) ReadPresetRequest
}

// Send marshals and sends the ReadPreset API request.
func (r ReadPresetRequest) Send(ctx context.Context) (*ReadPresetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ReadPresetResponse{
		ReadPresetOutput: r.Request.Data.(*ReadPresetOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ReadPresetResponse is the response type for the
// ReadPreset API operation.
type ReadPresetResponse struct {
	*ReadPresetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ReadPreset request.
func (r *ReadPresetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}