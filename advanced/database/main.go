package main

import (
	"fmt"
	"net/http"

	"github.com/humbertodias/go-course/advanced/database/handler"
	"github.com/humbertodias/go-course/advanced/database/repo"
)

func init() {
	fmt.Println("Vamos começar a subir o servidor...")
	_, err := repo.AbreConexaoComBancoDeDadosSQL()
	if err != nil {
		fmt.Println("Parando a carga do servidor. Erro ao abrir o banco de dados: ", err.Error())
		return
	}

	err = repo.AbreConexaoComMongo()
	if err != nil {
		fmt.Println("Parando a carga do servidor. Erro ao abrir a sessao com o MongoDB: ", err.Error())
		return
	}

}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá Mundo!")
	})
	http.HandleFunc("/local/", handler.Local)
	fmt.Println("Estou escutando, comandante...")
	http.ListenAndServe(":8181", nil)
}
