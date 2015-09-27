// Package websocket implements the websocket protocol according to RFC 6455.
// Websockets are used to send updates for all events through the notifications
// package.
package websocket

import (
	"github.com/fallenhitokiri/leeroyci/database"
)

// Message contains all fields consumed by differetn websockets to render
// different events.
type Message struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Event          string `json:"event"`
	RepositoryName string `json:"repository_name"`
	Branch         string `json:"branch"`
	Status         string `json:"status"`
}

// NewMessage converts a job and event type to a message that can be send
// through a websocket.
func NewMessage(job *database.Job, event string) *Message {
	return &Message{
		Name:           job.Name,
		Email:          job.Email,
		Event:          event,
		RepositoryName: job.Repository.Name,
		Branch:         job.Branch,
		Status:         job.Status(),
	}
}
