package repository

import (
	"go-kpl/internal/domain/models"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type (
	EntryHistoryRepository interface {
		AddEntry(ctx context.Context, tx *gorm.DB, entryHistory models.EntryHistory) error
		GetEntryAll(ctx context.Context, tx *gorm.DB) ([]models.EntryHistory, error)
		GetEntryByUserId(ctx context.Context, tx *gorm.DB, userId string) ([]models.EntryHistory, error)
	}

	entryHistoryRepository struct {
		db *gorm.DB
	}
)

func NewEntryHistory(db *gorm.DB) EntryHistoryRepository {
	return &entryHistoryRepository{db: db}
}

func (r *entryHistoryRepository) AddEntry(ctx context.Context, tx *gorm.DB, entryHistory models.EntryHistory) error {

	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&entryHistory).Error; err != nil {
		return err
	}

	return nil
}

func (r *entryHistoryRepository) GetEntryAll(ctx context.Context, tx *gorm.DB) ([]models.EntryHistory, error) {

	if tx == nil {
		tx = r.db
	}

	var entryHistory []models.EntryHistory
	if err := tx.WithContext(ctx).Find(&entryHistory).Error; err != nil {
		return []models.EntryHistory{}, err
	}

	return entryHistory, nil
}

func (r *entryHistoryRepository) GetEntryByUserId(ctx context.Context, tx *gorm.DB, userId string) ([]models.EntryHistory, error) {

	if tx == nil {
		tx = r.db
	}

	var entryHistory []models.EntryHistory
	if err := tx.WithContext(ctx).Where("user_id = ?", userId).
		Find(&entryHistory).Error; err != nil {
		return []models.EntryHistory{}, err
	}

	return entryHistory, nil
}
