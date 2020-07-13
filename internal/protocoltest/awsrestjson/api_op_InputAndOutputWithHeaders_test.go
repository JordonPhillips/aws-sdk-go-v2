// Code generated by smithy-go-codegen DO NOT EDIT.
package awsrestjson

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/awsrestjson/types"
	"github.com/awslabs/smithy-go/middleware"
	"github.com/awslabs/smithy-go/ptr"
	smithyrand "github.com/awslabs/smithy-go/rand"
	smithytesting "github.com/awslabs/smithy-go/testing"
	smithytime "github.com/awslabs/smithy-go/time"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func TestClient_InputAndOutputWithHeaders_awsRestjson1Serialize(t *testing.T) {
	cases := map[string]struct {
		Params        *InputAndOutputWithHeadersInput
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
		// Tests requests with string header bindings
		"RestJsonInputAndOutputWithStringHeaders": {
			Params: &InputAndOutputWithHeadersInput{
				HeaderString: ptr.String("Hello"),
				HeaderStringList: []*string{
					ptr.String("a"),
					ptr.String("b"),
					ptr.String("c"),
				},
				HeaderStringSet: []*string{
					ptr.String("a"),
					ptr.String("b"),
					ptr.String("c"),
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/InputAndOutputWithHeaders",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"X-String":     []string{"Hello"},
				"X-StringList": []string{"a, b, c"},
				"X-StringSet":  []string{"a, b, c"},
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
		// Tests requests with numeric header bindings
		"RestJsonInputAndOutputWithNumericHeaders": {
			Params: &InputAndOutputWithHeadersInput{
				HeaderByte:    ptr.Int8(1),
				HeaderShort:   ptr.Int16(123),
				HeaderInteger: ptr.Int32(123),
				HeaderLong:    ptr.Int64(123),
				HeaderFloat:   ptr.Float32(1.1),
				HeaderDouble:  ptr.Float64(1.1),
				HeaderIntegerList: []*int32{
					ptr.Int32(1),
					ptr.Int32(2),
					ptr.Int32(3),
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/InputAndOutputWithHeaders",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"X-Byte":        []string{"1"},
				"X-Double":      []string{"1.1"},
				"X-Float":       []string{"1.1"},
				"X-Integer":     []string{"123"},
				"X-IntegerList": []string{"1, 2, 3"},
				"X-Long":        []string{"123"},
				"X-Short":       []string{"123"},
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
		// Tests requests with boolean header bindings
		"RestJsonInputAndOutputWithBooleanHeaders": {
			Params: &InputAndOutputWithHeadersInput{
				HeaderTrueBool:  ptr.Bool(true),
				HeaderFalseBool: ptr.Bool(false),
				HeaderBooleanList: []*bool{
					ptr.Bool(true),
					ptr.Bool(false),
					ptr.Bool(true),
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/InputAndOutputWithHeaders",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"X-Boolean1":    []string{"true"},
				"X-Boolean2":    []string{"false"},
				"X-BooleanList": []string{"true, false, true"},
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
		// Tests requests with timestamp header bindings
		"RestJsonInputAndOutputWithTimestampHeaders": {
			Params: &InputAndOutputWithHeadersInput{
				HeaderTimestampList: []*time.Time{
					ptr.Time(smithytime.ParseEpochSeconds(1576540098)),
					ptr.Time(smithytime.ParseEpochSeconds(1576540098)),
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/InputAndOutputWithHeaders",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"X-TimestampList": []string{"Mon, 16 Dec 2019 23:48:18 GMT, Mon, 16 Dec 2019 23:48:18 GMT"},
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
		// Tests requests with enum header bindings
		"RestJsonInputAndOutputWithEnumHeaders": {
			Params: &InputAndOutputWithHeadersInput{
				HeaderEnum: types.FooEnum("Foo"),
				HeaderEnumList: []types.FooEnum{
					types.FooEnum("Foo"),
					types.FooEnum("Bar"),
					types.FooEnum("Baz"),
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/InputAndOutputWithHeaders",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"X-Enum":     []string{"Foo"},
				"X-EnumList": []string{"Foo, Bar, Baz"},
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
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
						s.Build.Clear()
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
			result, err := client.InputAndOutputWithHeaders(context.Background(), c.Params)
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

func TestClient_InputAndOutputWithHeaders_awsRestjson1Deserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *InputAndOutputWithHeadersOutput
	}{
		// Tests responses with string header bindings
		"RestJsonInputAndOutputWithStringHeaders": {
			StatusCode: 200,
			Header: http.Header{
				"X-String":     []string{"Hello"},
				"X-StringList": []string{"a, b, c"},
				"X-StringSet":  []string{"a, b, c"},
			},
			Body: []byte(``),
			ExpectResult: &InputAndOutputWithHeadersOutput{
				HeaderString: ptr.String("Hello"),
				HeaderStringList: []*string{
					ptr.String("a"),
					ptr.String("b"),
					ptr.String("c"),
				},
				HeaderStringSet: []*string{
					ptr.String("a"),
					ptr.String("b"),
					ptr.String("c"),
				},
			},
		},
		// Tests responses with numeric header bindings
		"RestJsonInputAndOutputWithNumericHeaders": {
			StatusCode: 200,
			Header: http.Header{
				"X-Byte":        []string{"1"},
				"X-Double":      []string{"1.1"},
				"X-Float":       []string{"1.1"},
				"X-Integer":     []string{"123"},
				"X-IntegerList": []string{"1, 2, 3"},
				"X-Long":        []string{"123"},
				"X-Short":       []string{"123"},
			},
			Body: []byte(``),
			ExpectResult: &InputAndOutputWithHeadersOutput{
				HeaderByte:    ptr.Int8(1),
				HeaderShort:   ptr.Int16(123),
				HeaderInteger: ptr.Int32(123),
				HeaderLong:    ptr.Int64(123),
				HeaderFloat:   ptr.Float32(1.1),
				HeaderDouble:  ptr.Float64(1.1),
				HeaderIntegerList: []*int32{
					ptr.Int32(1),
					ptr.Int32(2),
					ptr.Int32(3),
				},
			},
		},
		// Tests responses with boolean header bindings
		"RestJsonInputAndOutputWithBooleanHeaders": {
			StatusCode: 200,
			Header: http.Header{
				"X-Boolean1":    []string{"true"},
				"X-Boolean2":    []string{"false"},
				"X-BooleanList": []string{"true, false, true"},
			},
			Body: []byte(``),
			ExpectResult: &InputAndOutputWithHeadersOutput{
				HeaderTrueBool:  ptr.Bool(true),
				HeaderFalseBool: ptr.Bool(false),
				HeaderBooleanList: []*bool{
					ptr.Bool(true),
					ptr.Bool(false),
					ptr.Bool(true),
				},
			},
		},
		// Tests responses with timestamp header bindings
		"RestJsonInputAndOutputWithTimestampHeaders": {
			StatusCode: 200,
			Header: http.Header{
				"X-TimestampList": []string{"Mon, 16 Dec 2019 23:48:18 GMT, Mon, 16 Dec 2019 23:48:18 GMT"},
			},
			Body: []byte(``),
			ExpectResult: &InputAndOutputWithHeadersOutput{
				HeaderTimestampList: []*time.Time{
					ptr.Time(smithytime.ParseEpochSeconds(1576540098)),
					ptr.Time(smithytime.ParseEpochSeconds(1576540098)),
				},
			},
		},
		// Tests responses with enum header bindings
		"RestJsonInputAndOutputWithEnumHeaders": {
			StatusCode: 200,
			Header: http.Header{
				"X-Enum":     []string{"Foo"},
				"X-EnumList": []string{"Foo, Bar, Baz"},
			},
			Body: []byte(``),
			ExpectResult: &InputAndOutputWithHeadersOutput{
				HeaderEnum: types.FooEnum("Foo"),
				HeaderEnumList: []types.FooEnum{
					types.FooEnum("Foo"),
					types.FooEnum("Bar"),
					types.FooEnum("Baz"),
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
						s.Build.Clear()
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
			var params InputAndOutputWithHeadersInput
			result, err := client.InputAndOutputWithHeaders(context.Background(), &params)
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
