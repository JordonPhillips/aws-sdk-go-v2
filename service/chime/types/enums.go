// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AccountType string

// Enum values for AccountType
const (
	AccountTypeTeam                AccountType = "Team"
	AccountTypeEnterprisedirectory AccountType = "EnterpriseDirectory"
	AccountTypeEnterpriselwa       AccountType = "EnterpriseLWA"
	AccountTypeEnterpriseoidc      AccountType = "EnterpriseOIDC"
)

// Values returns all known values for AccountType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (AccountType) Values() []AccountType {
	return []AccountType{
		"Team",
		"EnterpriseDirectory",
		"EnterpriseLWA",
		"EnterpriseOIDC",
	}
}

type BotType string

// Enum values for BotType
const (
	BotTypeChatbot BotType = "ChatBot"
)

// Values returns all known values for BotType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (BotType) Values() []BotType {
	return []BotType{
		"ChatBot",
	}
}

type CallingNameStatus string

// Enum values for CallingNameStatus
const (
	CallingNameStatusUnassigned       CallingNameStatus = "Unassigned"
	CallingNameStatusUpdateinprogress CallingNameStatus = "UpdateInProgress"
	CallingNameStatusUpdatesucceeded  CallingNameStatus = "UpdateSucceeded"
	CallingNameStatusUpdatefailed     CallingNameStatus = "UpdateFailed"
)

// Values returns all known values for CallingNameStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CallingNameStatus) Values() []CallingNameStatus {
	return []CallingNameStatus{
		"Unassigned",
		"UpdateInProgress",
		"UpdateSucceeded",
		"UpdateFailed",
	}
}

type Capability string

// Enum values for Capability
const (
	CapabilityVoice Capability = "Voice"
	CapabilitySms   Capability = "SMS"
)

// Values returns all known values for Capability. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (Capability) Values() []Capability {
	return []Capability{
		"Voice",
		"SMS",
	}
}

type EmailStatus string

// Enum values for EmailStatus
const (
	EmailStatusNotsent EmailStatus = "NotSent"
	EmailStatusSent    EmailStatus = "Sent"
	EmailStatusFailed  EmailStatus = "Failed"
)

// Values returns all known values for EmailStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (EmailStatus) Values() []EmailStatus {
	return []EmailStatus{
		"NotSent",
		"Sent",
		"Failed",
	}
}

type ErrorCode string

// Enum values for ErrorCode
const (
	ErrorCodeBadrequest                           ErrorCode = "BadRequest"
	ErrorCodeConflict                             ErrorCode = "Conflict"
	ErrorCodeForbidden                            ErrorCode = "Forbidden"
	ErrorCodeNotfound                             ErrorCode = "NotFound"
	ErrorCodePreconditionfailed                   ErrorCode = "PreconditionFailed"
	ErrorCodeResourcelimitexceeded                ErrorCode = "ResourceLimitExceeded"
	ErrorCodeServicefailure                       ErrorCode = "ServiceFailure"
	ErrorCodeAccessdenied                         ErrorCode = "AccessDenied"
	ErrorCodeServiceunavailable                   ErrorCode = "ServiceUnavailable"
	ErrorCodeThrottled                            ErrorCode = "Throttled"
	ErrorCodeThrottling                           ErrorCode = "Throttling"
	ErrorCodeUnauthorized                         ErrorCode = "Unauthorized"
	ErrorCodeUnprocessable                        ErrorCode = "Unprocessable"
	ErrorCodeVoiceconnectorgroupassociationsexist ErrorCode = "VoiceConnectorGroupAssociationsExist"
	ErrorCodePhonenumberassociationsexist         ErrorCode = "PhoneNumberAssociationsExist"
)

// Values returns all known values for ErrorCode. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (ErrorCode) Values() []ErrorCode {
	return []ErrorCode{
		"BadRequest",
		"Conflict",
		"Forbidden",
		"NotFound",
		"PreconditionFailed",
		"ResourceLimitExceeded",
		"ServiceFailure",
		"AccessDenied",
		"ServiceUnavailable",
		"Throttled",
		"Throttling",
		"Unauthorized",
		"Unprocessable",
		"VoiceConnectorGroupAssociationsExist",
		"PhoneNumberAssociationsExist",
	}
}

type GeoMatchLevel string

// Enum values for GeoMatchLevel
const (
	GeoMatchLevelCountry  GeoMatchLevel = "Country"
	GeoMatchLevelAreacode GeoMatchLevel = "AreaCode"
)

// Values returns all known values for GeoMatchLevel. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (GeoMatchLevel) Values() []GeoMatchLevel {
	return []GeoMatchLevel{
		"Country",
		"AreaCode",
	}
}

type InviteStatus string

