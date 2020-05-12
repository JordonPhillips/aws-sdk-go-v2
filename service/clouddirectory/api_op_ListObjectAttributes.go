// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListObjectAttributesInput struct {
	_ struct{} `type:"structure"`

	// Represents the manner and timing in which the successful write or update
	// of an object is reflected in a subsequent read operation of that same object.
	ConsistencyLevel ConsistencyLevel `location:"header" locationName:"x-amz-consistency-level" type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) that is associated with the Directory where
	// the object resides. For more information, see arns.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// Used to filter the list of object attributes that are associated with a certain
	// facet.
	FacetFilter *SchemaFacet `type:"structure"`

	// The maximum number of items to be retrieved in a single call. This is an
	// approximate number.
	MaxResults *int64 `min:"1" type:"integer"`

	// The pagination token.
	NextToken *string `type:"string"`

	// The reference that identifies the object whose attributes will be listed.
	//
	// ObjectReference is a required field
	ObjectReference *ObjectReference `type:"structure" required:"true"`
}

// String returns the string representation
func (s ListObjectAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListObjectAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListObjectAttributesInput"}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.ObjectReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("ObjectReference"))
	}
	if s.FacetFilter != nil {
		if err := s.FacetFilter.Validate(); err != nil {
			invalidParams.AddNested("FacetFilter", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListObjectAttributesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.FacetFilter != nil {
		v := s.FacetFilter

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "FacetFilter", v, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ObjectReference != nil {
		v := s.ObjectReference

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ObjectReference", v, metadata)
	}
	if len(s.ConsistencyLevel) > 0 {
		v := s.ConsistencyLevel

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-consistency-level", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.DirectoryArn != nil {
		v := *s.DirectoryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListObjectAttributesOutput struct {
	_ struct{} `type:"structure"`

	// Attributes map that is associated with the object. AttributeArn is the key,
	// and attribute value is the value.
	Attributes []AttributeKeyAndValue `type:"list"`

	// The pagination token.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListObjectAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListObjectAttributesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Attributes != nil {
		v := s.Attributes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Attributes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListObjectAttributes = "ListObjectAttributes"

// ListObjectAttributesRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Lists all attributes that are associated with an object.
//
//    // Example sending a request using ListObjectAttributesRequest.
//    req := client.ListObjectAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListObjectAttributes
func (c *Client) ListObjectAttributesRequest(input *ListObjectAttributesInput) ListObjectAttributesRequest {
	op := &aws.Operation{
		Name:       opListObjectAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/object/attributes",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListObjectAttributesInput{}
	}

	req := c.newRequest(op, input, &ListObjectAttributesOutput{})

	return ListObjectAttributesRequest{Request: req, Input: input, Copy: c.ListObjectAttributesRequest}
}

// ListObjectAttributesRequest is the request type for the
// ListObjectAttributes API operation.
type ListObjectAttributesRequest struct {
	*aws.Request
	Input *ListObjectAttributesInput
	Copy  func(*ListObjectAttributesInput) ListObjectAttributesRequest
}

// Send marshals and sends the ListObjectAttributes API request.
func (r ListObjectAttributesRequest) Send(ctx context.Context) (*ListObjectAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListObjectAttributesResponse{
		ListObjectAttributesOutput: r.Request.Data.(*ListObjectAttributesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListObjectAttributesRequestPaginator returns a paginator for ListObjectAttributes.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListObjectAttributesRequest(input)
//   p := clouddirectory.NewListObjectAttributesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListObjectAttributesPaginator(req ListObjectAttributesRequest) ListObjectAttributesPaginator {
	return ListObjectAttributesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListObjectAttributesInput
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

// ListObjectAttributesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListObjectAttributesPaginator struct {
	aws.Pager
}

func (p *ListObjectAttributesPaginator) CurrentPage() *ListObjectAttributesOutput {
	return p.Pager.CurrentPage().(*ListObjectAttributesOutput)
}

// ListObjectAttributesResponse is the response type for the
// ListObjectAttributes API operation.
type ListObjectAttributesResponse struct {
	*ListObjectAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListObjectAttributes request.
func (r *ListObjectAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}