// Callbacks handles receiving notifications from repository sources like
// GitHub.
package callbacks

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Notification interface {
	Branch() string
	URL() string
	By() (string, string)
}

func Callback(rw http.ResponseWriter, req *http.Request, not chan Notification) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic("reading")
	}

	var cb GitHubCallback
	err = json.Unmarshal(body, &cb)

	if err != nil {
		panic("could not unmarshal request")
	}

	not <- &cb
}
