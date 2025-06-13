package dto

type UserRegistrationDto struct {
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	Age         int    `json:"age" binding:"required"`
	Gender      string `json:"gender" binding:"required"`
	PhoneNumber string `json:"no_hp" binding:"required"`
}

type UserLoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserGetByEmailDto struct {
	Email string `json:"email"`
}

type UserResponseDto struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}
