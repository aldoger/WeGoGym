package dto

type MembershipRequestDto struct {
	Type string `json:"type" binding:"required"`
	Duration int `json:"duration" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

type UpdateMembershipRequestDto struct {
	Id string `json:"id" binding:"required"`
	Duration int `json:"duration" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

type MembershipResponseDto struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Duration int `json:"duration"`
	Price float64 `json:"price"`
}