package notification

import (
	"github.com/fallenhitokiri/leeroyci/database"
	"github.com/fallenhitokiri/leeroyci/websocket"
)

func sendWebsocket(job *database.Job, event string) {
	msg := websocket.NewMessage(job, event)
	websocket.Send(msg)
}
