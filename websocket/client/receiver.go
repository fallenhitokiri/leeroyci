package main

import (
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

var origin = "http://localhost/"
var url = "ws://localhost:8082/websocket"

func main() {
	config, err := websocket.NewConfig(url, origin)
	config.Header["accesskey"] = []string{"gNV/bxhxG)IrvEeaZK_mA2HkCxC2yu!bHjGK!(MLNQ1tuDnUKyGBL9G/rGhXHNzU"}

	ws, err := websocket.DialConfig(config)
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
