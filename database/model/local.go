package model

// Local armazena dados da localidade pelo seu codigo telefonico
type Local struct {
	Country   string `json:"country" db:"country"`
	City      string `json:"city" db:"city"`
	PhoneCode string `json:"telcode" db:"telcode"`
}
