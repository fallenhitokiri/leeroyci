// Structs and methods used to process a push.
package github

import (
	"encoding/json"
	"ironman/logging"
	"log"
	"net/http"
	"strings"
	"time"
)

type PushCallback struct {
	Ref         string
	After       string
	Before      string
	Created     bool
	Deleted     bool
	Forced      bool
	Compare     string
	Commits     []PushCommit
	Head_commit PushCommit
	Repository  PushRepository
	Pusher      PushGitUser
}

type PushCommit struct {
	Id        string
	Distinct  bool
	Message   string
	Timestamp string
	Url       string
	Author    PushGitHubUser
	Committer PushGitHubUser
	Added     []string
	Removed   []string
	Modified  []string
}

type PushGitHubUser struct {
	Name     string
	Email    string
	Username string
}

type PushRepository struct {
	Id            int64
	Name          string
	Url           string
	Description   string
	Watchers      int
	Stargazers    int
	Forks         int
	Size          int
	Owner         PushGitUser
	Private       bool
	Open_issues   int
	Has_issues    bool
	Has_downloads bool
	Has_wiki      bool
	Created_at    int64
	Pushed_at     int64
	Master_branch string
}

type PushGitUser struct {
	Name  string
	Email string
}

// Branch returns the name of the branch.
func (p *PushCallback) Branch() string {
	s := strings.Split(p.Ref, "/")
	return s[2]
}

// URL returns the URL for the repository
func (p *PushCallback) URL() string {
	return p.Repository.Url
}

// By returns who pushed / triggered the callback. Format Name <email>.
func (p *PushCallback) By() (string, string) {
	return p.Pusher.Name, p.Pusher.Email
}

// CommitURL returns the URL to the head commit
func (p *PushCallback) CommitURL() string {
	return p.Head_commit.Url
}

// Returns if this commit should be build. Do not build if the branch was
// deleted for example.
func (p *PushCallback) ShouldBuild() bool {
	if p.Deleted == true {
		return false
	}
	return true
}

// Returns the ID of the head commit.
func (p *PushCallback) Commit() string {
	return p.Head_commit.Id
}

// Handle GitHub push events.
func handlePush(req *http.Request, jobs chan logging.Job) {
	b := parseBody(req)

	var cb PushCallback

	err := json.Unmarshal(b, &cb)
	if err != nil {
		log.Println(string(b))
		panic("Could not unmarshal request")
	}

	if cb.ShouldBuild() == true {
		pushToQueue(jobs, cb)
	} else {
		log.Println("Not adding", cb.URL(), cb.Branch(), "to build queue")
	}
}

// Convert a callback to a loggin.Job and push it to the build queue.
func pushToQueue(jobs chan logging.Job, cb PushCallback) {
	name, email := cb.By()

	j := logging.Job{
		URL:       cb.URL(),
		Branch:    cb.Branch(),
		Timestamp: time.Now(),
		Commit:    cb.Commit(),
		Name:      name,
		Email:     email,
		CommitURL: cb.CommitURL(),
	}

	jobs <- j
}
