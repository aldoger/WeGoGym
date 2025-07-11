package repository

import (
	"errors"
	"go-kpl/internal/domain/models"

	"github.com/google/uuid"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type (
	UserPersonalTrainerRepository interface {
		Create(ctx context.Context, tx *gorm.DB, userPersonalTrainer models.UserPersonalTrainer) (models.UserPersonalTrainer, error)
		GetByUserId(ctx context.Context, tx *gorm.DB, userId uuid.UUID) (*models.UserPersonalTrainer, error)
		AddSession(ctx context.Context, tx *gorm.DB, userPersonalTrainer *models.UserPersonalTrainer, sesi int) (models.UserPersonalTrainer, error)
		UseSession(ctx context.Context, tx *gorm.DB, userId uuid.UUID) (int, error)
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

	if err := tx.WithContext(ctx).Create(&userPersonalTrainer).Error; err != nil {
		return models.UserPersonalTrainer{}, err
	}

	return userPersonalTrainer, nil
}

func (r *userPersonalTrainerRepository) UseSession(ctx context.Context, tx *gorm.DB, userId uuid.UUID) (int, error) {
	if tx == nil {
		tx = r.db
	}

	var userPT models.UserPersonalTrainer
	if err := tx.WithContext(ctx).Where("user_id = ?", userId).First(&userPT).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return 0, err
	}

	if userPT.Sesi <= 0 {
		return 0, nil
	}

	userPT.UsedSession()

	if err := tx.WithContext(ctx).Save(&userPT).Error; err != nil {
		return 0, err
	}

	return userPT.Sesi, nil
}

func (r *userPersonalTrainerRepository) GetByUserId(ctx context.Context, tx *gorm.DB, userId uuid.UUID) (*models.UserPersonalTrainer, error) {
	if tx == nil {
		tx = r.db
	}

	var userPT models.UserPersonalTrainer
	if err := tx.WithContext(ctx).Take(&userPT, "user_id = ?", userId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &userPT, nil
}

func (r *userPersonalTrainerRepository) AddSession(ctx context.Context, tx *gorm.DB, userPersonalTrainer *models.UserPersonalTrainer, sesi int) (models.UserPersonalTrainer, error) {

	if tx == nil {
		tx = r.db
	}

	userPersonalTrainer.AddSession(sesi)

	if err := tx.WithContext(ctx).Save(userPersonalTrainer).Error; err != nil {
		return models.UserPersonalTrainer{}, err
	}

	return *userPersonalTrainer, nil
}
