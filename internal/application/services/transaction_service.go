package services

import (
	"go-kpl/internal/application/dto"

	"golang.org/x/net/context"
)

type (
	Transaction interface {
		CreateTransaction(ctx context.Context, req dto.TransactionDto)
	}

	transaction struct {
	}
)
