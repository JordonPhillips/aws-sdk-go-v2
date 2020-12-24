// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type Status string

// Enum values for Status
const (
	StatusPass             Status = "PASS"
	StatusFail             Status = "FAIL"
	StatusCanceled         Status = "CANCELED"
	StatusPending          Status = "PENDING"
	StatusRunning          Status = "RUNNING"
	StatusPassWithWarnings Status = "PASS_WITH_WARNINGS"
	StatusError            Status = "ERROR"
)

// Values returns all known values for Status. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Status) Values() []Status {
	return []Status{
		"PASS",
		"FAIL",
		"CANCELED",
		"PENDING",
		"RUNNING",
		"PASS_WITH_WARNINGS",
		"ERROR",
	}
}

type SuiteRunStatus string

// Enum values for SuiteRunStatus
const (
	SuiteRunStatusPass             SuiteRunStatus = "PASS"
	SuiteRunStatusFail             SuiteRunStatus = "FAIL"
	SuiteRunStatusCanceled         SuiteRunStatus = "CANCELED"
	SuiteRunStatusPending          SuiteRunStatus = "PENDING"
	SuiteRunStatusRunning          SuiteRunStatus = "RUNNING"
	SuiteRunStatusPassWithWarnings SuiteRunStatus = "PASS_WITH_WARNINGS"
	SuiteRunStatusError            SuiteRunStatus = "ERROR"
)

// Values returns all known values for SuiteRunStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SuiteRunStatus) Values() []SuiteRunStatus {
	return []SuiteRunStatus{
		"PASS",
		"FAIL",
		"CANCELED",
		"PENDING",
		"RUNNING",
		"PASS_WITH_WARNINGS",
		"ERROR",
	}
}
