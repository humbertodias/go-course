package main

import (
	"fmt"
	"net/http"

	"github.com/humbertodias/go-course/web-server/handler"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá Mundo")
	})
	http.HandleFunc("/funcao", handler.Funcao)
	http.HandleFunc("/hello", handler.Hello)

	fmt.Println("Estou escutando na porta 8081")
	http.ListenAndServe(":8081", nil)
}
