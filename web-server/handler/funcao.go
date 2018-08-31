package handler

import (
	"fmt"
	"net/http"
)

// Funcao é um manipulador de requisicao http
func Funcao(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Aqui é um manipulador de funcao")
}
