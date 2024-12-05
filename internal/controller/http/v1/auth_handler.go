package v1

import (
	"fmt"
	"net/http"

	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	httpcommon "github.com/VuKhoa23/advanced-web-be/internal/domain/http_common"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
	"github.com/VuKhoa23/advanced-web-be/internal/service"
	"github.com/VuKhoa23/advanced-web-be/internal/utils/validation"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	customerService service.CustomerService
}

func NewAuthHandler(customerService service.CustomerService) *AuthHandler {
	return &AuthHandler{customerService: customerService}
}

func (handler *AuthHandler) Register(ctx *gin.Context) {
	var registerRequest model.RegisterRequest

	if err := validation.BindJsonAndValidate(ctx, &registerRequest); err != nil {
		return
	}

	err := handler.customerService.Register(ctx, registerRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	ctx.AbortWithStatus(204)
}

func (handler *AuthHandler) Login(ctx *gin.Context){
	var loginRequest model.LoginRequest

	if err := validation.BindJsonAndValidate(ctx, &loginRequest); err != nil {
		return
	}

	customer, err := handler.customerService.Login(ctx, loginRequest)
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
		Username: customer.Username,
	}))
}

func (handler *AuthHandler) TestJWT(c *gin.Context) {
	fmt.Println("test login")
}
