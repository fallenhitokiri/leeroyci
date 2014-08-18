// Structs and methods to process a push.
package gitlab

import (
	"encoding/json"
	"leeroy/logging"
	"log"
	"time"
)

type PushCallback struct {
	Before     string
	After      string
	ref        string
	UserID     int    `json:"user_id"`
	UserName   string `json:"user_name"`
	ProjectId  int    `json:"project_id"`
	Repository Repository
	Commits    []Commit
	Total      int `json:"total_commits_count"`
}

type Repository struct {
	Name        string
	Url         string
	Description string
	Homepage    string
}

type Commit struct {
	Id        string
	Message   string
	Timestamp string
	Url       string
	Author    User
}

type User struct {
	Name  string
	Email string
}

// Branch returns the after commit id. GitLab does not push the branch
// as identifier, so we use the commit to identify this push.
func (p *PushCallback) Branch() string {
	return p.After
}

// URL returns the URL for the repositiory.
func (p *PushCallback) URL() string {
	return p.Repository.Url
}

// By returns who pushed (name and email) / triggered the callback.
func (p *PushCallback) By() (string, string) {
	l := p.lastCommit()
	return l.Author.Name, l.Author.Email
}

// CommitURL returns the URL to the head commit.
func (p *PushCallback) CommitURL() string {
	l := p.lastCommit()
	return l.Url
}

// Returns if this commit should be build.
func (p *PushCallback) ShouldBuild() bool {
	return true
}

// Returns the ID of the head commit.
func (p *PushCallback) Commit() string {
	return p.After
}

// Returns the last commit.
func (p *PushCallback) lastCommit() Commit {
	var fb Commit // empty fallback commit in case commit is not found

	for _, c := range p.Commits {
		if c.Id == p.After {
			return c
		}
	}

	return fb
}

// Handle GitLab push events.
func handlePush(body []byte, jobs chan logging.Job) {
	var pc PushCallback

	err := json.Unmarshal(body, &pc)
	if err != nil {
		log.Println(string(body))
		panic("Could not unmarshal request")
	}

	if pc.ShouldBuild() == true {
		pushToQueue(jobs, pc)
	} else {
		log.Println("Not adding", pc.URL(), pc.Branch(), "to build queue")
	}
}

// Convert a callback to a job and push it to the build queue.
func pushToQueue(jobs chan logging.Job, pc PushCallback) {
	n, e := pc.By()

	j := logging.Job{
		URL:       pc.URL(),
		Branch:    pc.Branch(),
		Timestamp: time.Now(),
		Commit:    pc.Commit(),
		Name:      n,
		Email:     e,
		CommitURL: pc.CommitURL(),
	}

	jobs <- j
}
