package services

import (
	"go-kpl/internal/application/dto"
	"go-kpl/internal/domain/models"
	"go-kpl/internal/domain/repository"
	"time"

	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type (
	UserMembershipService interface {
		CreateUserMembership(ctx context.Context, req dto.CreateUserMembershipRequestDto) (dto.UserMembershipResponseDto, error)
	}

	userMembershipService struct {
		userMembershipRepository repository.UserMembershipRepository
		membershipRepository     repository.MembershipRepository
	}
)

func NewUserMembershipService(userMembershipRepository repository.UserMembershipRepository, membershipRepository repository.MembershipRepository) UserMembershipService {
	return &userMembershipService{userMembershipRepository: userMembershipRepository, membershipRepository: membershipRepository}
}

func (s *userMembershipService) CreateUserMembership(ctx context.Context, req dto.CreateUserMembershipRequestDto) (dto.UserMembershipResponseDto, error) {

	membershipDetail, err := s.membershipRepository.GetById(ctx, nil, req.MembershipId)
	if err != nil {
		return dto.UserMembershipResponseDto{}, err
	}

	userUUID, err := uuid.Parse(req.UserId)
	if err != nil {
		return dto.UserMembershipResponseDto{}, err
	}

	memberUUID, err := uuid.Parse(req.MembershipId)
	if err != nil {
		return dto.UserMembershipResponseDto{}, err
	}

	userMembership := models.UserMembership{
		UserId:    userUUID,
		MemberId:  memberUUID,
		ExpiredAt: time.Now().Add(time.Hour * 24 * time.Duration(membershipDetail.Duration)),
	}

	newUserMembership, err := s.userMembershipRepository.Create(ctx, nil, userMembership)
	if err != nil {
		return dto.UserMembershipResponseDto{}, err
	}

	return dto.UserMembershipResponseDto{
		Id:        newUserMembership.Id.String(),
		ExpiredAt: newUserMembership.ExpiredAt,
	}, nil
}
