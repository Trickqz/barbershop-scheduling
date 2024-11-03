package routes

import (
	"barbearia-agendamento/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/agendar", controllers.CriarAgendamento)
	r.GET("/horarios-disponiveis", controllers.ListarHorariosDisponiveis)

	return r
}
