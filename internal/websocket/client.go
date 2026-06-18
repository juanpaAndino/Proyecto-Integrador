package websocket

import (
	"log"
	"net/http"

	//"github.com/golang-jwt/jwt/v5"
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
	// 1. Extraer el nombre de usuario directamente de la URL
	username := r.URL.Query().Get("username")
	if username == "" {
		username = "Usuario_Anonimo" // Red de seguridad por si acaso
	}

	/* === SEGURIDAD JWT APAGADA TEMPORALMENTE PARA LA DEFENSA ===
	tokenString := r.URL.Query().Get("token")
	if tokenString == "" {
	    http.Error(w, "Acceso denegado: falta token", http.StatusUnauthorized)
	    return
	}

	// Validar el JWT
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
	    return []byte("SecretoSuperSeguroParaElChat"), nil // Debe coincidir con el de handlers.go
	})

	if err != nil || !token.Valid {
	    http.Error(w, "Acceso denegado: token inválido", http.StatusUnauthorized)
	    return
	}

	// Extraer el username del token validado
	username = claims["username"].(string)
	============================================================ */

	// 2. Actualizar la conexión HTTP normal a una de WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// 3. Registramos al cliente con su nombre de usuario en el Hub
	client := &Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256), Username: username}
	client.Hub.Register <- client

	// 4. Encender los procesos de lectura y escritura en paralelo (Goroutines)
	go client.WritePump()
	go client.ReadPump()
}
