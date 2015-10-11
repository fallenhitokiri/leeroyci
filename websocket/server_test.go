package websocket

import (
	"testing"
)

func TestAddRemove(t *testing.T) {
	NewServer()
	c1 := NewClient(nil)
	c2 := NewClient(nil)
	c3 := NewClient(nil)

	socketServer.addClient(c1)
	socketServer.addClient(c2)
	socketServer.addClient(c3)

	if len(socketServer.clients) != 3 {
		t.Error("Wrong number of clients", len(socketServer.clients))
	}

	socketServer.removeClient(c2)

	for _, c := range socketServer.clients {
		if c == c2 {
			t.Error("client c2 not removed")
		}
	}
}
