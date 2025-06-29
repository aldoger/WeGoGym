package services

import (
	"context"
	"go-kpl/internal/application/dto"
	"go-kpl/internal/domain/models"
	"go-kpl/internal/domain/repository"
	valueobject "go-kpl/internal/domain/value_object"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = &repository.UserRepositoryMock{Mock: mock.Mock{}}
var userServiceTest = userService{userRepository: userRepository}

func TestUserGetById(t *testing.T) {
	ctx := context.Background()
	expectedUUID := uuid.New()

	mockUser := models.User{
		Id:    expectedUUID,
		Email: "test@example.com",
		Role:  valueobject.NewUserRole(valueobject.MEMBER_ROLE),
	}

	userRepository.Mock.On("GetById", mock.Anything, mock.Anything, expectedUUID.String()).
		Return(mockUser, nil)

	result, err := userServiceTest.GetMeDataById(ctx, expectedUUID.String())

	assert.NoError(t, err)
	assert.Equal(t, expectedUUID.String(), result.Id)
	assert.Equal(t, "test@example.com", result.Email)
	assert.Equal(t, "member", result.Role)

	userRepository.Mock.AssertExpectations(t)
}

func TestUserServiceRegister(t *testing.T) {
	ctx := context.Background()

	input := dto.UserRegistrationDto{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "securepassword",
		Gender:   "laki-laki",
	}

	createdId := uuid.New()

	mockUser := models.User{
		Id:       createdId,
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
		Gender:   models.Gender(input.Gender),
		Role:     valueobject.NewUserRole(valueobject.MEMBER_ROLE),
	}

	userRepository.Mock.On("Create", mock.Anything, mock.Anything, mock.MatchedBy(func(u models.User) bool {
		return u.Username == input.Username &&
			u.Email == input.Email &&
			u.Password == input.Password &&
			string(u.Gender) == input.Gender &&
			u.Role.Value() == valueobject.MEMBER_ROLE
	})).Return(mockUser, nil)

	result, err := userServiceTest.Register(ctx, input)

	assert.NoError(t, err)
	assert.Equal(t, createdId.String(), result.Id)
	assert.Equal(t, input.Email, result.Email)
	assert.Equal(t, input.Username, result.Username)

	userRepository.Mock.AssertExpectations(t)
}
