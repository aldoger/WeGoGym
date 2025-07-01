package repository

import (
	"go-kpl/internal/domain/models"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type (
	UserPersonalTrainerRepository interface {
		Create(ctx context.Context, tx *gorm.DB, userPersonalTrainer models.UserPersonalTrainer) (models.UserPersonalTrainer, error)
		GetUserTrainerById(ctx context.Context, tx *gorm.DB, userId string) (models.UserPersonalTrainer, error)
		UpdateUserPersonalTrainer(ctx context.Context, tx *gorm.DB, userId string, sesi int) (models.UserPersonalTrainer, error)
	}

	userPersonalTrainerRepository struct {
		db *gorm.DB
	}
)

func NewUserPersonalTrainerRepsitory(db *gorm.DB) UserPersonalTrainerRepository {
	return &userPersonalTrainerRepository{db: db}
}

func (r *userPersonalTrainerRepository) Create(ctx context.Context, tx *gorm.DB, userPersonalTrainer models.UserPersonalTrainer) (models.UserPersonalTrainer, error) {

	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(userPersonalTrainer).Error; err != nil {
		return models.UserPersonalTrainer{}, err
	}

	return userPersonalTrainer, nil
}

func (r *userPersonalTrainerRepository) GetUserTrainerById(ctx context.Context, tx *gorm.DB, userId string) (models.UserPersonalTrainer, error) {

	if tx == nil {
		tx = r.db
	}

	var userPT models.UserPersonalTrainer
	if err := tx.WithContext(ctx).Where("user_id = ?", userId).First(&userPT).Error; err != nil {
		return models.UserPersonalTrainer{}, err
	}

	return userPT, nil
}

func (r *userPersonalTrainerRepository) UpdateUserPersonalTrainer(ctx context.Context, tx *gorm.DB, userId string, sesi int) (models.UserPersonalTrainer, error) {

	if tx == nil {
		tx = r.db
	}

	var userPT models.UserPersonalTrainer
	if err := tx.WithContext(ctx).Where("user_id = ?", userId).First(&userPT).Error; err != nil {
		return models.UserPersonalTrainer{}, err
	}

	userPT.Sesi += sesi

	if err := tx.WithContext(ctx).Save(&userPT).Error; err != nil {
		return models.UserPersonalTrainer{}, err
	}

	return userPT, nil

}
