// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Contains an access policy that defines an identity's access to an AWS IoT
// SiteWise Monitor resource.
type AccessPolicySummary struct {

	// The ID of the access policy.
	//
	// This member is required.
	Id *string

	// The identity (an AWS SSO user, an AWS SSO group, or an IAM user).
	//
	// This member is required.
	Identity *Identity

	// The permissions for the access policy. Note that a project ADMINISTRATOR is also
	// known as a project owner.
	//
	// This member is required.
	Permission Permission

	// The AWS IoT SiteWise Monitor resource (a portal or project).
	//
	// This member is required.
	Resource *Resource

	// The date the access policy was created, in Unix epoch time.
	CreationDate *time.Time

	// The date the access policy was last updated, in Unix epoch time.
	LastUpdateDate *time.Time
}

// Contains aggregated asset property values (for example, average, minimum, and
// maximum).
type AggregatedValue struct {

	// The date the aggregating computations occurred, in Unix epoch time.
	//
	// This member is required.
	Timestamp *time.Time

	// The value of the aggregates.
	//
	// This member is required.
	Value *Aggregates

	// The quality of the aggregated data.
	Quality Quality
}

// Contains the (pre-calculated) aggregate values for an asset property.
type Aggregates struct {

	// The average (mean) value of the time series over a time interval window.
	Average *float64

	// The count of data points in the time series over a time interval window.
	Count *float64

	// The maximum value of the time series over a time interval window.
	Maximum *float64

	// The minimum value of the time series over a time interval window.
	Minimum *float64

	// The standard deviation of the time series over a time interval window.
	StandardDeviation *float64

	// The sum of the time series over a time interval window.
	Sum *float64
}

// Contains error details for the requested associate project asset action.
type AssetErrorDetails struct {

	// The ID of the asset.
	//
	// This member is required.
	AssetId *string

	// The error code.
	//
	// This member is required.
	Code AssetErrorCode

	// The error message.
	//
	// This member is required.
	Message *string
}

// Describes an asset hierarchy that contains a hierarchy's name and ID.
type AssetHierarchy struct {

	// The hierarchy name provided in the CreateAssetModel
	// (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_CreateAssetModel.html)
	// or UpdateAssetModel
	// (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_UpdateAssetModel.html)
	// API operation.
	//
	// This member is required.
	Name *string

	// The ID of the hierarchy. This ID is a hierarchyId.
	Id *string
}

// Describes an asset hierarchy that contains a hierarchy's name, ID, and child
// asset model ID that specifies the type of asset that can be in this hierarchy.
type AssetModelHierarchy struct {

	// The ID of the asset model. All assets in this hierarchy must be instances of the
	// childAssetModelId asset model.
	//
	// This member is required.
	ChildAssetModelId *string

	// The name of the asset model hierarchy that you specify by using the
	// CreateAssetModel
	// (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_CreateAssetModel.html)
	// or UpdateAssetModel
	// (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_UpdateAssetModel.html)
	// API operation.
	//
	// This member is required.
	Name *string

	// The ID of the asset model hierarchy. This ID is a hierarchyId.
	Id *string
}

// Contains an asset model hierarchy used in asset model creation. An asset model
// hierarchy determines the kind (or type) of asset that can belong to a hierarchy.
type AssetModelHierarchyDefinition struct {

	// The ID of an asset model for this hierarchy.
	//
	// This member is required.
	ChildAssetModelId *string

	// The name of the asset model hierarchy definition (as specified in the
	// CreateAssetModel
	// (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_CreateAssetModel.html)
	// or UpdateAssetModel
	// (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_UpdateAssetModel.html)
	// API operation).
	//
	// This member is required.
	Name *string
}

// Contains information about an asset model property.
type AssetModelProperty struct {

	// The data type of the asset model property.
	//
	// This member is required.
	DataType PropertyDataType

	// The name of the asset model property.
	//
	// This member is required.
	Name *string

	// The property type (see PropertyType).
	//
	// This member is required.
	Type *PropertyType

	// The ID of the asset model property.
	Id *string

	// The unit of the asset model property, such as Newtons or RPM.
	Unit *string
}

// Contains an asset model property definition. This property definition is applied
// to all assets created from the asset model.
type AssetModelPropertyDefinition struct {

	// The data type of the property definition.
	//
	// This member is required.
	DataType PropertyDataType

	// The name of the property definition.
	//
	// This member is required.
	Name *string

	// The property definition type (see PropertyType). You can only specify one type
	// in a property definition.
	//
	// This member is required.
	Type *PropertyType

	// The unit of the property definition, such as Newtons or RPM.
	Unit *string
}

