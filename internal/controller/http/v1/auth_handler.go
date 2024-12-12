package v1

import (
	"net/http"

	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/domain/entity"
	httpcommon "github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/domain/http_common"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/domain/model"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/service"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/utils/validation"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// @Summary Register
// @Description Register to account
// @Tags Auths
// @Accept json
// @Param request body model.RegisterRequest true "Auth payload"
// @Produce  json
// @Router /auth/register [post]
// @Success 204 "No Content"
// @Failure 400 {object} httpcommon.HttpResponse[any]
// @Failure 500 {object} httpcommon.HttpResponse[any]
func (handler *AuthHandler) Register(ctx *gin.Context) {
	var registerRequest model.RegisterRequest

	if err := validation.BindJsonAndValidate(ctx, &registerRequest); err != nil {
		return
	}

	err := handler.authService.Register(ctx, registerRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	ctx.AbortWithStatus(204)
}

// @Summary Login
// @Description Login to account
// @Tags Auths
// @Accept json
// @Param request body model.LoginRequest true "Auth payload"
// @Produce  json
// @Router /auth/login [post]
// @Success 200 {object} httpcommon.HttpResponse[entity.Customer]
// @Failure 400 {object} httpcommon.HttpResponse[any]
// @Failure 500 {object} httpcommon.HttpResponse[any]
func (handler *AuthHandler) Login(ctx *gin.Context) {
	var loginRequest model.LoginRequest

	if err := validation.BindJsonAndValidate(ctx, &loginRequest); err != nil {
		return
	}

	customer, err := handler.authService.Login(ctx, loginRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(
			httpcommon.Error{
				Message: err.Error(),
				Code:    httpcommon.ErrorResponseCode.InvalidRequest,
			},
		))
		return
	}

	ctx.JSON(200, httpcommon.NewSuccessResponse[entity.Customer](&entity.Customer{
		Email: customer.Email,
	}))
}

// @Summary Send OTP to Mail
// @Description Send OTP to user email
// @Tags Auths
// @Accept json
// @Param request body model.SendOTPRequest true "Send OTP payload"
// @Produce json
// @Router /auth/reset-password/otp [post]
// @Success 204 "No Content"
// @Failure 400 {object} httpcommon.HttpResponse[any]
// @Failure 500 {object} httpcommon.HttpResponse[any]
func (handler *AuthHandler) SendOTPToMail(ctx *gin.Context) {
	var sendOTPRequest model.SendOTPRequest

	if err := validation.BindJsonAndValidate(ctx, &sendOTPRequest); err != nil {
		return
	}

	err := handler.authService.SendOTPToEmail(ctx, sendOTPRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(
			httpcommon.Error{
				Message: err.Error(),
				Code:    httpcommon.ErrorResponseCode.InvalidRequest,
			},
		))
		return
	}
}

// @Summary Reset Password
// @Description Resets user password base on their email
// @Tags Auths
// @Accept json
// @Param request body model.ResetPasswordRequest true "Reset Password payload"
// @Produce json
// @Router /auth/reset-password [post]
// @Success 204 "No Content"
// @Failure 400 {object} httpcommon.HttpResponse[any]
// @Failure 500 {object} httpcommon.HttpResponse[any]
func (handler *AuthHandler) ResetPassword(ctx *gin.Context) {
	var resetPasswordRequest model.ResetPasswordRequest

	if err := validation.BindJsonAndValidate(ctx, &resetPasswordRequest); err != nil {
		return
	}

	err := handler.authService.ResetPassword(ctx, resetPasswordRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(
			httpcommon.Error{
				Message: err.Error(),
				Code:    httpcommon.ErrorResponseCode.InvalidRequest,
			},
		))
		return
	}
}