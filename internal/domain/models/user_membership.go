package models

import (
	"time"

	"github.com/google/uuid"
)

type UserMembership struct {
	Id        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserId    uuid.UUID `json:"user_id" gorm:"uniqueIndex"`
	MemberId  uuid.UUID `json:"membership_id"`
	Verified  bool      `json:"verified" gorm:"type:boolean;default:true"`
	ExpiredAt time.Time `json:"expired"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (UserMembership) Tablename() string {
	return "User Membership"
}
