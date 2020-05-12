// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateOriginEndpointInput struct {
	_ struct{} `type:"structure"`

	// CDN Authorization credentials
	Authorization *Authorization `locationName:"authorization" type:"structure"`

	// ChannelId is a required field
	ChannelId *string `locationName:"channelId" type:"string" required:"true"`

	// A Common Media Application Format (CMAF) packaging configuration.
	CmafPackage *CmafPackageCreateOrUpdateParameters `locationName:"cmafPackage" type:"structure"`

	// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
	DashPackage *DashPackage `locationName:"dashPackage" type:"structure"`

	Description *string `locationName:"description" type:"string"`

	// An HTTP Live Streaming (HLS) packaging configuration.
	HlsPackage *HlsPackage `locationName:"hlsPackage" type:"structure"`

	// Id is a required field
	Id *string `locationName:"id" type:"string" required:"true"`

	ManifestName *string `locationName:"manifestName" type:"string"`

	// A Microsoft Smooth Streaming (MSS) packaging configuration.
	MssPackage *MssPackage `locationName:"mssPackage" type:"structure"`

	Origination Origination `locationName:"origination" type:"string" enum:"true"`

	StartoverWindowSeconds *int64 `locationName:"startoverWindowSeconds" type:"integer"`

	// A collection of tags associated with a resource
	Tags map[string]string `locationName:"tags" type:"map"`

	TimeDelaySeconds *int64 `locationName:"timeDelaySeconds" type:"integer"`

	Whitelist []string `locationName:"whitelist" type:"list"`
}

