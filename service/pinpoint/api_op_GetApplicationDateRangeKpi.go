// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Retrieves (queries) pre-aggregated data for a standard metric that applies to an
// application.
func (c *Client) GetApplicationDateRangeKpi(ctx context.Context, params *GetApplicationDateRangeKpiInput, optFns ...func(*Options)) (*GetApplicationDateRangeKpiOutput, error) {
	stack := middleware.NewStack("GetApplicationDateRangeKpi", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetApplicationDateRangeKpiMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpGetApplicationDateRangeKpiValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetApplicationDateRangeKpi(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "GetApplicationDateRangeKpi",
			Err:           err,
		}
	}
	out := result.(*GetApplicationDateRangeKpiOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetApplicationDateRangeKpiInput struct {
	// The first date and time to retrieve data for, as part of an inclusive date range
	// that filters the query results. This value should be in extended ISO 8601 format
	// and use Coordinated Universal Time (UTC), for example: 2019-07-19T20:00:00Z for
	// 8:00 PM UTC July 19, 2019. This value should also be fewer than 90 days from the
	// current day.
	StartTime *time.Time
	// The last date and time to retrieve data for, as part of an inclusive date range
	// that filters the query results. This value should be in extended ISO 8601 format
	// and use Coordinated Universal Time (UTC), for example: 2019-07-26T20:00:00Z for
	// 8:00 PM UTC July 26, 2019.
	EndTime *time.Time
	// The maximum number of items to include in each page of a paginated response.
	// This parameter is not supported for application, campaign, and journey metrics.
	PageSize *string
	// The string that specifies which page of results to return in a paginated
	// response. This parameter is not supported for application, campaign, and journey
	// metrics.
	NextToken *string
	// The name of the metric, also referred to as a key performance indicator (KPI),
	// to retrieve data for. This value describes the associated metric and consists of
	// two or more terms, which are comprised of lowercase alphanumeric characters,
	// separated by a hyphen. Examples are email-open-rate and
	// successful-delivery-rate. For a list of valid values, see the Amazon Pinpoint
	// Developer Guide
	// (https://docs.aws.amazon.com/pinpoint/latest/developerguide/analytics-standard-metrics.html).
	KpiName *string
	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	ApplicationId *string
}

type GetApplicationDateRangeKpiOutput struct {
	// Provides the results of a query that retrieved the data for a standard metric
	// that applies to an application, and provides information about that query.
	ApplicationDateRangeKpiResponse *types.ApplicationDateRangeKpiResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetApplicationDateRangeKpiMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetApplicationDateRangeKpi{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetApplicationDateRangeKpi{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetApplicationDateRangeKpi(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "GetApplicationDateRangeKpi",
	}
}