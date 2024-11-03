package models

import (
	"gorm.io/gorm"
)

type Horario struct {
	gorm.Model
	Descricao  string `json:"descricao"`
	Disponivel bool   `json:"disponivel"`
}
