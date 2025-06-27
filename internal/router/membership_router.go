package router

import (
	"go-kpl/internal/presentation/controllers"
	"go-kpl/internal/presentation/middleware"

	"github.com/gin-gonic/gin"
)

func Membership(server *gin.Engine, controller controllers.MembershipController, middleware middleware.Middleware) {
	routes := server.Group("/api/membership")
	{
		routes.POST("/add-membership", middleware.OnlyAllow(), controller.CreateMembership)
		routes.GET("/", controller.GetAllMembership)
		routes.GET("/:id", controller.GetByIdMembership)
		routes.PUT("/update/:id", middleware.OnlyAllow(), controller.UpdateByIdMembership)
	}
}
