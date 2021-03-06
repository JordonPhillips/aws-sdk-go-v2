// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type CustomerGatewayAssociationState string

// Enum values for CustomerGatewayAssociationState
const (
	CustomerGatewayAssociationStatePending   CustomerGatewayAssociationState = "PENDING"
	CustomerGatewayAssociationStateAvailable CustomerGatewayAssociationState = "AVAILABLE"
	CustomerGatewayAssociationStateDeleting  CustomerGatewayAssociationState = "DELETING"
	CustomerGatewayAssociationStateDeleted   CustomerGatewayAssociationState = "DELETED"
)

// Values returns all known values for CustomerGatewayAssociationState. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (CustomerGatewayAssociationState) Values() []CustomerGatewayAssociationState {
	return []CustomerGatewayAssociationState{
		"PENDING",
		"AVAILABLE",
		"DELETING",
		"DELETED",
	}
}

type DeviceState string

// Enum values for DeviceState
const (
	DeviceStatePending   DeviceState = "PENDING"
	DeviceStateAvailable DeviceState = "AVAILABLE"
	DeviceStateDeleting  DeviceState = "DELETING"
	DeviceStateUpdating  DeviceState = "UPDATING"
)

// Values returns all known values for DeviceState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (DeviceState) Values() []DeviceState {
	return []DeviceState{
		"PENDING",
		"AVAILABLE",
		"DELETING",
		"UPDATING",
	}
}

type GlobalNetworkState string

// Enum values for GlobalNetworkState
const (
	GlobalNetworkStatePending   GlobalNetworkState = "PENDING"
	GlobalNetworkStateAvailable GlobalNetworkState = "AVAILABLE"
	GlobalNetworkStateDeleting  GlobalNetworkState = "DELETING"
	GlobalNetworkStateUpdating  GlobalNetworkState = "UPDATING"
)

// Values returns all known values for GlobalNetworkState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (GlobalNetworkState) Values() []GlobalNetworkState {
	return []GlobalNetworkState{
		"PENDING",
		"AVAILABLE",
		"DELETING",
		"UPDATING",
	}
}

type LinkAssociationState string

// Enum values for LinkAssociationState
const (
	LinkAssociationStatePending   LinkAssociationState = "PENDING"
	LinkAssociationStateAvailable LinkAssociationState = "AVAILABLE"
	LinkAssociationStateDeleting  LinkAssociationState = "DELETING"
	LinkAssociationStateDeleted   LinkAssociationState = "DELETED"
)

// Values returns all known values for LinkAssociationState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LinkAssociationState) Values() []LinkAssociationState {
	return []LinkAssociationState{
		"PENDING",
		"AVAILABLE",
		"DELETING",
		"DELETED",
	}
}

type LinkState string

// Enum values for LinkState
const (
	LinkStatePending   LinkState = "PENDING"
	LinkStateAvailable LinkState = "AVAILABLE"
	LinkStateDeleting  LinkState = "DELETING"
	LinkStateUpdating  LinkState = "UPDATING"
)

// Values returns all known values for LinkState. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (LinkState) Values() []LinkState {
	return []LinkState{
		"PENDING",
		"AVAILABLE",
		"DELETING",
		"UPDATING",
	}
}

type SiteState string

// Enum values for SiteState
const (
	SiteStatePending   SiteState = "PENDING"
	SiteStateAvailable SiteState = "AVAILABLE"
	SiteStateDeleting  SiteState = "DELETING"
	SiteStateUpdating  SiteState = "UPDATING"
)

// Values returns all known values for SiteState. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (SiteState) Values() []SiteState {
	return []SiteState{
		"PENDING",
		"AVAILABLE",
		"DELETING",
		"UPDATING",
	}
}

type TransitGatewayRegistrationState string

// Enum values for TransitGatewayRegistrationState
const (
	TransitGatewayRegistrationStatePending   TransitGatewayRegistrationState = "PENDING"
	TransitGatewayRegistrationStateAvailable TransitGatewayRegistrationState = "AVAILABLE"
	TransitGatewayRegistrationStateDeleting  TransitGatewayRegistrationState = "DELETING"
	TransitGatewayRegistrationStateDeleted   TransitGatewayRegistrationState = "DELETED"
	TransitGatewayRegistrationStateFailed    TransitGatewayRegistrationState = "FAILED"
)

// Values returns all known values for TransitGatewayRegistrationState. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (TransitGatewayRegistrationState) Values() []TransitGatewayRegistrationState {
	return []TransitGatewayRegistrationState{
		"PENDING",
		"AVAILABLE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonUnknownOperation      ValidationExceptionReason = "UnknownOperation"
	ValidationExceptionReasonCannotParse           ValidationExceptionReason = "CannotParse"
	ValidationExceptionReasonFieldValidationFailed ValidationExceptionReason = "FieldValidationFailed"
	ValidationExceptionReasonOther                 ValidationExceptionReason = "Other"
)

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"UnknownOperation",
		"CannotParse",
		"FieldValidationFailed",
		"Other",
	}
}
