// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateTemplatePermissionsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the AWS account that contains the template.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// A list of resource permissions to be granted on the template.
	GrantPermissions []ResourcePermission `type:"list"`

	// A list of resource permissions to be revoked from the template.
	RevokePermissions []ResourcePermission `type:"list"`

	// The ID for the template.
	//
	// TemplateId is a required field
	TemplateId *string `location:"uri" locationName:"TemplateId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateTemplatePermissionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateTemplatePermissionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateTemplatePermissionsInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.TemplateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateId"))
	}
	if s.TemplateId != nil && len(*s.TemplateId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateId", 1))
	}
	if s.GrantPermissions != nil {
		for i, v := range s.GrantPermissions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "GrantPermissions", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.RevokePermissions != nil {
		for i, v := range s.RevokePermissions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "RevokePermissions", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateTemplatePermissionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.GrantPermissions != nil {
		v := s.GrantPermissions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "GrantPermissions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.RevokePermissions != nil {
		v := s.RevokePermissions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "RevokePermissions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TemplateId != nil {
		v := *s.TemplateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "TemplateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateTemplatePermissionsOutput struct {
	_ struct{} `type:"structure"`

	// A list of resource permissions to be set on the template.
	Permissions []ResourcePermission `min:"1" type:"list"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The HTTP status of the request.
	Status *int64 `location:"statusCode" type:"integer"`

	// The Amazon Resource Name (ARN) of the template.
	TemplateArn *string `type:"string"`

	// The ID for the template.
	TemplateId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateTemplatePermissionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateTemplatePermissionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Permissions != nil {
		v := s.Permissions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Permissions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TemplateArn != nil {
		v := *s.TemplateArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TemplateArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TemplateId != nil {
		v := *s.TemplateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TemplateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opUpdateTemplatePermissions = "UpdateTemplatePermissions"

// UpdateTemplatePermissionsRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Updates the resource permissions for a template.
//
//    // Example sending a request using UpdateTemplatePermissionsRequest.
//    req := client.UpdateTemplatePermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/UpdateTemplatePermissions
func (c *Client) UpdateTemplatePermissionsRequest(input *UpdateTemplatePermissionsInput) UpdateTemplatePermissionsRequest {
	op := &aws.Operation{
		Name:       opUpdateTemplatePermissions,
		HTTPMethod: "PUT",
		HTTPPath:   "/accounts/{AwsAccountId}/templates/{TemplateId}/permissions",
	}

	if input == nil {
		input = &UpdateTemplatePermissionsInput{}
	}

	req := c.newRequest(op, input, &UpdateTemplatePermissionsOutput{})

	return UpdateTemplatePermissionsRequest{Request: req, Input: input, Copy: c.UpdateTemplatePermissionsRequest}
}

// UpdateTemplatePermissionsRequest is the request type for the
// UpdateTemplatePermissions API operation.
type UpdateTemplatePermissionsRequest struct {
	*aws.Request
	Input *UpdateTemplatePermissionsInput
	Copy  func(*UpdateTemplatePermissionsInput) UpdateTemplatePermissionsRequest
}

// Send marshals and sends the UpdateTemplatePermissions API request.
func (r UpdateTemplatePermissionsRequest) Send(ctx context.Context) (*UpdateTemplatePermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateTemplatePermissionsResponse{
		UpdateTemplatePermissionsOutput: r.Request.Data.(*UpdateTemplatePermissionsOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateTemplatePermissionsResponse is the response type for the
// UpdateTemplatePermissions API operation.
type UpdateTemplatePermissionsResponse struct {
	*UpdateTemplatePermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateTemplatePermissions request.
func (r *UpdateTemplatePermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}