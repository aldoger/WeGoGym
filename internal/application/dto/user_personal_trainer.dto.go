package dto

type CreateUserPersonalTrainerDto struct {
	Sesi int `json:"sesi" binding:"required"`
}

type UserPersonalTrainerResponse struct {
	Sesi   int    `json:"sesi"`
	UserId string `json:"user_id" binding:"omitempty"`
}
