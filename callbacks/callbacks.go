// Callbacks handles receiving notifications from repository sources like
// GitHub.
package callbacks

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Notification interface {
	Branch() string
	URL() string
	By() (string, string)
	ShouldBuild() bool
}

func Callback(rw http.ResponseWriter, req *http.Request, not chan Notification) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic("reading")
	}

	parse(not, body)
}

// Parse a request body and add it to the build queue.
func parse(not chan Notification, body []byte) {
	var cb GitHubCallback
	err := json.Unmarshal(body, &cb)

	if err != nil {
		log.Println(string(body))
		panic("Could not unmarshal request")
	}

	if cb.ShouldBuild() == true {
		not <- &cb
	} else {
		log.Println("Not adding", cb.URL(), cb.Branch(), "to build queue")
	}
}
