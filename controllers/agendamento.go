package controllers

import (
	"barbearia-agendamento/config"
	"barbearia-agendamento/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CriarAgendamento(c *gin.Context) {
	var agendamento models.Agendamento
	if err := c.ShouldBindJSON(&agendamento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&agendamento)
	c.JSON(http.StatusOK, agendamento)
}

func ListarHorariosDisponiveis(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Lista de horários disponíveis"})
}
