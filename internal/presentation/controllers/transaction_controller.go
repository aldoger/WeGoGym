package controllers

import (
	"fmt"
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
		TransactionNotification(ctx *gin.Context)
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

	if err := ctx.ShouldBind(&req); err != nil {
		response.NewFailed("failed get data from body", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	transaction, err := c.transactionService.CreateTransaction(ctx, req)
	if err != nil {
		response.NewFailed("Transaction failed to process", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	response.NewSuccess("Transation successfully process", transaction).Send(ctx)
}

func (c *transactionController) TransactionNotification(ctx *gin.Context) {
	var payload map[string]interface{}

	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON: " + err.Error(),
		})
		return
	}

	fmt.Println("Webhook diterima:")
	for key, val := range payload {
		fmt.Printf("%s: %v\n", key, val)
	}

	handleWebhook(payload)

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func handleWebhook(data map[string]interface{}) {
	// Contoh logika: ambil status transaksi
	if status, ok := data["transaction_status"]; ok {
		fmt.Println("Status Transaksi:", status)
	}
}
