package main

import (
	"fmt"
	"net/http"

	"github.com/humbertodias/go-course/advanced/database/handler"
	"github.com/humbertodias/go-course/advanced/database/repo"
)

func init() {
	initSQL()
	initNoSQL()
}

func initSQL() {
	fmt.Println("Vamos começar a subir o servidor...")
	_, err := repo.AbreConexaoComBancoDeDadosSQL()
	if err != nil {
		fmt.Println("Parando a carga do servidor. Erro ao abrir o banco de dados: ", err.Error())
		return
	}
	fmt.Println("Connected on SQL")
}

func initNoSQL() {
	err := repo.AbreConexaoComMongo()
	if err != nil {
		fmt.Println("Parando a carga do servidor. Erro ao abrir a sessao com o MongoDB: ", err.Error())
		return
	}
	fmt.Println("Connected on NoSQL")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá Mundo!")
	})
	http.HandleFunc("/sql/", handler.LocalSQL)
	http.HandleFunc("/nosql/", handler.LocalNoSQL)
	fmt.Println("Estou escutando em 8181, comandante...")
	http.ListenAndServe(":8181", nil)
}
