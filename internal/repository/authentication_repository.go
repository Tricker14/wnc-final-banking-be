package repository

import (
	"context"

	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
)

type AuthenticationRepository interface {
	CreateRefreshToken(ctx context.Context, refreshToken entity.Authentication) error
	UpdateRefreshToken(ctx context.Context, refreshToken entity.Authentication) error
	ValidateRefreshToken(ctx context.Context, customerId int64) (*entity.Authentication, error)
}