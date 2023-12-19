package repo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"context"
	"log"
	"time"
)

// SessionMongo armazena a sessao de conexao com o MongoDB
var MongoClient *mongo.Client
var ctx context.Context

// AbreConexaoComMongo faz a conexao com o Mongo
func AbreConexaoComMongo() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://go:go@localhost:27017/go-course")
	MongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return
}
