// Package integrations takes the payload of third party services and converts
// them to a LeeroyCI job. At the same time it takes care of updating PRs.
package integrations

// Service defines all values a third party service has to provide to be
// compatible with LeeroyCI.
type Service interface {
	RepositoryURL() string
	Branch() string
	Commit() string
	CommitURL() string
	Name() string
	Email() string
	Run() bool
}
