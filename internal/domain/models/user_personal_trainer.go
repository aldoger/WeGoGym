package models

import "github.com/google/uuid"

type UserPersonalTrainer struct {
	Id     uuid.UUID `json:"id" gorm:""`
	Sesi   int       `json:"sesi" gorm:""`
	UserId uuid.UUID `json:"id_user" gorm:""`
}

func (UserPersonalTrainer) TableName() string {
	return "UserPersonalTrainer"
}
