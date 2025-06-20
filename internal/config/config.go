package config

import (
	"go-kpl/database"
	"go-kpl/internal/application/services"
	"go-kpl/internal/domain/repository"
	"go-kpl/internal/presentation/controllers"
	"go-kpl/internal/presentation/middleware"
	"go-kpl/internal/router"

	"github.com/gin-gonic/gin"
)

type RestServer struct {
	Engine *gin.Engine
}

func NewGinServer() *RestServer {

	db := database.New()
	engine := gin.Default()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{
			"error": "Not Found",
		})
	})
	engine.Use(middleware.CORSMiddleware())

	var (
		userRepository repository.UserRepository = repository.NewUserRepository(db)

		userService services.UserService = services.NewUserService(userRepository)

		userController controllers.UserController = controllers.NewUserController(userService)
	)

	router.User(engine, userController)

	return &RestServer{
		Engine: engine,
	}
}

func (s *RestServer) Run(addr string) {
	if err := s.Engine.Run(addr); err != nil {
		panic("Failed to run server: " + err.Error())
	}
}
