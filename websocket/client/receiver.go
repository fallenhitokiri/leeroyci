package main

import (
	"fmt"
	"log"

	"code.google.com/p/go.net/websocket"
)

var origin = "http://localhost/"
var url = "ws://localhost:8082/websocket"

func main() {
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	for {
		var msg = make([]byte, 512)
		_, err = ws.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Receive: %s\n", msg)
	}
}
