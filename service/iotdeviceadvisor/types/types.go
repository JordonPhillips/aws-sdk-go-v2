// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Lists all the devices under test
type DeviceUnderTest struct {

	// Lists devices certificate arn
	CertificateArn *string

	// Lists devices thing arn
	ThingArn *string
}

// Show Group Result.
type GroupResult struct {

	// Show Group Result Id.
	GroupId *string

	// Show Group Result Name.
	GroupName *string

	// Show Group Result.
	Tests []TestCaseRun
}

// Gets Suite Definition Configuration.
type SuiteDefinitionConfiguration struct {

	// Gets device permission arn.
	DevicePermissionRoleArn *string

	// Gets the devices configured.
	Devices []DeviceUnderTest

	// Gets the tests intended for qualification in a suite.
	IntendedForQualification bool

	// Gets test suite root group.
	RootGroup *string

	// Gets Suite Definition Configuration name.
	SuiteDefinitionName *string
}

// Get suite definition information.
type SuiteDefinitionInformation struct {

	// Gets the information of when the test suite was created.
	CreatedAt *time.Time

	// Specifies the devices under test.
	DefaultDevices []DeviceUnderTest

	// Gets the test suites which will be used for qualification.
	IntendedForQualification bool

	// Get suite definition Id.
	SuiteDefinitionId *string

	// Get test suite name.
	SuiteDefinitionName *string
}

// Gets suite run configuration.
type SuiteRunConfiguration struct {

	// Gets the primary device for suite run.
	PrimaryDevice *DeviceUnderTest

	// Gets the secondary device for suite run.
	SecondaryDevice *DeviceUnderTest

	// Gets test case list.
	SelectedTestList []string
}

// Get suite run information.
type SuiteRunInformation struct {

	// Get suite run information based on time suite was created.
	CreatedAt *time.Time

	// Get suite run information based on end time of the run.
	EndAt *time.Time

	// Get suite run information based on result of the test suite run.
	Failed int32

	// Get suite run information based on result of the test suite run.
	Passed int32

	// Get suite run information based on start time of the run.
	StartedAt *time.Time

	// Get suite run information based on test run status.
	Status SuiteRunStatus

	// Get suite run information based on suite definition Id.
	SuiteDefinitionId *string

	// Get suite run information based on suite definition name.
	SuiteDefinitionName *string

	// Get suite run information based on suite definition version.
	SuiteDefinitionVersion *string

	// Get suite run information based on suite run Id.
	SuiteRunId *string
}

// Shows tests in a test group.
type TestCase struct {

	// Shows test case configuration.
	Configuration map[string]string

	// Shows test case name.
	Name *string

	// Specifies a test.
	Test *TestCaseDefinition
}

// Gets the test case category.
type TestCaseCategory struct {

	// Lists all the tests name in the specified category.
	Name *string

	// Lists all the tests in the specified category.
	Tests []TestCase
}

// Provides test case definition.
type TestCaseDefinition struct {

	// Provides test case definition Id.
	Id *string

	// Provides test case definition version.
	TestCaseVersion *string
}

// Provides test case run.
type TestCaseRun struct {

	// Provides test case run end time.
	EndTime *time.Time

	// Provides test case run failure result.
	Failure *string

	// Provides test case run log Url.
	LogUrl *string

	// Provides test case run start time.
	StartTime *time.Time

	// Provides test case run status.
	Status Status

	// Provides test case run definition Id.
	TestCaseDefinitionId *string

	// Provides test case run definition Name.
	TestCaseDefinitionName *string

	// Provides test case run Id.
	TestCaseRunId *string

	// Provides test case run warnings.
	Warnings *string
}

// Show each group result.
type TestResult struct {

	// Show each group of test results.
	Groups []GroupResult
}
