package repo

import (
	/*
		github.com/go-sql-driver/mysql não é usado diretamente pela aplicação
	*/
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

/**

// MySQL
docker run \
-e MYSQL_ALLOW_EMPTY_PASSWORD=yes \
-e MYSQL_ROOT_PASSWORD= \
-e MYSQL_DATABASE=go-course \
-p 3306:3306 \
-d mysql

// Client
sudo apt install mysql-client

// Connecting
mysql -h localhost -P 3306 -D go-course -u root -p < ddl.sql
mysql -h localhost -P 3306 -D go-course -u root -p < dml.sql

or Postgres

// Server
docker run \
-e POSTGRES_PASSWORD= \
-e POSTGRES_USER=root \
-e POSTGRES_DB=go-course \
-p 5432:5432 \
-d postgres

// Client
sudo apt install postgresql-client

// Connecting
psql -h localhost -p 5432 go-course root < ddl.sql
psql -h localhost -p 5432 go-course root < dml.sql
**/

//variavel singleton que armazena a conexao
var db *sqlx.DB

//AbreConexaoComBancoDeDadosSQL funcao que abre a conexao com o banco MYSQL
func AbreConexaoComBancoDeDadosSQL() (db *sqlx.DB, err error) {
	err = nil

	//db, err = sqlx.Open("mysql", "root@tcp(localhost:3306)/go-course?parseTime=true")
	db, err = sqlx.Open("postgres", "postgres://root@localhost/go-course?sslmode=disable")
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
