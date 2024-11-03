package models

import (
	"gorm.io/gorm"
)

type Agendamento struct {
	gorm.Model
	Nome    string `json:"nome"`
	Email   string `json:"email"`
	Horario string `json:"horario"`
	Status  string `json:"status"` // Ex: "confirmado", "cancelado"
}
