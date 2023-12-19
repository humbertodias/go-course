package repo

import (
	"github.com/humbertodias/go-course/advanced/database/model"
	"go.mongodb.org/mongo-driver/bson"
)

// GetLocal retorna um local do MongoDB
func GetLocal(codigoTelefone string) (local model.Local, err error) {
	collection := mongoClient.Database("go-course").Collection("local")
	filter := bson.M{"telcode": codigoTelefone}
	err = collection.FindOne(ctx, filter).Decode(&local)
	return
}

// WriteLog Salva Log
func WriteLog(reg model.RegistroLog) (err error) {
	collection := mongoClient.Database("go-course").Collection("logvisitas")
	_, err = collection.InsertOne(ctx, reg)
	return
}
