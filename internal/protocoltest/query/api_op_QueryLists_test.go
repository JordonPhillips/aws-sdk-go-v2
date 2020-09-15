// Code generated by smithy-go-codegen DO NOT EDIT.

package query

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/query/types"
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

func TestClient_QueryLists_awsAwsquerySerialize(t *testing.T) {
	cases := map[string]struct {
		Params        *QueryListsInput
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
		// Serializes query lists
		"QueryLists": {
			Params: &QueryListsInput{
				ListArg: []*string{
					ptr.String("foo"),
					ptr.String("bar"),
					ptr.String("baz"),
				},
				ComplexListArg: []*types.GreetingStruct{
					{
						Hi: ptr.String("hello"),
					},
					{
						Hi: ptr.String("hola"),
					},
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=QueryLists
			&Version=2020-01-08
			&ListArg.member.1=foo
			&ListArg.member.2=bar
			&ListArg.member.3=baz
			&ComplexListArg.member.1.hi=hello
			&ComplexListArg.member.2.hi=hola`))
			},
		},
		// Does not serialize empty query lists
		"EmptyQueryLists": {
			Params: &QueryListsInput{
				ListArg: []*string{},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=QueryLists
			&Version=2020-01-08`))
			},
		},
		// Flattens query lists by repeating the member name and removing the member
		// element
		"FlattenedQueryLists": {
			Params: &QueryListsInput{
				FlattenedListArg: []*string{
					ptr.String("A"),
					ptr.String("B"),
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=QueryLists
			&Version=2020-01-08
			&FlattenedListArg.1=A
			&FlattenedListArg.2=B`))
			},
		},
		// Changes the member of lists using xmlName trait
		"QueryListArgWithXmlNameMember": {
			Params: &QueryListsInput{
				ListArgWithXmlNameMember: []*string{
					ptr.String("A"),
					ptr.String("B"),
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=QueryLists
			&Version=2020-01-08
			&ListArgWithXmlNameMember.item.1=A
			&ListArgWithXmlNameMember.item.2=B`))
			},
		},
		// Changes the name of flattened lists using xmlName trait on the structure member
		"QueryFlattenedListArgWithXmlName": {
			Params: &QueryListsInput{
				FlattenedListArgWithXmlName: []*string{
					ptr.String("A"),
					ptr.String("B"),
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=QueryLists
			&Version=2020-01-08
			&Hi.1=A
			&Hi.2=B`))
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
			url := server.URL
			client := New(Options{
				APIOptions: []APIOptionFunc{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options ResolverOptions) (e aws.Endpoint, err error) {
					e.URL = url
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient:               aws.NewBuildableHTTPClient(),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			result, err := client.QueryLists(context.Background(), c.Params)
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
			if c.BodyAssert != nil {
				if err := c.BodyAssert(actualReq.Body); err != nil {
					t.Errorf("expect body equal, got %v", err)
				}
			}
		})
	}
}
