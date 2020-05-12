// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the result of a list device pools request.
type ListDevicePoolsInput struct {
	_ struct{} `type:"structure"`

	// The project ARN.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`

	// The device pools' type.
	//
	// Allowed values include:
	//
	//    * CURATED: A device pool that is created and managed by AWS Device Farm.
	//
	//    * PRIVATE: A device pool that is created and managed by the device pool
	//    developer.
	Type DevicePoolType `locationName:"type" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListDevicePoolsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDevicePoolsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDevicePoolsInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}
	if s.NextToken != nil && len(*s.NextToken) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 4))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the result of a list device pools request.
type ListDevicePoolsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the device pools.
	DevicePools []DevicePool `locationName:"devicePools" type:"list"`

	// If the number of items that are returned is significantly large, this is
	// an identifier that is also returned. It can be used in a subsequent call
	// to this operation to return the next set of items in the list.
	NextToken *string `locationName:"nextToken" min:"4" type:"string"`
}

// String returns the string representation
func (s ListDevicePoolsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListDevicePools = "ListDevicePools"

// ListDevicePoolsRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Gets information about device pools.
//
//    // Example sending a request using ListDevicePoolsRequest.
//    req := client.ListDevicePoolsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/ListDevicePools
func (c *Client) ListDevicePoolsRequest(input *ListDevicePoolsInput) ListDevicePoolsRequest {
	op := &aws.Operation{
		Name:       opListDevicePools,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDevicePoolsInput{}
	}

	req := c.newRequest(op, input, &ListDevicePoolsOutput{})

	return ListDevicePoolsRequest{Request: req, Input: input, Copy: c.ListDevicePoolsRequest}
}

// ListDevicePoolsRequest is the request type for the
// ListDevicePools API operation.
type ListDevicePoolsRequest struct {
	*aws.Request
	Input *ListDevicePoolsInput
	Copy  func(*ListDevicePoolsInput) ListDevicePoolsRequest
}

// Send marshals and sends the ListDevicePools API request.
func (r ListDevicePoolsRequest) Send(ctx context.Context) (*ListDevicePoolsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDevicePoolsResponse{
		ListDevicePoolsOutput: r.Request.Data.(*ListDevicePoolsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDevicePoolsRequestPaginator returns a paginator for ListDevicePools.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDevicePoolsRequest(input)
//   p := devicefarm.NewListDevicePoolsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDevicePoolsPaginator(req ListDevicePoolsRequest) ListDevicePoolsPaginator {
	return ListDevicePoolsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDevicePoolsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListDevicePoolsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDevicePoolsPaginator struct {
	aws.Pager
}

func (p *ListDevicePoolsPaginator) CurrentPage() *ListDevicePoolsOutput {
	return p.Pager.CurrentPage().(*ListDevicePoolsOutput)
}

// ListDevicePoolsResponse is the response type for the
// ListDevicePools API operation.
type ListDevicePoolsResponse struct {
	*ListDevicePoolsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDevicePools request.
func (r *ListDevicePoolsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}