package service

import "github.com/gin-gonic/gin"

type CoreService interface {
	EstimateTransferFee(ctx *gin.Context, amount int64) (int64, error)
}
