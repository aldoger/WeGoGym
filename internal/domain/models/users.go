package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Gender string
type UserRole string

const (
	GENDER_MALE   Gender = "laki-laki"
	GENDER_FEMALE Gender = "perempuan"
)

const (
	MEMBER_ROLE UserRole = "member"
	ADMIN_ROLE  UserRole = "admin"
)

type User struct {
	Id       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Gender   Gender    `json:"gender"`
	Role     UserRole  `json:"role" gorm:"default:member"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  gorm.DeletedAt
}

func (User) TableName() string {
	return "User"
}
