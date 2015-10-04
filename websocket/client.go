// Package websocket implements the websocket protocol according to RFC 6455.
// Websockets are used to send updates for all events through the notifications
// package.
package websocket

import (
	"io"
	"log"

	"golang.org/x/net/websocket"
)

type client struct {
	socket *websocket.Conn
}

// NewClient returns a new websocket client.
func NewClient(ws *websocket.Conn) *client {
	return &client{
		socket: ws,
	}
}

func (c *client) write(msg *Message) {
	websocket.JSON.Send(c.socket, msg)
}

func (c *client) listen() {
	for {
		select {

		default:
			var msg Message
			err := websocket.JSON.Receive(c.socket, &msg)
			if err == io.EOF {
				socketServer.removeClient(c)
				return
			} else if err != nil {
				log.Println(err)
			}
		}
	}
}
