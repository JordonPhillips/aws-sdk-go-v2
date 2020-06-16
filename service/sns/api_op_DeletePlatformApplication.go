// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Input for DeletePlatformApplication action.
type DeletePlatformApplicationInput struct {
	_ struct{} `type:"structure"`

	// PlatformApplicationArn of platform application object to delete.
	//
	// PlatformApplicationArn is a required field
	PlatformApplicationArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeletePlatformApplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeletePlatformApplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeletePlatformApplicationInput"}

	if s.PlatformApplicationArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("PlatformApplicationArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeletePlatformApplicationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeletePlatformApplicationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeletePlatformApplication = "DeletePlatformApplication"

// DeletePlatformApplicationRequest returns a request value for making API operation for
// Amazon Simple Notification Service.
//
// Deletes a platform application object for one of the supported push notification
// services, such as APNS and GCM (Firebase Cloud Messaging). For more information,
// see Using Amazon SNS Mobile Push Notifications (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html).
//
//    // Example sending a request using DeletePlatformApplicationRequest.
//    req := client.DeletePlatformApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/DeletePlatformApplication
func (c *Client) DeletePlatformApplicationRequest(input *DeletePlatformApplicationInput) DeletePlatformApplicationRequest {
	op := &aws.Operation{
		Name:       opDeletePlatformApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeletePlatformApplicationInput{}
	}

	req := c.newRequest(op, input, &DeletePlatformApplicationOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeletePlatformApplicationRequest{Request: req, Input: input, Copy: c.DeletePlatformApplicationRequest}
}

// DeletePlatformApplicationRequest is the request type for the
// DeletePlatformApplication API operation.
type DeletePlatformApplicationRequest struct {
	*aws.Request
	Input *DeletePlatformApplicationInput
	Copy  func(*DeletePlatformApplicationInput) DeletePlatformApplicationRequest
}

// Send marshals and sends the DeletePlatformApplication API request.
func (r DeletePlatformApplicationRequest) Send(ctx context.Context) (*DeletePlatformApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeletePlatformApplicationResponse{
		DeletePlatformApplicationOutput: r.Request.Data.(*DeletePlatformApplicationOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeletePlatformApplicationResponse is the response type for the
// DeletePlatformApplication API operation.
type DeletePlatformApplicationResponse struct {
	*DeletePlatformApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeletePlatformApplication request.
func (r *DeletePlatformApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}