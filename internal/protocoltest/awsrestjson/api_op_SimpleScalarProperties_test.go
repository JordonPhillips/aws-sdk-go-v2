// Code generated by smithy-go-codegen DO NOT EDIT.
package awsrestjson

import (
	"bytes"
	"context"
	"github.com/awslabs/smithy-go/middleware"
	"github.com/awslabs/smithy-go/ptr"
	smithytesting "github.com/awslabs/smithy-go/testing"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestClient_SimpleScalarProperties_awsRestjson1Serialize(t *testing.T) {
	cases := map[string]struct {
		Params        *SimpleScalarPropertiesInput
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
		// Serializes simple scalar properties
		"RestJsonSimpleScalarProperties": {
			Params: &SimpleScalarPropertiesInput{
				Foo:               ptr.String("Foo"),
				StringValue:       ptr.String("string"),
				TrueBooleanValue:  ptr.Bool(true),
				FalseBooleanValue: ptr.Bool(false),
				ByteValue:         ptr.Int8(1),
				ShortValue:        ptr.Int16(2),
				IntegerValue:      ptr.Int32(3),
				LongValue:         ptr.Int64(4),
				FloatValue:        ptr.Float32(5.5),
				DoubleValue:       ptr.Float64(6.5),
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/SimpleScalarProperties",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
				"X-Foo":        []string{"Foo"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "stringValue": "string",
			    "trueBooleanValue": true,
			    "falseBooleanValue": false,
			    "byteValue": 1,
			    "shortValue": 2,
			    "integerValue": 3,
			    "longValue": 4,
			    "floatValue": 5.5,
			    "DoubleDribble": 6.5
			}`))
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			var actualReq *http.Request
			client := New(Options{
				APIOptions: []APIOptionFunc{
					func(s *middleware.Stack) error {
						s.Build.Clear()
						s.Finalize.Clear()
						return nil
					},
				},
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					actualReq = r
					return &http.Response{
						StatusCode: 200,
						Header:     http.Header{},
						Body:       ioutil.NopCloser(strings.NewReader("")),
					}, nil
				}),
			})
			result, err := client.SimpleScalarProperties(context.Background(), c.Params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if e, a := c.ExpectMethod, actualReq.Method; e != a {
				t.Errorf("expect %v method, got %v", e, a)
			}
			if e, a := c.ExpectURIPath, actualReq.URL.Path; e != a {
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

func TestClient_SimpleScalarProperties_awsRestjson1Deserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *SimpleScalarPropertiesOutput
	}{
		// Serializes simple scalar properties
		"RestJsonSimpleScalarProperties": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
				"X-Foo":        []string{"Foo"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "stringValue": "string",
			    "trueBooleanValue": true,
			    "falseBooleanValue": false,
			    "byteValue": 1,
			    "shortValue": 2,
			    "integerValue": 3,
			    "longValue": 4,
			    "floatValue": 5.5,
			    "DoubleDribble": 6.5
			}`),
			ExpectResult: &SimpleScalarPropertiesOutput{
				Foo:               ptr.String("Foo"),
				StringValue:       ptr.String("string"),
				TrueBooleanValue:  ptr.Bool(true),
				FalseBooleanValue: ptr.Bool(false),
				ByteValue:         ptr.Int8(1),
				ShortValue:        ptr.Int16(2),
				IntegerValue:      ptr.Int32(3),
				LongValue:         ptr.Int64(4),
				FloatValue:        ptr.Float32(5.5),
				DoubleValue:       ptr.Float64(6.5),
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			client := New(Options{
				APIOptions: []APIOptionFunc{
					func(s *middleware.Stack) error {
						s.Build.Clear()
						s.Finalize.Clear()
						return nil
					},
				},
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: c.StatusCode,
						Header:     c.Header.Clone(),
						Body:       ioutil.NopCloser(bytes.NewReader(c.Body)),
					}, nil
				}),
			})
			var params SimpleScalarPropertiesInput
			result, err := client.SimpleScalarProperties(context.Background(), &params)
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