// Contains current status information for an asset model. For more information,
// see Asset and model states
// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-and-model-states.html)
// in the AWS IoT SiteWise User Guide.
type AssetModelStatus struct {

	// The current state of the asset model.
	//
	// This member is required.
	State AssetModelState

	// Contains associated error information, if any.
	Error *ErrorDetails
}

// Contains a summary of an asset model.
type AssetModelSummary struct {

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the asset model, which has the following format.
	// arn:${Partition}:iotsitewise:${Region}:${Account}:asset-model/${AssetModelId}
	//
	// This member is required.
	Arn *string

	// The date the asset model was created, in Unix epoch time.
	//
	// This member is required.
	CreationDate *time.Time

	// The asset model description.
	//
	// This member is required.
	Description *string

	// The ID of the asset model (used with AWS IoT SiteWise APIs).
	//
	// This member is required.
	Id *string

	// The date the asset model was last updated, in Unix epoch time.
	//
	// This member is required.
	LastUpdateDate *time.Time

	// The name of the asset model.
	//
	// This member is required.
	Name *string

	// The current status of the asset model.
	//
	// This member is required.
	Status *AssetModelStatus
}

// Contains asset property information.
type AssetProperty struct {

	// The data type of the asset property.
	//
	// This member is required.
	DataType PropertyDataType

	// The ID of the asset property.
	//
	// This member is required.
	Id *string

	// The name of the property.
	//
	// This member is required.
	Name *string

	// The property alias that identifies the property, such as an OPC-UA server data
	// stream path (for example, /company/windfarm/3/turbine/7/temperature). For more
	// information, see Mapping industrial data streams to asset properties
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/connect-data-streams.html)
	// in the AWS IoT SiteWise User Guide.
	Alias *string

	// The asset property's notification topic and state. For more information, see
	// UpdateAssetProperty
	// (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_UpdateAssetProperty.html).
	Notification *PropertyNotification

	// The unit (such as Newtons or RPM) of the asset property.
	Unit *string
}

// Contains asset property value information.
type AssetPropertyValue struct {

	// The timestamp of the asset property value.
	//
	// This member is required.
	Timestamp *TimeInNanos

	// The value of the asset property (see Variant).
	//
	// This member is required.
	Value *Variant

	// The quality of the asset property value.
	Quality Quality
}

// Contains information about the current status of an asset. For more information,
// see Asset and model states
// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-and-model-states.html)
// in the AWS IoT SiteWise User Guide.
type AssetStatus struct {

	// The current status of the asset.
	//
	// This member is required.
	State AssetState

	// Contains associated error information, if any.
	Error *ErrorDetails
}

// Contains a summary of an asset.
type AssetSummary struct {

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the asset, which has the following format.
	// arn:${Partition}:iotsitewise:${Region}:${Account}:asset/${AssetId}
	//
	// This member is required.
	Arn *string

	// The ID of the asset model used to create this asset.
	//
	// This member is required.
	AssetModelId *string

	// The date the asset was created, in Unix epoch time.
	//
	// This member is required.
	CreationDate *time.Time

	// A list of asset hierarchies that each contain a hierarchyId. A hierarchy
	// specifies allowed parent/child asset relationships.
	//
	// This member is required.
	Hierarchies []*AssetHierarchy

	// The ID of the asset.
	//
	// This member is required.
	Id *string

	// The date the asset was last updated, in Unix epoch time.
	//
	// This member is required.
	LastUpdateDate *time.Time

	// The name of the asset.
	//
	// This member is required.
	Name *string

	// The current status of the asset.
	//
	// This member is required.
	Status *AssetStatus
}

// Contains a summary of an associated asset.
type AssociatedAssetsSummary struct {

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the asset, which has the following format.
	// arn:${Partition}:iotsitewise:${Region}:${Account}:asset/${AssetId}
	//
	// This member is required.
	Arn *string

	// The ID of the asset model used to create the asset.
	//
	// This member is required.
	AssetModelId *string

	// The date the asset was created, in Unix epoch time.
	//
	// This member is required.
	CreationDate *time.Time

	// A list of asset hierarchies that each contain a hierarchyId. A hierarchy
	// specifies allowed parent/child asset relationships.
	//
	// This member is required.
	Hierarchies []*AssetHierarchy

	// The ID of the asset.
	//
	// This member is required.
	Id *string

	// The date the asset was last updated, in Unix epoch time.
	//
	// This member is required.
	LastUpdateDate *time.Time

	// The name of the asset.
	//
	// This member is required.
	Name *string

	// The current status of the asset.
	//
	// This member is required.
	Status *AssetStatus
}

