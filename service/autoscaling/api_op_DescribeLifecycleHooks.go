// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeLifecycleHooksInput struct {
	_ struct{} `type:"structure"`

	// The name of the Auto Scaling group.
	//
	// AutoScalingGroupName is a required field
	AutoScalingGroupName *string `min:"1" type:"string" required:"true"`

	// The names of one or more lifecycle hooks. If you omit this parameter, all
	// lifecycle hooks are described.
	LifecycleHookNames []string `type:"list"`
}

// String returns the string representation
func (s DescribeLifecycleHooksInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeLifecycleHooksInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeLifecycleHooksInput"}

	if s.AutoScalingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AutoScalingGroupName"))
	}
	if s.AutoScalingGroupName != nil && len(*s.AutoScalingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AutoScalingGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeLifecycleHooksOutput struct {
	_ struct{} `type:"structure"`

	// The lifecycle hooks for the specified group.
	LifecycleHooks []LifecycleHook `type:"list"`
}

// String returns the string representation
func (s DescribeLifecycleHooksOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeLifecycleHooks = "DescribeLifecycleHooks"

// DescribeLifecycleHooksRequest returns a request value for making API operation for
// Auto Scaling.
//
// Describes the lifecycle hooks for the specified Auto Scaling group.
//
//    // Example sending a request using DescribeLifecycleHooksRequest.
//    req := client.DescribeLifecycleHooksRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/DescribeLifecycleHooks
func (c *Client) DescribeLifecycleHooksRequest(input *DescribeLifecycleHooksInput) DescribeLifecycleHooksRequest {
	op := &aws.Operation{
		Name:       opDescribeLifecycleHooks,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeLifecycleHooksInput{}
	}

	req := c.newRequest(op, input, &DescribeLifecycleHooksOutput{})

	return DescribeLifecycleHooksRequest{Request: req, Input: input, Copy: c.DescribeLifecycleHooksRequest}
}

// DescribeLifecycleHooksRequest is the request type for the
// DescribeLifecycleHooks API operation.
type DescribeLifecycleHooksRequest struct {
	*aws.Request
	Input *DescribeLifecycleHooksInput
	Copy  func(*DescribeLifecycleHooksInput) DescribeLifecycleHooksRequest
}

// Send marshals and sends the DescribeLifecycleHooks API request.
func (r DescribeLifecycleHooksRequest) Send(ctx context.Context) (*DescribeLifecycleHooksResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeLifecycleHooksResponse{
		DescribeLifecycleHooksOutput: r.Request.Data.(*DescribeLifecycleHooksOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeLifecycleHooksResponse is the response type for the
// DescribeLifecycleHooks API operation.
type DescribeLifecycleHooksResponse struct {
	*DescribeLifecycleHooksOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeLifecycleHooks request.
func (r *DescribeLifecycleHooksResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}