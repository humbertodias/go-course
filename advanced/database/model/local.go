package model

// Local armazena dados da localidade pelo seu codigo telefonico
type Local struct {
	Country   string `json:"country" db:"country" bson:"country"`
	City      string `json:"city" db:"city" bson:"city"`
	PhoneCode string `json:"telcode" db:"telcode" bson:"telcode"`
}

// LogQuery armazena dados de query
type LogQuery struct {
	DateRequest string `json:"daterequest" db:"daterequest"`
}
