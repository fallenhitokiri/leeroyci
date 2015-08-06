// Package github integrates everything necessary to test commits, comment on
// pull requests and close them if the build failed.
package github

import "strings"

type pushCallback struct {
	Ref        string
	After      string
	Before     string
	Created    bool
	Deleted    bool
	Forced     bool
	Compare    string
	Commits    []pushCommit
	HeadCommit pushCommit `json:"head_commit"`
	Repository pushRepository
	Pusher     pushUser
}

type pushCommit struct {
	ID        string
	Distinct  bool
	Message   string
	Timestamp string
	URL       string
	Author    pushUser
	Committer pushUser
	Added     []string
	Removed   []string
	Modified  []string
}

type pushUser struct {
	Name  string
	Email string
}

type pushRepository struct {
	ID          int64
	Name        string
	URL         string
	Description string
	CreatedAt   int64 `json:"created_at"`
	PushedAt    int64 `json:"pushed_at"`
}

// repositoryURL returns the URL for the repository
func (p *pushCallback) repositoryURL() string {
	return p.Repository.URL
}

// branch returns the name of the branch.
func (p *pushCallback) branch() string {
	s := strings.Split(p.Ref, "/")
	return s[2]
}

// returns the ID of the head commit.
func (p *pushCallback) commit() string {
	return p.HeadCommit.ID
}

// commitURL returns the URL to the head commit.
func (p *pushCallback) commitURL() string {
	return p.HeadCommit.URL
}

// name returns the name of the git user.
func (p *pushCallback) name() string {
	return p.Pusher.Name
}

// email returns the email of the git user.
func (p *pushCallback) email() string {
	return p.Pusher.Email
}

// shouldRun returns if this push should create a job.
func (p *pushCallback) shouldRun() bool {
	if p.Deleted == true {
		return false
	}
	return true
}
