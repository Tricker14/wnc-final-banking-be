package repository

import (
	"context"

	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
)

type CustomerRepository interface {
	RegisterCommand(ctx context.Context, registerRequest model.RegisterRequest) error
	LoginCommand(ctx context.Context, loginRequest model.LoginRequest) (*entity.Customer, error)
}