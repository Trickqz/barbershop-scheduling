package routes

import (
	"barbearia-agendamento/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/agendar", controllers.CriarAgendamento)
	r.GET("/horarios-disponiveis", controllers.ListarHorariosDisponiveis)

	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)

	protected := r.Group("/admin")
	protected.Use(controllers.AuthMiddleware())
	{
		protected.POST("/criar-horario", controllers.CriarHorario)
		protected.POST("/criar-produto", controllers.CriarProduto)
	}

	return r
}
