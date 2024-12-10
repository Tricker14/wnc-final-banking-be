package repository

import (
	"context"

	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
)

type AuthenticationRepository interface {
	CreateCommand(ctx context.Context, refreshToken entity.Authentication) error
	UpdateCommand(ctx context.Context, refreshToken entity.Authentication) error
	GetOneByCustomerIdQuery(ctx context.Context, customerId int64) (*entity.Authentication, error)
}
