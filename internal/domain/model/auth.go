package model

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email,min=10,max=255"`
	Phone    string `json:"phone" binding:"required,min=10,max=255"`
	Password string `json:"password" binding:"required,min=8,max=255"`
}

type LoginRequest struct {
	Email          string `json:"email" binding:"required,email,min=10,max=255"`
	Password       string `json:"password" binding:"required,min=8,max=255"`
	RecaptchaToken string `json:"recaptcha_token" binding:"required"`
}

type AuthenticationRequest struct {
	CustomerID   int64  `json:"customerId" binding:"required,min=1,max=255"`
	RefreshToken string `json:"refreshToken" binding:"required"`
}