// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/query"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	presignedurlcust "github.com/aws/aws-sdk-go-v2/service/internal/presigned-url"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"net/http"
)

// Creates a new Amazon Aurora DB cluster. You can use the
// ReplicationSourceIdentifier parameter to create the DB cluster as a read replica
// of another DB cluster or Amazon RDS MySQL DB instance. For cross-region
// replication where the DB cluster identified by ReplicationSourceIdentifier is
// encrypted, you must also specify the PreSignedUrl parameter. For more
// information on Amazon Aurora, see
//
// What Is Amazon Aurora?
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide. This action only applies to Aurora DB clusters.
func (c *Client) CreateDBCluster(ctx context.Context, params *CreateDBClusterInput, optFns ...func(*Options)) (*CreateDBClusterOutput, error) {
	if params == nil {
		params = &CreateDBClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDBCluster", params, optFns, addOperationCreateDBClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateDBClusterInput struct {

	// The DB cluster identifier. This parameter is stored as a lowercase string.
	// Constraints:
	//
	// * Must contain from 1 to 63 letters, numbers, or hyphens.
	//
	// * First
	// character must be a letter.
	//
	// * Can't end with a hyphen or contain two
	// consecutive hyphens.
	//
	// Example: my-cluster1
	//
	// This member is required.
	DBClusterIdentifier *string

	// The name of the database engine to be used for this DB cluster. Valid Values:
	// aurora (for MySQL 5.6-compatible Aurora), aurora-mysql (for MySQL 5.7-compatible
	// Aurora), and aurora-postgresql
	//
	// This member is required.
	Engine *string

	// A list of Availability Zones (AZs) where instances in the DB cluster can be
	// created. For information on AWS Regions and Availability Zones, see Choosing the
	// Regions and Availability Zones
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.RegionsAndAvailabilityZones.html)
	// in the Amazon Aurora User Guide.
	AvailabilityZones []*string

	// The target backtrack window, in seconds. To disable backtracking, set this value
	// to 0. Currently, Backtrack is only supported for Aurora MySQL DB clusters.
	// Default: 0 Constraints:
	//
	// * If specified, this value must be set to a number from
	// 0 to 259,200 (72 hours).
	BacktrackWindow *int64

	// The number of days for which automated backups are retained. Default: 1
	// Constraints:
	//
	// * Must be a value from 1 to 35
	BackupRetentionPeriod *int32

	// A value that indicates that the DB cluster should be associated with the
	// specified CharacterSet.
	CharacterSetName *string

	// A value that indicates whether to copy all tags from the DB cluster to snapshots
	// of the DB cluster. The default is not to copy them.
	CopyTagsToSnapshot *bool

	// The name of the DB cluster parameter group to associate with this DB cluster. If
	// you do not specify a value, then the default DB cluster parameter group for the
	// specified DB engine and version is used. Constraints:
	//
	// * If supplied, must match
	// the name of an existing DB cluster parameter group.
	DBClusterParameterGroupName *string

	// A DB subnet group to associate with this DB cluster. Constraints: Must match the
	// name of an existing DBSubnetGroup. Must not be default. Example: mySubnetgroup
	DBSubnetGroupName *string

	// The name for your database of up to 64 alphanumeric characters. If you do not
	// provide a name, Amazon RDS doesn't create a database in the DB cluster you are
	// creating.
	DatabaseName *string

	// A value that indicates whether the DB cluster has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection is disabled.
	DeletionProtection *bool

	// The Active Directory directory ID to create the DB cluster in. For Amazon Aurora
	// DB clusters, Amazon RDS can use Kerberos Authentication to authenticate users
	// that connect to the DB cluster. For more information, see Kerberos
	// Authentication
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/kerberos-authentication.html)
	// in the Amazon Aurora User Guide.
	Domain *string

	// Specify the name of the IAM role to be used when making API calls to the
	// Directory Service.
	DomainIAMRoleName *string

	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	// The values in the list depend on the DB engine being used. For more information,
	// see Publishing Database Logs to Amazon CloudWatch Logs
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon Aurora User Guide. Aurora MySQL Possible values are audit, error,
	// general, and slowquery. Aurora PostgreSQL Possible values are postgresql and
	// upgrade.
	EnableCloudwatchLogsExports []*string

	// A value that indicates whether to enable write operations to be forwarded from
	// this cluster to the primary cluster in an Aurora global database. The resulting
	// changes are replicated back to this cluster. This parameter only applies to DB
	// clusters that are secondary clusters in an Aurora global database. By default,
	// Aurora disallows write operations for secondary clusters.
	EnableGlobalWriteForwarding *bool

	// A value that indicates whether to enable the HTTP endpoint for an Aurora
	// Serverless DB cluster. By default, the HTTP endpoint is disabled. When enabled,
	// the HTTP endpoint provides a connectionless web service API for running SQL
	// queries on the Aurora Serverless DB cluster. You can also query your database
	// from inside the RDS console with the query editor. For more information, see
	// Using the Data API for Aurora Serverless
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/data-api.html) in
	// the Amazon Aurora User Guide.
	EnableHttpEndpoint *bool

	// A value that indicates whether to enable mapping of AWS Identity and Access
	// Management (IAM) accounts to database accounts. By default, mapping is disabled.
	// For more information, see  IAM Database Authentication
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.IAMDBAuth.html)
	// in the Amazon Aurora User Guide.
	EnableIAMDatabaseAuthentication *bool

	// The DB engine mode of the DB cluster, either provisionedserverless,
	// parallelquery, global, or multimaster. The parallelquery engine mode isn't
	// required for Aurora MySQL version 1.23 and higher 1.x versions, and version 2.09
	// and higher 2.x versions. The global engine mode isn't required for Aurora MySQL
	// version 1.22 and higher 1.x versions, and global engine mode isn't required for
	// any 2.x versions. The multimaster engine mode only applies for DB clusters
	// created with Aurora MySQL version 5.6.10a. For Aurora PostgreSQL, the global
	// engine mode isn't required, and both the parallelquery and the multimaster
	// engine modes currently aren't supported. Limitations and requirements apply to
	// some DB engine modes. For more information, see the following sections in the
	// Amazon Aurora User Guide:
	//
	// * Limitations of Aurora Serverless
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.html#aurora-serverless.limitations)
	//
	// *
	// Limitations of Parallel Query
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-mysql-parallel-query.html#aurora-mysql-parallel-query-limitations)
	//
	// *
	// Limitations of Aurora Global Databases
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html#aurora-global-database.limitations)
	//
	// *
	// Limitations of Multi-Master Clusters
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-multi-master.html#aurora-multi-master-limitations)
	EngineMode *string

	// The version number of the database engine to use. To list all of the available
	// engine versions for aurora (for MySQL 5.6-compatible Aurora), use the following
	// command: aws rds describe-db-engine-versions --engine aurora --query
	// "DBEngineVersions[].EngineVersion" To list all of the available engine versions
	// for aurora-mysql (for MySQL 5.7-compatible Aurora), use the following command:
	// aws rds describe-db-engine-versions --engine aurora-mysql --query
	// "DBEngineVersions[].EngineVersion" To list all of the available engine versions
	// for aurora-postgresql, use the following command: aws rds
	// describe-db-engine-versions --engine aurora-postgresql --query
	// "DBEngineVersions[].EngineVersion" Aurora MySQL Example: 5.6.10a,
	// 5.6.mysql_aurora.1.19.2, 5.7.12, 5.7.mysql_aurora.2.04.5 Aurora PostgreSQL
	// Example: 9.6.3, 10.7
	EngineVersion *string

	// The global cluster ID of an Aurora cluster that becomes the primary cluster in
	// the new global database cluster.
	GlobalClusterIdentifier *string

	// The AWS KMS key identifier for an encrypted DB cluster. The KMS key identifier
	// is the Amazon Resource Name (ARN) for the KMS encryption key. If you are
	// creating a DB cluster with the same AWS account that owns the KMS encryption key
	// used to encrypt the new DB cluster, then you can use the KMS key alias instead
	// of the ARN for the KMS encryption key. If an encryption key isn't specified in
	// KmsKeyId:
	//
	// * If ReplicationSourceIdentifier identifies an encrypted source, then
	// Amazon RDS will use the encryption key used to encrypt the source. Otherwise,
	// Amazon RDS will use your default encryption key.
	//
	// * If the StorageEncrypted
	// parameter is enabled and ReplicationSourceIdentifier isn't specified, then
	// Amazon RDS will use your default encryption key.
	//
	// AWS KMS creates the default
	// encryption key for your AWS account. Your AWS account has a different default
	// encryption key for each AWS Region. If you create a read replica of an encrypted
	// DB cluster in another AWS Region, you must set KmsKeyId to a KMS key ID that is
	// valid in the destination AWS Region. This key is used to encrypt the read
	// replica in that AWS Region.
	KmsKeyId *string

	// The password for the master database user. This password can contain any
	// printable ASCII character except "/", """, or "@". Constraints: Must contain
	// from 8 to 41 characters.
	MasterUserPassword *string

	// The name of the master user for the DB cluster. Constraints:
	//
	// * Must be 1 to 16
	// letters or numbers.
	//
	// * First character must be a letter.
	//
	// * Can't be a reserved
	// word for the chosen database engine.
	MasterUsername *string

	// A value that indicates that the DB cluster should be associated with the
	// specified option group. Permanent options can't be removed from an option group.
	// The option group can't be removed from a DB cluster once it is associated with a
	// DB cluster.
	OptionGroupName *string

	// The port number on which the instances in the DB cluster accept connections.
	// Default: 3306 if engine is set as aurora or 5432 if set to aurora-postgresql.
	Port *int32

	// A URL that contains a Signature Version 4 signed request for the CreateDBCluster
	// action to be called in the source AWS Region where the DB cluster is replicated
	// from. You only need to specify PreSignedUrl when you are performing cross-region
	// replication from an encrypted DB cluster. The pre-signed URL must be a valid
	// request for the CreateDBCluster API action that can be executed in the source
	// AWS Region that contains the encrypted DB cluster to be copied. The pre-signed
	// URL request must contain the following parameter values:
	//
	// * KmsKeyId - The AWS
	// KMS key identifier for the key to use to encrypt the copy of the DB cluster in
	// the destination AWS Region. This should refer to the same KMS key for both the
	// CreateDBCluster action that is called in the destination AWS Region, and the
	// action contained in the pre-signed URL.
	//
	// * DestinationRegion - The name of the
	// AWS Region that Aurora read replica will be created in.
	//
	// *
	// ReplicationSourceIdentifier - The DB cluster identifier for the encrypted DB
	// cluster to be copied. This identifier must be in the Amazon Resource Name (ARN)
	// format for the source AWS Region. For example, if you are copying an encrypted
	// DB cluster from the us-west-2 AWS Region, then your ReplicationSourceIdentifier
	// would look like Example:
	// arn:aws:rds:us-west-2:123456789012:cluster:aurora-cluster1.
	//
	// To learn how to
	// generate a Signature Version 4 signed request, see  Authenticating Requests:
	// Using Query Parameters (AWS Signature Version 4)
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html)
	// and  Signature Version 4 Signing Process
	// (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html). If you
	// are using an AWS SDK tool or the AWS CLI, you can specify SourceRegion (or
	// --source-region for the AWS CLI) instead of specifying PreSignedUrl manually.
	// Specifying SourceRegion autogenerates a pre-signed URL that is a valid request
	// for the operation that can be executed in the source AWS Region.
	PreSignedUrl *string

	// The daily time range during which automated backups are created if automated
	// backups are enabled using the BackupRetentionPeriod parameter. The default is a
	// 30-minute window selected at random from an 8-hour block of time for each AWS
	// Region. To see the time blocks available, see  Adjusting the Preferred DB
	// Cluster Maintenance Window
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow.Aurora)
	// in the Amazon Aurora User Guide. Constraints:
	//
	// * Must be in the format
	// hh24:mi-hh24:mi.
	//
	// * Must be in Universal Coordinated Time (UTC).
	//
	// * Must not
	// conflict with the preferred maintenance window.
	//
	// * Must be at least 30 minutes.
	PreferredBackupWindow *string

	// The weekly time range during which system maintenance can occur, in Universal
	// Coordinated Time (UTC). Format: ddd:hh24:mi-ddd:hh24:mi The default is a
	// 30-minute window selected at random from an 8-hour block of time for each AWS
	// Region, occurring on a random day of the week. To see the time blocks available,
	// see  Adjusting the Preferred DB Cluster Maintenance Window
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow.Aurora)
	// in the Amazon Aurora User Guide. Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string

	// The Amazon Resource Name (ARN) of the source DB instance or DB cluster if this
	// DB cluster is created as a read replica.
	ReplicationSourceIdentifier *string

	// For DB clusters in serverless DB engine mode, the scaling properties of the DB
	// cluster.
	ScalingConfiguration *types.ScalingConfiguration

	// The AWS region the resource is in. The presigned URL will be created with this
	// region, if the PresignURL member is empty set.
	SourceRegion *string

	// A value that indicates whether the DB cluster is encrypted.
	StorageEncrypted *bool

	// Tags to assign to the DB cluster.
	Tags []*types.Tag

	// A list of EC2 VPC security groups to associate with this DB cluster.
	VpcSecurityGroupIds []*string

	// Used by the SDK's PresignURL autofill customization to specify the region the of
	// the client's request.
	destinationRegion *string
}

