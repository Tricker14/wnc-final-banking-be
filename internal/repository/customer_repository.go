package repository

import (
	"context"

	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
)

type CustomerRepository interface {
	RegisterCommand(ctx context.Context, registerRequest model.RegisterRequest) error
	LoginCommand(ctx context.Context, loginRequest model.LoginRequest) (*entity.Customer, error)

	CreateRefreshToken(ctx context.Context, refreshToken entity.RefreshToken) error
	UpdateRefreshToken(ctx context.Context, refreshToken entity.RefreshToken) error
	ValidateRefreshToken(ctx context.Context, customerId int64) (*entity.RefreshToken, error)
}