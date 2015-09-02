// Package github integrates everything necessary to test commits, comment on
// pull requests and close them if the build failed.
package github

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// githubRequest handles HTTP requests to GitHubs API.
// If the API endpoint does not expect any information nil should be passed as payload.
func makeRequest(method string, url string, token string, payload []byte) ([]byte, error) {
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

// AddHeaders adds all headers to a request to conform to GitHubs API.
// token is the API token that will be used for the request.
func addHeaders(token string, req *http.Request) {
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "token "+token)
}
