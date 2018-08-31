package main

import (
	"fmt"
	"github.com/humbertodias/go-course/advanced/web-server/handler"
	"net/http"
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
