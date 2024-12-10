package repository

import (
	"context"

	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
)

type CustomerRepository interface {
	CreateCommand(ctx context.Context, customer *entity.Customer) error
	GetOneByEmailQuery(ctx context.Context, email string) (*entity.Customer, error)
}
