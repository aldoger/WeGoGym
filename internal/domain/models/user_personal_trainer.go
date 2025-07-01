package models

import "github.com/google/uuid"

type UserPersonalTrainer struct {
	Id     uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Sesi   int       `json:"sesi" gorm:"type:integer"`
	UserId uuid.UUID `json:"user_id" gorm:"uniqueIndex"`
}

func (UserPersonalTrainer) TableName() string {
	return "UserPersonalTrainer"
}