// Enum values for InviteStatus
const (
	InviteStatusPending  InviteStatus = "Pending"
	InviteStatusAccepted InviteStatus = "Accepted"
	InviteStatusFailed   InviteStatus = "Failed"
)

// Values returns all known values for InviteStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (InviteStatus) Values() []InviteStatus {
	return []InviteStatus{
		"Pending",
		"Accepted",
		"Failed",
	}
}

type License string

// Enum values for License
const (
	LicenseBasic    License = "Basic"
	LicensePlus     License = "Plus"
	LicensePro      License = "Pro"
	LicenseProtrial License = "ProTrial"
)

// Values returns all known values for License. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (License) Values() []License {
	return []License{
		"Basic",
		"Plus",
		"Pro",
		"ProTrial",
	}
}

type MemberType string

// Enum values for MemberType
const (
	MemberTypeUser    MemberType = "User"
	MemberTypeBot     MemberType = "Bot"
	MemberTypeWebhook MemberType = "Webhook"
)

// Values returns all known values for MemberType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (MemberType) Values() []MemberType {
	return []MemberType{
		"User",
		"Bot",
		"Webhook",
	}
}

type NotificationTarget string

// Enum values for NotificationTarget
const (
	NotificationTargetEventbridge NotificationTarget = "EventBridge"
	NotificationTargetSns         NotificationTarget = "SNS"
	NotificationTargetSqs         NotificationTarget = "SQS"
)

// Values returns all known values for NotificationTarget. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NotificationTarget) Values() []NotificationTarget {
	return []NotificationTarget{
		"EventBridge",
		"SNS",
		"SQS",
	}
}

type NumberSelectionBehavior string

// Enum values for NumberSelectionBehavior
const (
	NumberSelectionBehaviorPrefersticky NumberSelectionBehavior = "PreferSticky"
	NumberSelectionBehaviorAvoidsticky  NumberSelectionBehavior = "AvoidSticky"
)

// Values returns all known values for NumberSelectionBehavior. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NumberSelectionBehavior) Values() []NumberSelectionBehavior {
	return []NumberSelectionBehavior{
		"PreferSticky",
		"AvoidSticky",
	}
}

type OrderedPhoneNumberStatus string

// Enum values for OrderedPhoneNumberStatus
const (
	OrderedPhoneNumberStatusProcessing OrderedPhoneNumberStatus = "Processing"
	OrderedPhoneNumberStatusAcquired   OrderedPhoneNumberStatus = "Acquired"
	OrderedPhoneNumberStatusFailed     OrderedPhoneNumberStatus = "Failed"
)

// Values returns all known values for OrderedPhoneNumberStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OrderedPhoneNumberStatus) Values() []OrderedPhoneNumberStatus {
	return []OrderedPhoneNumberStatus{
		"Processing",
		"Acquired",
		"Failed",
	}
}

type OriginationRouteProtocol string

// Enum values for OriginationRouteProtocol
const (
	OriginationRouteProtocolTcp OriginationRouteProtocol = "TCP"
	OriginationRouteProtocolUdp OriginationRouteProtocol = "UDP"
)

// Values returns all known values for OriginationRouteProtocol. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OriginationRouteProtocol) Values() []OriginationRouteProtocol {
	return []OriginationRouteProtocol{
		"TCP",
		"UDP",
	}
}

type PhoneNumberAssociationName string

// Enum values for PhoneNumberAssociationName
const (
	PhoneNumberAssociationNameAccountid             PhoneNumberAssociationName = "AccountId"
	PhoneNumberAssociationNameUserid                PhoneNumberAssociationName = "UserId"
	PhoneNumberAssociationNameVoiceconnectorid      PhoneNumberAssociationName = "VoiceConnectorId"
	PhoneNumberAssociationNameVoiceconnectorgroupid PhoneNumberAssociationName = "VoiceConnectorGroupId"
)

// Values returns all known values for PhoneNumberAssociationName. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (PhoneNumberAssociationName) Values() []PhoneNumberAssociationName {
	return []PhoneNumberAssociationName{
		"AccountId",
		"UserId",
		"VoiceConnectorId",
		"VoiceConnectorGroupId",
	}
}

type PhoneNumberOrderStatus string

// Enum values for PhoneNumberOrderStatus
const (
	PhoneNumberOrderStatusProcessing PhoneNumberOrderStatus = "Processing"
	PhoneNumberOrderStatusSuccessful PhoneNumberOrderStatus = "Successful"
	PhoneNumberOrderStatusFailed     PhoneNumberOrderStatus = "Failed"
	PhoneNumberOrderStatusPartial    PhoneNumberOrderStatus = "Partial"
)

