package main

import (
	"barbearia-agendamento/config"
	"barbearia-agendamento/models"
	"barbearia-agendamento/routes"
	"log"
)

func main() {
	config.ConnectDatabase()

	err := config.DB.AutoMigrate(&models.Agendamento{}, &models.Usuario{}, &models.Horario{}, &models.Produto{})
	if err != nil {
		log.Fatal("Erro ao realizar a migração:", err)
	}

	r := routes.SetupRouter()

	log.Fatal(r.Run(":8080"))
}
