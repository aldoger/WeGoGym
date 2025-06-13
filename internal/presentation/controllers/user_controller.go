package controllers

import (
	"go-kpl/internal/application/services"

	"github.com/gin-gonic/gin"
)

type (
	UserController interface {
		Register(ctx *gin.Context)
		Login(ctx *gin.Context)
		GetByEmail(ctx *gin.Context)
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

}

func (c *userController) Login(ctx *gin.Context) {

}

func (c *userController) GetByEmail(ctx *gin.Context) {

}
