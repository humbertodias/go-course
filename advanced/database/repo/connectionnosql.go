package repo

import (
	"gopkg.in/mgo.v2"
)

// SessionMongo armazena a sessao de conexao com o MongoDB
var SessionMongo *mgo.Session

// AbreConexaoComMongo faz a conexao com o Mongo
func AbreConexaoComMongo() (err error) {
	err = nil
	SessionMongo, err = mgo.Dial("mongodb://go:go@localhost:27017/go-course")
	return
}