// Values returns all known values for PhoneNumberOrderStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PhoneNumberOrderStatus) Values() []PhoneNumberOrderStatus {
	return []PhoneNumberOrderStatus{
		"Processing",
		"Successful",
		"Failed",
		"Partial",
	}
}

type PhoneNumberProductType string

// Enum values for PhoneNumberProductType
const (
	PhoneNumberProductTypeBusinesscalling PhoneNumberProductType = "BusinessCalling"
	PhoneNumberProductTypeVoiceconnector  PhoneNumberProductType = "VoiceConnector"
)

// Values returns all known values for PhoneNumberProductType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PhoneNumberProductType) Values() []PhoneNumberProductType {
	return []PhoneNumberProductType{
		"BusinessCalling",
		"VoiceConnector",
	}
}

type PhoneNumberStatus string

// Enum values for PhoneNumberStatus
const (
	PhoneNumberStatusAcquireinprogress PhoneNumberStatus = "AcquireInProgress"
	PhoneNumberStatusAcquirefailed     PhoneNumberStatus = "AcquireFailed"
	PhoneNumberStatusUnassigned        PhoneNumberStatus = "Unassigned"
	PhoneNumberStatusAssigned          PhoneNumberStatus = "Assigned"
	PhoneNumberStatusReleaseinprogress PhoneNumberStatus = "ReleaseInProgress"
	PhoneNumberStatusDeleteinprogress  PhoneNumberStatus = "DeleteInProgress"
	PhoneNumberStatusReleasefailed     PhoneNumberStatus = "ReleaseFailed"
	PhoneNumberStatusDeletefailed      PhoneNumberStatus = "DeleteFailed"
)

// Values returns all known values for PhoneNumberStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PhoneNumberStatus) Values() []PhoneNumberStatus {
	return []PhoneNumberStatus{
		"AcquireInProgress",
		"AcquireFailed",
		"Unassigned",
		"Assigned",
		"ReleaseInProgress",
		"DeleteInProgress",
		"ReleaseFailed",
		"DeleteFailed",
	}
}

type PhoneNumberType string

// Enum values for PhoneNumberType
const (
	PhoneNumberTypeLocal    PhoneNumberType = "Local"
	PhoneNumberTypeTollfree PhoneNumberType = "TollFree"
)

// Values returns all known values for PhoneNumberType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PhoneNumberType) Values() []PhoneNumberType {
	return []PhoneNumberType{
		"Local",
		"TollFree",
	}
}

type ProxySessionStatus string

// Enum values for ProxySessionStatus
const (
	ProxySessionStatusOpen       ProxySessionStatus = "Open"
	ProxySessionStatusInprogress ProxySessionStatus = "InProgress"
	ProxySessionStatusClosed     ProxySessionStatus = "Closed"
)

// Values returns all known values for ProxySessionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ProxySessionStatus) Values() []ProxySessionStatus {
	return []ProxySessionStatus{
		"Open",
		"InProgress",
		"Closed",
	}
}

type RegistrationStatus string

// Enum values for RegistrationStatus
const (
	RegistrationStatusUnregistered RegistrationStatus = "Unregistered"
	RegistrationStatusRegistered   RegistrationStatus = "Registered"
	RegistrationStatusSuspended    RegistrationStatus = "Suspended"
)

// Values returns all known values for RegistrationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RegistrationStatus) Values() []RegistrationStatus {
	return []RegistrationStatus{
		"Unregistered",
		"Registered",
		"Suspended",
	}
}

type RoomMembershipRole string

// Enum values for RoomMembershipRole
const (
	RoomMembershipRoleAdministrator RoomMembershipRole = "Administrator"
	RoomMembershipRoleMember        RoomMembershipRole = "Member"
)

// Values returns all known values for RoomMembershipRole. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RoomMembershipRole) Values() []RoomMembershipRole {
	return []RoomMembershipRole{
		"Administrator",
		"Member",
	}
}

type UserType string

// Enum values for UserType
const (
	UserTypePrivateuser  UserType = "PrivateUser"
	UserTypeShareddevice UserType = "SharedDevice"
)

// Values returns all known values for UserType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (UserType) Values() []UserType {
	return []UserType{
		"PrivateUser",
		"SharedDevice",
	}
}

type VoiceConnectorAwsRegion string

// Enum values for VoiceConnectorAwsRegion
const (
	VoiceConnectorAwsRegionUsEast1 VoiceConnectorAwsRegion = "us-east-1"
	VoiceConnectorAwsRegionUsWest2 VoiceConnectorAwsRegion = "us-west-2"
)

// Values returns all known values for VoiceConnectorAwsRegion. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VoiceConnectorAwsRegion) Values() []VoiceConnectorAwsRegion {
	return []VoiceConnectorAwsRegion{
		"us-east-1",
		"us-west-2",
	}
}