// Contains an asset attribute property. For more information, see Attributes
// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-properties.html#attributes)
// in the AWS IoT SiteWise User Guide.
type Attribute struct {

	// The default value of the asset model property attribute. All assets that you
	// create from the asset model contain this attribute value. You can update an
	// attribute's value after you create an asset. For more information, see Updating
	// attribute values
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/update-attribute-values.html)
	// in the AWS IoT SiteWise User Guide.
	DefaultValue *string
}

// Contains error information from updating a batch of asset property values.
type BatchPutAssetPropertyError struct {

	// The error code.
	//
	// This member is required.
	ErrorCode BatchPutAssetPropertyValueErrorCode

	// The associated error message.
	//
	// This member is required.
	ErrorMessage *string

	// A list of timestamps for each error, if any.
	//
	// This member is required.
	Timestamps []*TimeInNanos
}

// Contains error information for asset property value entries that are associated
// with the BatchPutAssetPropertyValue
// (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_BatchPutAssetPropertyValue.html)
// API.
type BatchPutAssetPropertyErrorEntry struct {

	// The ID of the failed entry.
	//
	// This member is required.
	EntryId *string

	// The list of update property value errors.
	//
	// This member is required.
	Errors []*BatchPutAssetPropertyError
}

// Contains a dashboard summary.
type DashboardSummary struct {

	// The ID of the dashboard.
	//
	// This member is required.
	Id *string

	// The name of the dashboard
	//
	// This member is required.
	Name *string

	// The date the dashboard was created, in Unix epoch time.
	CreationDate *time.Time

	// The dashboard's description.
	Description *string

	// The date the dashboard was last updated, in Unix epoch time.
	LastUpdateDate *time.Time
}

// Contains the details of an AWS IoT SiteWise error.
type ErrorDetails struct {

	// The error code.
	//
	// This member is required.
	Code ErrorCode

	// The error message.
	//
	// This member is required.
	Message *string
}

// Contains expression variable information.
type ExpressionVariable struct {

	// The friendly name of the variable to be used in the expression.
	//
	// This member is required.
	Name *string

	// The variable that identifies an asset property from which to use values.
	//
	// This member is required.
	Value *VariableValue
}

// Contains a summary of a gateway capability configuration.
type GatewayCapabilitySummary struct {

	// The namespace of the capability configuration. For example, if you configure
	// OPC-UA sources from the AWS IoT SiteWise console, your OPC-UA capability
	// configuration has the namespace iotsitewise:opcuacollector:version, where
	// version is a number such as 1.
	//
	// This member is required.
	CapabilityNamespace *string

	// The synchronization status of the capability configuration. The sync status can
	// be one of the following:
	//
	// * IN_SYNC – The gateway is running the capability
	// configuration.
	//
	// * OUT_OF_SYNC – The gateway hasn't received the capability
	// configuration.
	//
	// * SYNC_FAILED – The gateway rejected the capability
	// configuration.
	//
	// This member is required.
	CapabilitySyncStatus CapabilitySyncStatus
}

// Contains a gateway's platform information.
type GatewayPlatform struct {

	// A gateway that runs on AWS IoT Greengrass.
	//
	// This member is required.
	Greengrass *Greengrass
}

// Contains a summary of a gateway.
type GatewaySummary struct {

	// The date the gateway was created, in Unix epoch time.
	//
	// This member is required.
	CreationDate *time.Time

	// The ID of the gateway device.
	//
	// This member is required.
	GatewayId *string

	// The name of the asset.
	//
	// This member is required.
	GatewayName *string

	// The date the gateway was last updated, in Unix epoch time.
	//
	// This member is required.
	LastUpdateDate *time.Time

	// A list of gateway capability summaries that each contain a namespace and status.
	// Each gateway capability defines data sources for the gateway. To retrieve a
	// capability configuration's definition, use
	// DescribeGatewayCapabilityConfiguration
	// (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_DescribeGatewayCapabilityConfiguration.html).
	GatewayCapabilitySummaries []*GatewayCapabilitySummary
}

