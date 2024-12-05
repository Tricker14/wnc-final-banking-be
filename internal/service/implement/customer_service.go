package serviceimplement

import (
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

type CustomertService struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerService(customerRepository repository.CustomerRepository) service.CustomerService {
	return &CustomertService{customerRepository: customerRepository}
}

func (service CustomertService) Register(ctx *gin.Context, registerRequest model.RegisterRequest) error {
	err := service.customerRepository.RegisterCommand(ctx, registerRequest)
	return err
}

func (service CustomertService) Login(ctx *gin.Context, loginRequest model.LoginRequest) (entity.Customer, error) {
	// validate captcha
	isValid, err := google_recaptcha.ValidateRecaptcha(ctx, loginRequest.RecaptchaToken)
    if err != nil || !isValid {
        return entity.Customer{}, fmt.Errorf("invalid reCAPTCHA token")
    }

	customer, err := service.customerRepository.LoginCommand(ctx, loginRequest)

	jwtSecret, err := env.GetEnv("JWT_SECRET")
	if err != nil {
		return entity.Customer{}, err
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
	if err == nil {
		ctx.SetCookie(
			"refresh_token",
			refreshToken,
			constants.COOKIE_DURATION,
			"/",
			"",
			false,
			true,
		)
	}
	err = service.customerRepository.UpdateRefreshToken(ctx, customer.ID, refreshToken)
	if err != nil {
		return entity.Customer{}, err
	}

	return customer, nil
}
