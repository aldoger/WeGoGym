package router

import (
	"go-kpl/internal/presentation/controllers"

	"github.com/gin-gonic/gin"
)

func User(server *gin.Engine, controller controllers.UserController) {
	routes := server.Group("/api/user")
	{
		routes.POST("/register", controller.Register)
		routes.POST("/login", controller.Login)
		routes.GET("/me", controller.GetMe)
		routes.GET("/generate-qr", controller.GenerateQrMe)
	}
}
