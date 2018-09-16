package repo

import (
	"gopkg.in/mgo.v2"
)

// SessionMongo armazena a sessao de conexao com o MongoDB
var SessionMongo *mgo.Session

// AbreConexaoComMongo faz a conexao com o Mongo
func AbreConexaoComMongo() (err error) {
	err = nil
	SessionMongo, err = mgo.Dial("mongodb://go:go@127.0.0.1:27017/go-course")
	return
}

/*
// Server
docker run \
-e MONGODB_USERNAME=go \
-e MONGODB_PASSWORD=go \
-e MONGODB_DATABASE=go-course \
-p 27017:27017 \
-d bitnami/mongodb

//Client
sudo apt install mongodb-clients

// Connecting iterative
mongo "mongodb://go:go@127.0.0.1:27017/go-course"

// Importing
mongoimport --uri mongodb://go:go@127.0.0.1:27017/go-course --collection local --jsonArray --file places.json

*/
