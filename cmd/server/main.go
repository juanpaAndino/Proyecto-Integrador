package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Candied666/Avance-de-proyecto-de-GO/internal/api"
)

func main() {
	http.HandleFunc("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok"}`))
	})

	http.HandleFunc("/api/v1/register", api.RegisterHandler)
	http.HandleFunc("/api/v1/login", api.LoginHandler)

	fmt.Println("🚀 Servidor corriendo en el puerto 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
