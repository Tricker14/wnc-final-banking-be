package repositoryimplement

import (
	"context"
	"time"

	"github.com/VuKhoa23/advanced-web-be/internal/database"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/VuKhoa23/advanced-web-be/internal/repository"
	"github.com/jmoiron/sqlx"
)

type AuthenticationRepository struct {
	db *sqlx.DB
}

func NewAuthenticationRepository(db database.Db) repository.AuthenticationRepository {
	return &AuthenticationRepository{db: db}
}

func (repo *AuthenticationRepository) CreateRefreshToken(ctx context.Context, authentication entity.Authentication) error {
	query := `
		INSERT INTO authentications (customer_id, refresh_token)
		VALUES (:customer_id, :refresh_token)
	`
	_, err := repo.db.NamedExecContext(ctx, query, authentication)
	return err
}

func (repo *AuthenticationRepository) UpdateRefreshToken(ctx context.Context, authentication entity.Authentication) error {
	query := `
		UPDATE authentications
		SET refresh_token = :refresh_token
		WHERE customer_id = :customer_id
	`
	_, err := repo.db.NamedExecContext(ctx, query, authentication)
	return err
}

// find refresh token and validate its expiration time
func (repo *AuthenticationRepository) ValidateRefreshToken(ctx context.Context, customerId int64) (*entity.Authentication, error) {
	var authentication entity.Authentication
	query := `
		SELECT customer_id, refresh_token, created_at
		FROM authentications
		WHERE customer_id = ?
	`
	err := repo.db.GetContext(ctx, &authentication, query, customerId)
	if err != nil {
		return nil, err
	}

	// Check if the token has expired
	if authentication.CreatedAt.Before(time.Now()) {
		return nil, nil
	}
	return &authentication, nil
}