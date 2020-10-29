// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type JobState string

// Enum values for JobState
const (
	JobStateCompleted JobState = "Completed"
	JobStatePending   JobState = "Pending"
	JobStateFailed    JobState = "Failed"
	JobStateDeleting  JobState = "Deleting"
)

// Values returns all known values for JobState. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (JobState) Values() []JobState {
	return []JobState{
		"Completed",
		"Pending",
		"Failed",
		"Deleting",
	}
}

type ProviderType string

// Enum values for ProviderType
const (
	ProviderTypeCodeCommit             ProviderType = "CodeCommit"
	ProviderTypeGitHub                 ProviderType = "GitHub"
	ProviderTypeBitbucket              ProviderType = "Bitbucket"
	ProviderTypeGitHubEnterpriseServer ProviderType = "GitHubEnterpriseServer"
)

// Values returns all known values for ProviderType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ProviderType) Values() []ProviderType {
	return []ProviderType{
		"CodeCommit",
		"GitHub",
		"Bitbucket",
		"GitHubEnterpriseServer",
	}
}

type Reaction string

// Enum values for Reaction
const (
	ReactionThumbsUp   Reaction = "ThumbsUp"
	ReactionThumbsDown Reaction = "ThumbsDown"
)

// Values returns all known values for Reaction. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Reaction) Values() []Reaction {
	return []Reaction{
		"ThumbsUp",
		"ThumbsDown",
	}
}

type RepositoryAssociationState string

// Enum values for RepositoryAssociationState
const (
	RepositoryAssociationStateAssociated     RepositoryAssociationState = "Associated"
	RepositoryAssociationStateAssociating    RepositoryAssociationState = "Associating"
	RepositoryAssociationStateFailed         RepositoryAssociationState = "Failed"
	RepositoryAssociationStateDisassociating RepositoryAssociationState = "Disassociating"
)

// Values returns all known values for RepositoryAssociationState. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (RepositoryAssociationState) Values() []RepositoryAssociationState {
	return []RepositoryAssociationState{
		"Associated",
		"Associating",
		"Failed",
		"Disassociating",
	}
}

type Type string

// Enum values for Type
const (
	TypePullRequest        Type = "PullRequest"
	TypeRepositoryAnalysis Type = "RepositoryAnalysis"
)

// Values returns all known values for Type. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Type) Values() []Type {
	return []Type{
		"PullRequest",
		"RepositoryAnalysis",
	}
}
