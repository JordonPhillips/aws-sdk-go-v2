// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteIPSetInput struct {
	_ struct{} `type:"structure"`

	// The value returned by the most recent call to GetChangeToken.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`

	// The IPSetId of the IPSet that you want to delete. IPSetId is returned by
	// CreateIPSet and by ListIPSets.
	//
	// IPSetId is a required field
	IPSetId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteIPSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteIPSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteIPSetInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if s.IPSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IPSetId"))
	}
	if s.IPSetId != nil && len(*s.IPSetId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IPSetId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteIPSetOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used to submit the DeleteIPSet request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus.
	ChangeToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteIPSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteIPSet = "DeleteIPSet"

// DeleteIPSetRequest returns a request value for making API operation for
// AWS WAF.
//
//
// This is AWS WAF Classic documentation. For more information, see AWS WAF
// Classic (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the AWS
// WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
// With the latest version, AWS WAF has a single set of endpoints for regional
// and global use.
//
// Permanently deletes an IPSet. You can't delete an IPSet if it's still used
// in any Rules or if it still includes any IP addresses.
//
// If you just want to remove an IPSet from a Rule, use UpdateRule.
//
// To permanently delete an IPSet from AWS WAF, perform the following steps:
//
// Update the IPSet to remove IP address ranges, if any. For more information,
// see UpdateIPSet.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of a DeleteIPSet request.
//
// Submit a DeleteIPSet request.
//
//    // Example sending a request using DeleteIPSetRequest.
//    req := client.DeleteIPSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/DeleteIPSet
func (c *Client) DeleteIPSetRequest(input *DeleteIPSetInput) DeleteIPSetRequest {
	op := &aws.Operation{
		Name:       opDeleteIPSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteIPSetInput{}
	}

	req := c.newRequest(op, input, &DeleteIPSetOutput{})

	return DeleteIPSetRequest{Request: req, Input: input, Copy: c.DeleteIPSetRequest}
}

// DeleteIPSetRequest is the request type for the
// DeleteIPSet API operation.
type DeleteIPSetRequest struct {
	*aws.Request
	Input *DeleteIPSetInput
	Copy  func(*DeleteIPSetInput) DeleteIPSetRequest
}

// Send marshals and sends the DeleteIPSet API request.
func (r DeleteIPSetRequest) Send(ctx context.Context) (*DeleteIPSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteIPSetResponse{
		DeleteIPSetOutput: r.Request.Data.(*DeleteIPSetOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteIPSetResponse is the response type for the
// DeleteIPSet API operation.
type DeleteIPSetResponse struct {
	*DeleteIPSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteIPSet request.
func (r *DeleteIPSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}