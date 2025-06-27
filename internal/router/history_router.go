package router

import (
	"go-kpl/internal/presentation/controllers"
	"go-kpl/internal/presentation/middleware"

	"github.com/gin-gonic/gin"
)

func EntryHistory(server *gin.Engine, controller controllers.EntryHistoryController, middleware middleware.Middleware) {
	routes := server.Group("/api/history")
	{
		routes.GET("/all-entry", middleware.OnlyAllow(), controller.GetEntryAll)
		routes.GET("/user/:id", controller.GetEntryByUserId)
	}
}
