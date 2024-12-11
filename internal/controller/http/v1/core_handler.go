package v1

import (
	httpcommon "github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/domain/http_common"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CoreHandler struct {
	coreService service.CoreService
}

func NewCoreHandler(coreService service.CoreService) *CoreHandler {
	return &CoreHandler{coreService: coreService}
}

func (handler *CoreHandler) EstimateTransferFee(ctx *gin.Context) {
	amount := ctx.Query("amount")
	if amount == "" {
		ctx.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: httpcommon.ErrorMessage.InvalidRequest, Field: "amount", Code: httpcommon.ErrorResponseCode.InvalidRequest,
		}))
		return
	}
	amountInt, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: httpcommon.ErrorMessage.InvalidRequest, Field: "amount", Code: httpcommon.ErrorResponseCode.InvalidRequest,
		}))
		return
	}
	fee, err := handler.coreService.EstimateTransferFee(ctx, amountInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httpcommon.NewErrorResponse(httpcommon.Error{
			Message: err.Error(), Field: "fee", Code: httpcommon.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	ctx.JSON(200, httpcommon.NewSuccessResponse[int64](&fee))
}
