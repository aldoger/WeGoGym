package main

import (
	"go-kpl/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	ginServer := config.NewGinServer()

	ginServer.Engine.Any("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	ginServer.Run(":8080")
}
