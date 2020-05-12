// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListImportsInput struct {
	_ struct{} `type:"structure"`

	// The name of the exported output value. AWS CloudFormation returns the stack
	// names that are importing this value.
	//
	// ExportName is a required field
	ExportName *string `type:"string" required:"true"`

	// A string (provided by the ListImports response output) that identifies the
	// next page of stacks that are importing the specified exported output value.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListImportsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListImportsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListImportsInput"}

	if s.ExportName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ExportName"))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListImportsOutput struct {
	_ struct{} `type:"structure"`

	// A list of stack names that are importing the specified exported output value.
	Imports []string `type:"list"`

	// A string that identifies the next page of exports. If there is no additional
	// page, this value is null.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListImportsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListImports = "ListImports"

// ListImportsRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Lists all stacks that are importing an exported output value. To modify or
// remove an exported output value, first use this action to see which stacks
// are using it. To see the exported output values in your account, see ListExports.
//
// For more information about importing an exported output value, see the Fn::ImportValue
// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-importvalue.html)
// function.
//
//    // Example sending a request using ListImportsRequest.
//    req := client.ListImportsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/ListImports
func (c *Client) ListImportsRequest(input *ListImportsInput) ListImportsRequest {
	op := &aws.Operation{
		Name:       opListImports,
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
		input = &ListImportsInput{}
	}

	req := c.newRequest(op, input, &ListImportsOutput{})

	return ListImportsRequest{Request: req, Input: input, Copy: c.ListImportsRequest}
}

// ListImportsRequest is the request type for the
// ListImports API operation.
type ListImportsRequest struct {
	*aws.Request
	Input *ListImportsInput
	Copy  func(*ListImportsInput) ListImportsRequest
}

// Send marshals and sends the ListImports API request.
func (r ListImportsRequest) Send(ctx context.Context) (*ListImportsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListImportsResponse{
		ListImportsOutput: r.Request.Data.(*ListImportsOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListImportsRequestPaginator returns a paginator for ListImports.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListImportsRequest(input)
//   p := cloudformation.NewListImportsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListImportsPaginator(req ListImportsRequest) ListImportsPaginator {
	return ListImportsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListImportsInput
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

// ListImportsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListImportsPaginator struct {
	aws.Pager
}

func (p *ListImportsPaginator) CurrentPage() *ListImportsOutput {
	return p.Pager.CurrentPage().(*ListImportsOutput)
}

// ListImportsResponse is the response type for the
// ListImports API operation.
type ListImportsResponse struct {
	*ListImportsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListImports request.
func (r *ListImportsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}