package router

import (
	"go-kpl/internal/presentation/controllers"

	"github.com/gin-gonic/gin"
)

func UserMembership(server *gin.Engine, controller controllers.UserMembershipController) {
	routes := server.Group("/api/user-membership")
	{
		routes.POST("/new-membership", controller.CreateUserMembership)
	}
}
