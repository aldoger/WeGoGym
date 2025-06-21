package services

import (
	"go-kpl/infrastructure/externals/midtrans"
	"go-kpl/internal/application/dto"

	"golang.org/x/net/context"
)

type (
	TransactionService interface {
		CreateTransaction(ctx context.Context, req dto.TransactionRequestDto) (dto.TransactionResponseDto, error)
	}

	transactionService struct {
		midtrans midtrans.MidtransClient
	}
)

func NewTransactionService(midtrans midtrans.MidtransClient) TransactionService {
	return &transactionService{midtrans: midtrans}
}

func (m *transactionService) CreateTransaction(ctx context.Context, req dto.TransactionRequestDto) (dto.TransactionResponseDto, error) {

	transaction, err := m.midtrans.CreateTransaction(req.Kode)
	if err != nil {
		return dto.TransactionResponseDto{}, err
	}

	return dto.TransactionResponseDto{
		TokenTransaksi: transaction.Token,
		UrlRedirect:    transaction.RedirectURL,
	}, nil
}
