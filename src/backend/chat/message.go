package chat

import (
	"backend/utils"
	"time"
)

type Message struct {
	ID     int64  `json: "id"`
	Body   string `json:"body"`
	Sender string `json:"sender"`
	Time   string `json:"time"`
}

func NewMessage(body, sender string) *Message {
	return &Message{
		ID:     utils.GetRandomI64(),
		Body:   body,
		Sender: sender,
		Time:   time.Now().Format("15:04:05"),
	}
}
