package serviceimplement

import (
	"database/sql"
	"errors"

	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/bean"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/domain/entity"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/domain/model"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/repository"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/service"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/utils/constants"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/utils/env"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/utils/google_recaptcha"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/utils/jwt"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/utils/mail"
	stringutils "github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/utils/string_utils"
	"github.com/gin-gonic/gin"
)

type AuthService struct {
	customerRepository       repository.CustomerRepository
	authenticationRepository repository.AuthenticationRepository
	passwordEncoder          bean.PasswordEncoder
	redisCLient 			 bean.RedisCLient
}

func NewAuthService(customerRepository repository.CustomerRepository, 
	authenticationRepository repository.AuthenticationRepository, 
	encoder bean.PasswordEncoder,
	redisCLient bean.RedisCLient,
	) service.AuthService {
	return &AuthService{
		customerRepository:       customerRepository,
		authenticationRepository: authenticationRepository,
		passwordEncoder:          encoder,
		redisCLient:			  redisCLient,
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
	hashPW, err := service.passwordEncoder.Encrypt(registerRequest.Password)
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
		return nil, err
	}

	existsCustomer, err := service.customerRepository.GetOneByEmailQuery(ctx, loginRequest.Email)
	if err != nil {
		return nil, err
	}
	if existsCustomer == nil {
		return nil, errors.New("Email not found")
	}
	checkPw := service.passwordEncoder.Compare(existsCustomer.Password, loginRequest.Password)
	if checkPw == false {
		return nil, errors.New("invalid password")
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

func (service *AuthService) SendOTPToMail(ctx *gin.Context) error {
	// generate otp
	otp := mail.GenerateOTP(6)

	// store otp in redis for 1 minute
	customerId, _ := ctx.Get("customerId")
	customerIdInt64, _ := customerId.(int64)
	
	baseKey := constants.REDIS_KEY
	key := stringutils.Concat(baseKey, customerIdInt64)

	err := service.redisCLient.Set(ctx, key, otp, constants.REDIS_EXP_TIME)
	if err != nil {
		return err
	}

	// send otp to user email
	customerMail, err := service.customerRepository.GetMailByIdQuery(ctx, customerIdInt64)
	err = mail.SendEmail(customerMail, "test otp", otp)
	if err != nil {
		return err
	}

	return nil
}

func (service *AuthService) ResetPassword(ctx *gin.Context, resetPasswordRequest model.ResetPasswordRequest) error {
	customerId, _ := ctx.Get("customerId")
	customerIdInt64, _ := customerId.(int64)
	
	baseKey := constants.REDIS_KEY
	key := stringutils.Concat(baseKey, customerIdInt64)
	
	val, err := service.redisCLient.Get(ctx, key)
	if err != nil {
		return err
	}

	service.redisCLient.Delete(ctx, key)

	hashedPW, err := service.passwordEncoder.Encrypt(resetPasswordRequest.Password)
	if err != nil {
		return err
	}

	if(val == resetPasswordRequest.OTP){
		err = service.customerRepository.UpdatePasswordByIdQuery(ctx, customerIdInt64, hashedPW)
			if err != nil {
				return err
			}
	} else {
		return errors.New("Invalid OTP")
	}

	return nil
}
