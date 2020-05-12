// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetKeyPairsInput struct {
	_ struct{} `type:"structure"`

	// The token to advance to the next page of results from your request.
	//
	// To get a page token, perform an initial GetKeyPairs request. If your results
	// are paginated, the response will return a next page token that you can specify
	// as the page token in a subsequent request.
	PageToken *string `locationName:"pageToken" type:"string"`
}

// String returns the string representation
func (s GetKeyPairsInput) String() string {
	return awsutil.Prettify(s)
}

type GetKeyPairsOutput struct {
	_ struct{} `type:"structure"`

	// An array of key-value pairs containing information about the key pairs.
	KeyPairs []KeyPair `locationName:"keyPairs" type:"list"`

	// The token to advance to the next page of resutls from your request.
	//
	// A next page token is not returned if there are no more results to display.
	//
	// To get the next page of results, perform another GetKeyPairs request and
	// specify the next page token using the pageToken parameter.
	NextPageToken *string `locationName:"nextPageToken" type:"string"`
}

// String returns the string representation
func (s GetKeyPairsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetKeyPairs = "GetKeyPairs"

// GetKeyPairsRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns information about all key pairs in the user's account.
//
//    // Example sending a request using GetKeyPairsRequest.
//    req := client.GetKeyPairsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetKeyPairs
func (c *Client) GetKeyPairsRequest(input *GetKeyPairsInput) GetKeyPairsRequest {
	op := &aws.Operation{
		Name:       opGetKeyPairs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetKeyPairsInput{}
	}

	req := c.newRequest(op, input, &GetKeyPairsOutput{})

	return GetKeyPairsRequest{Request: req, Input: input, Copy: c.GetKeyPairsRequest}
}

// GetKeyPairsRequest is the request type for the
// GetKeyPairs API operation.
type GetKeyPairsRequest struct {
	*aws.Request
	Input *GetKeyPairsInput
	Copy  func(*GetKeyPairsInput) GetKeyPairsRequest
}

// Send marshals and sends the GetKeyPairs API request.
func (r GetKeyPairsRequest) Send(ctx context.Context) (*GetKeyPairsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetKeyPairsResponse{
		GetKeyPairsOutput: r.Request.Data.(*GetKeyPairsOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetKeyPairsResponse is the response type for the
// GetKeyPairs API operation.
type GetKeyPairsResponse struct {
	*GetKeyPairsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetKeyPairs request.
func (r *GetKeyPairsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}