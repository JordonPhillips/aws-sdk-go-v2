// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation configures notifications that will be sent when specific events
// happen to a vault. By default, you don't get any notifications.  <p>To configure
// vault notifications, send a PUT request to the
// <code>notification-configuration</code> subresource of the vault. The request
// should include a JSON document that provides an Amazon SNS topic and specific
// events for which you want Amazon S3 Glacier to send notifications to the
// topic.</p> <p>Amazon SNS topics must grant permission to the vault to be allowed
// to publish notifications to the topic. You can configure a vault to publish a
// notification for the following vault events:</p> <ul> <li> <p>
// <b>ArchiveRetrievalCompleted</b> This event occurs when a job that was initiated
// for an archive retrieval is completed (<a>InitiateJob</a>). The status of the
// completed job can be "Succeeded" or "Failed". The notification sent to the SNS
// topic is the same output as returned from <a>DescribeJob</a>. </p> </li> <li>
// <p> <b>InventoryRetrievalCompleted</b> This event occurs when a job that was
// initiated for an inventory retrieval is completed (<a>InitiateJob</a>). The
// status of the completed job can be "Succeeded" or "Failed". The notification
// sent to the SNS topic is the same output as returned from <a>DescribeJob</a>.
// </p> </li> </ul> <p>An AWS account has full permission to perform all operations
// (actions). However, AWS Identity and Access Management (IAM) users don't have
// any permissions by default. You must grant them explicit permission to perform
// specific actions. For more information, see <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html">Access
// Control Using AWS Identity and Access Management (IAM)</a>.</p> <p>For
// conceptual information and underlying REST API, see <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/configuring-notifications.html">Configuring
// Vault Notifications in Amazon S3 Glacier</a> and <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vault-notifications-put.html">Set
// Vault Notification Configuration </a> in the <i>Amazon Glacier Developer
// Guide</i>. </p>
func (c *Client) SetVaultNotifications(ctx context.Context, params *SetVaultNotificationsInput, optFns ...func(*Options)) (*SetVaultNotificationsOutput, error) {
	stack := middleware.NewStack("SetVaultNotifications", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpSetVaultNotificationsMiddlewares(stack)
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
	addOpSetVaultNotificationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetVaultNotifications(options.Region), middleware.Before)
	glaciercust.AddTreeHashMiddleware(stack)

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
			OperationName: "SetVaultNotifications",
			Err:           err,
		}
	}
	out := result.(*SetVaultNotificationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides options to configure notifications that will be sent when specific
// events happen to a vault.
type SetVaultNotificationsInput struct {
	// The name of the vault.
	VaultName *string
	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen), in
	// which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	AccountId *string
	// Provides options for specifying notification configuration.
	VaultNotificationConfig *types.VaultNotificationConfig
}

type SetVaultNotificationsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpSetVaultNotificationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpSetVaultNotifications{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpSetVaultNotifications{}, middleware.After)
}

func newServiceMetadataMiddleware_opSetVaultNotifications(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "SetVaultNotifications",
	}
}
