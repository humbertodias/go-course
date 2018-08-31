package model

import "database/sql"

// Local armazena dados da localidade pelo seu codigo telefonico
type Local struct {
	Country   string         `json:"country" db:"country"`
	City      sql.NullString `json:"city" db:"city"`
	PhoneCode string         `json:"telcode" db:"telcode"`
}
