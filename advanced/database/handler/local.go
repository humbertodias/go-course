package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/humbertodias/go-course/advanced/database/model"
	"github.com/humbertodias/go-course/advanced/database/repo"
)

//Local é o manipulador da requisição de rota /local/
func Local(w http.ResponseWriter, r *http.Request) {

	local := model.Local{}
	codigoTelefone := r.URL.Path[7:]

	// codigoTelefone, err := strconv.Atoi(parametroCodigoTelefone)
	// if err != nil {
	// 	http.Error(w, "Não foi enviado um numero válido. Verifique.", http.StatusBadRequest)
	// 	fmt.Println("[local] erro ao converter o numero enviado: ", err.Error())
	// 	return
	// }

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

	localMongo, err := repo.GetLocal(codigoTelefone)
	if err != nil {
		http.Error(w, "Não foi possível pesquisar esse numero.", http.StatusInternalServerError)
		fmt.Println("[local] nao foi possível executar a query no MongoDB: ", sql, " Erro: ", err.Error())
		return
	}

	// if err := ModeloLocal.ExecuteTemplate(w, "local.html", local); err != nil {
	if err := ModeloLocal.ExecuteTemplate(w, "local.html", localMongo); err != nil {
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

	log := model.RegistroLog{}
	log.DataVisita = time.Now().Format("02/01/2006 15:04:05")
	err = repo.WriteLog(log)
	if err != nil {
		fmt.Println("[local] Erro na inclusao do log no MongoDB: ", err.Error())
	}

	fmt.Println("Sucesso! ", linhasAfetadas, " linha(s) afetada(s)")
}
