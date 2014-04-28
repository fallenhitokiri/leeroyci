package callbacks

import (
	"encoding/json"
	//"io/ioutil"
	"net/http"
)

type Notification interface {
	Branch() string
	URL() string
	By() (string, string)
}

func Callback(rw http.ResponseWriter, req *http.Request, not chan Notification) {
	/*body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic("reading")
	}*/

	body := []byte(`{ "ref":"refs/heads/master", "after":"2c7f7accbcf73b7b4c98ee7e1f213eb46a885042", "before":"b369a904428a1d5d671e3f740443590d3db55fb0", "created":false, "deleted":false, "forced":false, "compare":"https://github.com/fallenhitokiri/pushtest/compare/b369a904428a...2c7f7accbcf7", "commits":[ { "id":"2c7f7accbcf73b7b4c98ee7e1f213eb46a885042", "distinct":true, "message":"asdf", "timestamp":"2014-04-27T11:53:14+02:00", "url":"https://github.com/fallenhitokiri/pushtest/commit/2c7f7accbcf73b7b4c98ee7e1f213eb46a885042", "author": { "name":"Timo Zimmermann", "email":"timo@screamingatmyscreen.com", "username":"fallenhitokiri" }, "committer":{ "name":"Timo Zimmermann", "email":"timo@screamingatmyscreen.com", "username":"fallenhitokiri" }, "added":[ "test2.md" ], "removed": [ ], "modified": [ ] } ], "head_commit": { "id":"2c7f7accbcf73b7b4c98ee7e1f213eb46a885042", "distinct":true, "message":"asdf", "timestamp":"2014-04-27T11:53:14+02:00", "url":"https://github.com/fallenhitokiri/pushtest/commit/2c7f7accbcf73b7b4c98ee7e1f213eb46a885042", "author": { "name":"Timo Zimmermann", "email":"timo@screamingatmyscreen.com", "username":"fallenhitokiri" }, "committer": { "name":"Timo Zimmermann", "email":"timo@screamingatmyscreen.com", "username":"fallenhitokiri" }, "added": [ "test2.md" ], "removed":[ ], "modified":[ ] }, "repository":{ "id":19200766, "name":"pushtest", "url":"https://github.com/fallenhitokiri/pushtest", "description":"nothing to see here", "watchers":0, "stargazers":0, "forks":0, "fork":false, "size":0, "owner":{ "name":"fallenhitokiri", "email":"timo@screamingatmyscreen.com" }, "private":false, "open_issues":0, "has_issues":true, "has_downloads":true, "has_wiki":true, "created_at":1398591886, "pushed_at":1398592401, "master_branch":"master" }, "pusher": { "name":"fallenhitokiri", "email":"timo@screamingatmyscreen.com"} }`)

	var cb GitHubCallback
	err := json.Unmarshal(body, &cb)

	if err != nil {
		panic("could not unmarshal request")
	}

	not <- &cb
}
