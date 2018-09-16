package repo

import (
	"github.com/humbertodias/go-course/advanced/database/model"
	"gopkg.in/mgo.v2/bson"
)

// GetLocal retorna um local do MongoDB
func GetLocal(codigoTelefone string) (local model.Local, err error) {
	session := SessionMongo.Copy()
	defer session.Close()
	collection := session.DB("go-course").C("local")
	err = collection.Find(bson.M{"telcode": codigoTelefone}).One(&local)
	return
}

// WriteLog Salva Log
func WriteLog(reg model.RegistroLog) (err error) {
	session := SessionMongo.Copy()
	defer session.Close()
	collection := session.DB("go-course").C("logvisitas")
	err = collection.Insert(reg)
	return
}
