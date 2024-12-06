package repositoryimplement

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/VuKhoa23/advanced-web-be/internal/database"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	httpcommon "github.com/VuKhoa23/advanced-web-be/internal/domain/http_common"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
	"github.com/VuKhoa23/advanced-web-be/internal/repository"
	"github.com/jmoiron/sqlx"
)

type CustomerRepository struct {
	db *sqlx.DB
}

func NewCustomerRepository(db database.Db) repository.CustomerRepository {
	return &CustomerRepository{db: db}
}

func (repo *CustomerRepository) RegisterCommand(ctx context.Context, registerRequest model.RegisterRequest) error {
	// Check if username already exists
	var existingCustomer entity.Customer
	query := "SELECT id FROM customers WHERE username = ?"
	err := repo.db.GetContext(ctx, &existingCustomer, query, registerRequest.Username)
	if err != nil && err.Error() != httpcommon.ErrorMessage.SqlxNoRow {
		return err
	}
	if err == nil {
		return errors.New("username already exists")
	}

	// Insert the new customer
	insertQuery := `INSERT INTO customers(username, email, phone, password) VALUES (:username, :email, :phone, :password)`
	_, err = repo.db.NamedExecContext(ctx, insertQuery, registerRequest)
	if err != nil {
		return err
	}

	return nil
}

func (repo *CustomerRepository) LoginCommand(c context.Context, loginRequest model.LoginRequest) (*entity.Customer, error) {
	var customer entity.Customer
	query := "SELECT id, username, password FROM customers WHERE username = ?"
	err := repo.db.GetContext(c, &customer, query, loginRequest.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &entity.Customer{}, errors.New(httpcommon.ErrorMessage.BadCredential)
		}
		return &entity.Customer{}, err
	}

	if customer.Password != loginRequest.Password {
		return &entity.Customer{}, errors.New(httpcommon.ErrorMessage.BadCredential)
	}

	return &customer, nil
}

func (repo *CustomerRepository) CreateRefreshToken(ctx context.Context, refreshToken entity.RefreshToken) error {
	query := `
		INSERT INTO refresh_tokens (customer_id, value)
		VALUES (:customer_id, :value)
	`
	_, err := repo.db.NamedExecContext(ctx, query, refreshToken)
	return err
}

func (repo *CustomerRepository) UpdateRefreshToken(ctx context.Context, refreshToken entity.RefreshToken) error {
	query := `
		UPDATE refresh_tokens
		SET value = :value
		WHERE customer_id = :customer_id
	`
	_, err := repo.db.NamedExecContext(ctx, query, refreshToken)
	return err
}

// find refresh token and validate its expiration time
func (repo *CustomerRepository) ValidateRefreshToken(ctx context.Context, customerId int64) (*entity.RefreshToken, error) {
	var refreshToken entity.RefreshToken
	query := `
		SELECT customer_id, value, created_at
		FROM refresh_tokens
		WHERE customer_id = ?
	`
	err := repo.db.GetContext(ctx, &refreshToken, query, customerId)
	if err != nil {
		return nil, err
	}

	// Check if the token has expired
	if refreshToken.CreatedAt.Before(time.Now()) {
		return nil, nil
	}
	return &refreshToken, nil
}
