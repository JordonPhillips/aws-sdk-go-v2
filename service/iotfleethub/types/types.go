// Code generated by smithy-go-codegen DO NOT EDIT.

package types

// A summary of information about a AWS IoT Device Management web application.
// Fleet Hub for AWS IoT Device Management is in public preview and is subject to
// change.
type ApplicationSummary struct {

	// The unique Id of the web application.
	//
	// This member is required.
	ApplicationId *string

	// The name of the web application.
	//
	// This member is required.
	ApplicationName *string

	// The URL of the web application.
	//
	// This member is required.
	ApplicationUrl *string

	// The date (in Unix epoch time) when the web application was created.
	ApplicationCreationDate int64

	// An optional description of the web application.
	ApplicationDescription *string

	// The date (in Unix epoch time) when the web application was last updated.
	ApplicationLastUpdateDate int64

	// The current state of the web application.
	ApplicationState ApplicationState
}
