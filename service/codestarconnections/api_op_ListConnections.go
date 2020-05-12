// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestarconnections

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListConnectionsInput struct {
	_ struct{} `type:"structure"`

	// Filters the list of connections to those associated with a specified host.
	HostArnFilter *string `type:"string"`

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `type:"integer"`

	// The token that was returned from the previous ListConnections call, which
	// can be used to return the next set of connections in the list.
	NextToken *string `min:"1" type:"string"`

	// Filters the list of connections to those associated with a specified provider,
	// such as Bitbucket.
	ProviderTypeFilter ProviderType `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListConnectionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListConnectionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListConnectionsInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListConnectionsOutput struct {
	_ struct{} `type:"structure"`

	// A list of connections and the details for each connection, such as status,
	// owner, and provider type.
	Connections []Connection `type:"list"`

	// A token that can be used in the next ListConnections call. To view all items
	// in the list, continue to call this operation with each subsequent token until
	// no more nextToken values are returned.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListConnectionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListConnections = "ListConnections"

// ListConnectionsRequest returns a request value for making API operation for
// AWS CodeStar connections.
//
// Lists the connections associated with your account.
//
//    // Example sending a request using ListConnectionsRequest.
//    req := client.ListConnectionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-connections-2019-12-01/ListConnections
func (c *Client) ListConnectionsRequest(input *ListConnectionsInput) ListConnectionsRequest {
	op := &aws.Operation{
		Name:       opListConnections,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListConnectionsInput{}
	}

	req := c.newRequest(op, input, &ListConnectionsOutput{})

	return ListConnectionsRequest{Request: req, Input: input, Copy: c.ListConnectionsRequest}
}

// ListConnectionsRequest is the request type for the
// ListConnections API operation.
type ListConnectionsRequest struct {
	*aws.Request
	Input *ListConnectionsInput
	Copy  func(*ListConnectionsInput) ListConnectionsRequest
}

// Send marshals and sends the ListConnections API request.
func (r ListConnectionsRequest) Send(ctx context.Context) (*ListConnectionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListConnectionsResponse{
		ListConnectionsOutput: r.Request.Data.(*ListConnectionsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListConnectionsRequestPaginator returns a paginator for ListConnections.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListConnectionsRequest(input)
//   p := codestarconnections.NewListConnectionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListConnectionsPaginator(req ListConnectionsRequest) ListConnectionsPaginator {
	return ListConnectionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListConnectionsInput
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

// ListConnectionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListConnectionsPaginator struct {
	aws.Pager
}

func (p *ListConnectionsPaginator) CurrentPage() *ListConnectionsOutput {
	return p.Pager.CurrentPage().(*ListConnectionsOutput)
}

// ListConnectionsResponse is the response type for the
// ListConnections API operation.
type ListConnectionsResponse struct {
	*ListConnectionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListConnections request.
func (r *ListConnectionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}