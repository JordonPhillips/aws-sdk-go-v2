// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// CreateSMBFileShareInput
type CreateSMBFileShareInput struct {
	_ struct{} `type:"structure"`

	// A list of users or groups in the Active Directory that will be granted administrator
	// privileges on the file share. These users can do all file operations as the
	// super-user. Acceptable formats include: DOMAIN\User1, user1, @group1, and
	// @DOMAIN\group1.
	//
	// Use this option very carefully, because any user in this list can do anything
	// they like on the file share, regardless of file permissions.
	AdminUserList []string `type:"list"`

	// The Amazon Resource Name (ARN) of the storage used for the audit logs.
	AuditDestinationARN *string `type:"string"`

	// The authentication method that users use to access the file share. The default
	// is ActiveDirectory.
	//
	// Valid Values: ActiveDirectory | GuestAccess
	Authentication *string `min:"5" type:"string"`

	// Refresh cache information.
	CacheAttributes *CacheAttributes `type:"structure"`

	// The case of an object name in an Amazon S3 bucket. For ClientSpecified, the
	// client determines the case sensitivity. For CaseSensitive, the gateway determines
	// the case sensitivity. The default value is ClientSpecified.
	CaseSensitivity CaseSensitivity `type:"string" enum:"true"`

	// A unique string value that you supply that is used by file gateway to ensure
	// idempotent file share creation.
	//
	// ClientToken is a required field
	ClientToken *string `min:"5" type:"string" required:"true"`

	// The default storage class for objects put into an Amazon S3 bucket by the
	// file gateway. The default value is S3_INTELLIGENT_TIERING. Optional.
	//
	// Valid Values: S3_STANDARD | S3_INTELLIGENT_TIERING | S3_STANDARD_IA | S3_ONEZONE_IA
	DefaultStorageClass *string `min:"5" type:"string"`

	// The name of the file share. Optional.
	//
	// FileShareName must be set if an S3 prefix name is set in LocationARN.
	FileShareName *string `min:"1" type:"string"`

	// The ARN of the file gateway on which you want to create a file share.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`

	// A value that enables guessing of the MIME type for uploaded objects based
	// on file extensions. Set this value to true to enable MIME type guessing,
	// otherwise set to false. The default value is true.
	//
	// Valid Values: true | false
	GuessMIMETypeEnabled *bool `type:"boolean"`

	// A list of users or groups in the Active Directory that are not allowed to
	// access the file share. A group must be prefixed with the @ character. Acceptable
	// formats include: DOMAIN\User1, user1, @group1, and @DOMAIN\group1. Can only
	// be set if Authentication is set to ActiveDirectory.
	InvalidUserList []string `type:"list"`

	// Set to true to use Amazon S3 server-side encryption with your own AWS KMS
	// key, or false to use a key managed by Amazon S3. Optional.
	//
	// Valid Values: true | false
	KMSEncrypted *bool `type:"boolean"`

	// The Amazon Resource Name (ARN) of a symmetric customer master key (CMK) used
	// for Amazon S3 server-side encryption. Storage Gateway does not support asymmetric
	// CMKs. This value can only be set when KMSEncrypted is true. Optional.
	KMSKey *string `min:"7" type:"string"`

	// The ARN of the backend storage used for storing file data. A prefix name
	// can be added to the S3 bucket name. It must end with a "/".
	//
	// LocationARN is a required field
	LocationARN *string `min:"16" type:"string" required:"true"`

	// A value that sets the access control list (ACL) permission for objects in
	// the S3 bucket that a file gateway puts objects into. The default value is
	// private.
	ObjectACL ObjectACL `type:"string" enum:"true"`

	// A value that sets the write status of a file share. Set this value to true
	// to set the write status to read-only, otherwise set to false.
	//
	// Valid Values: true | false
	ReadOnly *bool `type:"boolean"`

	// A value that sets who pays the cost of the request and the cost associated
	// with data download from the S3 bucket. If this value is set to true, the
	// requester pays the costs; otherwise, the S3 bucket owner pays. However, the
	// S3 bucket owner always pays the cost of storing data.
	//
	// RequesterPays is a configuration for the S3 bucket that backs the file share,
	// so make sure that the configuration on the file share is the same as the
	// S3 bucket configuration.
	//
	// Valid Values: true | false
	RequesterPays *bool `type:"boolean"`

	// The ARN of the AWS Identity and Access Management (IAM) role that a file
	// gateway assumes when it accesses the underlying storage.
	//
	// Role is a required field
	Role *string `min:"20" type:"string" required:"true"`

	// Set this value to true to enable access control list (ACL) on the SMB file
	// share. Set it to false to map file and directory permissions to the POSIX
	// permissions.
	//
	// For more information, see Using Microsoft Windows ACLs to control access
	// to an SMB file share (https://docs.aws.amazon.com/storagegateway/latest/userguide/smb-acl.html)
	// in the AWS Storage Gateway User Guide.
	//
	// Valid Values: true | false
	SMBACLEnabled *bool `type:"boolean"`

	// A list of up to 50 tags that can be assigned to the NFS file share. Each
	// tag is a key-value pair.
	//
	// Valid characters for key and value are letters, spaces, and numbers representable
	// in UTF-8 format, and the following special characters: + - = . _ : / @. The
	// maximum length of a tag's key is 128 characters, and the maximum length for
	// a tag's value is 256.
	Tags []Tag `type:"list"`

	// A list of users or groups in the Active Directory that are allowed to access
	// the file share. A group must be prefixed with the @ character. Acceptable
	// formats include: DOMAIN\User1, user1, @group1, and @DOMAIN\group1. Can only
	// be set if Authentication is set to ActiveDirectory.
	ValidUserList []string `type:"list"`
}

