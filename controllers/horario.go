package controllers

import (
	"barbearia-agendamento/config"
	"barbearia-agendamento/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CriarHorario(c *gin.Context) {
	var horario models.Horario
	if err := c.ShouldBindJSON(&horario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&horario)
	c.JSON(http.StatusOK, horario)
}

func CriarProduto(c *gin.Context) {
	var produto models.Produto
	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&produto)
	c.JSON(http.StatusOK, produto)
}
