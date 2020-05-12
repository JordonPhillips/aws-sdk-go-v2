// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListTagsForResourceInput struct {
	_ struct{} `type:"structure"`

	Limit *int64 `type:"integer"`

	NextMarker *string `min:"1" type:"string"`

	// ResourceARN is a required field
	ResourceARN *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListTagsForResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTagsForResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTagsForResourceInput"}
	if s.NextMarker != nil && len(*s.NextMarker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextMarker", 1))
	}

	if s.ResourceARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceARN"))
	}
	if s.ResourceARN != nil && len(*s.ResourceARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceARN", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListTagsForResourceOutput struct {
	_ struct{} `type:"structure"`

	NextMarker *string `min:"1" type:"string"`

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
	// Information for a tag associated with an AWS resource. Tags are key:value
	// pairs that you can use to categorize and manage your resources, for purposes
	// like billing. For example, you might set the tag key to "customer" and the
	// value to the customer name or ID. You can specify one or more tags to add
	// to each AWS resource, up to 50 tags for a resource.
	//
	// Tagging is only available through the API, SDKs, and CLI. You can't manage
	// or view tags through the AWS WAF Classic console. You can tag the AWS resources
	// that you manage through AWS WAF Classic: web ACLs, rule groups, and rules.
	TagInfoForResource *TagInfoForResource `type:"structure"`
}

// String returns the string representation
func (s ListTagsForResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opListTagsForResource = "ListTagsForResource"

// ListTagsForResourceRequest returns a request value for making API operation for
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
// Retrieves the tags associated with the specified AWS resource. Tags are key:value
// pairs that you can use to categorize and manage your resources, for purposes
// like billing. For example, you might set the tag key to "customer" and the
// value to the customer name or ID. You can specify one or more tags to add
// to each AWS resource, up to 50 tags for a resource.
//
// Tagging is only available through the API, SDKs, and CLI. You can't manage
// or view tags through the AWS WAF Classic console. You can tag the AWS resources
// that you manage through AWS WAF Classic: web ACLs, rule groups, and rules.
//
//    // Example sending a request using ListTagsForResourceRequest.
//    req := client.ListTagsForResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/ListTagsForResource
func (c *Client) ListTagsForResourceRequest(input *ListTagsForResourceInput) ListTagsForResourceRequest {
	op := &aws.Operation{
		Name:       opListTagsForResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListTagsForResourceInput{}
	}

	req := c.newRequest(op, input, &ListTagsForResourceOutput{})

	return ListTagsForResourceRequest{Request: req, Input: input, Copy: c.ListTagsForResourceRequest}
}

// ListTagsForResourceRequest is the request type for the
// ListTagsForResource API operation.
type ListTagsForResourceRequest struct {
	*aws.Request
	Input *ListTagsForResourceInput
	Copy  func(*ListTagsForResourceInput) ListTagsForResourceRequest
}

// Send marshals and sends the ListTagsForResource API request.
func (r ListTagsForResourceRequest) Send(ctx context.Context) (*ListTagsForResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTagsForResourceResponse{
		ListTagsForResourceOutput: r.Request.Data.(*ListTagsForResourceOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListTagsForResourceResponse is the response type for the
// ListTagsForResource API operation.
type ListTagsForResourceResponse struct {
	*ListTagsForResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTagsForResource request.
func (r *ListTagsForResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}