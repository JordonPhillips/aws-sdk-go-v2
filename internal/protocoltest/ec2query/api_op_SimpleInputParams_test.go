// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2query

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/ec2query/types"
	"github.com/awslabs/smithy-go/middleware"
	"github.com/awslabs/smithy-go/ptr"
	smithyrand "github.com/awslabs/smithy-go/rand"
	smithytesting "github.com/awslabs/smithy-go/testing"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestClient_SimpleInputParams_awsEc2querySerialize(t *testing.T) {
	cases := map[string]struct {
		Params        *SimpleInputParamsInput
		ExpectMethod  string
		ExpectURIPath string
		ExpectQuery   []smithytesting.QueryItem
		RequireQuery  []string
		ForbidQuery   []string
		ExpectHeader  http.Header
		RequireHeader []string
		ForbidHeader  []string
		BodyMediaType string
		BodyAssert    func(io.Reader) error
	}{
		// Serializes strings
		"Ec2SimpleInputParamsStrings": {
			Params: &SimpleInputParamsInput{
				Foo: ptr.String("val1"),
				Bar: ptr.String("val2"),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=SimpleInputParams
			&Version=2020-01-08
			&Foo=val1
			&Bar=val2`))
			},
		},
		// Serializes booleans that are true
		"Ec2SimpleInputParamsStringAndBooleanTrue": {
			Params: &SimpleInputParamsInput{
				Foo: ptr.String("val1"),
				Baz: ptr.Bool(true),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=SimpleInputParams
			&Version=2020-01-08
			&Foo=val1
			&Baz=true`))
			},
		},
		// Serializes booleans that are false
		"Ec2SimpleInputParamsStringsAndBooleanFalse": {
			Params: &SimpleInputParamsInput{
				Baz: ptr.Bool(false),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=SimpleInputParams
			&Version=2020-01-08
			&Baz=false`))
			},
		},
		// Serializes integers
		"Ec2SimpleInputParamsInteger": {
			Params: &SimpleInputParamsInput{
				Bam: ptr.Int32(10),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=SimpleInputParams
			&Version=2020-01-08
			&Bam=10`))
			},
		},
		// Serializes floats
		"Ec2SimpleInputParamsFloat": {
			Params: &SimpleInputParamsInput{
				Boo: ptr.Float64(10.8),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=SimpleInputParams
			&Version=2020-01-08
			&Boo=10.8`))
			},
		},
		// Blobs are base64 encoded in the query string
		"Ec2SimpleInputParamsBlob": {
			Params: &SimpleInputParamsInput{
				Qux: []byte("value"),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=SimpleInputParams
			&Version=2020-01-08
			&Qux=dmFsdWU%3D`))
			},
		},
		// Serializes enums in the query string
		"Ec2Enums": {
			Params: &SimpleInputParamsInput{
				FooEnum: types.FooEnum("Foo"),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=SimpleInputParams
			&Version=2020-01-08
			&FooEnum=Foo`))
			},
		},
		// Serializes query using ec2QueryName trait.
		"Ec2Query": {
			Params: &SimpleInputParamsInput{
				HasQueryName: ptr.String("Hi"),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=SimpleInputParams
			&Version=2020-01-08
			&A=Hi`))
			},
		},
		// ec2QueryName trait is preferred over xmlName.
		"Ec2QueryIsPreferred": {
			Params: &SimpleInputParamsInput{
				HasQueryAndXmlName: ptr.String("Hi"),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=SimpleInputParams
			&Version=2020-01-08
			&B=Hi`))
			},
		},
		// xmlName is used with the ec2 protocol, but the first character is uppercased
		"Ec2XmlNameIsUppercased": {
			Params: &SimpleInputParamsInput{
				UsesXmlName: ptr.String("Hi"),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=SimpleInputParams
			&Version=2020-01-08
			&C=Hi`))
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			var actualReq *http.Request
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				actualReq = r.Clone(r.Context())
				if len(actualReq.URL.RawPath) == 0 {
					actualReq.URL.RawPath = actualReq.URL.Path
				}
				if v := actualReq.ContentLength; v != 0 {
					actualReq.Header.Set("Content-Length", strconv.FormatInt(v, 10))
				}
				var buf bytes.Buffer
				if _, err := io.Copy(&buf, r.Body); err != nil {
					t.Errorf("failed to read request body, %v", err)
				}
				actualReq.Body = ioutil.NopCloser(&buf)

				w.WriteHeader(200)
			}))
			defer server.Close()
			client := New(Options{
				APIOptions: []APIOptionFunc{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options ResolverOptions) (e aws.Endpoint, err error) {
					e.URL = server.URL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient:               aws.NewBuildableHTTPClient(),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			result, err := client.SimpleInputParams(context.Background(), c.Params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if e, a := c.ExpectMethod, actualReq.Method; e != a {
				t.Errorf("expect %v method, got %v", e, a)
			}
			if e, a := c.ExpectURIPath, actualReq.URL.RawPath; e != a {
				t.Errorf("expect %v path, got %v", e, a)
			}
			queryItems := smithytesting.ParseRawQuery(actualReq.URL.RawQuery)
			smithytesting.AssertHasQuery(t, c.ExpectQuery, queryItems)
			smithytesting.AssertHasQueryKeys(t, c.RequireQuery, queryItems)
			smithytesting.AssertNotHaveQueryKeys(t, c.ForbidQuery, queryItems)
			smithytesting.AssertHasHeader(t, c.ExpectHeader, actualReq.Header)
			smithytesting.AssertHasHeaderKeys(t, c.RequireHeader, actualReq.Header)
			smithytesting.AssertNotHaveHeaderKeys(t, c.ForbidHeader, actualReq.Header)
			if actualReq.Body != nil {
				defer actualReq.Body.Close()
			}
			if c.BodyAssert != nil {
				if err := c.BodyAssert(actualReq.Body); err != nil {
					t.Errorf("expect body equal, got %v", err)
				}
			}
		})
	}
}
