package router

import (
	"go-kpl/internal/presentation/controllers"
	"go-kpl/internal/presentation/middleware"

	"github.com/gin-gonic/gin"
)

func UserMembership(server *gin.Engine, controller controllers.UserMembershipController, middleware middleware.Middleware) {
	routes := server.Group("/api/user-membership")
	{
		routes.POST("/new-membership", middleware.Authenticate(), controller.CreateUserMembership)
		routes.GET("/search-membership/:id", middleware.OnlyAllow(), controller.SearchMembership)
	}
}
