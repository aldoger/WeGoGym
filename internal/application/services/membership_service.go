package services

import (
	"go-kpl/internal/application/dto"
	"go-kpl/internal/domain/models"
	"go-kpl/internal/domain/repository"

	"golang.org/x/net/context"
)

type (
	MembershipService interface {
		CreateMembership(ctx context.Context, req dto.MembershipRequestDto) (dto.MembershipResponseDto, error)
		GetAllMembership(ctx context.Context) ([]dto.MembershipResponseDto, error)
		GetByIdMembership(ctx context.Context, id string) (dto.MembershipResponseDto, error)
		UpdateByIdMembership(ctx context.Context,  req dto.UpdateMembershipRequestDto) (dto.MembershipResponseDto, error)
	}

	membershipService struct {
		membershipRepository repository.MembershipRepository
	}
)

func NewMembershipService(membershipRepository repository.MembershipRepository) MembershipService {
	return &membershipService{membershipRepository: membershipRepository}
}

func (s *membershipService) CreateMembership(ctx context.Context, req dto.MembershipRequestDto) (dto.MembershipResponseDto, error) {
	membership := models.Membership{
		Type: req.Type,
		Price: req.Price,
		Duration: req.Duration,
	}
	
	newMembership, err := s.membershipRepository.Create(ctx, nil, membership)

	if err != nil {
		return dto.MembershipResponseDto{}, err
	}
	return dto.MembershipResponseDto{
		Id: newMembership.Id.String(),
		Type: newMembership.Type,
		Price: newMembership.Price,
		Duration: newMembership.Duration,
	}, nil
}

func (s *membershipService) GetAllMembership(ctx context.Context) ([]dto.MembershipResponseDto, error) {
	
	allMembership, err := s.membershipRepository.GetAll(ctx, nil)

	if err != nil {
		return []dto.MembershipResponseDto{}, err
	}

	var responses []dto.MembershipResponseDto
	for _, m := range allMembership {
		responses = append(responses, dto.MembershipResponseDto{
			Id:       m.Id.String(),
			Type:     m.Type,
			Duration: m.Duration,
			Price:    m.Price,
		})
	}

	return responses, nil
}

func (s *membershipService) GetByIdMembership(ctx context.Context, id string) (dto.MembershipResponseDto, error) {
	 
	membershipById, err := s.membershipRepository.GetById(ctx, nil, id)

	if err != nil {
		return dto.MembershipResponseDto{}, err
	}

	response := dto.MembershipResponseDto{
		Id:       membershipById.Id.String(),
		Type:     membershipById.Type,
		Duration: membershipById.Duration,
		Price:    membershipById.Price,
	}

	return response, nil

}

func (s *membershipService) UpdateByIdMembership(ctx context.Context, req dto.UpdateMembershipRequestDto) (dto.MembershipResponseDto, error) {
	
	updateMembership, err := s.membershipRepository.UpdateById(ctx, nil, req.Id, &req.Duration, &req.Price)

	if err != nil {
		return dto.MembershipResponseDto{}, err
	}

	return dto.MembershipResponseDto{
		Id: updateMembership.Id.String(),
		Type: updateMembership.Type,
		Duration: updateMembership.Duration,
		Price: updateMembership.Price,
	}, nil
	
}

