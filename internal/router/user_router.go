package router

import (
	"go-kpl/internal/presentation/controllers"

	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine, c controllers.UserController) {
	routes := r.Group("/api/user")
	{
		routes.POST("/register", c.Register)
		routes.POST("/login", c.Login)
		routes.GET("/me", c.GetMe)
	}
}
