package router

import (
	"go-kpl/internal/presentation/controllers"

	"github.com/gin-gonic/gin"
)

func UserPersonalTrainer(server *gin.Engine, controller controllers.UserPersonalTrainerController) {
	routes := server.Group("/api/user-trainer/")
	{
		routes.POST("/new-trainer", controller.CreateUserPersonalTrainer)
	}
}