// String returns the string representation
func (s CreateSMBFileShareInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSMBFileShareInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSMBFileShareInput"}
	if s.Authentication != nil && len(*s.Authentication) < 5 {
		invalidParams.Add(aws.NewErrParamMinLen("Authentication", 5))
	}

	if s.ClientToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientToken"))
	}
	if s.ClientToken != nil && len(*s.ClientToken) < 5 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientToken", 5))
	}
	if s.DefaultStorageClass != nil && len(*s.DefaultStorageClass) < 5 {
		invalidParams.Add(aws.NewErrParamMinLen("DefaultStorageClass", 5))
	}
	if s.FileShareName != nil && len(*s.FileShareName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FileShareName", 1))
	}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}
	if s.KMSKey != nil && len(*s.KMSKey) < 7 {
		invalidParams.Add(aws.NewErrParamMinLen("KMSKey", 7))
	}

	if s.LocationARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("LocationARN"))
	}
	if s.LocationARN != nil && len(*s.LocationARN) < 16 {
		invalidParams.Add(aws.NewErrParamMinLen("LocationARN", 16))
	}

	if s.Role == nil {
		invalidParams.Add(aws.NewErrParamRequired("Role"))
	}
	if s.Role != nil && len(*s.Role) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("Role", 20))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// CreateSMBFileShareOutput
type CreateSMBFileShareOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the newly created file share.
	FileShareARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s CreateSMBFileShareOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateSMBFileShare = "CreateSMBFileShare"

// CreateSMBFileShareRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Creates a Server Message Block (SMB) file share on an existing file gateway.
// In Storage Gateway, a file share is a file system mount point backed by Amazon
// S3 cloud storage. Storage Gateway expose file shares using an SMB interface.
// This operation is only supported for file gateways.
//
// File gateways require AWS Security Token Service (AWS STS) to be activated
// to enable you to create a file share. Make sure that AWS STS is activated
// in the AWS Region you are creating your file gateway in. If AWS STS is not
// activated in this AWS Region, activate it. For information about how to activate
// AWS STS, see Activating and deactivating AWS STS in an AWS Region (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html)
// in the AWS Identity and Access Management User Guide.
//
// File gateways don't support creating hard or symbolic links on a file share.
//
//    // Example sending a request using CreateSMBFileShareRequest.
//    req := client.CreateSMBFileShareRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/CreateSMBFileShare
func (c *Client) CreateSMBFileShareRequest(input *CreateSMBFileShareInput) CreateSMBFileShareRequest {
	op := &aws.Operation{
		Name:       opCreateSMBFileShare,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateSMBFileShareInput{}
	}

	req := c.newRequest(op, input, &CreateSMBFileShareOutput{})

	return CreateSMBFileShareRequest{Request: req, Input: input, Copy: c.CreateSMBFileShareRequest}
}

// CreateSMBFileShareRequest is the request type for the
// CreateSMBFileShare API operation.
type CreateSMBFileShareRequest struct {
	*aws.Request
	Input *CreateSMBFileShareInput
	Copy  func(*CreateSMBFileShareInput) CreateSMBFileShareRequest
}

// Send marshals and sends the CreateSMBFileShare API request.
func (r CreateSMBFileShareRequest) Send(ctx context.Context) (*CreateSMBFileShareResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSMBFileShareResponse{
		CreateSMBFileShareOutput: r.Request.Data.(*CreateSMBFileShareOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSMBFileShareResponse is the response type for the
// CreateSMBFileShare API operation.
type CreateSMBFileShareResponse struct {
	*CreateSMBFileShareOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSMBFileShare request.
func (r *CreateSMBFileShareResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}