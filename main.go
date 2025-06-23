package main

import (
	"go-kpl/cmd"
	"go-kpl/internal/config"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("RAILWAY_ENVIRONMENT") == "" {
		if err := godotenv.Load(); err != nil {
			log.Println("Warning: .env file not found, using system env variables")
		}
	}

	if err := cmd.Commands(); err != nil {
		panic("Failed to get Commands: " + err.Error())
	}

	ginServer := config.NewGinServer()

	ginServer.Engine.Any("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	ginServer.Run(":8080")
}
