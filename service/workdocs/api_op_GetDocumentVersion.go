// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetDocumentVersionInput struct {
	_ struct{} `type:"structure"`

	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string `location:"header" locationName:"Authentication" min:"1" type:"string" sensitive:"true"`

	// The ID of the document.
	//
	// DocumentId is a required field
	DocumentId *string `location:"uri" locationName:"DocumentId" min:"1" type:"string" required:"true"`

	// A comma-separated list of values. Specify "SOURCE" to include a URL for the
	// source document.
	Fields *string `location:"querystring" locationName:"fields" min:"1" type:"string"`

	// Set this to TRUE to include custom metadata in the response.
	IncludeCustomMetadata *bool `location:"querystring" locationName:"includeCustomMetadata" type:"boolean"`

	// The version ID of the document.
	//
	// VersionId is a required field
	VersionId *string `location:"uri" locationName:"VersionId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDocumentVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDocumentVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDocumentVersionInput"}
	if s.AuthenticationToken != nil && len(*s.AuthenticationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthenticationToken", 1))
	}

	if s.DocumentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DocumentId"))
	}
	if s.DocumentId != nil && len(*s.DocumentId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DocumentId", 1))
	}
	if s.Fields != nil && len(*s.Fields) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Fields", 1))
	}

	if s.VersionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VersionId"))
	}
	if s.VersionId != nil && len(*s.VersionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VersionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDocumentVersionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AuthenticationToken != nil {
		v := *s.AuthenticationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Authentication", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DocumentId != nil {
		v := *s.DocumentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DocumentId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "VersionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Fields != nil {
		v := *s.Fields

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "fields", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IncludeCustomMetadata != nil {
		v := *s.IncludeCustomMetadata

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "includeCustomMetadata", protocol.BoolValue(v), metadata)
	}
	return nil
}

type GetDocumentVersionOutput struct {
	_ struct{} `type:"structure"`

	// The custom metadata on the document version.
	CustomMetadata map[string]string `min:"1" type:"map"`

	// The version metadata.
	Metadata *DocumentVersionMetadata `type:"structure"`
}

// String returns the string representation
func (s GetDocumentVersionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDocumentVersionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CustomMetadata != nil {
		v := s.CustomMetadata

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "CustomMetadata", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.Metadata != nil {
		v := s.Metadata

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Metadata", v, metadata)
	}
	return nil
}

const opGetDocumentVersion = "GetDocumentVersion"

// GetDocumentVersionRequest returns a request value for making API operation for
// Amazon WorkDocs.
//
// Retrieves version metadata for the specified document.
//
//    // Example sending a request using GetDocumentVersionRequest.
//    req := client.GetDocumentVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/GetDocumentVersion
func (c *Client) GetDocumentVersionRequest(input *GetDocumentVersionInput) GetDocumentVersionRequest {
	op := &aws.Operation{
		Name:       opGetDocumentVersion,
		HTTPMethod: "GET",
		HTTPPath:   "/api/v1/documents/{DocumentId}/versions/{VersionId}",
	}

	if input == nil {
		input = &GetDocumentVersionInput{}
	}

	req := c.newRequest(op, input, &GetDocumentVersionOutput{})

	return GetDocumentVersionRequest{Request: req, Input: input, Copy: c.GetDocumentVersionRequest}
}

// GetDocumentVersionRequest is the request type for the
// GetDocumentVersion API operation.
type GetDocumentVersionRequest struct {
	*aws.Request
	Input *GetDocumentVersionInput
	Copy  func(*GetDocumentVersionInput) GetDocumentVersionRequest
}

// Send marshals and sends the GetDocumentVersion API request.
func (r GetDocumentVersionRequest) Send(ctx context.Context) (*GetDocumentVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDocumentVersionResponse{
		GetDocumentVersionOutput: r.Request.Data.(*GetDocumentVersionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDocumentVersionResponse is the response type for the
// GetDocumentVersion API operation.
type GetDocumentVersionResponse struct {
	*GetDocumentVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDocumentVersion request.
func (r *GetDocumentVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}