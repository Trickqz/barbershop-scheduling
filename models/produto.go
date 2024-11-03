package models

import (
	"gorm.io/gorm"
)

type Produto struct {
	gorm.Model
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
}
