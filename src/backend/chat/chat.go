package chat

import (
	"backend/utils"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

type Chat struct {
	users    map[string]*User
	messages chan *Message
	join     chan *User
	leave    chan *User
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		log.Printf("%s %s %s %v\n", r.Method, r.Host, r.RequestURI, r.Proto)
		return r.Method == http.MethodGet
	},
}

func (c *Chat) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Error on websocket connection:", err.Error())
	}

	keys := r.URL.Query()
	username := keys.Get("username")
	if strings.TrimSpace(username) == "" {
		username = fmt.Sprintf("anom-%d", utils.GetRandomI64())
	}

	user := &User{
		Username: username,
		Conn:     conn,
		Global:   c,
	}

	c.join <- user

	user.Read()
}

func (c *Chat) Run() {
	for {
		select {
		case user := <-c.join:
			c.add(user)
		case message := <-c.messages:
			c.boardcasts(message)
		case user := <-c.leave:
			c.disconnect(user)
		}
	}
}

func (c *Chat) add(user *User) {
	if _, ok := c.users[user.Username]; !ok {
		c.users[user.Username] = user
		body := fmt.Sprintf("%s 加入了聊天\n", user.Username)
		c.boardcasts(NewMessage(body, "系统"))
		log.Printf("Added user: %s, Total: %d\n", user.Username, len(c.users))
	}
}

func (c *Chat) boardcasts(message *Message) {
	log.Println("Boardcast message", message)
	for _, user := range c.users {
		user.Write(message)
	}
}

func (c *Chat) disconnect(user *User) {
	if _, ok := c.users[user.Username]; ok {
		defer user.Conn.Close()
		delete(c.users, user.Username)

		body := fmt.Sprintf("%s 离开了聊天\n", user.Username)
		c.boardcasts(NewMessage(body, "系统"))
		log.Printf("%s left the chat, Total: %d\n", user.Username, len(c.users))
	}
}

func Start(port string) {
	fmt.Println("Chat listening on http://localhost", port)

	c := &Chat{
		users:    make(map[string]*User),
		messages: make(chan *Message),
		join:     make(chan *User),
		leave:    make(chan *User),
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Welcome to WebChat"))
	})

	http.HandleFunc("/chat", c.Handler)

	// 启动Goroutine，一直监听
	go c.Run()

	log.Fatal(http.ListenAndServe(port, nil))
}
