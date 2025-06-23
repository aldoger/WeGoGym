package services

import (
	"go-kpl/internal/application/dto"
	"go-kpl/internal/domain/models"
	"go-kpl/internal/domain/repository"
	valueobject "go-kpl/internal/domain/value_object"

	"golang.org/x/net/context"
)

type (
	UserService interface {
		Register(ctx context.Context, req dto.UserRegistrationDto) (dto.UserResponseDto, error)
		Login(ctx context.Context, req dto.UserLoginDto) (dto.UserResponseDto, error)
		GetMeData(ctx context.Context, userId string) (dto.UserResponseDto, error)
	}

	userService struct {
		userRepository repository.UserRepository
	}
)

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) Register(ctx context.Context, req dto.UserRegistrationDto) (dto.UserResponseDto, error) {

	role, err := valueobject.NewUserRole(valueobject.MEMBER_ROLE)
	if err != nil {
		return dto.UserResponseDto{}, err
	}

	createUser, err := s.userRepository.Create(ctx, nil, models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Gender:   models.Gender(req.Gender),
		Role:     role,
	})

	if err != nil {
		return dto.UserResponseDto{}, err
	}

	return dto.UserResponseDto{
		Id:       createUser.Id.String(),
		Email:    createUser.Email,
		Username: createUser.Username,
	}, nil
}

func (s *userService) Login(ctx context.Context, req dto.UserLoginDto) (dto.UserResponseDto, error) {

	findUser, err := s.userRepository.GetByEmail(ctx, nil, req.Email, req.Password)
	if err != nil {
		return dto.UserResponseDto{}, err
	}

	return dto.UserResponseDto{
		Id:       findUser.Id.String(),
		Email:    findUser.Email,
		Role:     findUser.Role.GetRole(),
		Username: findUser.Username,
	}, nil
}

func (s *userService) GetMeData(ctx context.Context, userId string) (dto.UserResponseDto, error) {

	userData, err := s.userRepository.GetById(ctx, nil, userId)
	if err != nil {
		return dto.UserResponseDto{}, err
	}
	//TODO jangan lupa buat interface baru untuk ambil usermembership sekalian
	return dto.UserResponseDto{
		Id:       userData.Id.String(),
		Email:    userData.Email,
		Role:     userData.Role.GetRole(),
		Username: userData.Username,
		UserMembership: dto.UserMembershipResponseDto{
			Id:        userData.UserMembership.Id.String(),
			ExpiredAt: userData.UserMembership.ExpiredAt,
		},
	}, nil
}
