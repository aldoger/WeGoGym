package repository

import (
	"context"

	"go-kpl/internal/domain/models"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MembershipRepositoryMock struct {
	Mock mock.Mock
}

func (m *MembershipRepositoryMock) Create(ctx context.Context, tx *gorm.DB, membership models.Membership) (models.Membership, error) {
	args := m.Mock.Called(ctx, tx, membership)
	return args.Get(0).(models.Membership), args.Error(1)
}

func (m *MembershipRepositoryMock) GetAll(ctx context.Context, tx *gorm.DB) ([]models.Membership, error) {
	args := m.Mock.Called(ctx, tx)
	return args.Get(0).([]models.Membership), args.Error(1)
}

func (m *MembershipRepositoryMock) GetById(ctx context.Context, tx *gorm.DB, id string) (models.Membership, error) {
	args := m.Mock.Called(ctx, tx, id)
	return args.Get(0).(models.Membership), args.Error(1)
}

func (m *MembershipRepositoryMock) UpdateById(ctx context.Context, tx *gorm.DB, id string, duration *int, price *int) (models.Membership, error) {
	args := m.Mock.Called(ctx, tx, id, duration, price)
	return args.Get(0).(models.Membership), args.Error(1)
}
