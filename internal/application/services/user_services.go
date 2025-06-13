package services

import (
	"go-kpl/internal/application/dto"
	"go-kpl/internal/domain/repository"

	"golang.org/x/net/context"
)

type (
	UserService interface {
		Register(ctx context.Context, req dto.UserRegistrationDto) (dto.UserResponseDto, error)
		Login(ctx context.Context, req dto.UserLoginDto) (dto.UserResponseDto, error)
		GetByEmail(ctx context.Context, req dto.UserGetByEmailDto) (dto.UserResponseDto, error)
	}

	userService struct {
		userRepository repository.UserRepository
	}
)

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) Register(ctx context.Context, req dto.UserRegistrationDto) (dto.UserResponseDto, error) {
	return dto.UserResponseDto{}, nil
}

func (s *userService) Login(ctx context.Context, req dto.UserLoginDto) (dto.UserResponseDto, error) {
	return dto.UserResponseDto{}, nil
}

func (s *userService) GetByEmail(ctx context.Context, req dto.UserGetByEmailDto) (dto.UserResponseDto, error) {
	return dto.UserResponseDto{}, nil
}
