package websocket

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"code.google.com/p/go.net/websocket"
)

func TestClient(t *testing.T) {
	http.Handle("/websocket", GetHandler())
	server := httptest.NewServer(nil)
	defer server.Close()
	uri := fmt.Sprintf("ws://%s%s", server.Listener.Addr(), "/websocket")

	ws, err := websocket.Dial(uri, "", "http://localhost")

	if err != nil {
		t.Error(err)
	}

	client := NewClient(ws)

	msg := &Message{
		Event: "foo",
	}
	client.write(msg)
}
