// Code generated by smithy-go-codegen DO NOT EDIT.
package query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/query/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

func (c *Client) XmlNamespaces(ctx context.Context, params *XmlNamespacesInput, optFns ...func(*Options)) (*XmlNamespacesOutput, error) {
	stack := middleware.NewStack("XmlNamespaces", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	awsmiddleware.AddResolveServiceEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opXmlNamespaces(options.Region), middleware.Before)

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
			OperationName: "XmlNamespaces",
			Err:           err,
		}
	}
	out := result.(*XmlNamespacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type XmlNamespacesInput struct {
}

type XmlNamespacesOutput struct {
	Nested *types.XmlNamespaceNested

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func newServiceMetadataMiddleware_opXmlNamespaces(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Query Protocol",
		ServiceID:      "queryprotocol",
		EndpointPrefix: "queryprotocol",
		OperationName:  "XmlNamespaces",
	}
}
