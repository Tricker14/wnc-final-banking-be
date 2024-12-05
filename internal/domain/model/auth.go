package model

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=1,max=255"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required,min=10,max=255"`
	Password string `json:"password" binding:"required,min=8,max=255"`
}

type LoginRequest struct {
	Username       string `json:"username" binding:"required,min=1,max=255"`
	Password       string `json:"password" binding:"required,min=8,max=255"`
	RecaptchaToken string `json:"recaptcha_token" binding:"required"`
}