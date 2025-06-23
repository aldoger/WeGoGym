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
	MembershipController interface {
		CreateMembership(ctx *gin.Context)
		GetAllMembership(ctx *gin.Context)
		GetByIdMembership(ctx *gin.Context)
		UpdateByIdMembership(ctx *gin.Context)
	}
	membershipController struct {
		membershipService services.MembershipService
	}
)

func NewMembershipController(membershipService services.MembershipService) MembershipController {
	return &membershipController{
		membershipService: membershipService,
	}
}

func (c *membershipController) CreateMembership(ctx *gin.Context) {
	var req dto.MembershipRequestDto

	if err := ctx.ShouldBind(&req); err != nil {
		response.NewFailed("failed get data from body",
			myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	membership, err := c.membershipService.CreateMembership(ctx, req)

	if err != nil {
		response.NewFailed("failed to create new membership", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	response.NewSuccess("create new membership success", membership).Send(ctx)
}

func (c *membershipController) GetAllMembership(ctx *gin.Context) {
	memberships, err := c.membershipService.GetAllMembership(ctx)

	if err != nil {
		response.NewFailed("failed to get all memberships", myerror.New(err.Error(), http.StatusInternalServerError)).Send(ctx)
		return
	}

	response.NewSuccess("get all memberships success", memberships).Send(ctx)
}

func (c *membershipController) GetByIdMembership(ctx *gin.Context) {
	id := ctx.Param("id")

	membership, err := c.membershipService.GetByIdMembership(ctx, id)
	if err != nil {
		response.NewFailed(
			"failed to get membership by id",
			myerror.New(err.Error(), http.StatusNotFound),
		).Send(ctx)
		return
	}

	response.NewSuccess("get membership by id success", membership).Send(ctx)
}

func (c *membershipController) UpdateByIdMembership(ctx *gin.Context) {
	id := ctx.Param("id")

	var req dto.UpdateMembershipRequestDto
	if err := ctx.ShouldBind(&req); err != nil {
		response.NewFailed(
			"failed to bind request data",
			myerror.New(err.Error(), http.StatusBadRequest),
		).Send(ctx)
		return
	}

	req.Id = id

	updatedMembership, err := c.membershipService.UpdateByIdMembership(ctx, req)
	if err != nil {
		response.NewFailed(
			"failed to update membership",
			myerror.New(err.Error(), http.StatusBadRequest),
		).Send(ctx)
		return
	}

	response.NewSuccess("update membership success", updatedMembership).Send(ctx)
}
