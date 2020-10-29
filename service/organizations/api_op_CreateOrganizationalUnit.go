// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an organizational unit (OU) within a root or parent OU. An OU is a
// container for accounts that enables you to organize your accounts to apply
// policies according to your business requirements. The number of levels deep that
// you can nest OUs is dependent upon the policy types enabled for that root. For
// service control policies, the limit is five. For more information about OUs, see
// Managing Organizational Units
// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_ous.html)
// in the AWS Organizations User Guide. If the request includes tags, then the
// requester must have the organizations:TagResource permission. This operation can
// be called only from the organization's management account.
func (c *Client) CreateOrganizationalUnit(ctx context.Context, params *CreateOrganizationalUnitInput, optFns ...func(*Options)) (*CreateOrganizationalUnitOutput, error) {
	if params == nil {
		params = &CreateOrganizationalUnitInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateOrganizationalUnit", params, optFns, addOperationCreateOrganizationalUnitMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateOrganizationalUnitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateOrganizationalUnitInput struct {

	// The friendly name to assign to the new OU.
	//
	// This member is required.
	Name *string

	// The unique identifier (ID) of the parent root or OU that you want to create the
	// new OU in. The regex pattern (http://wikipedia.org/wiki/regex) for a parent ID
	// string requires one of the following:
	//
	// * Root - A string that begins with "r-"
	// followed by from 4 to 32 lowercase letters or digits.
	//
	// * Organizational unit
	// (OU) - A string that begins with "ou-" followed by from 4 to 32 lowercase
	// letters or digits (the ID of the root that the OU is in). This string is
	// followed by a second "-" dash and from 8 to 32 additional lowercase letters or
	// digits.
	//
	// This member is required.
	ParentId *string

	// A list of tags that you want to attach to the newly created OU. For each tag in
	// the list, you must specify both a tag key and a value. You can set the value to
	// an empty string, but you can't set it to null. For more information about
	// tagging, see Tagging AWS Organizations resources
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_tagging.html)
	// in the AWS Organizations User Guide. If any one of the tags is invalid or if you
	// exceed the allowed number of tags for an OU, then the entire request fails and
	// the OU is not created.
	Tags []*types.Tag
}

type CreateOrganizationalUnitOutput struct {

	// A structure that contains details about the newly created OU.
	OrganizationalUnit *types.OrganizationalUnit

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateOrganizationalUnitMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateOrganizationalUnit{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateOrganizationalUnit{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpCreateOrganizationalUnitValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateOrganizationalUnit(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateOrganizationalUnit(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "CreateOrganizationalUnit",
	}
}
