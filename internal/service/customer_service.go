package service

import (
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
	"github.com/gin-gonic/gin"
)

type CustomerService interface {
	Register(ctx *gin.Context, customerRequest model.RegisterRequest) error
	Login(ctx *gin.Context, customerRequest model.LoginRequest) (entity.Customer, error)
}
