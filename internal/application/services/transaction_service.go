package services

import (
	"go-kpl/infrastructure/externals/midtrans"
	"go-kpl/internal/application/dto"
	"go-kpl/internal/domain/repository"

	"golang.org/x/net/context"
)

type (
	TransactionService interface {
		CreateTransaction(ctx context.Context, req dto.TransactionRequestDto) (dto.TransactionResponseDto, error)
	}

	transactionService struct {
		midtrans             *midtrans.MidtransClient
		membershipRepository repository.MembershipRepository
	}
)

func NewTransactionService(midtrans *midtrans.MidtransClient, membershipRepository repository.MembershipRepository) TransactionService {
	return &transactionService{midtrans: midtrans, membershipRepository: membershipRepository}
}

func (m *transactionService) CreateTransaction(ctx context.Context, req dto.TransactionRequestDto) (dto.TransactionResponseDto, error) {

	membershipDetail, err := m.membershipRepository.GetById(ctx, nil, req.MembershipId)
	if err != nil {
		return dto.TransactionResponseDto{}, err
	}

	transaction, err := m.midtrans.CreateMemberTransaction(req.Email, req.Kode, membershipDetail)
	if err != nil {
		return dto.TransactionResponseDto{}, err
	}

	return dto.TransactionResponseDto{
		TokenTransaksi: transaction.Token,
		RedirectUrl:    transaction.RedirectURL,
	}, nil
}
