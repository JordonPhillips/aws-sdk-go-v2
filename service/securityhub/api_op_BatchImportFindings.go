// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Imports security findings generated from an integrated third-party product into
// Security Hub. This action is requested by the integrated product to import its
// findings into Security Hub. The maximum allowed size for a finding is 240 Kb. An
// error is returned for any finding larger than 240 Kb. After a finding is
// created, BatchImportFindings cannot be used to update the following finding
// fields and objects, which Security Hub customers use to manage their
// investigation workflow.
//
// * Note
//
// * UserDefinedFields
//
// * VerificationState
//
// *
// Workflow
//
// BatchImportFindings can be used to update the following finding fields
// and objects only if they have not been updated using BatchUpdateFindings. After
// they are updated using BatchUpdateFindings, these fields cannot be updated using
// BatchImportFindings.
//
// * Confidence
//
// * Criticality
//
// * RelatedFindings
//
// *
// Severity
//
// * Types
func (c *Client) BatchImportFindings(ctx context.Context, params *BatchImportFindingsInput, optFns ...func(*Options)) (*BatchImportFindingsOutput, error) {
	if params == nil {
		params = &BatchImportFindingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchImportFindings", params, optFns, addOperationBatchImportFindingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchImportFindingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchImportFindingsInput struct {

	// A list of findings to import. To successfully import a finding, it must follow
	// the AWS Security Finding Format
	// (https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-findings-format.html).
	// Maximum of 100 findings per request.
	//
	// This member is required.
	Findings []types.AwsSecurityFinding
}

type BatchImportFindingsOutput struct {

	// The number of findings that failed to import.
	//
	// This member is required.
	FailedCount int32

	// The number of findings that were successfully imported.
	//
	// This member is required.
	SuccessCount int32

	// The list of findings that failed to import.
	FailedFindings []types.ImportFindingsError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchImportFindingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchImportFindings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchImportFindings{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpBatchImportFindingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchImportFindings(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opBatchImportFindings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "BatchImportFindings",
	}
}
