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
}

func Callback(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic("reading")
	}

	var cb GitHubCallback
	err = json.Unmarshal(body, &cb)

	if err != nil {
		panic("could not unmarshal request")
	}

	log.Println(cb.Branch())
	log.Println(cb.URL())
}
