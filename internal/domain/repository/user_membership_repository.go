package repository

import (
	"go-kpl/internal/domain/models"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type (
	UserMembershipRepository interface {
		Create(ctx context.Context, tx *gorm.DB, userMembership models.UserMembership) (models.UserMembership, error)
		SearchMember(ctx context.Context, tx *gorm.DB, userId string) (string, error)
		UpdateMember(ctx context.Context, tx *gorm.DB, userId string) (models.UserMembership, error)
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

func (r *userMembershipRepository) SearchMember(ctx context.Context, tx *gorm.DB, userId string) (string, error) {

	if tx == nil {
		tx = r.db
	}

	var UserMembership models.UserMembership
	if err := tx.WithContext(ctx).Where("user_id = ?", userId).First(&UserMembership).Error; err != nil {
		return "", err
	}

	return UserMembership.UserId.String(), nil
}

func (r *userMembershipRepository) UpdateMember(ctx context.Context, tx *gorm.DB, userId string) (models.UserMembership, error) {
	if tx == nil {
		tx = r.db
	}

	var userMembership models.UserMembership
	if err := tx.WithContext(ctx).Where("user_id = ?", userId).First(&userMembership).Error; err != nil {
		return models.UserMembership{}, err
	}

	userMembership.Verified = true

	if err := tx.WithContext(ctx).Save(&userMembership).Error; err != nil {
		return models.UserMembership{}, err
	}

	return userMembership, nil
}
