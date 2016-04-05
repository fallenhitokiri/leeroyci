package github

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type github interface {
	// makeRequest handles HTTP requests to GitHubs API.
	// If the API endpoint does not expect any information nil should be passed as payload.
	makeRequest(method string, url string, token string, payload []byte) ([]byte, error)
}

type githubAPI struct{}

// makeRequest handles HTTP requests to GitHubs API.
// If the API endpoint does not expect any information nil should be passed as payload.
func (g githubAPI) makeRequest(method string, url string, token string, payload []byte) ([]byte, error) {
	r, err := http.NewRequest(method, url, bytes.NewReader(payload))

	if err != nil {
		return nil, err
	}

	addHeaders(token, r)

	c := http.Client{}

	re, err := c.Do(r)

	if err != nil {
		return nil, err
	}

	defer re.Body.Close()

	b, err := ioutil.ReadAll(re.Body)

	if err != nil {
		return nil, err
	}

	return b, nil
}
