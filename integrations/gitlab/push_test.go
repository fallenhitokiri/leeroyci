package gitlab

import (
	"encoding/json"
	"testing"
)

var payload = []byte(`{ "before": "95790bf891e76fee5e1747ab589903a6a1f80f22", "after": "da1560886d4f094c3e6c9ef40349f7d38b5d27d7", "ref": "refs/heads/master", "user_id": 4, "user_name": "John Smith", "project_id": 15, "repository": { "name": "Diaspora", "url": "git@example.com:diaspora.git", "description": "", "homepage": "http://example.com/diaspora" }, "commits": [ { "id": "b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327", "message": "Update Catalan translation to e38cb41.", "timestamp": "2011-12-12T14:27:31+02:00", "url": "http://example.com/diaspora/commits/b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327", "author": { "name": "Jordi Mallach", "email": "jordi@softcatala.org" } }, { "id": "da1560886d4f094c3e6c9ef40349f7d38b5d27d7", "message": "fixed readme", "timestamp": "2012-01-03T23:36:29+02:00", "url": "http://example.com/diaspora/commits/da1560886d4f094c3e6c9ef40349f7d38b5d27d7", "author": { "name": "GitLab dev user", "email": "gitlabdev@dv6700.(none)" } } ], "total_commits_count": 4}`)

func TestUnmarshal(t *testing.T) {
	var pc PushCallback
	json.Unmarshal(payload, &pc)

	if len(pc.Commits) != 2 {
		t.Error("wrong number of commits", len(pc.Commits))
	}

	if pc.ProjectId != 15 {
		t.Error("wrong project id", pc.ProjectId)
	}
}

func TestBranch(t *testing.T) {
	var pc PushCallback
	json.Unmarshal(payload, &pc)

	if pc.Branch() != "da1560886d4f094c3e6c9ef40349f7d38b5d27d7" {
		t.Error("wrong branch", pc.Branch())
	}
}

func TestURL(t *testing.T) {
	var pc PushCallback
	json.Unmarshal(payload, &pc)

	if pc.URL() != "git@example.com:diaspora.git" {
		t.Error("wrong URL", pc.URL())
	}
}

func TestBy(t *testing.T) {
	var pc PushCallback
	json.Unmarshal(payload, &pc)

	n, e := pc.By()

	if n != "GitLab dev user" {
		t.Error("wrong name", n)
	}

	if e != "gitlabdev@dv6700.(none)" {
		t.Error("wrong email", e)
	}
}

func TestCommitURL(t *testing.T) {
	var pc PushCallback
	json.Unmarshal(payload, &pc)

	if pc.CommitURL() != "http://example.com/diaspora/commits/da1560886d4f094c3e6c9ef40349f7d38b5d27d7" {
		t.Error("wrong URL", pc.CommitURL())
	}
}

func TestShouldBuild(t *testing.T) {
	var pc PushCallback
	json.Unmarshal(payload, &pc)

	if pc.ShouldBuild() != true {
		t.Error("wrong ShouldBuild", pc.ShouldBuild())
	}
}

func TestCommit(t *testing.T) {
	var pc PushCallback
	json.Unmarshal(payload, &pc)

	if pc.Commit() != "da1560886d4f094c3e6c9ef40349f7d38b5d27d7" {
		t.Error("wrong commit", pc.Commit())
	}
}

func TestLastCommit(t *testing.T) {
	var pc PushCallback
	json.Unmarshal(payload, &pc)

	if pc.lastCommit().Id != "da1560886d4f094c3e6c9ef40349f7d38b5d27d7" {
		t.Error("wrong commit", pc.lastCommit().Id)
	}
}
