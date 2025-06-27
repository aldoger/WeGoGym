package dto

import "time"

type EntryHistoryAllResponseDto struct {
	UserName  string    `json:"username"`
	EntryTime time.Time `json:"entry_time"`
}

type EntryHistoryUserResponseDto struct {
	EntryTime time.Time `json:"entry_time"`
}
