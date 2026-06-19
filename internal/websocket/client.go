package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Client struct {
	Hub      *Hub
	Conn     *websocket.Conn
	Send     chan []byte
	Username string
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error: %v", err)
			}
			break
		}
		finalMessage := []byte(c.Username + ": " + string(message))
		c.Hub.Broadcast <- finalMessage
	}
}

func (c *Client) WritePump() {
	defer func() {
		c.Conn.Close()
	}()
	for message := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			return
		}
	}
}

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		username = "Usuario_Anonimo"
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256), Username: username}
	client.Hub.Register <- client

	go client.WritePump()
	go client.ReadPump()
}
