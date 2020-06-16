// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudhsm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the inputs for the CreateHapgRequest action.
type CreateHapgInput struct {
	_ struct{} `type:"structure"`

	// The label of the new high-availability partition group.
	//
	// Label is a required field
	Label *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateHapgInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateHapgInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateHapgInput"}

	if s.Label == nil {
		invalidParams.Add(aws.NewErrParamRequired("Label"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of the CreateHAPartitionGroup action.
type CreateHapgOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the high-availability partition group.
	HapgArn *string `type:"string"`
}

// String returns the string representation
func (s CreateHapgOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateHapg = "CreateHapg"

// CreateHapgRequest returns a request value for making API operation for
// Amazon CloudHSM.
//
// This is documentation for AWS CloudHSM Classic. For more information, see
// AWS CloudHSM Classic FAQs (http://aws.amazon.com/cloudhsm/faqs-classic/),
// the AWS CloudHSM Classic User Guide (http://docs.aws.amazon.com/cloudhsm/classic/userguide/),
// and the AWS CloudHSM Classic API Reference (http://docs.aws.amazon.com/cloudhsm/classic/APIReference/).
//
// For information about the current version of AWS CloudHSM, see AWS CloudHSM
// (http://aws.amazon.com/cloudhsm/), the AWS CloudHSM User Guide (http://docs.aws.amazon.com/cloudhsm/latest/userguide/),
// and the AWS CloudHSM API Reference (http://docs.aws.amazon.com/cloudhsm/latest/APIReference/).
//
// Creates a high-availability partition group. A high-availability partition
// group is a group of partitions that spans multiple physical HSMs.
//
//    // Example sending a request using CreateHapgRequest.
//    req := client.CreateHapgRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudhsm-2014-05-30/CreateHapg
func (c *Client) CreateHapgRequest(input *CreateHapgInput) CreateHapgRequest {
	op := &aws.Operation{
		Name:       opCreateHapg,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateHapgInput{}
	}

	req := c.newRequest(op, input, &CreateHapgOutput{})

	return CreateHapgRequest{Request: req, Input: input, Copy: c.CreateHapgRequest}
}

// CreateHapgRequest is the request type for the
// CreateHapg API operation.
type CreateHapgRequest struct {
	*aws.Request
	Input *CreateHapgInput
	Copy  func(*CreateHapgInput) CreateHapgRequest
}

// Send marshals and sends the CreateHapg API request.
func (r CreateHapgRequest) Send(ctx context.Context) (*CreateHapgResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateHapgResponse{
		CreateHapgOutput: r.Request.Data.(*CreateHapgOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateHapgResponse is the response type for the
// CreateHapg API operation.
type CreateHapgResponse struct {
	*CreateHapgOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateHapg request.
func (r *CreateHapgResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}