// Code generated by smithy-go-codegen DO NOT EDIT.
package restxml

import (
	"context"
	cryptorand "crypto/rand"
	"fmt"
	"github.com/awslabs/smithy-go/middleware"
	smithyrand "github.com/awslabs/smithy-go/rand"
	"io"
)

var randReader io.Reader = cryptorand.Reader

type idempotencyToken_initializeOpQueryIdempotencyTokenAutoFill struct {
}

func (*idempotencyToken_initializeOpQueryIdempotencyTokenAutoFill) ID() string {
	return "idempotencyToken_initializeOpQueryIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpQueryIdempotencyTokenAutoFill) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*QueryIdempotencyTokenAutoFillInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *QueryIdempotencyTokenAutoFillInput ")
	}
	if input.Token == nil {
		uuid := smithyrand.NewUUID(randReader)
		v, err := uuid.GetUUID()
		if err != nil {
			return out, metadata, err
		}
		input.Token = &v
	}
	return next.HandleInitialize(ctx, in)
}
