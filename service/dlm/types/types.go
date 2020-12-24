// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Specifies an action for an event-based policy.
type Action struct {

	// The rule for copying shared snapshots across Regions.
	//
	// This member is required.
	CrossRegionCopy []CrossRegionCopyAction

	// A descriptive name for the action.
	//
	// This member is required.
	Name *string
}

// Specifies when to create snapshots of EBS volumes. You must specify either a
// Cron expression or an interval, interval unit, and start time. You cannot
// specify both.
type CreateRule struct {

	// The schedule, as a Cron expression. The schedule interval must be between 1 hour
	// and 1 year. For more information, see Cron expressions
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html#CronExpressions)
	// in the Amazon CloudWatch User Guide.
	CronExpression *string

	// The interval between snapshots. The supported values are 1, 2, 3, 4, 6, 8, 12,
	// and 24.
	Interval int32

	// The interval unit.
	IntervalUnit IntervalUnitValues

	// The time, in UTC, to start the operation. The supported format is hh:mm. The
	// operation occurs within a one-hour window following the specified time. If you
	// do not specify a time, Amazon DLM selects a time within the next 24 hours.
	Times []string
}

// Specifies a rule for copying shared snapshots across Regions.
type CrossRegionCopyAction struct {

	// The encryption settings for the copied snapshot.
	//
	// This member is required.
	EncryptionConfiguration *EncryptionConfiguration

	// The target Region.
	//
	// This member is required.
	Target *string

	// Specifies the retention rule for cross-Region snapshot copies.
	RetainRule *CrossRegionCopyRetainRule
}

// Specifies the retention rule for cross-Region snapshot copies.
type CrossRegionCopyRetainRule struct {

	// The amount of time to retain each snapshot. The maximum is 100 years. This is
	// equivalent to 1200 months, 5200 weeks, or 36500 days.
	Interval int32

	// The unit of time for time-based retention.
	IntervalUnit RetentionIntervalUnitValues
}

// Specifies a rule for cross-Region snapshot copies.
type CrossRegionCopyRule struct {

	// To encrypt a copy of an unencrypted snapshot if encryption by default is not
	// enabled, enable encryption using this parameter. Copies of encrypted snapshots
	// are encrypted, even if this parameter is false or if encryption by default is
	// not enabled.
	//
	// This member is required.
	Encrypted bool

	// The target Region.
	//
	// This member is required.
	TargetRegion *string

	// The Amazon Resource Name (ARN) of the AWS KMS customer master key (CMK) to use
	// for EBS encryption. If this parameter is not specified, your AWS managed CMK for
	// EBS is used.
	CmkArn *string

	// Copy all user-defined tags from the source snapshot to the copied snapshot.
	CopyTags bool

	// The retention rule.
	RetainRule *CrossRegionCopyRetainRule
}

// Specifies the encryption settings for shared snapshots that are copied across
// Regions.
type EncryptionConfiguration struct {

	// To encrypt a copy of an unencrypted snapshot when encryption by default is not
	// enabled, enable encryption using this parameter. Copies of encrypted snapshots
	// are encrypted, even if this parameter is false or when encryption by default is
	// not enabled.
	//
	// This member is required.
	Encrypted bool

	// The Amazon Resource Name (ARN) of the AWS KMS customer master key (CMK) to use
	// for EBS encryption. If this parameter is not specified, your AWS managed CMK for
	// EBS is used.
	CmkArn *string
}

// Specifies an event that triggers an event-based policy.
type EventParameters struct {

	// The snapshot description that can trigger the policy. The description pattern is
	// specified using a regular expression. The policy runs only if a snapshot with a
	// description that matches the specified pattern is shared with your account. For
	// example, specifying ^.*Created for policy: policy-1234567890abcdef0.*$
	// configures the policy to run only if snapshots created by policy
	// policy-1234567890abcdef0 are shared with your account.
	//
	// This member is required.
	DescriptionRegex *string

	// The type of event. Currently, only snapshot sharing events are supported.
	//
	// This member is required.
	EventType EventTypeValues

	// The IDs of the AWS accounts that can trigger policy by sharing snapshots with
	// your account. The policy only runs if one of the specified AWS accounts shares a
	// snapshot with your account.
	//
	// This member is required.
	SnapshotOwner []string
}

