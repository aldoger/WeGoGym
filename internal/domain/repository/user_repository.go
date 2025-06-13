package repository

import (
	"go-kpl/internal/domain/models"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		Create(ctx context.Context, tx *gorm.DB, user models.User) (models.User, error)
		GetById(ctx context.Context, tx *gorm.DB, userId string) (models.User, error)
		GetByEmail(ctx context.Context, tx *gorm.DB, email string) (models.User, error)
		Delete(ctx context.Context, tx *gorm.DB, userId string) (models.User, error)
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, tx *gorm.DB, user models.User) (models.User, error) {
	return models.User{}, nil
}

func (r *userRepository) GetById(ctx context.Context, tx *gorm.DB, userId string) (models.User, error) {
	return models.User{}, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, tx *gorm.DB, email string) (models.User, error) {
	return models.User{}, nil
}

func (r *userRepository) Delete(ctx context.Context, tx *gorm.DB, userId string) (models.User, error) {
	return models.User{}, nil
}
