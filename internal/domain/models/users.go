package models

import (
	"time"

	valueobject "go-kpl/internal/domain/value_object"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Gender string

const (
	GENDER_MALE   Gender = "laki-laki"
	GENDER_FEMALE Gender = "perempuan"
)

type User struct {
	Id             uuid.UUID            `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Username       string               `json:"username"`
	Email          string               `json:"email"`
	Password       string               `json:"password"`
	Gender         Gender               `json:"gender"`
	Role           valueobject.UserRole `json:"role"`
	UserMembership *UserMembership      `json:"user_membership_id" gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserPT         *UserPersonalTrainer `json:"user_pt_id" gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  gorm.DeletedAt
}

func (user *User) IsMember() bool {
	if user.UserMembership == nil {
		return false
	}
	return user.UserMembership.Verified
}

func (User) TableName() string {
	return "User"
}
