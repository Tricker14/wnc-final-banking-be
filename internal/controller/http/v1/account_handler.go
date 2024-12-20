package v1

import (
	httpcommon "github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/domain/http_common"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/domain/model"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/service"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/utils/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountHandler struct {
	accountService service.AccountService
}

func NewAccountHandler(accountService service.AccountService) *AccountHandler {
	return &AccountHandler{accountService: accountService}
}

// @Summary Transfer
// @Description Transfer from internal account to internal account
// @Tags Accounts
// @Accept json
// @Param request body model.InternalTransferRequest true "Account payload"
// @Produce  json
// @Router /account/internal-transfer [post]
// @Success 204 "No Content"
// @Failure 400 {object} httpcommon.HttpResponse[any]
// @Failure 500 {object} httpcommon.HttpResponse[any]
func (handler *AccountHandler) InternalTransfer(ctx *gin.Context) {
	var transfer model.InternalTransferRequest

	if err := validation.BindJsonAndValidate(ctx, &transfer); err != nil {
		return
	}
	err := handler.accountService.InternalTransfer(ctx, transfer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	ctx.AbortWithStatus(204)
}

// @Summary Get Customer Name by Account Number
// @Description Get Customer Name by Account Number
// @Tags Accounts
// @Param accountNumber query string true "Account payload"
// @Produce  json
// @Router /account/customer-name [get]
// @Success 200 {object} httpcommon.HttpResponse[model.GetCustomerNameByAccountNumberResponse]
// @Failure 500 {object} httpcommon.HttpResponse[any]
func (handler *AccountHandler) GetCustomerNameByAccountNumber(ctx *gin.Context) {
	accountNumber := ctx.Query("accountNumber")
	customer, err := handler.accountService.GetCustomerByAccountNumber(ctx, accountNumber)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	ctx.JSON(http.StatusOK, httpcommon.NewSuccessResponse(&model.GetCustomerNameByAccountNumberResponse{
		Name: customer.Name,
	}))
}
