package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/humbertodias/go-course/advanced/database/model"
	"github.com/humbertodias/go-course/advanced/database/repo"
)

// LocalSQL é o manipulador da requisição de rota /local/
func LocalSQL(w http.ResponseWriter, r *http.Request) {

	local := model.Local{}
	//sql/ = length 5
	codigoTelefone := r.URL.Path[5:]
	fmt.Printf("codigoTelefone: %s\n", codigoTelefone)
	db, err := repo.GetDBConnection()
	if err != nil {
		log.Println("[Local] Erro na conexao: ", err.Error())
		return
	}
	p := model.Local{PhoneCode: codigoTelefone}
	sql := "select country, city, telcode from place where telcode = :telcode"
	linha, err := db.NamedQuery(sql, p)
	if err != nil {
		http.Error(w, "Não foi possível pesquisar esse numero.", http.StatusInternalServerError)
		fmt.Println("[local] nao foi possível executar a query: ", sql, " Erro: ", err.Error())
		return
	}
	for linha.Next() {
		err = linha.StructScan(&local)
		if err != nil {
			http.Error(w, "Não foi possível pesquisar esse numero.", http.StatusInternalServerError)
			fmt.Println("[local] nao foi fazer o binding dos dados do banco na struct: ", err.Error())
			return
		}
	}

	if err := ModeloLocal.Execute(w, local); err != nil {
		http.Error(w, "Houve um erro na renderização da página.", http.StatusInternalServerError)
		fmt.Println("[local] Erro na execucao do modelo: ", err.Error())
	}
	sql = "insert into logquery (daterequest) values (:daterequest)"
	daterequest := time.Now().Format("02/01/2006 15:04:05")
	plq := model.LogQuery{DateRequest: daterequest}
	resultado, err := db.NamedExec(sql, plq)
	if err != nil {
		fmt.Println("[local] Erro na inclusao do log: ", sql, " - ", err.Error())
	}
	linhasAfetadas, err := resultado.RowsAffected()
	if err != nil {
		fmt.Println("[local] Erro ao pegar o numero de linhas afetadas pelo comando: ", sql, " - ", err.Error())
	}

	fmt.Println("Sucesso! ", linhasAfetadas, " linha(s) afetada(s)")
}

// LocalNoSQL é o manipulador da requisição de rota /local/
func LocalNoSQL(w http.ResponseWriter, r *http.Request) {
	//nosql/ = length 7
	codigoTelefone := r.URL.Path[7:]
	localMongo, err := repo.GetLocal(codigoTelefone)
	if err != nil {
		http.Error(w, "Não foi possível pesquisar esse numero.", http.StatusInternalServerError)
		fmt.Println("[local] nao foi possível executar a query no MongoDB:", "Erro: ", err.Error())
		return
	}

	if err := ModeloLocal.Execute(w, localMongo); err != nil {
		http.Error(w, "Houve um erro na renderização da página.", http.StatusInternalServerError)
		fmt.Println("[local] Erro na execucao do modelo: ", err.Error())
	}

	log := model.RegistroLog{}
	log.DataVisita = time.Now().Format("02/01/2006 15:04:05")
	err = repo.WriteLog(log)
	if err != nil {
		fmt.Println("[local] Erro na inclusao do log no MongoDB: ", err.Error())
	}

	fmt.Println("Sucesso!")
}
