package dto

type SignupInput struct {
	Name     string `json:"name" binding:"required" example:"John Doe"`
	Email    string `json:"email" binding:"required,email" example:"test1@example.com"`
	Password string `json:"password" binding:"required,min=8" example:"password123"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email" example:"test1@example.com"`
	Password string `json:"password" binding:"required,min=8" example:"password123"`
}
