package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Gender string

const (
	GENDER_MALE   Gender = "laki-laki"
	GENDER_FEMALE Gender = "perempuan"
)

type User struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Age         int       `json:"age"`
	Gender      Gender    `json:"gender"`
	PhoneNumber string    `json:"no_hp"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  gorm.DeletedAt
}

func (User) TableName() string {
	return "User"
}
