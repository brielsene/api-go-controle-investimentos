package controllers

import (
	"api-go-controle-investimentos/database"
	"api-go-controle-investimentos/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CriaUsuario(c *gin.Context) {
	var usuario models.Usuario

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error:": err.Error,
		})
		return
	}

	database.DB.Create(&usuario)
	c.JSON(http.StatusCreated, &usuario)
}

func CriaInvestimento(c *gin.Context) {

}
