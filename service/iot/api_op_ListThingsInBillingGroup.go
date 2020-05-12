// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListThingsInBillingGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the billing group.
	//
	// BillingGroupName is a required field
	BillingGroupName *string `location:"uri" locationName:"billingGroupName" min:"1" type:"string" required:"true"`

	// The maximum number of results to return per request.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token to retrieve the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListThingsInBillingGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListThingsInBillingGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListThingsInBillingGroupInput"}

	if s.BillingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BillingGroupName"))
	}
	if s.BillingGroupName != nil && len(*s.BillingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BillingGroupName", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListThingsInBillingGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.BillingGroupName != nil {
		v := *s.BillingGroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "billingGroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListThingsInBillingGroupOutput struct {
	_ struct{} `type:"structure"`

	// The token used to get the next set of results, or null if there are no additional
	// results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// A list of things in the billing group.
	Things []string `locationName:"things" type:"list"`
}

// String returns the string representation
func (s ListThingsInBillingGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListThingsInBillingGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Things != nil {
		v := s.Things

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "things", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opListThingsInBillingGroup = "ListThingsInBillingGroup"

// ListThingsInBillingGroupRequest returns a request value for making API operation for
// AWS IoT.
//
// Lists the things you have added to the given billing group.
//
//    // Example sending a request using ListThingsInBillingGroupRequest.
//    req := client.ListThingsInBillingGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListThingsInBillingGroupRequest(input *ListThingsInBillingGroupInput) ListThingsInBillingGroupRequest {
	op := &aws.Operation{
		Name:       opListThingsInBillingGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/billing-groups/{billingGroupName}/things",
	}

	if input == nil {
		input = &ListThingsInBillingGroupInput{}
	}

	req := c.newRequest(op, input, &ListThingsInBillingGroupOutput{})

	return ListThingsInBillingGroupRequest{Request: req, Input: input, Copy: c.ListThingsInBillingGroupRequest}
}

// ListThingsInBillingGroupRequest is the request type for the
// ListThingsInBillingGroup API operation.
type ListThingsInBillingGroupRequest struct {
	*aws.Request
	Input *ListThingsInBillingGroupInput
	Copy  func(*ListThingsInBillingGroupInput) ListThingsInBillingGroupRequest
}

// Send marshals and sends the ListThingsInBillingGroup API request.
func (r ListThingsInBillingGroupRequest) Send(ctx context.Context) (*ListThingsInBillingGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListThingsInBillingGroupResponse{
		ListThingsInBillingGroupOutput: r.Request.Data.(*ListThingsInBillingGroupOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListThingsInBillingGroupResponse is the response type for the
// ListThingsInBillingGroup API operation.
type ListThingsInBillingGroupResponse struct {
	*ListThingsInBillingGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListThingsInBillingGroup request.
func (r *ListThingsInBillingGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}