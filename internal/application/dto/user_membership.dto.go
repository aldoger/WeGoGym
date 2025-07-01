package dto

import "time"

type CreateUserMembershipRequestDto struct {
	MembershipId string `json:"membership_id" binding:"required"`
}

type UpdateUserMembershipRequestDto struct {
	UserId string `json:"user_id" binding:"required"`
}

type UserMembershipResponseDto struct {
	Id        string    `json:"id"`
	ExpiredAt time.Time `json:"expired"`
}

type UserSearchMembershipResponse struct {
	Id             string                    `json:"id"`
	Username       string                    `json:"username"`
	Email          string                    `json:"email"`
	Role           string                    `json:"role,omitempty"`
	UserMembership UserMembershipResponseDto `json:"user_membership,omitempty"`
	Sesi           int                       `json:"sesi"`
}
