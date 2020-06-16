// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListSubscribedWorkteamsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of work teams to return in each page of the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// A string in the work team name. This filter returns only work teams whose
	// name contains the specified string.
	NameContains *string `min:"1" type:"string"`

	// If the result of the previous ListSubscribedWorkteams request was truncated,
	// the response includes a NextToken. To retrieve the next set of labeling jobs,
	// use the token in the next request.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListSubscribedWorkteamsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListSubscribedWorkteamsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListSubscribedWorkteamsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NameContains != nil && len(*s.NameContains) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NameContains", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListSubscribedWorkteamsOutput struct {
	_ struct{} `type:"structure"`

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of work teams, use it in the subsequent request.
	NextToken *string `type:"string"`

	// An array of Workteam objects, each describing a work team.
	//
	// SubscribedWorkteams is a required field
	SubscribedWorkteams []SubscribedWorkteam `type:"list" required:"true"`
}

// String returns the string representation
func (s ListSubscribedWorkteamsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListSubscribedWorkteams = "ListSubscribedWorkteams"

// ListSubscribedWorkteamsRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Gets a list of the work teams that you are subscribed to in the AWS Marketplace.
// The list may be empty if no work team satisfies the filter specified in the
// NameContains parameter.
//
//    // Example sending a request using ListSubscribedWorkteamsRequest.
//    req := client.ListSubscribedWorkteamsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListSubscribedWorkteams
func (c *Client) ListSubscribedWorkteamsRequest(input *ListSubscribedWorkteamsInput) ListSubscribedWorkteamsRequest {
	op := &aws.Operation{
		Name:       opListSubscribedWorkteams,
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
		input = &ListSubscribedWorkteamsInput{}
	}

	req := c.newRequest(op, input, &ListSubscribedWorkteamsOutput{})

	return ListSubscribedWorkteamsRequest{Request: req, Input: input, Copy: c.ListSubscribedWorkteamsRequest}
}

// ListSubscribedWorkteamsRequest is the request type for the
// ListSubscribedWorkteams API operation.
type ListSubscribedWorkteamsRequest struct {
	*aws.Request
	Input *ListSubscribedWorkteamsInput
	Copy  func(*ListSubscribedWorkteamsInput) ListSubscribedWorkteamsRequest
}

// Send marshals and sends the ListSubscribedWorkteams API request.
func (r ListSubscribedWorkteamsRequest) Send(ctx context.Context) (*ListSubscribedWorkteamsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListSubscribedWorkteamsResponse{
		ListSubscribedWorkteamsOutput: r.Request.Data.(*ListSubscribedWorkteamsOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListSubscribedWorkteamsRequestPaginator returns a paginator for ListSubscribedWorkteams.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListSubscribedWorkteamsRequest(input)
//   p := sagemaker.NewListSubscribedWorkteamsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListSubscribedWorkteamsPaginator(req ListSubscribedWorkteamsRequest) ListSubscribedWorkteamsPaginator {
	return ListSubscribedWorkteamsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListSubscribedWorkteamsInput
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

// ListSubscribedWorkteamsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListSubscribedWorkteamsPaginator struct {
	aws.Pager
}

func (p *ListSubscribedWorkteamsPaginator) CurrentPage() *ListSubscribedWorkteamsOutput {
	return p.Pager.CurrentPage().(*ListSubscribedWorkteamsOutput)
}

// ListSubscribedWorkteamsResponse is the response type for the
// ListSubscribedWorkteams API operation.
type ListSubscribedWorkteamsResponse struct {
	*ListSubscribedWorkteamsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListSubscribedWorkteams request.
func (r *ListSubscribedWorkteamsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}