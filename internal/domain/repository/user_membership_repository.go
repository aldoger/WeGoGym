package repository

import (
	"go-kpl/internal/domain/models"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type (
	UserMembershipRepository interface {
		Create(ctx context.Context, tx *gorm.DB, userMembership models.UserMembership) (models.UserMembership, error)
	}

	userMembershipRepository struct {
		db *gorm.DB
	}
)

func NewUserMembershipRepository(db *gorm.DB) UserMembershipRepository {
	return &userMembershipRepository{db: db}
}

func (r *userMembershipRepository) Create(ctx context.Context, tx *gorm.DB, userMembership models.UserMembership) (models.UserMembership, error) {

	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&userMembership).Error; err != nil {
		return models.UserMembership{}, err
	}

	return userMembership, nil
}
