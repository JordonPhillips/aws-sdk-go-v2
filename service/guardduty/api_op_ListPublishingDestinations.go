// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListPublishingDestinationsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the detector to retrieve publishing destinations for.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`

	// The maximum number of results to return in the response.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// A token to use for paginating results that are returned in the response.
	// Set the value of this parameter to null for the first request to a list action.
	// For subsequent calls, use the NextToken value returned from the previous
	// request to continue listing results after the first page.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListPublishingDestinationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPublishingDestinationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListPublishingDestinationsInput"}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
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
func (s ListPublishingDestinationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DetectorId != nil {
		v := *s.DetectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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

type ListPublishingDestinationsOutput struct {
	_ struct{} `type:"structure"`

	// A Destinations object that includes information about each publishing destination
	// returned.
	//
	// Destinations is a required field
	Destinations []Destination `locationName:"destinations" type:"list" required:"true"`

	// A token to use for paginating results that are returned in the response.
	// Set the value of this parameter to null for the first request to a list action.
	// For subsequent calls, use the NextToken value returned from the previous
	// request to continue listing results after the first page.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListPublishingDestinationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPublishingDestinationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Destinations != nil {
		v := s.Destinations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "destinations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListPublishingDestinations = "ListPublishingDestinations"

// ListPublishingDestinationsRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Returns a list of publishing destinations associated with the specified dectectorId.
//
//    // Example sending a request using ListPublishingDestinationsRequest.
//    req := client.ListPublishingDestinationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/ListPublishingDestinations
func (c *Client) ListPublishingDestinationsRequest(input *ListPublishingDestinationsInput) ListPublishingDestinationsRequest {
	op := &aws.Operation{
		Name:       opListPublishingDestinations,
		HTTPMethod: "GET",
		HTTPPath:   "/detector/{detectorId}/publishingDestination",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListPublishingDestinationsInput{}
	}

	req := c.newRequest(op, input, &ListPublishingDestinationsOutput{})

	return ListPublishingDestinationsRequest{Request: req, Input: input, Copy: c.ListPublishingDestinationsRequest}
}

// ListPublishingDestinationsRequest is the request type for the
// ListPublishingDestinations API operation.
type ListPublishingDestinationsRequest struct {
	*aws.Request
	Input *ListPublishingDestinationsInput
	Copy  func(*ListPublishingDestinationsInput) ListPublishingDestinationsRequest
}

// Send marshals and sends the ListPublishingDestinations API request.
func (r ListPublishingDestinationsRequest) Send(ctx context.Context) (*ListPublishingDestinationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPublishingDestinationsResponse{
		ListPublishingDestinationsOutput: r.Request.Data.(*ListPublishingDestinationsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListPublishingDestinationsRequestPaginator returns a paginator for ListPublishingDestinations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListPublishingDestinationsRequest(input)
//   p := guardduty.NewListPublishingDestinationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListPublishingDestinationsPaginator(req ListPublishingDestinationsRequest) ListPublishingDestinationsPaginator {
	return ListPublishingDestinationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListPublishingDestinationsInput
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

// ListPublishingDestinationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListPublishingDestinationsPaginator struct {
	aws.Pager
}

func (p *ListPublishingDestinationsPaginator) CurrentPage() *ListPublishingDestinationsOutput {
	return p.Pager.CurrentPage().(*ListPublishingDestinationsOutput)
}

// ListPublishingDestinationsResponse is the response type for the
// ListPublishingDestinations API operation.
type ListPublishingDestinationsResponse struct {
	*ListPublishingDestinationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPublishingDestinations request.
func (r *ListPublishingDestinationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}