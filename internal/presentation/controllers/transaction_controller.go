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
	TransactionController interface {
		CreateMemberTransaction(ctx *gin.Context)
		CreatePersonalTrainerTransaction(ctx *gin.Context)
	}

	transactionController struct {
		transactionService services.TransactionService
	}
)

func NewTransactionController(transactionService services.TransactionService) TransactionController {
	return &transactionController{transactionService: transactionService}
}

func (c *transactionController) CreateMemberTransaction(ctx *gin.Context) {

	userEmail, err := ctx.Cookie("email")
	if err != nil {
		response.NewFailed("failed to get data from cookie", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	var req dto.TransactionMemberRequestDto

	if err := ctx.ShouldBind(&req); err != nil {
		response.NewFailed("failed get data from body", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	transaction, err := c.transactionService.CreateMemberTransaction(ctx, req, userEmail)
	if err != nil {
		response.NewFailed("Transaction failed to process", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	response.NewSuccess("Transation successfully process", transaction).Send(ctx)
}

func (c *transactionController) CreatePersonalTrainerTransaction(ctx *gin.Context) {

	userEmail, err := ctx.Cookie("email")
	if err != nil {
		response.NewFailed("failed to get data from cookie", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	var req dto.TransactionPersonalTrainerRequestDto

	if err := ctx.ShouldBind(&req); err != nil {
		response.NewFailed("failed get data from body", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	transaction, err := c.transactionService.CreatePersonalTrainerTransaction(ctx, req, userEmail)
	if err != nil {
		response.NewFailed("Transaction failed to process", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	response.NewSuccess("Transaction successfully process", transaction).Send(ctx)
}