// Specifies an event that triggers an event-based policy.
type EventSource struct {

	// The source of the event. Currently only managed AWS CloudWatch Events rules are
	// supported.
	//
	// This member is required.
	Type EventSourceValues

	// Information about the event.
	Parameters *EventParameters
}

// Specifies a rule for enabling fast snapshot restore. You can enable fast
// snapshot restore based on either a count or a time interval.
type FastRestoreRule struct {

	// The Availability Zones in which to enable fast snapshot restore.
	//
	// This member is required.
	AvailabilityZones []string

	// The number of snapshots to be enabled with fast snapshot restore.
	Count int32

	// The amount of time to enable fast snapshot restore. The maximum is 100 years.
	// This is equivalent to 1200 months, 5200 weeks, or 36500 days.
	Interval int32

	// The unit of time for enabling fast snapshot restore.
	IntervalUnit RetentionIntervalUnitValues
}

// Detailed information about a lifecycle policy.
type LifecyclePolicy struct {

	// The local date and time when the lifecycle policy was created.
	DateCreated *time.Time

	// The local date and time when the lifecycle policy was last modified.
	DateModified *time.Time

	// The description of the lifecycle policy.
	Description *string

	// The Amazon Resource Name (ARN) of the IAM role used to run the operations
	// specified by the lifecycle policy.
	ExecutionRoleArn *string

	// The Amazon Resource Name (ARN) of the policy.
	PolicyArn *string

	// The configuration of the lifecycle policy
	PolicyDetails *PolicyDetails

	// The identifier of the lifecycle policy.
	PolicyId *string

	// The activation state of the lifecycle policy.
	State GettablePolicyStateValues

	// The description of the status.
	StatusMessage *string

	// The tags.
	Tags map[string]string
}

// Summary information about a lifecycle policy.
type LifecyclePolicySummary struct {

	// The description of the lifecycle policy.
	Description *string

	// The identifier of the lifecycle policy.
	PolicyId *string

	// The type of policy. EBS_SNAPSHOT_MANAGEMENT indicates that the policy manages
	// the lifecycle of Amazon EBS snapshots. IMAGE_MANAGEMENT indicates that the
	// policy manages the lifecycle of EBS-backed AMIs.
	PolicyType PolicyTypeValues

	// The activation state of the lifecycle policy.
	State GettablePolicyStateValues

	// The tags.
	Tags map[string]string
}

// Specifies optional parameters to add to a policy. The set of valid parameters
// depends on the combination of policy type and resource type.
type Parameters struct {

	// [EBS Snapshot Management – Instance policies only] Indicates whether to exclude
	// the root volume from snapshots created using CreateSnapshots
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateSnapshots.html).
	// The default is false.
	ExcludeBootVolume bool

	// Applies to AMI lifecycle policies only. Indicates whether targeted instances are
	// rebooted when the lifecycle policy runs. true indicates that targeted instances
	// are not rebooted when the policy runs. false indicates that target instances are
	// rebooted when the policy runs. The default is true (instances are not rebooted).
	NoReboot bool
}

