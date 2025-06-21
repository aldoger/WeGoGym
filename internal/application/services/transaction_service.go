package services

import (
	"go-kpl/infrastructure/externals/midtrans"
	"go-kpl/internal/application/dto"

	"golang.org/x/net/context"
)

type (
	TransactionService interface {
		CreateTransaction(ctx context.Context, req dto.TransactionRequestDto, email string) (dto.TransactionResponseDto, error)
	}

	transactionService struct {
		midtrans midtrans.MidtransClient
	}
)

func NewTransactionService(midtrans midtrans.MidtransClient) TransactionService {
	return &transactionService{midtrans: midtrans}
}

func (m *transactionService) CreateTransaction(ctx context.Context, req dto.TransactionRequestDto, email string) (dto.TransactionResponseDto, error) {

	transactionToken, err := m.midtrans.CreateTransaction(req.Price, email, req.Kode)
	if err != nil {
		return dto.TransactionResponseDto{}, err
	}

	return dto.TransactionResponseDto{
		TokenTransaksi: transactionToken,
	}, nil
}
