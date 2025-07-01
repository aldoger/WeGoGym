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
	UserPersonalTrainerController interface {
		CreateUserPersonalTrainer(ctx *gin.Context)
	}

	userPersonalTrainerController struct {
		userPersonalService services.UserPersonalTrainerService
	}
)

func NewUserPersonalTrainerController(userPersonalTrainerService services.UserPersonalTrainerService) UserPersonalTrainerController {
	return &userPersonalTrainerController{userPersonalService: userPersonalTrainerService}
}

func (c *userPersonalTrainerController) CreateUserPersonalTrainer(ctx *gin.Context) {

	var req dto.CreateUserPersonalTrainerDto
	if err := ctx.ShouldBind(&req); err != nil {
		response.NewFailed("failed get data from body", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	userPT, err := c.userPersonalService.CreateUserPersonalTrainer(ctx, req)
	if err != nil {
		response.NewFailed("failed to make user personal trainer", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	response.NewSuccess("successfuly make user personal trainer", userPT).Send(ctx)
}
