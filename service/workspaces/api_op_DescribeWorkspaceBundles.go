// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workspaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeWorkspaceBundlesInput struct {
	_ struct{} `type:"structure"`

	// The identifiers of the bundles. You cannot combine this parameter with any
	// other filter.
	BundleIds []string `min:"1" type:"list"`

	// The token for the next set of results. (You received this token from a previous
	// call.)
	NextToken *string `min:"1" type:"string"`

	// The owner of the bundles. You cannot combine this parameter with any other
	// filter.
	//
	// Specify AMAZON to describe the bundles provided by AWS or null to describe
	// the bundles that belong to your account.
	Owner *string `type:"string"`
}

// String returns the string representation
func (s DescribeWorkspaceBundlesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeWorkspaceBundlesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeWorkspaceBundlesInput"}
	if s.BundleIds != nil && len(s.BundleIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BundleIds", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeWorkspaceBundlesOutput struct {
	_ struct{} `type:"structure"`

	// Information about the bundles.
	Bundles []WorkspaceBundle `type:"list"`

	// The token to use to retrieve the next set of results, or null if there are
	// no more results available. This token is valid for one day and must be used
	// within that time frame.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeWorkspaceBundlesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeWorkspaceBundles = "DescribeWorkspaceBundles"

// DescribeWorkspaceBundlesRequest returns a request value for making API operation for
// Amazon WorkSpaces.
//
// Retrieves a list that describes the available WorkSpace bundles.
//
// You can filter the results using either bundle ID or owner, but not both.
//
//    // Example sending a request using DescribeWorkspaceBundlesRequest.
//    req := client.DescribeWorkspaceBundlesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/DescribeWorkspaceBundles
func (c *Client) DescribeWorkspaceBundlesRequest(input *DescribeWorkspaceBundlesInput) DescribeWorkspaceBundlesRequest {
	op := &aws.Operation{
		Name:       opDescribeWorkspaceBundles,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeWorkspaceBundlesInput{}
	}

	req := c.newRequest(op, input, &DescribeWorkspaceBundlesOutput{})

	return DescribeWorkspaceBundlesRequest{Request: req, Input: input, Copy: c.DescribeWorkspaceBundlesRequest}
}

// DescribeWorkspaceBundlesRequest is the request type for the
// DescribeWorkspaceBundles API operation.
type DescribeWorkspaceBundlesRequest struct {
	*aws.Request
	Input *DescribeWorkspaceBundlesInput
	Copy  func(*DescribeWorkspaceBundlesInput) DescribeWorkspaceBundlesRequest
}

// Send marshals and sends the DescribeWorkspaceBundles API request.
func (r DescribeWorkspaceBundlesRequest) Send(ctx context.Context) (*DescribeWorkspaceBundlesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeWorkspaceBundlesResponse{
		DescribeWorkspaceBundlesOutput: r.Request.Data.(*DescribeWorkspaceBundlesOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeWorkspaceBundlesRequestPaginator returns a paginator for DescribeWorkspaceBundles.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeWorkspaceBundlesRequest(input)
//   p := workspaces.NewDescribeWorkspaceBundlesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeWorkspaceBundlesPaginator(req DescribeWorkspaceBundlesRequest) DescribeWorkspaceBundlesPaginator {
	return DescribeWorkspaceBundlesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeWorkspaceBundlesInput
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

// DescribeWorkspaceBundlesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeWorkspaceBundlesPaginator struct {
	aws.Pager
}

func (p *DescribeWorkspaceBundlesPaginator) CurrentPage() *DescribeWorkspaceBundlesOutput {
	return p.Pager.CurrentPage().(*DescribeWorkspaceBundlesOutput)
}

// DescribeWorkspaceBundlesResponse is the response type for the
// DescribeWorkspaceBundles API operation.
type DescribeWorkspaceBundlesResponse struct {
	*DescribeWorkspaceBundlesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeWorkspaceBundles request.
func (r *DescribeWorkspaceBundlesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}