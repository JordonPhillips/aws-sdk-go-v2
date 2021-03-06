// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/jsonrpc/types"
)

func ExampleMyUnion_outputUsage() {
	var union types.MyUnion
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.MyUnionMemberBlobValue:
		_ = v.Value // Value is []byte

	case *types.MyUnionMemberBooleanValue:
		_ = v.Value // Value is bool

	case *types.MyUnionMemberEnumValue:
		_ = v.Value // Value is FooEnum

	case *types.MyUnionMemberListValue:
		_ = v.Value // Value is StringList

	case *types.MyUnionMemberMapValue:
		_ = v.Value // Value is StringMap

	case *types.MyUnionMemberNumberValue:
		_ = v.Value // Value is int32

	case *types.MyUnionMemberStringValue:
		_ = v.Value // Value is string

	case *types.MyUnionMemberStructureValue:
		_ = v.Value // Value is GreetingStruct

	case *types.MyUnionMemberTimestampValue:
		_ = v.Value // Value is Time

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}
