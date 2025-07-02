package repository

import (
	"go-kpl/internal/domain/models"

	"github.com/google/uuid"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		Create(ctx context.Context, tx *gorm.DB, user models.User) (models.User, error)
		GetById(ctx context.Context, tx *gorm.DB, userId string) (models.User, error)
		GetByEmail(ctx context.Context, tx *gorm.DB, email string, password string) (models.User, error)
		GetByEmailNoPassword(ctx context.Context, tx *gorm.DB, email string) (models.User, error)
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, tx *gorm.DB, user models.User) (models.User, error) {

	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) GetById(ctx context.Context, tx *gorm.DB, userId string) (models.User, error) {

	if tx == nil {
		tx = r.db
	}

	var user models.User
	if err := tx.WithContext(ctx).Preload("UserPT").
		Preload("UserMembership").
		Take(&user, "id = ?", userId).Error; err != nil {
		return models.User{}, err
	}

	if user.UserMembership != nil && user.UserMembership.Id == uuid.Nil {
		user.UserMembership = nil
	}

	if user.UserPT != nil && user.UserPT.Id == uuid.Nil {
		user.UserPT = nil
	}

	return user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, tx *gorm.DB, email string, password string) (models.User, error) {

	if tx == nil {
		tx = r.db
	}

	var user models.User
	if err := tx.WithContext(ctx).Take(&user, "email = ? AND password = ?", email, password).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetByEmailNoPassword(ctx context.Context, tx *gorm.DB, email string) (models.User, error) {
	if tx == nil {
		tx = r.db
	}

	var user models.User
	if err := tx.WithContext(ctx).Preload("UserMembership").Take(&user, "email = ?", email).Error; err != nil {
		return models.User{}, err
	}

	if user.UserMembership != nil && user.UserMembership.Id == uuid.Nil {
		user.UserMembership = nil
	}

	return user, nil
}
