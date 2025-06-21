package repository

import (
	"context"

	"go-kpl/internal/domain/models"

	"gorm.io/gorm"
)

//TODO Create membership

//TODO read membership (ambil semua)

//TODO read membership by Id

//TODO update membership by id dengan harga atau durasi yang di update

type (
	MembershipRepository interface {
		Create(ctx context.Context, tx *gorm.DB, membership models.Membership) (models.Membership, error)
		GetAll(ctx context.Context, tx *gorm.DB) ([]models.Membership, error)
		GetById(ctx context.Context, tx *gorm.DB, id string) (models.Membership, error)
		UpdateById(ctx context.Context, tx *gorm.DB, id string, duration *int, price *float64) (models.Membership, error)
	}
	membershipRepository struct {
		db *gorm.DB
	}	
)

func NewMembershipRepository(db *gorm.DB) MembershipRepository {
	return &membershipRepository{db: db}
}

func (r *membershipRepository) Create(ctx context.Context, tx *gorm.DB, membership models.Membership) (models.Membership, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&membership).Error; err != nil {
		return membership, err
	}

	return membership, nil
}

func (r *membershipRepository) GetAll(ctx context.Context, tx *gorm.DB) ([]models.Membership, error) {
	if tx == nil {
		tx = r.db
	}

	var memberships []models.Membership
	if err := tx.WithContext(ctx).Find(&memberships).Error; err != nil {
		return nil, err
	}

	return memberships, nil
}

func (r *membershipRepository) GetById(ctx context.Context, tx *gorm.DB, id string) (models.Membership, error) {
	if tx == nil {
		tx = r.db
	}

	var membership models.Membership
	if err := tx.WithContext(ctx).Take(&membership, "id = ?", id).Error; err != nil {
		return models.Membership{}, err
	}

	return membership, nil
}

func (r *membershipRepository) UpdateById(ctx context.Context, tx *gorm.DB, id string, duration *int, price *float64) (models.Membership, error) {
	if tx == nil {
		tx = r.db
	}

	var membership models.Membership
	if err := tx.WithContext(ctx).Take(&membership, "id = ?", id).Error; err != nil {
		return  models.Membership{}, err
	}

	if duration != nil {
		membership.Duration = *duration
		
	}
	if price != nil {
		membership.Price = *price
	}

	if err := tx.WithContext(ctx).Save(&membership).Error; err != nil {
		return models.Membership{}, err
	}
	
	return membership, nil
}