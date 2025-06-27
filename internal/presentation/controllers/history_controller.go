package controllers

import (
	"go-kpl/internal/application/services"
	myerror "go-kpl/internal/pkg/errors"
	"go-kpl/internal/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	EntryHistoryController interface {
		GetEntryAll(ctx *gin.Context)
		GetEntryByUserId(ctx *gin.Context)
	}

	entryHistoryController struct {
		entryHistoryService services.EntryHistoryService
	}
)

func NewEntryHistoryController(entryHistoryService services.EntryHistoryService) EntryHistoryController {
	return &entryHistoryController{entryHistoryService: entryHistoryService}
}

func (c *entryHistoryController) GetEntryAll(ctx *gin.Context) {

	allEntry, err := c.entryHistoryService.GetEntryAll(ctx)
	if err != nil {
		response.NewFailed("failed to get all entry", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	response.NewSuccess("all entry succesfuly retrieve", allEntry).Send(ctx)
}

func (c *entryHistoryController) GetEntryByUserId(ctx *gin.Context) {

	userId := ctx.Param("id")
	if userId == "" {
		response.NewFailed("failed to get id", myerror.New("id not provided", http.StatusBadRequest)).Send(ctx)
		return
	}

	userEntry, err := c.entryHistoryService.GetEntryByUserId(ctx, userId)
	if err != nil {
		response.NewFailed("failed to get user entry", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	response.NewSuccess("user entry successfuly retrieve", userEntry).Send(ctx)
}
