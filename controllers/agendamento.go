package controllers

import (
	"barbearia-agendamento/config"
	"barbearia-agendamento/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

func CriarAgendamento(c *gin.Context) {
	var agendamento models.Agendamento
	if err := c.ShouldBindJSON(&agendamento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingAgendamento models.Agendamento
	if err := config.DB.Where("horario = ?", agendamento.Horario).First(&existingAgendamento).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Horário já está ocupado"})
		return
	}

	config.DB.Create(&agendamento)

	err := enviarEmailConfirmacao(agendamento)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar e-mail de confirmação"})
		return
	}

	c.JSON(http.StatusOK, agendamento)
}

func enviarEmailConfirmacao(agendamento models.Agendamento) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "seu_email@example.com")
	m.SetHeader("To", agendamento.Email)
	m.SetHeader("Subject", "Confirmação de Agendamento")
	m.SetBody("text/plain", "Olá "+agendamento.Nome+",\n\nSeu agendamento foi confirmado para o horário: "+agendamento.Horario+".\n\nObrigado!")

	d := gomail.NewDialer("smtp.example.com", 587, "seu_email@example.com", "sua_senha")
	return d.DialAndSend(m)
}

func ListarHorariosDisponiveis(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Lista de horários disponíveis"})
}
