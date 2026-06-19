package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/juanpaAndino/Proyecto-Integrador/internal/api"
	"github.com/juanpaAndino/Proyecto-Integrador/internal/websocket"
)

func main() {
	hub := websocket.NewHub()
	go hub.Run()

	http.HandleFunc("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok"}`))
	})
	http.HandleFunc("/api/v1/register", api.RegisterHandler)
	http.HandleFunc("/api/v1/login", api.LoginHandler)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.ServeWs(hub, w, r)
	})

	fmt.Println("Servidor corriendo en el puerto 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
