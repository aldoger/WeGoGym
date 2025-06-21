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
		CreateTransaction(ctx *gin.Context)
	}

	transactionController struct {
		transactionService services.TransactionService
	}
)

func NewTransactionController(transactionService services.TransactionService) TransactionController {
	return &transactionController{transactionService: transactionService}
}

func (c *transactionController) CreateTransaction(ctx *gin.Context) {

	var req dto.TransactionRequestDto

	UserEmail, err := ctx.Cookie("email")
	if err != nil {
		response.NewFailed("Id user not found in cookie", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	if err := ctx.ShouldBind(&req); err != nil {
		response.NewFailed("failed get data from body", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	transaction, err := c.transactionService.CreateTransaction(ctx, req, UserEmail)
	if err != nil {
		response.NewFailed("Transaction failed to process", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	response.NewSuccess("Transation successfully process", transaction)
}
