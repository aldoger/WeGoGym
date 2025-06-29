package repository

import (
	"go-kpl/internal/domain/models"

	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (m *UserRepositoryMock) Create(ctx context.Context, tx *gorm.DB, user models.User) (models.User, error) {
	args := m.Mock.Called(ctx, tx, user)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *UserRepositoryMock) GetById(ctx context.Context, tx *gorm.DB, userId string) (models.User, error) {
	args := m.Mock.Called(ctx, tx, userId)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *UserRepositoryMock) GetByEmail(ctx context.Context, tx *gorm.DB, email string, password string) (models.User, error) {
	args := m.Mock.Called(ctx, tx, email, password)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *UserRepositoryMock) GetByEmailNoPassword(ctx context.Context, tx *gorm.DB, email string) (models.User, error) {
	args := m.Mock.Called(ctx, tx, email)
	return args.Get(0).(models.User), args.Error(1)
}
