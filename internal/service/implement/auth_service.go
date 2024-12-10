package serviceimplement

import (
	"database/sql"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"

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
	customerRepository       repository.CustomerRepository
	authenticationRepository repository.AuthenticationRepository
}

func NewAuthService(customerRepository repository.CustomerRepository, authenticationRepository repository.AuthenticationRepository) service.AuthService {
	return &AuthService{
		customerRepository:       customerRepository,
		authenticationRepository: authenticationRepository,
	}
}

func (service *AuthService) Register(ctx *gin.Context, registerRequest model.RegisterRequest) error {
	existsCustomer, err := service.customerRepository.GetOneByEmailQuery(ctx, registerRequest.Email)
	if err != nil {
		return err
	}
	if existsCustomer != nil {
		return errors.New("Email have already registered")
	}
	hashPW, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newCustomer := &entity.Customer{
		Email:       registerRequest.Email,
		PhoneNumber: registerRequest.PhoneNumber,
		Password:    string(hashPW),
	}
	err = service.customerRepository.CreateCommand(ctx, newCustomer)
	if err != nil {
		return err
	}
	return nil
}

func (service *AuthService) Login(ctx *gin.Context, loginRequest model.LoginRequest) (*entity.Customer, error) {
	// validate captcha
	isValid, err := google_recaptcha.ValidateRecaptcha(ctx, loginRequest.RecaptchaToken)
	if err != nil || !isValid {
		return &entity.Customer{}, fmt.Errorf("invalid reCAPTCHA token")
	}

	existsCustomer, err := service.customerRepository.GetOneByEmailQuery(ctx, loginRequest.Email)
	if err != nil {
		return nil, err
	}
	if existsCustomer == nil {
		return nil, errors.New("Email not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(existsCustomer.Password), []byte(loginRequest.Password))
	if err != nil {
		return nil, err
	}

	jwtSecret, err := env.GetEnv("JWT_SECRET")
	if err != nil {
		return nil, err
	}
	accessToken, err := jwt.GenerateToken(constants.ACCESS_TOKEN_DURATION, jwtSecret, map[string]interface{}{
		"id": existsCustomer.ID,
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
		"id": existsCustomer.ID,
	})
	if err != nil {
		return nil, err
	}

	// Check if a refresh token already exists
	existingRefreshToken, err := service.authenticationRepository.GetOneByCustomerIdQuery(ctx, existsCustomer.ID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if existingRefreshToken == nil {
		// Create a new refresh token
		err = service.authenticationRepository.CreateCommand(ctx, entity.Authentication{
			CustomerID:   existsCustomer.ID,
			RefreshToken: refreshToken,
		})
		if err != nil {
			return &entity.Customer{}, err
		}
	} else {
		// Update the existing refresh token
		err = service.authenticationRepository.UpdateCommand(ctx, entity.Authentication{
			CustomerID:   existsCustomer.ID,
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

	return existsCustomer, nil
}

func (service *AuthService) ValidateRefreshToken(ctx *gin.Context, customerId int64) (*entity.Authentication, error) {
	refreshToken, err := service.authenticationRepository.GetOneByCustomerIdQuery(ctx, customerId)
	if err != nil {
		return nil, err
	}
	return refreshToken, nil
}
