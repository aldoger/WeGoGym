package config

import (
	"go-kpl/infrastructure/database"
	"go-kpl/infrastructure/externals/midtrans"
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
	midtransClient := midtrans.NewMidtrans()
	serverMiddleware := middleware.New()
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
		userRepository                repository.UserRepository                = repository.NewUserRepository(db)
		membershipRepository          repository.MembershipRepository          = repository.NewMembershipRepository(db)
		userMembershipRepository      repository.UserMembershipRepository      = repository.NewUserMembershipRepository(db)
		historyRepository             repository.EntryHistoryRepository        = repository.NewEntryHistory(db)
		userPersonalTrainerRepository repository.UserPersonalTrainerRepository = repository.NewUserPersonalTrainerRepsitory(db)

		userService                services.UserService                = services.NewUserService(userRepository)
		membersipService           services.MembershipService          = services.NewMembershipService(membershipRepository)
		userMembershipService      services.UserMembershipService      = services.NewUserMembershipService(userMembershipRepository, membershipRepository, userRepository, historyRepository, userPersonalTrainerRepository)
		transactionService         services.TransactionService         = services.NewTransactionService(midtransClient, membershipRepository, userMembershipRepository, userRepository)
		historyService             services.EntryHistoryService        = services.NewEntryHistoryService(historyRepository, userRepository)
		userPersonalTrainerService services.UserPersonalTrainerService = services.NewUserPersonalTrainerService(userPersonalTrainerRepository, userRepository)

		userController                controllers.UserController                = controllers.NewUserController(userService)
		membershipController          controllers.MembershipController          = controllers.NewMembershipController(membersipService)
		userMembershipController      controllers.UserMembershipController      = controllers.NewUserMembershipController(userMembershipService)
		transactionController         controllers.TransactionController         = controllers.NewTransactionController(transactionService)
		historyController             controllers.EntryHistoryController        = controllers.NewEntryHistoryController(historyService)
		userPersonalTrainerController controllers.UserPersonalTrainerController = controllers.NewUserPersonalTrainerController(userPersonalTrainerService)
	)

	router.User(engine, userController)
	router.Membership(engine, membershipController, serverMiddleware)
	router.Transaction(engine, transactionController)
	router.UserMembership(engine, userMembershipController, serverMiddleware)
	router.EntryHistory(engine, historyController, serverMiddleware)
	router.UserPersonalTrainer(engine, userPersonalTrainerController)

	return &RestServer{
		Engine: engine,
	}
}

func (s *RestServer) Run(addr string) {
	if err := s.Engine.Run(addr); err != nil {
		panic("Failed to run server: " + err.Error())
	}
}
