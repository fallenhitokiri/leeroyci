// GitHub provides all structs to unmarshal a GitHub webhook.
package callbacks

import (
	"strings"
)

type GitHubCallback struct {
	Ref         string
	After       string
	Before      string
	Created     bool
	Deleted     bool
	Forced      bool
	Compare     string
	Commits     []Commit
	Head_commit Commit
	Repository  Repository
	Pusher      GitUser
}

type Commit struct {
	Id        string
	Distinct  bool
	Message   string
	Timestamp string
	Url       string
	Author    GitHubUser
	Committer GitHubUser
	Added     []string
	Removed   []string
	Modified  []string
}

type GitHubUser struct {
	Name     string
	Email    string
	Username string
}

type Repository struct {
	Id            int64
	Name          string
	Url           string
	Description   string
	Watchers      int
	Stargazers    int
	Forks         int
	Size          int
	Owner         GitUser
	Private       bool
	Open_issues   int
	Has_issues    bool
	Has_downloads bool
	Has_wiki      bool
	Created_at    int64
	Pushed_at     int64
	Master_branch string
}

type GitUser struct {
	Name  string
	Email string
}

// Branch returns the name of the branch.
func (g GitHubCallback) Branch() string {
	s := strings.Split(g.Ref, "/")
	return s[2]
}

// URL returns the URL for the repository
func (g GitHubCallback) URL() string {
	return g.Repository.Url
}
