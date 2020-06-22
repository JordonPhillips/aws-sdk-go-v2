// Code generated by smithy-go-codegen DO NOT EDIT.
package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/restxml/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This example serializes enums as top level properties, in lists, sets, and maps.
func (c *Client) XmlEnums(ctx context.Context, params *XmlEnumsInput, optFns ...func(*Options)) (*XmlEnumsOutput, error) {
	stack := middleware.NewStack("XmlEnums", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	stack.Build.Add(awsmiddleware.RequestInvocationIDMiddleware{}, middleware.After)
	stack.Deserialize.Add(awsmiddleware.AttemptClockSkewMiddleware{}, middleware.After)
	stack.Finalize.Add(retry.NewAttemptMiddleware(options.Retryer, smithyhttp.RequestCloner), middleware.After)
	stack.Finalize.Add(retry.MetricsHeaderMiddleware{}, middleware.After)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "XmlEnums",
			Err:           err,
		}
	}
	out := result.(*XmlEnumsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type XmlEnumsInput struct {
	FooEnum1    types.FooEnum
	FooEnum2    types.FooEnum
	FooEnum3    types.FooEnum
	FooEnumList []types.FooEnum
	FooEnumMap  map[string]types.FooEnum
	FooEnumSet  []types.FooEnum
}

type XmlEnumsOutput struct {
	FooEnum1    types.FooEnum
	FooEnum2    types.FooEnum
	FooEnum3    types.FooEnum
	FooEnumList []types.FooEnum
	FooEnumMap  map[string]types.FooEnum
	FooEnumSet  []types.FooEnum

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}
