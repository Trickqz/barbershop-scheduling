package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env:", err)
	}

	databaseURL := os.Getenv("DATABASE_URL")

	DB, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
}
