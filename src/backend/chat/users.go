package chat

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type User struct {
	Username string
	Conn     *websocket.Conn
	Global   *Chat
}

func (u *User) Read() {
	for {
		// ReadMessage()是阻塞的
		if _, message, err := u.Conn.ReadMessage(); err != nil {
			log.Println("Error on read message:", err.Error())
			u.Global.leave <- u
			return
		} else {
			u.Global.messages <- NewMessage(string(message), u.Username, MESSAGE, len(u.Global.users))
		}
	}
}

func (u *User) Write(message *Message) {
	b, _ := json.Marshal(message)

	if err := u.Conn.WriteMessage(websocket.TextMessage, b); err != nil {
		log.Println("Error on write message:", err.Error())
	}
}
