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
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opXmlEnums(options.Region), middleware.Before)
	addawsRestxml_serdeOpXmlEnumsMiddlewares(stack)

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
	FooEnumSet  []types.FooEnum
	FooEnumMap  map[string]types.FooEnum
}

type XmlEnumsOutput struct {
	FooEnum1    types.FooEnum
	FooEnum2    types.FooEnum
	FooEnum3    types.FooEnum
	FooEnumList []types.FooEnum
	FooEnumSet  []types.FooEnum
	FooEnumMap  map[string]types.FooEnum

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpXmlEnumsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpXmlEnums{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpXmlEnums{}, middleware.After)
}

func newServiceMetadataMiddleware_opXmlEnums(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Rest Xml Protocol",
		ServiceID:      "restxmlprotocol",
		EndpointPrefix: "restxmlprotocol",
		OperationName:  "XmlEnums",
	}
}
