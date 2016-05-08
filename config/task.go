// Package config contains all data models used for LeeroyCI.
package config

const (
	// TaskKindTest is used to represent tests.
	TaskKindTest = "test"

	// TaskKindBuild is used for tasks that build a binary / archive / ...
	TaskKindBuild = "build"

	// TaskKindDeploy is used for tasks that deploy the code / results of 'build'
	// to a server.
	TaskKindDeploy = "deploy"
)

// Task represents one task for a project to run when new code was pushed to
// the repository.
type Task struct {
	Name    string
	Kind    string
	Branch  string
	Execute string
}
