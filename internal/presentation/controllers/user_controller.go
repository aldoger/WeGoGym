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
		GetMe(ctx *gin.Context)
	}

	userController struct {
		userService services.UserService
	}
)

const MAX_AGE = 259200 // 3 hari

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

	ctx.SetCookie("id", user.Id, MAX_AGE, "/", "", false, true)
	ctx.SetCookie("email", user.Email, MAX_AGE, "/", "", false, true)
	ctx.SetCookie("role", user.Role, MAX_AGE, "/", "", false, true)
	ctx.SetCookie("username", user.Username, MAX_AGE, "/", "", false, true)

	response.NewSuccess("login successfully", user).Send(ctx)
}

func (c *userController) GetMe(ctx *gin.Context) {

	userId, err := ctx.Cookie("id")
	if err != nil {
		response.NewFailed("Id user not found in cookie", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	user, err := c.userService.GetMeData(ctx, userId)
	if err != nil {
		response.NewFailed("failed to retrive user data", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
	}

	response.NewSuccess("data successfuly retrive", user).Send(ctx)
}