type CreateDBClusterOutput struct {

	// Contains the details of an Amazon Aurora DB cluster. This data type is used as a
	// response element in the DescribeDBClusters, StopDBCluster, and StartDBCluster
	// actions.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateDBClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBCluster{}, middleware.After)
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
	addCreateDBClusterPresignURLMiddleware(stack, options)
	addOpCreateDBClusterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBCluster(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func copyCreateDBClusterInputForPresign(params interface{}) (interface{}, error) {
	input, ok := params.(*CreateDBClusterInput)
	if !ok {
		return nil, fmt.Errorf("expect *CreateDBClusterInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func getCreateDBClusterPreSignedUrl(params interface{}) (string, bool, error) {
	input, ok := params.(*CreateDBClusterInput)
	if !ok {
		return ``, false, fmt.Errorf("expect *CreateDBClusterInput type, got %T", params)
	}
	if input.PreSignedUrl == nil || len(*input.PreSignedUrl) == 0 {
		return ``, false, nil
	}
	return *input.PreSignedUrl, true, nil
}
func getCreateDBClusterSourceRegion(params interface{}) (string, bool, error) {
	input, ok := params.(*CreateDBClusterInput)
	if !ok {
		return ``, false, fmt.Errorf("expect *CreateDBClusterInput type, got %T", params)
	}
	if input.SourceRegion == nil || len(*input.SourceRegion) == 0 {
		return ``, false, nil
	}
	return *input.SourceRegion, true, nil
}
func setCreateDBClusterPreSignedUrl(params interface{}, value string) error {
	input, ok := params.(*CreateDBClusterInput)
	if !ok {
		return fmt.Errorf("expect *CreateDBClusterInput type, got %T", params)
	}
	input.PreSignedUrl = &value
	return nil
}
func setCreateDBClusterdestinationRegion(params interface{}, value string) error {
	input, ok := params.(*CreateDBClusterInput)
	if !ok {
		return fmt.Errorf("expect *CreateDBClusterInput type, got %T", params)
	}
	input.destinationRegion = &value
	return nil
}

type createDBClusterHTTPPresignURLClient struct {
	client    *Client
	presigner *v4.Signer
}

func newCreateDBClusterHTTPPresignURLClient(options Options, optFns ...func(*Options)) *createDBClusterHTTPPresignURLClient {
	return &createDBClusterHTTPPresignURLClient{
		client:    New(options, optFns...),
		presigner: v4.NewSigner(),
	}
}
func (c *createDBClusterHTTPPresignURLClient) PresignCreateDBCluster(ctx context.Context, params *CreateDBClusterInput, optFns ...func(*Options)) (string, http.Header, error) {
	if params == nil {
		params = &CreateDBClusterInput{}
	}

	optFns = append(optFns, func(o *Options) {
		o.HTTPClient = &smithyhttp.NopClient{}
	})

	ctx = presignedurlcust.WithIsPresigning(ctx)
	result, _, err := c.client.invokeOperation(ctx, "CreateDBCluster", params, optFns,
		addOperationCreateDBClusterMiddlewares,
		c.convertToPresignMiddleware,
	)
	if err != nil {
		return ``, nil, err
	}

	out := result.(*v4.PresignedHTTPRequest)
	return out.URL, out.SignedHeader, nil
}
func (c *createDBClusterHTTPPresignURLClient) convertToPresignMiddleware(stack *middleware.Stack, options Options) (err error) {
	stack.Finalize.Clear()
	stack.Deserialize.Clear()
	stack.Build.Remove(awsmiddleware.RequestInvocationIDMiddleware{}.ID())
	err = stack.Finalize.Add(v4.NewPresignHTTPRequestMiddleware(options.Credentials, c.presigner), middleware.After)
	if err != nil {
		return err
	}
	err = query.AddAsGetRequestMiddleware(stack)
	if err != nil {
		return err
	}
	return nil
}
func addCreateDBClusterPresignURLMiddleware(stack *middleware.Stack, options Options) error {
	return presignedurlcust.AddMiddleware(stack, presignedurlcust.Options{
		Accessor: presignedurlcust.ParameterAccessor{
			GetPresignedURL:      getCreateDBClusterPreSignedUrl,
			GetSourceRegion:      getCreateDBClusterSourceRegion,
			CopyInput:            copyCreateDBClusterInputForPresign,
			SetDestinationRegion: setCreateDBClusterdestinationRegion,
			SetPresignedURL:      setCreateDBClusterPreSignedUrl,
		},
		Presigner: &presignAutoFillCreateDBClusterClient{client: newCreateDBClusterHTTPPresignURLClient(options)},
	})
}

type presignAutoFillCreateDBClusterClient struct {
	client *createDBClusterHTTPPresignURLClient
}

func (c *presignAutoFillCreateDBClusterClient) PresignURL(ctx context.Context, region string, params interface{}) (string, http.Header, error) {
	input, ok := params.(*CreateDBClusterInput)
	if !ok {
		return ``, nil, fmt.Errorf("expect *CreateDBClusterInput type, got %T", params)
	}
	optFn := func(o *Options) {
		o.Region = region
		o.APIOptions = append(o.APIOptions, presignedurlcust.RemoveMiddleware)
	}
	return c.client.PresignCreateDBCluster(ctx, input, optFn)
}

func newServiceMetadataMiddleware_opCreateDBCluster(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateDBCluster",
	}
}
