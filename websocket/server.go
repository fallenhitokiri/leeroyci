// Package websocket implements the websocket protocol according to RFC 6455.
// Websockets are used to send updates for all events through the notifications
// package.
package websocket

import (
	"log"

	"golang.org/x/net/websocket"
)

var socketServer *server

type server struct {
	clients []*client
}

// NewServer initializes a new websocket server.
func NewServer() {
	socketServer = &server{
		clients: make([]*client, 0),
	}
}

func (s *server) addClient(c *client) {
	log.Println("Websocket added client")
	s.clients = append(s.clients, c)
}

func (s *server) removeClient(c *client) {
	log.Println("Websocket deleted client")
	for index, client := range s.clients {
		if client == c {
			s.clients[index] = s.clients[len(s.clients)-1]
			s.clients[len(s.clients)-1] = nil
			s.clients = s.clients[:len(s.clients)-1]
			return
		}
	}
}

// Send sends a message to all connected clients.
func Send(msg *Message) {
	for _, c := range socketServer.clients {
		c.write(msg)
	}
}

func connectHandler(ws *websocket.Conn) {
	defer func() {
		err := ws.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	client := NewClient(ws)
	socketServer.addClient(client)
	client.listen()
}

// GetHandler returns a net/http.Handler compatible handler for websockets.
func GetHandler() websocket.Handler {
	return websocket.Handler(connectHandler)
}
