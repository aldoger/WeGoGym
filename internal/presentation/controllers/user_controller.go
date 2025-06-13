package controllers

import (
	"go-kpl/internal/application/dto"
	"go-kpl/internal/application/services"
	myerror "go-kpl/internal/pkg/errors"
	"go-kpl/internal/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	UserController interface {
		Register(ctx *gin.Context)
		Login(ctx *gin.Context)
	}

	userController struct {
		userService services.UserService
	}
)

func NewUserController(userService services.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (c *userController) Register(ctx *gin.Context) {
	var req dto.UserRegistrationDto
	if err := ctx.ShouldBind(&req); err != nil {
		response.NewFailed("failed get data from body", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	user, err := c.userService.Register(ctx, req)
	if err != nil {
		response.NewFailed("failed to register", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	response.NewSuccess("registration success", user).Send(ctx)
}

func (c *userController) Login(ctx *gin.Context) {
	var req dto.UserLoginDto
	if err := ctx.ShouldBind(&req); err != nil {
		response.NewFailed("failed get data from body", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	user, err := c.userService.Login(ctx, req)
	if err != nil {
		response.NewFailed("failed to login", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	//TODO cookie send

	response.NewSuccess("login successfully", user).Send(ctx)
}
