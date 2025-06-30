package models

import "github.com/google/uuid"

type UserPersonalTrainer struct {
	Id     uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Sesi   int       `json:"sesi" gorm:"type:integer"`
	UserId uuid.UUID `json:"id_user" gorm:"uniqueIndex"`
}

func (UserPersonalTrainer) TableName() string {
	return "UserPersonalTrainer"
}
