// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AttackLayer string

// Enum values for AttackLayer
const (
	AttackLayerNetwork     AttackLayer = "NETWORK"
	AttackLayerApplication AttackLayer = "APPLICATION"
)

// Values returns all known values for AttackLayer. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (AttackLayer) Values() []AttackLayer {
	return []AttackLayer{
		"NETWORK",
		"APPLICATION",
	}
}

type AttackPropertyIdentifier string

// Enum values for AttackPropertyIdentifier
const (
	AttackPropertyIdentifierDestinationUrl             AttackPropertyIdentifier = "DESTINATION_URL"
	AttackPropertyIdentifierReferrer                   AttackPropertyIdentifier = "REFERRER"
	AttackPropertyIdentifierSourceAsn                  AttackPropertyIdentifier = "SOURCE_ASN"
	AttackPropertyIdentifierSourceCountry              AttackPropertyIdentifier = "SOURCE_COUNTRY"
	AttackPropertyIdentifierSourceIpAddress            AttackPropertyIdentifier = "SOURCE_IP_ADDRESS"
	AttackPropertyIdentifierSourceUserAgent            AttackPropertyIdentifier = "SOURCE_USER_AGENT"
	AttackPropertyIdentifierWordpressPingbackReflector AttackPropertyIdentifier = "WORDPRESS_PINGBACK_REFLECTOR"
	AttackPropertyIdentifierWordpressPingbackSource    AttackPropertyIdentifier = "WORDPRESS_PINGBACK_SOURCE"
)

// Values returns all known values for AttackPropertyIdentifier. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AttackPropertyIdentifier) Values() []AttackPropertyIdentifier {
	return []AttackPropertyIdentifier{
		"DESTINATION_URL",
		"REFERRER",
		"SOURCE_ASN",
		"SOURCE_COUNTRY",
		"SOURCE_IP_ADDRESS",
		"SOURCE_USER_AGENT",
		"WORDPRESS_PINGBACK_REFLECTOR",
		"WORDPRESS_PINGBACK_SOURCE",
	}
}

type AutoRenew string

// Enum values for AutoRenew
const (
	AutoRenewEnabled  AutoRenew = "ENABLED"
	AutoRenewDisabled AutoRenew = "DISABLED"
)

// Values returns all known values for AutoRenew. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (AutoRenew) Values() []AutoRenew {
	return []AutoRenew{
		"ENABLED",
		"DISABLED",
	}
}

type ProactiveEngagementStatus string

// Enum values for ProactiveEngagementStatus
const (
	ProactiveEngagementStatusEnabled  ProactiveEngagementStatus = "ENABLED"
	ProactiveEngagementStatusDisabled ProactiveEngagementStatus = "DISABLED"
	ProactiveEngagementStatusPending  ProactiveEngagementStatus = "PENDING"
)

// Values returns all known values for ProactiveEngagementStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ProactiveEngagementStatus) Values() []ProactiveEngagementStatus {
	return []ProactiveEngagementStatus{
		"ENABLED",
		"DISABLED",
		"PENDING",
	}
}

type SubResourceType string

// Enum values for SubResourceType
const (
	SubResourceTypeIp  SubResourceType = "IP"
	SubResourceTypeUrl SubResourceType = "URL"
)

// Values returns all known values for SubResourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SubResourceType) Values() []SubResourceType {
	return []SubResourceType{
		"IP",
		"URL",
	}
}

type SubscriptionState string

// Enum values for SubscriptionState
const (
	SubscriptionStateActive   SubscriptionState = "ACTIVE"
	SubscriptionStateInactive SubscriptionState = "INACTIVE"
)

// Values returns all known values for SubscriptionState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SubscriptionState) Values() []SubscriptionState {
	return []SubscriptionState{
		"ACTIVE",
		"INACTIVE",
	}
}

type Unit string

// Enum values for Unit
const (
	UnitBits     Unit = "BITS"
	UnitBytes    Unit = "BYTES"
	UnitPackets  Unit = "PACKETS"
	UnitRequests Unit = "REQUESTS"
)

// Values returns all known values for Unit. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Unit) Values() []Unit {
	return []Unit{
		"BITS",
		"BYTES",
		"PACKETS",
		"REQUESTS",
	}
}
