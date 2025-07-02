package services

import (
	"errors"
	"go-kpl/internal/application/dto"
	"go-kpl/internal/domain/models"
	"go-kpl/internal/domain/repository"

	"golang.org/x/net/context"
)

type (
	UserPersonalTrainerService interface {
		NewUserPersonalTrainerSesi(ctx context.Context, req dto.CreateUserPersonalTrainerDto, userId string) (dto.UserPersonalTrainerResponse, error)
	}

	userPersonalTrainerService struct {
		userPersonalTrainerRepository repository.UserPersonalTrainerRepository
		userRepository                repository.UserRepository
	}
)

func NewUserPersonalTrainerService(userPersonalRepository repository.UserPersonalTrainerRepository, userRepository repository.UserRepository) UserPersonalTrainerService {
	return &userPersonalTrainerService{userPersonalTrainerRepository: userPersonalRepository, userRepository: userRepository}
}

func (s *userPersonalTrainerService) NewUserPersonalTrainerSesi(ctx context.Context, req dto.CreateUserPersonalTrainerDto, userId string) (dto.UserPersonalTrainerResponse, error) {

	userData, err := s.userRepository.GetById(ctx, nil, userId)
	if err != nil {
		return dto.UserPersonalTrainerResponse{}, err
	}

	if !userData.IsMember() {
		return dto.UserPersonalTrainerResponse{}, errors.New("user is not a member")
	}

	existingUserPT, err := s.userPersonalTrainerRepository.GetByUserId(ctx, nil, userData.Id)
	if err != nil {
		return dto.UserPersonalTrainerResponse{}, err
	}

	if existingUserPT != nil {
		addUserSesi, err := s.userPersonalTrainerRepository.AddSession(ctx, nil, existingUserPT, req.Sesi)
		if err != nil {
			return dto.UserPersonalTrainerResponse{}, err
		}

		return dto.UserPersonalTrainerResponse{
			Sesi:   addUserSesi.Sesi,
			UserId: addUserSesi.UserId.String(),
		}, nil
	}

	userPT := models.UserPersonalTrainer{
		Sesi:   req.Sesi,
		UserId: userData.Id,
	}
	userPersonalTrainer, err := s.userPersonalTrainerRepository.Create(ctx, nil, userPT)
	if err != nil {
		return dto.UserPersonalTrainerResponse{}, err
	}

	return dto.UserPersonalTrainerResponse{
		Sesi:   userPersonalTrainer.Sesi,
		UserId: userPersonalTrainer.UserId.String(),
	}, nil
}
