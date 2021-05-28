package chat

import (
	"backend/utils"
	"time"
)

type MessageType int

const (
	MESSAGE MessageType = iota
	JOIN
	LEVAE
)

type Message struct {
	ID          int64       `json:"id"`
	Body        string      `json:"body"`
	Sender      string      `json:"sender"`
	MessageType MessageType `json:"type"`
	Time        string      `json:"time"`
	UserCount   int         `json:"usercount"`
}

func NewMessage(body, sender string, messageType MessageType, userCount int) *Message {
	return &Message{
		ID:          utils.GetRandomI64(),
		Body:        body,
		Sender:      sender,
		MessageType: messageType,
		Time:        time.Now().Format("15:04:05"),
		UserCount:   userCount,
	}
}
