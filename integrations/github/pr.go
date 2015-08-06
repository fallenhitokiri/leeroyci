// Package github integrates everything necessary to test commits, comment on
// pull requests and close them if the build failed.
package github

type pullRequestCallback struct {
	Number int
	Action string
	PR     pullRequest `json:"pull_request"`
}

type pullRequest struct {
	URL         string `json:"url"`
	State       string
	CommentsURL string `json:"comments_url"`
	StatusURL   string `json:"statuses_url"`
	Head        pullRequestCommit
}

type pullRequestCommit struct {
	Commit     string                `json:"sha"`
	Repository pullRequestRepository `json:"repo"`
}

type pullRequestRepository struct {
	HTMLURL string `json:"html_url"`
}

func (p *pullRequestCallback) repositoryURL() string {
	return p.PR.Head.Repository.HTMLURL
}
