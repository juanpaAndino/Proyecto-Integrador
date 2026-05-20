package api

import (
	"fmt"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Endpoint para registrar usuarios (Etapa 3)")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Endpoint para login con bcrypt (Etapa 3)")
}
