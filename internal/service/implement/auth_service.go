package serviceimplement

import (
	"database/sql"
	"fmt"

	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
	"github.com/VuKhoa23/advanced-web-be/internal/repository"
	"github.com/VuKhoa23/advanced-web-be/internal/service"
	"github.com/VuKhoa23/advanced-web-be/internal/utils/constants"
	"github.com/VuKhoa23/advanced-web-be/internal/utils/env"
	"github.com/VuKhoa23/advanced-web-be/internal/utils/google_recaptcha"
	"github.com/VuKhoa23/advanced-web-be/internal/utils/jwt"
	"github.com/gin-gonic/gin"
)

type AuthService struct {
	customerRepository repository.CustomerRepository
	authenticationRepository repository.AuthenticationRepository
}

func NewAuthService(customerRepository repository.CustomerRepository, authenticationRepository repository.AuthenticationRepository) service.AuthService {
	return &AuthService{
		customerRepository: customerRepository,
		authenticationRepository: authenticationRepository,
	}
}

func (service *AuthService) Register(ctx *gin.Context, registerRequest model.RegisterRequest) error {
	err := service.customerRepository.RegisterCommand(ctx, registerRequest)
	return err
}

func (service *AuthService) Login(ctx *gin.Context, loginRequest model.LoginRequest) (*entity.Customer, error) {
	// validate captcha
	isValid, err := google_recaptcha.ValidateRecaptcha(ctx, loginRequest.RecaptchaToken)
    if err != nil || !isValid {
        return &entity.Customer{}, fmt.Errorf("invalid reCAPTCHA token")
    }

	customer, err := service.customerRepository.LoginCommand(ctx, loginRequest)

	jwtSecret, err := env.GetEnv("JWT_SECRET")
	if err != nil {
		return &entity.Customer{}, err
	}
	accessToken, err := jwt.GenerateToken(constants.ACCESS_TOKEN_DURATION, jwtSecret, map[string]interface{}{
		"id": customer.ID,
	})

	if err == nil {
		ctx.SetCookie(
			"access_token",
			accessToken,
			constants.COOKIE_DURATION,
			"/",
			"",
			false,
			true,
		)
	}

	refreshToken, err := jwt.GenerateToken(constants.REFRESH_TOKEN_DURATION, jwtSecret, map[string]interface{}{
		"id": customer.ID,
	})
	if err != nil {
		return &entity.Customer{}, err
	}

	// Check if a refresh token already exists
	existingRefreshToken, err := service.authenticationRepository.ValidateRefreshToken(ctx, customer.ID)
	if err != nil && err != sql.ErrNoRows {
		return &entity.Customer{}, err
	}

	if existingRefreshToken == nil {
		// Create a new refresh token
		err = service.authenticationRepository.CreateRefreshToken(ctx, entity.Authentication{
			CustomerID: customer.ID,
			RefreshToken: refreshToken,
		})
		if err != nil {
			return &entity.Customer{}, err
		}
	} else {
		// Update the existing refresh token
		err = service.authenticationRepository.UpdateRefreshToken(ctx, entity.Authentication{
			CustomerID: customer.ID,
			RefreshToken: refreshToken,
		})
		if err != nil {
			return &entity.Customer{}, err
		}
	}

	// Set refresh token as a cookie
	ctx.SetCookie(
		"refresh_token",
		refreshToken,
		constants.COOKIE_DURATION,
		"/",
		"",
		false,
		true,
	)

	return customer, nil
}

func (service *AuthService) ValidateRefreshToken(ctx *gin.Context, customerId int64) (*entity.Authentication, error) {
	refreshToken, err := service.authenticationRepository.ValidateRefreshToken(ctx, customerId)
	if err != nil {
		return nil, err
	}
	return refreshToken, nil
}