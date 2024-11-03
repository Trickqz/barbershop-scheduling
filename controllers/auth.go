package controllers

import (
	"barbearia-agendamento/config"
	"barbearia-agendamento/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("sua_chave_secreta")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			c.Abort()
			return
		}

		tokenString = tokenString[len("Bearer "):]

		claims := &jwt.StandardClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}

// Register cria um novo usuário
func Register(c *gin.Context) {
	var usuario models.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usuario.Senha), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário"})
		return
	}
	usuario.Senha = string(hashedPassword)

	config.DB.Create(&usuario)
	c.JSON(http.StatusOK, usuario)
}

func Login(c *gin.Context) {
	var usuario models.Usuario
	var foundUsuario models.Usuario

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Where("email = ?", usuario.Email).First(&foundUsuario).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário ou senha inválidos"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUsuario.Senha), []byte(usuario.Senha)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário ou senha inválidos"})
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		Subject:   foundUsuario.Email,
		ExpiresAt: expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
