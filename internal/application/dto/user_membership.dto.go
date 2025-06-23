package dto

import "time"

type CreateUserMembershipRequestDto struct {
	UserId       string `json:"user_id" binding:"required"`
	MembershipId string `json:"membership_id" binding:"required"`
}

type UserMembershipResponseDto struct {
	Id        string    `json:"id"`
	ExpiredAt time.Time `json:"expired"`
}
