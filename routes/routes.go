package routes

import (
	"api-go-controle-investimentos/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.POST("/usuario", controllers.CriaUsuario)
	r.Run(":8000")
}
