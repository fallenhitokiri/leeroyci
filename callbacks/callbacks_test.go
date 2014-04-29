package callbacks

import (
	"testing"
)

func TestParse(t *testing.T) {
	not := make(chan Notification, 5)
	body1 := []byte(`{ "ref":"refs/heads/master", "after":"2c7f7accbcf73b7b4c98ee7e1f213eb46a885042", "before":"b369a904428a1d5d671e3f740443590d3db55fb0", "created":false, "deleted":false, "forced":false, "compare":"https://github.com/fallenhitokiri/pushtest/compare/b369a904428a...2c7f7accbcf7", "commits":[ { "id":"2c7f7accbcf73b7b4c98ee7e1f213eb46a885042", "distinct":true, "message":"asdf", "timestamp":"2014-04-27T11:53:14+02:00", "url":"https://github.com/fallenhitokiri/pushtest/commit/2c7f7accbcf73b7b4c98ee7e1f213eb46a885042", "author": { "name":"Timo Zimmermann", "email":"timo@screamingatmyscreen.com", "username":"fallenhitokiri" }, "committer":{ "name":"Timo Zimmermann", "email":"timo@screamingatmyscreen.com", "username":"fallenhitokiri" }, "added":[ "test2.md" ], "removed": [ ], "modified": [ ] } ], "head_commit": { "id":"2c7f7accbcf73b7b4c98ee7e1f213eb46a885042", "distinct":true, "message":"asdf", "timestamp":"2014-04-27T11:53:14+02:00", "url":"https://github.com/fallenhitokiri/pushtest/commit/2c7f7accbcf73b7b4c98ee7e1f213eb46a885042", "author": { "name":"Timo Zimmermann", "email":"timo@screamingatmyscreen.com", "username":"fallenhitokiri" }, "committer": { "name":"Timo Zimmermann", "email":"timo@screamingatmyscreen.com", "username":"fallenhitokiri" }, "added": [ "test2.md" ], "removed":[ ], "modified":[ ] }, "repository":{ "id":19200766, "name":"pushtest", "url":"https://github.com/fallenhitokiri/pushtest", "description":"nothing to see here", "watchers":0, "stargazers":0, "forks":0, "fork":false, "size":0, "owner":{ "name":"fallenhitokiri", "email":"timo@screamingatmyscreen.com" }, "private":false, "open_issues":0, "has_issues":true, "has_downloads":true, "has_wiki":true, "created_at":1398591886, "pushed_at":1398592401, "master_branch":"master" }, "pusher": { "name":"fallenhitokiri", "email":"timo@screamingatmyscreen.com"} }`)
	body2 := []byte(`{ "ref":"refs/heads/master", "after":"2c7f7accbcf73b7b4c98ee7e1f213eb46a885042", "before":"b369a904428a1d5d671e3f740443590d3db55fb0", "created":false, "deleted":true, "forced":false, "compare":"https://github.com/fallenhitokiri/pushtest/compare/b369a904428a...2c7f7accbcf7", "commits":[ { "id":"2c7f7accbcf73b7b4c98ee7e1f213eb46a885042", "distinct":true, "message":"asdf", "timestamp":"2014-04-27T11:53:14+02:00", "url":"https://github.com/fallenhitokiri/pushtest/commit/2c7f7accbcf73b7b4c98ee7e1f213eb46a885042", "author": { "name":"Timo Zimmermann", "email":"timo@screamingatmyscreen.com", "username":"fallenhitokiri" }, "committer":{ "name":"Timo Zimmermann", "email":"timo@screamingatmyscreen.com", "username":"fallenhitokiri" }, "added":[ "test2.md" ], "removed": [ ], "modified": [ ] } ], "head_commit": { "id":"2c7f7accbcf73b7b4c98ee7e1f213eb46a885042", "distinct":true, "message":"asdf", "timestamp":"2014-04-27T11:53:14+02:00", "url":"https://github.com/fallenhitokiri/pushtest/commit/2c7f7accbcf73b7b4c98ee7e1f213eb46a885042", "author": { "name":"Timo Zimmermann", "email":"timo@screamingatmyscreen.com", "username":"fallenhitokiri" }, "committer": { "name":"Timo Zimmermann", "email":"timo@screamingatmyscreen.com", "username":"fallenhitokiri" }, "added": [ "test2.md" ], "removed":[ ], "modified":[ ] }, "repository":{ "id":19200766, "name":"pushtest", "url":"https://github.com/fallenhitokiri/pushtest", "description":"nothing to see here", "watchers":0, "stargazers":0, "forks":0, "fork":false, "size":0, "owner":{ "name":"fallenhitokiri", "email":"timo@screamingatmyscreen.com" }, "private":false, "open_issues":0, "has_issues":true, "has_downloads":true, "has_wiki":true, "created_at":1398591886, "pushed_at":1398592401, "master_branch":"master" }, "pusher": { "name":"fallenhitokiri", "email":"timo@screamingatmyscreen.com"} }`)

	parse(not, body1)

	if len(not) != 1 {
		t.Error("Wrong length of queue", len(not))
	}

	parse(not, body2)

	if len(not) != 1 {
		t.Error("Wrong length of queue", len(not))
	}
}
