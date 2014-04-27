package callbacks

import (
	"encoding/json"
	"testing"
)

func TestJsonUnmarshal(t *testing.T) {
	payload := []byte(`{ "ref":"refs/heads/master", "after":"2c7f7accbcf73b7b4c98ee7e1f213eb46a885042", "before":"b369a904428a1d5d671e3f740443590d3db55fb0", "created":false, "deleted":false, "forced":false, "compare":"https://github.com/fallenhitokiri/pushtest/compare/b369a904428a...2c7f7accbcf7", "commits":[ { "id":"2c7f7accbcf73b7b4c98ee7e1f213eb46a885042", "distinct":true, "message":"asdf", "timestamp":"2014-04-27T11:53:14+02:00", "url":"https://github.com/fallenhitokiri/pushtest/commit/2c7f7accbcf73b7b4c98ee7e1f213eb46a885042", "author": { "name":"Timo Zimmermann", "email":"timo@screamingatmyscreen.com", "username":"fallenhitokiri" }, "committer":{ "name":"Timo Zimmermann", "email":"timo@screamingatmyscreen.com", "username":"fallenhitokiri" }, "added":[ "test2.md" ], "removed": [ ], "modified": [ ] } ], "head_commit": { "id":"2c7f7accbcf73b7b4c98ee7e1f213eb46a885042", "distinct":true, "message":"asdf", "timestamp":"2014-04-27T11:53:14+02:00", "url":"https://github.com/fallenhitokiri/pushtest/commit/2c7f7accbcf73b7b4c98ee7e1f213eb46a885042", "author": { "name":"Timo Zimmermann", "email":"timo@screamingatmyscreen.com", "username":"fallenhitokiri" }, "committer": { "name":"Timo Zimmermann", "email":"timo@screamingatmyscreen.com", "username":"fallenhitokiri" }, "added": [ "test2.md" ], "removed":[ ], "modified":[ ] }, "repository":{ "id":19200766, "name":"pushtest", "url":"https://github.com/fallenhitokiri/pushtest", "description":"nothing to see here", "watchers":0, "stargazers":0, "forks":0, "fork":false, "size":0, "owner":{ "name":"fallenhitokiri", "email":"timo@screamingatmyscreen.com" }, "private":false, "open_issues":0, "has_issues":true, "has_downloads":true, "has_wiki":true, "created_at":1398591886, "pushed_at":1398592401, "master_branch":"master" }, "pusher": { "name":"fallenhitokiri", "email":"timo@screamingatmyscreen.com"} }`)
	var cb GitHubCallback
	err := json.Unmarshal(payload, &cb)

	if err != nil {
		t.Error(err)
	}

	if cb.Pusher.Name != "fallenhitokiri" {
		t.Error("wrong pusher")
	}

	if cb.Commits[0].Author.Email != "timo@screamingatmyscreen.com" {
		t.Error("wrong email")
	}

	// TODO: test all fields
}

func TestGitHubCallbackBranch(t *testing.T) {
	var cb GitHubCallback

	cb = GitHubCallback{
		Ref: "refs/heads/master",
	}

	if cb.Branch() != "master" {
		t.Error("Wrong branch name", cb.Branch())
	}
}

func TestGitHubCallbackURL(t *testing.T) {
	var cb GitHubCallback
	var r Repository

	r = Repository{
		Url: "foobar",
	}

	cb = GitHubCallback{
		Repository: r,
	}

	if cb.URL() != "foobar" {
		t.Error("Wrong URL", cb.URL())
	}
}
