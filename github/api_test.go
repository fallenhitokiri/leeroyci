package github

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type githubMock struct {
	method  string
	url     string
	token   string
	payload []byte
}

func (g *githubMock) makeRequest(method string, url string, token string, payload []byte) ([]byte, error) {
	g.method = method
	g.url = url
	g.token = token
	g.payload = payload

	return nil, nil
}

func TestGithubAPIMakeRequest(t *testing.T) {
	var request *http.Request

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		request = r
	}))
	defer ts.Close()

	githubAPI{}.makeRequest("POST", ts.URL, "foo", nil)

	if request.Header["Authorization"][0] != "token foo" {
		t.Error("Wrong auth token", request.Header["Authorization"][0])
	}

	if request.Method != "POST" {
		t.Error("Wrong request method", request.Method)
	}
}
