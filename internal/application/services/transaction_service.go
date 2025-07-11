package services

import (
	"errors"
	"go-kpl/infrastructure/externals/midtrans"
	"go-kpl/internal/application/dto"
	"go-kpl/internal/domain/repository"

	"golang.org/x/net/context"
)

type (
	TransactionService interface {
		CreateMemberTransaction(ctx context.Context, req dto.TransactionMemberRequestDto, email string) (dto.TransactionResponseDto, error)
		CreatePersonalTrainerTransaction(ctx context.Context, req dto.TransactionPersonalTrainerRequestDto, email string) (dto.TransactionResponseDto, error)
	}

	transactionService struct {
		midtrans             *midtrans.MidtransClient
		membershipRepository repository.MembershipRepository
		userRepository       repository.UserRepository
	}
)

func NewTransactionService(midtrans *midtrans.MidtransClient, membershipRepository repository.MembershipRepository, userMembershipRepository repository.UserMembershipRepository,
	userRepository repository.UserRepository) TransactionService {
	return &transactionService{midtrans: midtrans, membershipRepository: membershipRepository, userRepository: userRepository}
}

func (m *transactionService) CreateMemberTransaction(ctx context.Context, req dto.TransactionMemberRequestDto, email string) (dto.TransactionResponseDto, error) {

	membershipDetail, err := m.membershipRepository.GetById(ctx, nil, req.MembershipId)
	if err != nil {
		return dto.TransactionResponseDto{}, err
	}

	userData, err := m.userRepository.GetByEmailNoPassword(ctx, nil, email)
	if err != nil {
		return dto.TransactionResponseDto{}, err
	}

	if userData.IsMember() {
		return dto.TransactionResponseDto{}, errors.New("user already a member")
	}

	transaction, err := m.midtrans.CreateMemberTransaction(email, req.Kode, membershipDetail)
	if err != nil {
		return dto.TransactionResponseDto{}, err
	}

	return dto.TransactionResponseDto{
		TokenTransaksi: transaction.Token,
		RedirectUrl:    transaction.RedirectURL,
	}, nil
}

func (m *transactionService) CreatePersonalTrainerTransaction(ctx context.Context, req dto.TransactionPersonalTrainerRequestDto, email string) (dto.TransactionResponseDto, error) {

	userData, err := m.userRepository.GetByEmailNoPassword(ctx, nil, email)
	if err != nil {
		return dto.TransactionResponseDto{}, err
	}

	if !userData.IsMember() {
		return dto.TransactionResponseDto{}, errors.New("user is not a member")
	}

	transaction, err := m.midtrans.CreatePersonalTrainerTransaction(email, req.Harga, req.Sesi)
	if err != nil {
		return dto.TransactionResponseDto{}, err
	}

	return dto.TransactionResponseDto{
		TokenTransaksi: transaction.Token,
		RedirectUrl:    transaction.RedirectURL,
	}, nil
}
