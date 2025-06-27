package models

import (
	"time"

	"github.com/google/uuid"
)

type EntryHistory struct {
	UserId    uuid.UUID `json:"user_id" gorm:"type:uuid"`
	EntryTime time.Time `json:"entry_time"`
}

func (EntryHistory) TableName() string {
	return "EntryHistory"
}
