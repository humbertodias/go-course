package repo

import (
	/*
	github.com/go-sql-driver/mysql não é usado diretamente pela aplicação
	*/
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

/**
docker run \
-e MYSQL_ALLOW_EMPTY_PASSWORD=yes \
-e MYSQL_ROOT_PASSWORD= \
-e MYSQL_DATABASE=go-course \
-p 3306:3306 \
-d mysql
**/

//variavel singleton que armazena a conexao
var db *sqlx.DB

//AbreConexaoComBancoDeDadosSQL funcao que abre a conexao com o banco MYSQL
func AbreConexaoComBancoDeDadosSQL() (db *sqlx.DB, err error) {
	err = nil
	db, err = sqlx.Open("mysql", "root@tcp(localhost:3306)/go-course?parseTime=true")
	if err != nil {
		log.Println("[AbreConexaoComBancoDeDadosSQL] Erro na conexao: ", err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		log.Println("[AbreConexaoComBancoDeDadosSQL] Erro no ping na conexao: ", err.Error())
		return
	}
	return
}

//GetDBConnection Obtem a conexao com o banco de dados
func GetDBConnection() (localdb *sqlx.DB, err error) {
	if db == nil {
		db, err = AbreConexaoComBancoDeDadosSQL()
		if err != nil {
			log.Println("[GetDBConnection] Erro na conexao: ", err.Error())
			return
		}
	}
	err = db.Ping()
	if err != nil {
		log.Println("[GetDBConnection] Erro no ping na conexao: ", err.Error())
		return
	}
	localdb = db
	return
}
