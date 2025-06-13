package config

import (
	"go-kpl/internal/presentation/middleware"

	"github.com/gin-gonic/gin"
)

type RestServer struct {
	Engine *gin.Engine
}

func NewGinServer() *RestServer {

	engine := gin.Default()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{
			"error": "Not Found",
		})
	})
	engine.Use(middleware.CORSMiddleware())

	return &RestServer{
		Engine: engine,
	}
}

func (s *RestServer) Run(addr string) {
	if err := s.Engine.Run(addr); err != nil {
		panic("Failed to run server: " + err.Error())
	}
}
