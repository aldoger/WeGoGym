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
		TransactionNotification(ctx *gin.Context)
	}

	transactionController struct {
		transactionService    services.TransactionService
		userService           services.UserService
		userMembershipService services.UserMembershipService
	}
)

func NewTransactionController(transactionService services.TransactionService, userMembershipService services.UserMembershipService,
	userService services.UserService) TransactionController {
	return &transactionController{transactionService: transactionService, userMembershipService: userMembershipService, userService: userService}
}

func (c *transactionController) CreateTransaction(ctx *gin.Context) {

	userEmail, err := ctx.Cookie("email")
	if err != nil {
		response.NewFailed("failed to get data from cookie", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	var req dto.TransactionRequestDto

	if err := ctx.ShouldBind(&req); err != nil {
		response.NewFailed("failed get data from body", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	transaction, err := c.transactionService.CreateTransaction(ctx, req, userEmail)
	if err != nil {
		response.NewFailed("Transaction failed to process", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	response.NewSuccess("Transation successfully process", transaction).Send(ctx)
}

func (c *transactionController) TransactionNotification(ctx *gin.Context) {
	var payload map[string]interface{}

	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	status, _ := payload["transaction_status"].(string)
	if status != "settlement" {
		ctx.JSON(http.StatusOK, gin.H{"status": "not processed"})
		return
	}

	orderId, ok := payload["order_id"].(string)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "order_id missing"})
		return
	}

	err := c.userMembershipService.UpdateMembership(ctx, orderId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "update user membership failed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "user membership is verified"})
}
