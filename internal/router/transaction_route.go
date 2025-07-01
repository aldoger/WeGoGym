package router

import (
	"go-kpl/internal/presentation/controllers"

	"github.com/gin-gonic/gin"
)

func Transaction(server *gin.Engine, controller controllers.TransactionController) {
	routes := server.Group("/api/transaction")
	{
		routes.POST("/membership", controller.CreateMemberTransaction)
		routes.POST("/personal-trainer", controller.CreatePersonalTrainerTransaction)
	}
}
