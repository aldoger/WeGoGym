package repository

import (
	"go-kpl/internal/domain/models"

	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type UserMembershipRepositoryMock struct {
	Mock mock.Mock
}

func (m *UserMembershipRepositoryMock) Create(ctx context.Context, tx *gorm.DB, userMembership models.UserMembership) (models.UserMembership, error) {
	args := m.Mock.Called(ctx, tx, userMembership)
	return args.Get(0).(models.UserMembership), args.Error(1)
}

func (m *UserMembershipRepositoryMock) SearchMember(ctx context.Context, tx *gorm.DB, userId string) (string, error) {
	args := m.Mock.Called(ctx, tx, userId)
	return args.String(0), args.Error(1)
}
