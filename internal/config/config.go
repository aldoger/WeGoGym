package config

import (
	"go-kpl/infrastructure/database"
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
		userRepository       repository.UserRepository       = repository.NewUserRepository(db)
		membershipRepository repository.MembershipRepository = repository.NewMembershipRepository(db)

		userService      services.UserService       = services.NewUserService(userRepository)
		membersipService services.MembershipService = services.NewMembershipService(membershipRepository)

		userController       controllers.UserController       = controllers.NewUserController(userService)
		membershipController controllers.MembershipController = controllers.NewMembershipController(membersipService)
	)

	router.User(engine, userController)
	router.Membership(engine, membershipController)

	return &RestServer{
		Engine: engine,
	}
}

func (s *RestServer) Run(addr string) {
	if err := s.Engine.Run(addr); err != nil {
		panic("Failed to run server: " + err.Error())
	}
}