// Contains details for a gateway that runs on AWS IoT Greengrass. To create a
// gateway that runs on AWS IoT Greengrass, you must add the IoT SiteWise connector
// to a Greengrass group and deploy it. Your Greengrass group must also have
// permissions to upload data to AWS IoT SiteWise. For more information, see
// Ingesting data using a gateway
// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/gateway-connector.html)
// in the AWS IoT SiteWise User Guide.
type Greengrass struct {

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the Greengrass group. For more information about how to find a group's ARN, see
	// ListGroups
	// (https://docs.aws.amazon.com/greengrass/latest/apireference/listgroups-get.html)
	// and GetGroup
	// (https://docs.aws.amazon.com/greengrass/latest/apireference/getgroup-get.html)
	// in the AWS IoT Greengrass API Reference.
	//
	// This member is required.
	GroupArn *string
}

// Contains information for a group identity in an access policy.
type GroupIdentity struct {

	// The AWS SSO ID of the group.
	//
	// This member is required.
	Id *string
}

// Contains information about an AWS Identity and Access Management (IAM) user.
type IAMUserIdentity struct {

	// The ARN of the IAM user. IAM users must have the
	// iotsitewise:CreatePresignedPortalUrl permission to sign in to the portal. For
	// more information, see IAM ARNs
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html) in
	// the IAM User Guide. If you delete the IAM user, access policies that contain
	// this identity include an empty arn. You can delete the access policy for the IAM
	// user that no longer exists.
	//
	// This member is required.
	Arn *string
}

// Contains an identity that can access an AWS IoT SiteWise Monitor resource.
// Currently, you can't use AWS APIs to retrieve AWS SSO identity IDs. You can find
// the AWS SSO identity IDs in the URL of user and group pages in the AWS SSO
// console (https://console.aws.amazon.com/singlesignon).
type Identity struct {

	// An AWS SSO group identity.
	Group *GroupIdentity

	// An IAM user identity.
	IamUser *IAMUserIdentity

	// An AWS SSO user identity.
	User *UserIdentity
}

// Contains an image that is one of the following:
//
// * An image file. Choose this
// option to upload a new image.
//
// * The ID of an existing image. Choose this option
// to keep an existing image.
type Image struct {

	// Contains an image file.
	File *ImageFile

	// The ID of an existing image. Specify this parameter to keep an existing image.
	Id *string
}

// Contains an image file.
type ImageFile struct {

	// The image file contents, represented as a base64-encoded string. The file size
	// must be less than 1 MB.
	//
	// This member is required.
	Data []byte

	// The file type of the image.
	//
	// This member is required.
	Type ImageFileType
}

// Contains an image that is uploaded to AWS IoT SiteWise and available at a URL.
type ImageLocation struct {

	// The ID of the image.
	//
	// This member is required.
	Id *string

	// The URL where the image is available. The URL is valid for 15 minutes so that
	// you can view and download the image
	//
	// This member is required.
	Url *string
}

// Contains logging options.
type LoggingOptions struct {

	// The AWS IoT SiteWise logging verbosity level.
	//
	// This member is required.
	Level LoggingLevel
}

// Contains an asset measurement property. This structure is empty. For more
// information, see Measurements
// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-properties.html#measurements)
// in the AWS IoT SiteWise User Guide.
type Measurement struct {
}

// Contains an asset metric property. With metrics, you can calculate aggregate
// functions, such as an average, maximum, or minimum, as specified through an
// expression. A metric maps several values to a single value (such as a sum). The
// maximum number of dependent/cascading variables used in any one metric
// calculation is 10. Therefore, a root metric can have up to 10 cascading metrics
// in its computational dependency tree. Additionally, a metric can only have a
// data type of DOUBLE and consume properties with data types of INTEGER or DOUBLE.
// For more information, see Metrics
// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-properties.html#metrics)
// in the AWS IoT SiteWise User Guide.
type Metric struct {

	// The mathematical expression that defines the metric aggregation function. You
	// can specify up to 10 variables per expression. You can specify up to 10
	// functions per expression. For more information, see Quotas
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/quotas.html) in the
	// AWS IoT SiteWise User Guide.
	//
	// This member is required.
	Expression *string

	// The list of variables used in the expression.
	//
	// This member is required.
	Variables []*ExpressionVariable

	// The window (time interval) over which AWS IoT SiteWise computes the metric's
	// aggregation expression. AWS IoT SiteWise computes one data point per window.
	//
	// This member is required.
	Window *MetricWindow
}

// Contains a time interval window used for data aggregate computations (for
// example, average, sum, count, and so on).
type MetricWindow struct {

	// The tumbling time interval window.
	Tumbling *TumblingWindow
}

