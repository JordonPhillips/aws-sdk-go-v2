// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakerruntime

import (
	"bytes"
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsRestjson1_serializeOpInvokeEndpoint struct {
}

func (*awsRestjson1_serializeOpInvokeEndpoint) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpInvokeEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*InvokeEndpointInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/endpoints/{EndpointName}/invocations")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsInvokeEndpointInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if input.Body != nil {
		if !restEncoder.HasHeader("Content-Type") {
			restEncoder.SetHeader("Content-Type").String("application/octet-stream")
		}

		payload := bytes.NewReader(input.Body)
		if request, err = request.SetStream(payload); err != nil {
			return out, metadata, &smithy.SerializationError{Err: err}
		}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsInvokeEndpointInput(v *InvokeEndpointInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.Accept != nil && len(*v.Accept) > 0 {
		locationName := "Accept"
		encoder.SetHeader(locationName).String(*v.Accept)
	}

	if v.ContentType != nil && len(*v.ContentType) > 0 {
		locationName := "Content-Type"
		encoder.SetHeader(locationName).String(*v.ContentType)
	}

	if v.CustomAttributes != nil && len(*v.CustomAttributes) > 0 {
		locationName := "X-Amzn-Sagemaker-Custom-Attributes"
		encoder.SetHeader(locationName).String(*v.CustomAttributes)
	}

	if v.EndpointName == nil || len(*v.EndpointName) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member EndpointName must not be empty")}
	}
	if v.EndpointName != nil {
		if err := encoder.SetURI("EndpointName").String(*v.EndpointName); err != nil {
			return err
		}
	}

	if v.InferenceId != nil && len(*v.InferenceId) > 0 {
		locationName := "X-Amzn-Sagemaker-Inference-Id"
		encoder.SetHeader(locationName).String(*v.InferenceId)
	}

	if v.TargetModel != nil && len(*v.TargetModel) > 0 {
		locationName := "X-Amzn-Sagemaker-Target-Model"
		encoder.SetHeader(locationName).String(*v.TargetModel)
	}

	if v.TargetVariant != nil && len(*v.TargetVariant) > 0 {
		locationName := "X-Amzn-Sagemaker-Target-Variant"
		encoder.SetHeader(locationName).String(*v.TargetVariant)
	}

	return nil
}
