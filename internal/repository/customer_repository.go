package repository

import (
	"context"

	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
)

type CustomerRepository interface {
	FindCustomerByUsername(ctx context.Context, username string) (*entity.Customer, error)
	CreateCustomer(ctx context.Context, customer *entity.Customer) error

	RegisterCommand(ctx context.Context, registerRequest model.RegisterRequest) error
	LoginCommand(ctx context.Context, loginRequest model.LoginRequest) (entity.Customer, error)
	UpdateRefreshToken(ctx context.Context, customerId int64, token string) error
	ValidateRefreshToken(ctx context.Context, customerId int64, token string) bool
}