// Contains AWS IoT SiteWise Monitor error details.
type MonitorErrorDetails struct {

	// The error code.
	Code MonitorErrorCode

	// The error message.
	Message *string
}

// Identifies an AWS IoT SiteWise Monitor portal.
type PortalResource struct {

	// The ID of the portal.
	//
	// This member is required.
	Id *string
}

// Contains information about the current status of a portal.
type PortalStatus struct {

	// The current state of the portal.
	//
	// This member is required.
	State PortalState

	// Contains associated error information, if any.
	Error *MonitorErrorDetails
}

// Contains a portal summary.
type PortalSummary struct {

	// The ID of the portal.
	//
	// This member is required.
	Id *string

	// The name of the portal.
	//
	// This member is required.
	Name *string

	// The URL for the AWS IoT SiteWise Monitor portal. You can use this URL to access
	// portals that use AWS SSO for authentication. For portals that use IAM for
	// authentication, you must use the CreatePresignedPortalUrl
	// (https://docs.aws.amazon.com/AWS IoT SiteWise API
	// ReferenceAPI_CreatePresignedPortalUrl.html) operation to create a URL that you
	// can use to access the portal.
	//
	// This member is required.
	StartUrl *string

	// Contains information about the current status of a portal.
	//
	// This member is required.
	Status *PortalStatus

	// The date the portal was created, in Unix epoch time.
	CreationDate *time.Time

	// The portal's description.
	Description *string

	// The date the portal was last updated, in Unix epoch time.
	LastUpdateDate *time.Time

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the service role that allows the portal's users to access your AWS IoT SiteWise
	// resources on your behalf. For more information, see Using service roles for AWS
	// IoT SiteWise Monitor
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/monitor-service-role.html)
	// in the AWS IoT SiteWise User Guide.
	RoleArn *string
}

// Identifies a specific AWS IoT SiteWise Monitor project.
type ProjectResource struct {

	// The ID of the project.
	//
	// This member is required.
	Id *string
}

// Contains project summary information.
type ProjectSummary struct {

	// The ID of the project.
	//
	// This member is required.
	Id *string

	// The name of the project.
	//
	// This member is required.
	Name *string

	// The date the project was created, in Unix epoch time.
	CreationDate *time.Time

	// The project's description.
	Description *string

	// The date the project was last updated, in Unix epoch time.
	LastUpdateDate *time.Time
}

// Contains asset property information.
type Property struct {

	// The property data type.
	//
	// This member is required.
	DataType PropertyDataType

	// The ID of the asset property.
	//
	// This member is required.
	Id *string

	// The name of the property.
	//
	// This member is required.
	Name *string

	// The property alias that identifies the property, such as an OPC-UA server data
	// stream path (for example, /company/windfarm/3/turbine/7/temperature). For more
	// information, see Mapping industrial data streams to asset properties
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/connect-data-streams.html)
	// in the AWS IoT SiteWise User Guide.
	Alias *string

	// The asset property's notification topic and state. For more information, see
	// UpdateAssetProperty
	// (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_UpdateAssetProperty.html).
	Notification *PropertyNotification

	// The property type (see PropertyType). A property contains one type.
	Type *PropertyType

	// The unit (such as Newtons or RPM) of the asset property.
	Unit *string
}

// Contains asset property value notification information. When the notification
// state is enabled, AWS IoT SiteWise publishes property value updates to a unique
// MQTT topic. For more information, see Interacting with other services
// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/interact-with-other-services.html)
// in the AWS IoT SiteWise User Guide.
type PropertyNotification struct {

	// The current notification state.
	//
	// This member is required.
	State PropertyNotificationState

	// The MQTT topic to which AWS IoT SiteWise publishes property value update
	// notifications.
	//
	// This member is required.
	Topic *string
}

// Contains a property type, which can be one of attribute, measurement, metric, or
// transform.
type PropertyType struct {

	// Specifies an asset attribute property. An attribute generally contains static
	// information, such as the serial number of an IIoT
	// (https://en.wikipedia.org/wiki/Internet_of_things#Industrial_applications) wind
	// turbine.
	Attribute *Attribute

	// Specifies an asset measurement property. A measurement represents a device's raw
	// sensor data stream, such as timestamped temperature values or timestamped power
	// values.
	Measurement *Measurement

	// Specifies an asset metric property. A metric contains a mathematical expression
	// that uses aggregate functions to process all input data points over a time
	// interval and output a single data point, such as to calculate the average hourly
	// temperature.
	Metric *Metric

	// Specifies an asset transform property. A transform contains a mathematical
	// expression that maps a property's data points from one form to another, such as
	// a unit conversion from Celsius to Fahrenheit.
	Transform *Transform
}

