package services

import (
	"errors"
	"go-kpl/internal/application/dto"
	"go-kpl/internal/domain/models"
	"go-kpl/internal/domain/repository"
	"time"

	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type (
	UserMembershipService interface {
		CreateUserMembership(ctx context.Context, req dto.CreateUserMembershipRequestDto, userId string) (dto.UserMembershipResponseDto, error)
		SearchMembership(ctx context.Context, userId string) (dto.UserSearchMembershipResponse, error)
	}

	userMembershipService struct {
		userMembershipRepository repository.UserMembershipRepository
		membershipRepository     repository.MembershipRepository
		userRepository           repository.UserRepository
		entryHistoryRepository   repository.EntryHistoryRepository
		userPTRepository         repository.UserPersonalTrainerRepository
	}
)

func NewUserMembershipService(userMembershipRepository repository.UserMembershipRepository, membershipRepository repository.MembershipRepository,
	userRepository repository.UserRepository, entryHistoryRepository repository.EntryHistoryRepository, userPTRepository repository.UserPersonalTrainerRepository) UserMembershipService {
	return &userMembershipService{userMembershipRepository: userMembershipRepository, membershipRepository: membershipRepository, userRepository: userRepository, entryHistoryRepository: entryHistoryRepository,
		userPTRepository: userPTRepository}
}

func (s *userMembershipService) CreateUserMembership(ctx context.Context, req dto.CreateUserMembershipRequestDto, userId string) (dto.UserMembershipResponseDto, error) {

	membershipDetail, err := s.membershipRepository.GetById(ctx, nil, req.MembershipId)
	if err != nil {
		return dto.UserMembershipResponseDto{}, err
	}

	userData, err := s.userRepository.GetById(ctx, nil, userId)
	if err != nil {
		return dto.UserMembershipResponseDto{}, err
	}

	if userData.IsMember() {
		return dto.UserMembershipResponseDto{}, errors.New("user already a member")
	}

	memberUUID, err := uuid.Parse(req.MembershipId)
	if err != nil {
		return dto.UserMembershipResponseDto{}, err
	}

	userMembership := models.UserMembership{
		UserId:    userData.Id,
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

func (s *userMembershipService) SearchMembership(ctx context.Context, userId string) (dto.UserSearchMembershipResponse, error) {

	UserId, err := s.userMembershipRepository.SearchMember(ctx, nil, userId)
	if err != nil {
		return dto.UserSearchMembershipResponse{}, err
	}

	parsedUserId, err := uuid.Parse(userId)
	if err != nil {
		return dto.UserSearchMembershipResponse{}, err
	}

	userPT, err := s.userPTRepository.UseSession(ctx, nil, parsedUserId)
	if err != nil {
		return dto.UserSearchMembershipResponse{}, err
	}

	if err := s.entryHistoryRepository.AddEntry(ctx, nil, models.EntryHistory{
		UserId:    parsedUserId,
		EntryTime: time.Now(),
	}); err != nil {
		return dto.UserSearchMembershipResponse{}, err
	}

	User, err := s.userRepository.GetById(ctx, nil, UserId)
	if err != nil {
		return dto.UserSearchMembershipResponse{}, err
	}

	return dto.UserSearchMembershipResponse{
		Id:       User.Id.String(),
		Username: User.Username,
		Email:    User.Email,
		Role:     User.Role.GetRole(),
		UserMembership: dto.UserMembershipResponseDto{
			Id:        User.UserMembership.Id.String(),
			ExpiredAt: User.UserMembership.ExpiredAt,
		},
		Sesi: userPT,
	}, nil
}