// Specifies the configuration of a lifecycle policy.
type PolicyDetails struct {

	// The actions to be performed when the event-based policy is triggered. You can
	// specify only one action per policy. This parameter is required for event-based
	// policies only. If you are creating a snapshot or AMI policy, omit this
	// parameter.
	Actions []Action

	// The event that triggers the event-based policy. This parameter is required for
	// event-based policies only. If you are creating a snapshot or AMI policy, omit
	// this parameter.
	EventSource *EventSource

	// A set of optional parameters for snapshot and AMI lifecycle policies. This
	// parameter is required for snapshot and AMI policies only. If you are creating an
	// event-based policy, omit this parameter.
	Parameters *Parameters

	// The valid target resource types and actions a policy can manage. Specify
	// EBS_SNAPSHOT_MANAGEMENT to create a lifecycle policy that manages the lifecycle
	// of Amazon EBS snapshots. Specify IMAGE_MANAGEMENT to create a lifecycle policy
	// that manages the lifecycle of EBS-backed AMIs. Specify EVENT_BASED_POLICY  to
	// create an event-based policy that performs specific actions when a defined event
	// occurs in your AWS account. The default is EBS_SNAPSHOT_MANAGEMENT.
	PolicyType PolicyTypeValues

	// The target resource type for snapshot and AMI lifecycle policies. Use VOLUME to
	// create snapshots of individual volumes or use INSTANCE to create multi-volume
	// snapshots from the volumes for an instance. This parameter is required for
	// snapshot and AMI policies only. If you are creating an event-based policy, omit
	// this parameter.
	ResourceTypes []ResourceTypeValues

	// The schedules of policy-defined actions for snapshot and AMI lifecycle policies.
	// A policy can have up to four schedules—one mandatory schedule and up to three
	// optional schedules. This parameter is required for snapshot and AMI policies
	// only. If you are creating an event-based policy, omit this parameter.
	Schedules []Schedule

	// The single tag that identifies targeted resources for this policy. This
	// parameter is required for snapshot and AMI policies only. If you are creating an
	// event-based policy, omit this parameter.
	TargetTags []Tag
}

// Specifies the retention rule for a lifecycle policy. You can retain snapshots
// based on either a count or a time interval.
type RetainRule struct {

	// The number of snapshots to retain for each volume, up to a maximum of 1000.
	Count int32

	// The amount of time to retain each snapshot. The maximum is 100 years. This is
	// equivalent to 1200 months, 5200 weeks, or 36500 days.
	Interval int32

	// The unit of time for time-based retention.
	IntervalUnit RetentionIntervalUnitValues
}

// Specifies a backup schedule for a snapshot or AMI lifecycle policy.
type Schedule struct {

	// Copy all user-defined tags on a source volume to snapshots of the volume created
	// by this policy.
	CopyTags bool

	// The creation rule.
	CreateRule *CreateRule

	// The rule for cross-Region snapshot copies.
	CrossRegionCopyRules []CrossRegionCopyRule

	// The rule for enabling fast snapshot restore.
	FastRestoreRule *FastRestoreRule

	// The name of the schedule.
	Name *string

	// The retention rule.
	RetainRule *RetainRule

	// The rule for sharing snapshots with other AWS accounts.
	ShareRules []ShareRule

	// The tags to apply to policy-created resources. These user-defined tags are in
	// addition to the AWS-added lifecycle tags.
	TagsToAdd []Tag

	// A collection of key/value pairs with values determined dynamically when the
	// policy is executed. Keys may be any valid Amazon EC2 tag key. Values must be in
	// one of the two following formats: $(instance-id) or $(timestamp). Variable tags
	// are only valid for EBS Snapshot Management – Instance policies.
	VariableTags []Tag
}

// Specifies a rule for sharing snapshots across AWS accounts.
type ShareRule struct {

	// The IDs of the AWS accounts with which to share the snapshots.
	//
	// This member is required.
	TargetAccounts []string

	// The period after which snapshots that are shared with other AWS accounts are
	// automatically unshared.
	UnshareInterval int32

	// The unit of time for the automatic unsharing interval.
	UnshareIntervalUnit RetentionIntervalUnitValues
}

// Specifies a tag for a resource.
type Tag struct {

	// The tag key.
	//
	// This member is required.
	Key *string

	// The tag value.
	//
	// This member is required.
	Value *string
}
