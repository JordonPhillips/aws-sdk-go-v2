// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type UpdateGatewayInput struct {
	_ struct{} `type:"structure"`

	// The ID of the gateway to update.
	//
	// GatewayId is a required field
	GatewayId *string `location:"uri" locationName:"gatewayId" min:"36" type:"string" required:"true"`

	// A unique, friendly name for the gateway.
	//
	// GatewayName is a required field
	GatewayName *string `locationName:"gatewayName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateGatewayInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateGatewayInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateGatewayInput"}

	if s.GatewayId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayId"))
	}
	if s.GatewayId != nil && len(*s.GatewayId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayId", 36))
	}

	if s.GatewayName == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayName"))
	}
	if s.GatewayName != nil && len(*s.GatewayName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateGatewayInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.GatewayName != nil {
		v := *s.GatewayName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "gatewayName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.GatewayId != nil {
		v := *s.GatewayId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "gatewayId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateGatewayOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateGatewayOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateGatewayOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateGateway = "UpdateGateway"

// UpdateGatewayRequest returns a request value for making API operation for
// AWS IoT SiteWise.
//
// Updates a gateway's name.
//
//    // Example sending a request using UpdateGatewayRequest.
//    req := client.UpdateGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotsitewise-2019-12-02/UpdateGateway
func (c *Client) UpdateGatewayRequest(input *UpdateGatewayInput) UpdateGatewayRequest {
	op := &aws.Operation{
		Name:       opUpdateGateway,
		HTTPMethod: "PUT",
		HTTPPath:   "/20200301/gateways/{gatewayId}",
	}

	if input == nil {
		input = &UpdateGatewayInput{}
	}

	req := c.newRequest(op, input, &UpdateGatewayOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	req.Handlers.Build.PushBackNamed(protocol.NewHostPrefixHandler("edge.", nil))
	req.Handlers.Build.PushBackNamed(protocol.ValidateEndpointHostHandler)

	return UpdateGatewayRequest{Request: req, Input: input, Copy: c.UpdateGatewayRequest}
}

// UpdateGatewayRequest is the request type for the
// UpdateGateway API operation.
type UpdateGatewayRequest struct {
	*aws.Request
	Input *UpdateGatewayInput
	Copy  func(*UpdateGatewayInput) UpdateGatewayRequest
}

// Send marshals and sends the UpdateGateway API request.
func (r UpdateGatewayRequest) Send(ctx context.Context) (*UpdateGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateGatewayResponse{
		UpdateGatewayOutput: r.Request.Data.(*UpdateGatewayOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateGatewayResponse is the response type for the
// UpdateGateway API operation.
type UpdateGatewayResponse struct {
	*UpdateGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateGateway request.
func (r *UpdateGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}