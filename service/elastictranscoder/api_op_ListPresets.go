// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elastictranscoder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The ListPresetsRequest structure.
type ListPresetsInput struct {
	_ struct{} `type:"structure"`

	// To list presets in chronological order by the date and time that they were
	// created, enter true. To list presets in reverse chronological order, enter
	// false.
	Ascending *string `location:"querystring" locationName:"Ascending" type:"string"`

	// When Elastic Transcoder returns more than one page of results, use pageToken
	// in subsequent GET requests to get each successive page of results.
	PageToken *string `location:"querystring" locationName:"PageToken" type:"string"`
}

// String returns the string representation
func (s ListPresetsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPresetsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Ascending != nil {
		v := *s.Ascending

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "Ascending", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PageToken != nil {
		v := *s.PageToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "PageToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The ListPresetsResponse structure.
type ListPresetsOutput struct {
	_ struct{} `type:"structure"`

	// A value that you use to access the second and subsequent pages of results,
	// if any. When the presets fit on one page or when you've reached the last
	// page of results, the value of NextPageToken is null.
	NextPageToken *string `type:"string"`

	// An array of Preset objects.
	Presets []Preset `type:"list"`
}

// String returns the string representation
func (s ListPresetsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPresetsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextPageToken != nil {
		v := *s.NextPageToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextPageToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Presets != nil {
		v := s.Presets

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Presets", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListPresets = "ListPresets"

// ListPresetsRequest returns a request value for making API operation for
// Amazon Elastic Transcoder.
//
// The ListPresets operation gets a list of the default presets included with
// Elastic Transcoder and the presets that you've added in an AWS region.
//
//    // Example sending a request using ListPresetsRequest.
//    req := client.ListPresetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListPresetsRequest(input *ListPresetsInput) ListPresetsRequest {
	op := &aws.Operation{
		Name:       opListPresets,
		HTTPMethod: "GET",
		HTTPPath:   "/2012-09-25/presets",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"PageToken"},
			OutputTokens:    []string{"NextPageToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListPresetsInput{}
	}

	req := c.newRequest(op, input, &ListPresetsOutput{})

	return ListPresetsRequest{Request: req, Input: input, Copy: c.ListPresetsRequest}
}

// ListPresetsRequest is the request type for the
// ListPresets API operation.
type ListPresetsRequest struct {
	*aws.Request
	Input *ListPresetsInput
	Copy  func(*ListPresetsInput) ListPresetsRequest
}

// Send marshals and sends the ListPresets API request.
func (r ListPresetsRequest) Send(ctx context.Context) (*ListPresetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPresetsResponse{
		ListPresetsOutput: r.Request.Data.(*ListPresetsOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListPresetsRequestPaginator returns a paginator for ListPresets.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListPresetsRequest(input)
//   p := elastictranscoder.NewListPresetsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListPresetsPaginator(req ListPresetsRequest) ListPresetsPaginator {
	return ListPresetsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListPresetsInput
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

// ListPresetsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListPresetsPaginator struct {
	aws.Pager
}

func (p *ListPresetsPaginator) CurrentPage() *ListPresetsOutput {
	return p.Pager.CurrentPage().(*ListPresetsOutput)
}

// ListPresetsResponse is the response type for the
// ListPresets API operation.
type ListPresetsResponse struct {
	*ListPresetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPresets request.
func (r *ListPresetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}