// String returns the string representation
func (s CreateOriginEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateOriginEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateOriginEndpointInput"}

	if s.ChannelId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChannelId"))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Authorization != nil {
		if err := s.Authorization.Validate(); err != nil {
			invalidParams.AddNested("Authorization", err.(aws.ErrInvalidParams))
		}
	}
	if s.CmafPackage != nil {
		if err := s.CmafPackage.Validate(); err != nil {
			invalidParams.AddNested("CmafPackage", err.(aws.ErrInvalidParams))
		}
	}
	if s.DashPackage != nil {
		if err := s.DashPackage.Validate(); err != nil {
			invalidParams.AddNested("DashPackage", err.(aws.ErrInvalidParams))
		}
	}
	if s.HlsPackage != nil {
		if err := s.HlsPackage.Validate(); err != nil {
			invalidParams.AddNested("HlsPackage", err.(aws.ErrInvalidParams))
		}
	}
	if s.MssPackage != nil {
		if err := s.MssPackage.Validate(); err != nil {
			invalidParams.AddNested("MssPackage", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateOriginEndpointInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Authorization != nil {
		v := s.Authorization

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "authorization", v, metadata)
	}
	if s.ChannelId != nil {
		v := *s.ChannelId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "channelId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CmafPackage != nil {
		v := s.CmafPackage

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "cmafPackage", v, metadata)
	}
	if s.DashPackage != nil {
		v := s.DashPackage

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "dashPackage", v, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.HlsPackage != nil {
		v := s.HlsPackage

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "hlsPackage", v, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ManifestName != nil {
		v := *s.ManifestName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "manifestName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MssPackage != nil {
		v := s.MssPackage

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "mssPackage", v, metadata)
	}
	if len(s.Origination) > 0 {
		v := s.Origination

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "origination", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.StartoverWindowSeconds != nil {
		v := *s.StartoverWindowSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "startoverWindowSeconds", protocol.Int64Value(v), metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.TimeDelaySeconds != nil {
		v := *s.TimeDelaySeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "timeDelaySeconds", protocol.Int64Value(v), metadata)
	}
	if s.Whitelist != nil {
		v := s.Whitelist

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "whitelist", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

type CreateOriginEndpointOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `locationName:"arn" type:"string"`

	// CDN Authorization credentials
	Authorization *Authorization `locationName:"authorization" type:"structure"`

	ChannelId *string `locationName:"channelId" type:"string"`

	// A Common Media Application Format (CMAF) packaging configuration.
	CmafPackage *CmafPackage `locationName:"cmafPackage" type:"structure"`

	// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
	DashPackage *DashPackage `locationName:"dashPackage" type:"structure"`

	Description *string `locationName:"description" type:"string"`

	// An HTTP Live Streaming (HLS) packaging configuration.
	HlsPackage *HlsPackage `locationName:"hlsPackage" type:"structure"`

	Id *string `locationName:"id" type:"string"`

	ManifestName *string `locationName:"manifestName" type:"string"`

	// A Microsoft Smooth Streaming (MSS) packaging configuration.
	MssPackage *MssPackage `locationName:"mssPackage" type:"structure"`

	Origination Origination `locationName:"origination" type:"string" enum:"true"`

	StartoverWindowSeconds *int64 `locationName:"startoverWindowSeconds" type:"integer"`

	// A collection of tags associated with a resource
	Tags map[string]string `locationName:"tags" type:"map"`

	TimeDelaySeconds *int64 `locationName:"timeDelaySeconds" type:"integer"`

	Url *string `locationName:"url" type:"string"`

	Whitelist []string `locationName:"whitelist" type:"list"`
}

// String returns the string representation
func (s CreateOriginEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateOriginEndpointOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Authorization != nil {
		v := s.Authorization

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "authorization", v, metadata)
	}
	if s.ChannelId != nil {
		v := *s.ChannelId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "channelId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CmafPackage != nil {
		v := s.CmafPackage

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "cmafPackage", v, metadata)
	}
	if s.DashPackage != nil {
		v := s.DashPackage

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "dashPackage", v, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.HlsPackage != nil {
		v := s.HlsPackage

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "hlsPackage", v, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ManifestName != nil {
		v := *s.ManifestName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "manifestName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MssPackage != nil {
		v := s.MssPackage

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "mssPackage", v, metadata)
	}
	if len(s.Origination) > 0 {
		v := s.Origination

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "origination", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.StartoverWindowSeconds != nil {
		v := *s.StartoverWindowSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "startoverWindowSeconds", protocol.Int64Value(v), metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.TimeDelaySeconds != nil {
		v := *s.TimeDelaySeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "timeDelaySeconds", protocol.Int64Value(v), metadata)
	}
	if s.Url != nil {
		v := *s.Url

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "url", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Whitelist != nil {
		v := s.Whitelist

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "whitelist", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opCreateOriginEndpoint = "CreateOriginEndpoint"

// CreateOriginEndpointRequest returns a request value for making API operation for
// AWS Elemental MediaPackage.
//
// Creates a new OriginEndpoint record.
//
//    // Example sending a request using CreateOriginEndpointRequest.
//    req := client.CreateOriginEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-2017-10-12/CreateOriginEndpoint
func (c *Client) CreateOriginEndpointRequest(input *CreateOriginEndpointInput) CreateOriginEndpointRequest {
	op := &aws.Operation{
		Name:       opCreateOriginEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/origin_endpoints",
	}

	if input == nil {
		input = &CreateOriginEndpointInput{}
	}

	req := c.newRequest(op, input, &CreateOriginEndpointOutput{})

	return CreateOriginEndpointRequest{Request: req, Input: input, Copy: c.CreateOriginEndpointRequest}
}

// CreateOriginEndpointRequest is the request type for the
// CreateOriginEndpoint API operation.
type CreateOriginEndpointRequest struct {
	*aws.Request
	Input *CreateOriginEndpointInput
	Copy  func(*CreateOriginEndpointInput) CreateOriginEndpointRequest
}

// Send marshals and sends the CreateOriginEndpoint API request.
func (r CreateOriginEndpointRequest) Send(ctx context.Context) (*CreateOriginEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateOriginEndpointResponse{
		CreateOriginEndpointOutput: r.Request.Data.(*CreateOriginEndpointOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateOriginEndpointResponse is the response type for the
// CreateOriginEndpoint API operation.
type CreateOriginEndpointResponse struct {
	*CreateOriginEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateOriginEndpoint request.
func (r *CreateOriginEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}