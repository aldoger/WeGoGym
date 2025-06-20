package models

import (
	"github.com/google/uuid"
)

type Membership struct {
	Id       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Type     string    `json:"type" gorm:"type:varchar(15)"`
	Price    float64   `json:"harga" gorm:"type:decimal(18,2)"`
	Duration int       `json:"durasi" gorm:"type:int"` // durasi masa berlaku
}

func (Membership) TableName() string {
	return "membeship"
}
