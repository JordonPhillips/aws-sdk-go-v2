// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appmesh

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteVirtualNodeInput struct {
	_ struct{} `type:"structure"`

	// MeshName is a required field
	MeshName *string `location:"uri" locationName:"meshName" min:"1" type:"string" required:"true"`

	MeshOwner *string `location:"querystring" locationName:"meshOwner" min:"12" type:"string"`

	// VirtualNodeName is a required field
	VirtualNodeName *string `location:"uri" locationName:"virtualNodeName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteVirtualNodeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVirtualNodeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteVirtualNodeInput"}

	if s.MeshName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeshName"))
	}
	if s.MeshName != nil && len(*s.MeshName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MeshName", 1))
	}
	if s.MeshOwner != nil && len(*s.MeshOwner) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("MeshOwner", 12))
	}

	if s.VirtualNodeName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VirtualNodeName"))
	}
	if s.VirtualNodeName != nil && len(*s.VirtualNodeName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VirtualNodeName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteVirtualNodeInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MeshName != nil {
		v := *s.MeshName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "meshName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VirtualNodeName != nil {
		v := *s.VirtualNodeName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "virtualNodeName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MeshOwner != nil {
		v := *s.MeshOwner

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "meshOwner", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteVirtualNodeOutput struct {
	_ struct{} `type:"structure" payload:"VirtualNode"`

	// An object that represents a virtual node returned by a describe operation.
	//
	// VirtualNode is a required field
	VirtualNode *VirtualNodeData `locationName:"virtualNode" type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteVirtualNodeOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteVirtualNodeOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.VirtualNode != nil {
		v := s.VirtualNode

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "virtualNode", v, metadata)
	}
	return nil
}

const opDeleteVirtualNode = "DeleteVirtualNode"

// DeleteVirtualNodeRequest returns a request value for making API operation for
// AWS App Mesh.
//
// Deletes an existing virtual node.
//
// You must delete any virtual services that list a virtual node as a service
// provider before you can delete the virtual node itself.
//
//    // Example sending a request using DeleteVirtualNodeRequest.
//    req := client.DeleteVirtualNodeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/DeleteVirtualNode
func (c *Client) DeleteVirtualNodeRequest(input *DeleteVirtualNodeInput) DeleteVirtualNodeRequest {
	op := &aws.Operation{
		Name:       opDeleteVirtualNode,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v20190125/meshes/{meshName}/virtualNodes/{virtualNodeName}",
	}

	if input == nil {
		input = &DeleteVirtualNodeInput{}
	}

	req := c.newRequest(op, input, &DeleteVirtualNodeOutput{})

	return DeleteVirtualNodeRequest{Request: req, Input: input, Copy: c.DeleteVirtualNodeRequest}
}

// DeleteVirtualNodeRequest is the request type for the
// DeleteVirtualNode API operation.
type DeleteVirtualNodeRequest struct {
	*aws.Request
	Input *DeleteVirtualNodeInput
	Copy  func(*DeleteVirtualNodeInput) DeleteVirtualNodeRequest
}

// Send marshals and sends the DeleteVirtualNode API request.
func (r DeleteVirtualNodeRequest) Send(ctx context.Context) (*DeleteVirtualNodeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVirtualNodeResponse{
		DeleteVirtualNodeOutput: r.Request.Data.(*DeleteVirtualNodeOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVirtualNodeResponse is the response type for the
// DeleteVirtualNode API operation.
type DeleteVirtualNodeResponse struct {
	*DeleteVirtualNodeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVirtualNode request.
func (r *DeleteVirtualNodeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}