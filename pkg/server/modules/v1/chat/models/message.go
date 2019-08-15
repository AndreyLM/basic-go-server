package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// ChatMessage - chat message
type ChatMessage struct {
	UUID    uuid.UUID `json:"uuid"`
	Token   string    `json:"token"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}
