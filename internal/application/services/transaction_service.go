package services

import (
	"go-kpl/infrastructure/externals/midtrans"
	"go-kpl/internal/application/dto"
	"go-kpl/internal/domain/models"
	"go-kpl/internal/domain/repository"
	"time"

	"golang.org/x/net/context"
)

type (
	TransactionService interface {
		CreateTransaction(ctx context.Context, req dto.TransactionRequestDto, email string) (dto.TransactionResponseDto, error)
	}

	transactionService struct {
		midtrans                 *midtrans.MidtransClient
		membershipRepository     repository.MembershipRepository
		userMembershipRepository repository.UserMembershipRepository
		userRepository           repository.UserRepository
	}
)

func NewTransactionService(midtrans *midtrans.MidtransClient, membershipRepository repository.MembershipRepository, userMembershipRepository repository.UserMembershipRepository,
	userRepository repository.UserRepository) TransactionService {
	return &transactionService{midtrans: midtrans, membershipRepository: membershipRepository, userMembershipRepository: userMembershipRepository, userRepository: userRepository}
}

func (m *transactionService) CreateTransaction(ctx context.Context, req dto.TransactionRequestDto, email string) (dto.TransactionResponseDto, error) {

	membershipDetail, err := m.membershipRepository.GetById(ctx, nil, req.MembershipId)
	if err != nil {
		return dto.TransactionResponseDto{}, err
	}

	userData, err := m.userRepository.GetByEmailNoPassword(ctx, nil, email)
	if err != nil {
		return dto.TransactionResponseDto{}, err
	}

	userMembership := models.UserMembership{
		UserId:    userData.Id,
		MemberId:  membershipDetail.Id,
		ExpiredAt: time.Now().Add(time.Hour * 24 * time.Duration(membershipDetail.Duration)),
	}

	_, err = m.userMembershipRepository.Create(ctx, nil, userMembership)
	if err != nil {
		return dto.TransactionResponseDto{}, err
	}

	transaction, err := m.midtrans.CreateMemberTransaction(userData.Id.String(), email, req.Kode, membershipDetail)
	if err != nil {
		return dto.TransactionResponseDto{}, err
	}

	return dto.TransactionResponseDto{
		TokenTransaksi: transaction.Token,
		RedirectUrl:    transaction.RedirectURL,
	}, nil
}