// Contains a list of value updates for an asset property in the list of asset
// entries consumed by the BatchPutAssetPropertyValue
// (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_BatchPutAssetPropertyValue.html)
// API operation.
type PutAssetPropertyValueEntry struct {

	// The user specified ID for the entry. You can use this ID to identify which
	// entries failed.
	//
	// This member is required.
	EntryId *string

	// The list of property values to upload. You can specify up to 10 propertyValues
	// array elements.
	//
	// This member is required.
	PropertyValues []*AssetPropertyValue

	// The ID of the asset to update.
	AssetId *string

	// The property alias that identifies the property, such as an OPC-UA server data
	// stream path (for example, /company/windfarm/3/turbine/7/temperature). For more
	// information, see Mapping industrial data streams to asset properties
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/connect-data-streams.html)
	// in the AWS IoT SiteWise User Guide.
	PropertyAlias *string

	// The ID of the asset property for this entry.
	PropertyId *string
}

// Contains an AWS IoT SiteWise Monitor resource ID for a portal or project.
type Resource struct {

	// A portal resource.
	Portal *PortalResource

	// A project resource.
	Project *ProjectResource
}

// Contains a timestamp with optional nanosecond granularity.
type TimeInNanos struct {

	// The timestamp date, in seconds, in the Unix epoch format. Fractional nanosecond
	// data is provided by offsetInNanos.
	//
	// This member is required.
	TimeInSeconds *int64

	// The nanosecond offset from timeInSeconds.
	OffsetInNanos *int32
}

// Contains an asset transform property. A transform is a one-to-one mapping of a
// property's data points from one form to another. For example, you can use a
// transform to convert a Celsius data stream to Fahrenheit by applying the
// transformation expression to each data point of the Celsius stream. A transform
// can only have a data type of DOUBLE and consume properties with data types of
// INTEGER or DOUBLE. For more information, see Transforms
// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-properties.html#transforms)
// in the AWS IoT SiteWise User Guide.
type Transform struct {

	// The mathematical expression that defines the transformation function. You can
	// specify up to 10 variables per expression. You can specify up to 10 functions
	// per expression. For more information, see Quotas
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/quotas.html) in the
	// AWS IoT SiteWise User Guide.
	//
	// This member is required.
	Expression *string

	// The list of variables used in the expression.
	//
	// This member is required.
	Variables []*ExpressionVariable
}

// Contains a tumbling window, which is a repeating fixed-sized, non-overlapping,
// and contiguous time interval. This window is used in metric and aggregation
// computations.
type TumblingWindow struct {

	// The time interval for the tumbling window. Note that w represents weeks, d
	// represents days, h represents hours, and m represents minutes. AWS IoT SiteWise
	// computes the 1w interval the end of Sunday at midnight each week (UTC), the 1d
	// interval at the end of each day at midnight (UTC), the 1h interval at the end of
	// each hour, and so on. When AWS IoT SiteWise aggregates data points for metric
	// computations, the start of each interval is exclusive and the end of each
	// interval is inclusive. AWS IoT SiteWise places the computed data point at the
	// end of the interval.
	//
	// This member is required.
	Interval *string
}

// Contains information for a user identity in an access policy.
type UserIdentity struct {

	// The AWS SSO ID of the user.
	//
	// This member is required.
	Id *string
}

// Identifies a property value used in an expression.
type VariableValue struct {

	// The ID of the property to use as the variable. You can use the property name if
	// it's from the same asset model.
	//
	// This member is required.
	PropertyId *string

	// The ID of the hierarchy to query for the property ID. You can use the
	// hierarchy's name instead of the hierarchy's ID. You use a hierarchy ID instead
	// of a model ID because you can have several hierarchies using the same model and
	// therefore the same propertyId. For example, you might have separately grouped
	// assets that come from the same asset model. For more information, see Asset
	// hierarchies
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-hierarchies.html)
	// in the AWS IoT SiteWise User Guide.
	HierarchyId *string
}

// Contains an asset property value (of a single type only).
type Variant struct {

	// Asset property data of type Boolean (true or false).
	BooleanValue *bool

	// Asset property data of type double (floating point number).
	DoubleValue *float64

	// Asset property data of type integer (whole number).
	IntegerValue *int32

	// Asset property data of type string (sequence of characters).
	StringValue *string
}
