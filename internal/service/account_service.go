package service

import (
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/domain/model"
	"github.com/gin-gonic/gin"
)

type AccountService interface {
	InternalTransfer(ctx *gin.Context, tranferReq model.InternalTransferRequest) error
}
