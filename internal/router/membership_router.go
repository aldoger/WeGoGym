package router

import (
	"go-kpl/internal/presentation/controllers"

	"github.com/gin-gonic/gin"
)

func Membership(server *gin.Engine, controller controllers.MembershipController) {
	routes := server.Group("/api/membership")
	{
		routes.POST("/add-membership", controller.CreateMembership)
		routes.GET("/", controller.GetAllMembership)
		routes.GET("/:id", controller.GetByIdMembership)
		routes.PUT("/update/:id", controller.UpdateByIdMembership)
	}
}
