// Code generated by smithy-go-codegen DO NOT EDIT.
package restxml

import (
	"bytes"
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/restxml/types"
	"github.com/awslabs/smithy-go/middleware"
	"github.com/awslabs/smithy-go/ptr"
	smithyrand "github.com/awslabs/smithy-go/rand"
	smithytesting "github.com/awslabs/smithy-go/testing"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestClient_GreetingWithErrors_awsRestxmlDeserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *GreetingWithErrorsOutput
	}{
		// Ensures that operations with errors successfully know how to deserialize the
		// successful response
		"GreetingWithErrors": {
			StatusCode: 200,
			Header: http.Header{
				"X-Greeting": []string{"Hello"},
			},
			Body: []byte(``),
			ExpectResult: &GreetingWithErrorsOutput{
				Greeting: ptr.String("Hello"),
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				for k, vs := range c.Header {
					for _, v := range vs {
						w.Header().Add(k, v)
					}
				}
				if len(c.BodyMediaType) != 0 && len(w.Header().Values("Content-Type")) == 0 {
					w.Header().Set("Content-Type", c.BodyMediaType)
				}
				if len(c.Body) != 0 {
					w.Header().Set("Content-Length", strconv.Itoa(len(c.Body)))
				}
				w.WriteHeader(c.StatusCode)
				if len(c.Body) != 0 {
					if _, err := io.Copy(w, bytes.NewReader(c.Body)); err != nil {
						t.Errorf("failed to write response body, %v", err)
					}
				}
			}))
			defer server.Close()
			client := New(Options{
				APIOptions: []APIOptionFunc{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: aws.EndpointResolverFunc(func(service, region string) (e aws.Endpoint, err error) {
					e.URL = server.URL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient:               aws.NewBuildableHTTPClient(),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params GreetingWithErrorsInput
			result, err := client.GreetingWithErrors(context.Background(), &params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if diff := cmp.Diff(c.ExpectResult, result, cmpopts.IgnoreUnexported(middleware.Metadata{})); len(diff) != 0 {
				t.Errorf("expect c.ExpectResult value match:\n%s", diff)
			}
		})
	}
}

func TestClient_GreetingWithErrors_InvalidGreeting_awsRestxmlDeserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectError   *types.InvalidGreeting
	}{
		// Parses simple XML errors
		"InvalidGreetingError": {
			StatusCode: 400,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<ErrorResponse>
			   <Error>
			      <Type>Sender</Type>
			      <Code>InvalidGreeting</Code>
			      <Message>Hi</Message>
			      <AnotherSetting>setting</Message>
			   </Error>
			   <RequestId>foo-id</RequestId>
			</ErrorResponse>
			`),
			ExpectError: &types.InvalidGreeting{
				Message: ptr.String("Hi"),
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				for k, vs := range c.Header {
					for _, v := range vs {
						w.Header().Add(k, v)
					}
				}
				if len(c.BodyMediaType) != 0 && len(w.Header().Values("Content-Type")) == 0 {
					w.Header().Set("Content-Type", c.BodyMediaType)
				}
				if len(c.Body) != 0 {
					w.Header().Set("Content-Length", strconv.Itoa(len(c.Body)))
				}
				w.WriteHeader(c.StatusCode)
				if len(c.Body) != 0 {
					if _, err := io.Copy(w, bytes.NewReader(c.Body)); err != nil {
						t.Errorf("failed to write response body, %v", err)
					}
				}
			}))
			defer server.Close()
			client := New(Options{
				APIOptions: []APIOptionFunc{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: aws.EndpointResolverFunc(func(service, region string) (e aws.Endpoint, err error) {
					e.URL = server.URL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient:               aws.NewBuildableHTTPClient(),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params GreetingWithErrorsInput
			result, err := client.GreetingWithErrors(context.Background(), &params)
			if err == nil {
				t.Fatalf("expect not nil err")
			}
			if result != nil {
				t.Fatalf("expect nil result, got %v", result)
			}
			var opErr interface {
				Service() string
				Operation() string
			}
			if !errors.As(err, &opErr) {
				t.Fatalf("expect *types.InvalidGreeting operation error, got %T", err)
			}
			if e, a := client.ServiceID(), opErr.Service(); e != a {
				t.Errorf("expect %v operation service name, got %v", e, a)
			}
			if e, a := "GreetingWithErrors", opErr.Operation(); e != a {
				t.Errorf("expect %v operation service name, got %v", e, a)
			}
			var actualErr *types.InvalidGreeting
			if !errors.As(err, &actualErr) {
				t.Fatalf("expect *types.InvalidGreeting result error, got %T", err)
			}
			if diff := cmp.Diff(c.ExpectError, actualErr, cmpopts.IgnoreUnexported(middleware.Metadata{})); len(diff) != 0 {
				t.Errorf("expect c.ExpectError value match:\n%s", diff)
			}
		})
	}
}

func TestClient_GreetingWithErrors_ComplexError_awsRestxmlDeserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectError   *types.ComplexError
	}{
		"ComplexError": {
			StatusCode: 400,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
				"X-Header":     []string{"Header"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<ErrorResponse>
			   <Error>
			      <Type>Sender</Type>
			      <Code>ComplexError</Code>
			      <Message>Hi</Message>
			      <TopLevel>Top level</TopLevel>
			      <Nested>
			          <Foo>bar</Foo>
			      </Nested>
			   </Error>
			   <RequestId>foo-id</RequestId>
			</ErrorResponse>
			`),
			ExpectError: &types.ComplexError{
				Header:   ptr.String("Header"),
				TopLevel: ptr.String("Top level"),
				Nested: &types.ComplexNestedErrorData{
					Foo: ptr.String("bar"),
				},
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				for k, vs := range c.Header {
					for _, v := range vs {
						w.Header().Add(k, v)
					}
				}
				if len(c.BodyMediaType) != 0 && len(w.Header().Values("Content-Type")) == 0 {
					w.Header().Set("Content-Type", c.BodyMediaType)
				}
				if len(c.Body) != 0 {
					w.Header().Set("Content-Length", strconv.Itoa(len(c.Body)))
				}
				w.WriteHeader(c.StatusCode)
				if len(c.Body) != 0 {
					if _, err := io.Copy(w, bytes.NewReader(c.Body)); err != nil {
						t.Errorf("failed to write response body, %v", err)
					}
				}
			}))
			defer server.Close()
			client := New(Options{
				APIOptions: []APIOptionFunc{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: aws.EndpointResolverFunc(func(service, region string) (e aws.Endpoint, err error) {
					e.URL = server.URL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient:               aws.NewBuildableHTTPClient(),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params GreetingWithErrorsInput
			result, err := client.GreetingWithErrors(context.Background(), &params)
			if err == nil {
				t.Fatalf("expect not nil err")
			}
			if result != nil {
				t.Fatalf("expect nil result, got %v", result)
			}
			var opErr interface {
				Service() string
				Operation() string
			}
			if !errors.As(err, &opErr) {
				t.Fatalf("expect *types.ComplexError operation error, got %T", err)
			}
			if e, a := client.ServiceID(), opErr.Service(); e != a {
				t.Errorf("expect %v operation service name, got %v", e, a)
			}
			if e, a := "GreetingWithErrors", opErr.Operation(); e != a {
				t.Errorf("expect %v operation service name, got %v", e, a)
			}
			var actualErr *types.ComplexError
			if !errors.As(err, &actualErr) {
				t.Fatalf("expect *types.ComplexError result error, got %T", err)
			}
			if diff := cmp.Diff(c.ExpectError, actualErr, cmpopts.IgnoreUnexported(middleware.Metadata{})); len(diff) != 0 {
				t.Errorf("expect c.ExpectError value match:\n%s", diff)
			}
		})
	}
}
