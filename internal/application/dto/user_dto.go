package dto

type UserRegistrationDto struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
}

type UserLoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponseDto struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}
