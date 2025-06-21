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
	UserMembershipController interface {
		CreateUserMembership(ctx *gin.Context)
	}

	userMembershipController struct {
		userMembershipService services.UserMembershipService
	}
)

func NewUserMembershipController(userMembershipService services.UserMembershipService) UserMembershipController {
	return &userMembershipController{userMembershipService: userMembershipService}
}

func (c *userMembershipController) CreateUserMembership(ctx *gin.Context) {
	var req dto.CreateUserMembershipRequestDto
	if err := ctx.ShouldBind(&req); err != nil {
		response.NewFailed("failed get data from body", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	userMembership, err := c.userMembershipService.CreateUserMembership(ctx, req)
	if err != nil {
		response.NewFailed("failed to create user membership", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	response.NewSuccess("user membership created", userMembership).Send(ctx)
}
