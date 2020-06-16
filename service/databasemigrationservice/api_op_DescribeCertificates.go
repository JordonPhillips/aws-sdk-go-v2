// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeCertificatesInput struct {
	_ struct{} `type:"structure"`

	// Filters applied to the certificate described in the form of key-value pairs.
	Filters []Filter `type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 10
	MaxRecords *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeCertificatesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeCertificatesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeCertificatesInput"}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeCertificatesOutput struct {
	_ struct{} `type:"structure"`

	// The Secure Sockets Layer (SSL) certificates associated with the replication
	// instance.
	Certificates []Certificate `type:"list"`

	// The pagination token.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeCertificatesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeCertificates = "DescribeCertificates"

// DescribeCertificatesRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Provides a description of the certificate.
//
//    // Example sending a request using DescribeCertificatesRequest.
//    req := client.DescribeCertificatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DescribeCertificates
func (c *Client) DescribeCertificatesRequest(input *DescribeCertificatesInput) DescribeCertificatesRequest {
	op := &aws.Operation{
		Name:       opDescribeCertificates,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeCertificatesInput{}
	}

	req := c.newRequest(op, input, &DescribeCertificatesOutput{})

	return DescribeCertificatesRequest{Request: req, Input: input, Copy: c.DescribeCertificatesRequest}
}

// DescribeCertificatesRequest is the request type for the
// DescribeCertificates API operation.
type DescribeCertificatesRequest struct {
	*aws.Request
	Input *DescribeCertificatesInput
	Copy  func(*DescribeCertificatesInput) DescribeCertificatesRequest
}

// Send marshals and sends the DescribeCertificates API request.
func (r DescribeCertificatesRequest) Send(ctx context.Context) (*DescribeCertificatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCertificatesResponse{
		DescribeCertificatesOutput: r.Request.Data.(*DescribeCertificatesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeCertificatesRequestPaginator returns a paginator for DescribeCertificates.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeCertificatesRequest(input)
//   p := databasemigrationservice.NewDescribeCertificatesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeCertificatesPaginator(req DescribeCertificatesRequest) DescribeCertificatesPaginator {
	return DescribeCertificatesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeCertificatesInput
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

// DescribeCertificatesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeCertificatesPaginator struct {
	aws.Pager
}

func (p *DescribeCertificatesPaginator) CurrentPage() *DescribeCertificatesOutput {
	return p.Pager.CurrentPage().(*DescribeCertificatesOutput)
}

// DescribeCertificatesResponse is the response type for the
// DescribeCertificates API operation.
type DescribeCertificatesResponse struct {
	*DescribeCertificatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCertificates request.
func (r *DescribeCertificatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}