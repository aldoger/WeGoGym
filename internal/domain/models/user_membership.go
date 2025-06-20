package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserMembership struct {
	Id       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserId   uuid.UUID `json:"user_id"`
	MemberId uuid.UUID `json:"membership_id"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  gorm.DeletedAt
}
