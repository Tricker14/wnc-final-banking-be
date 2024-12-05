package repositoryimplement

import (
	"context"
	"database/sql"
	"errors"

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

func (repo *CustomerRepository) CreateCustomer(ctx context.Context, customer *entity.Customer) error {
	query := `INSERT INTO customers(username, email, phone, password) VALUES (:username, :email, :phone, :password)`

	_, err := repo.db.NamedExecContext(ctx, query, customer)
	if err != nil {
		return err
	}

	return nil
}

func (repo *CustomerRepository) FindCustomerByUsername(ctx context.Context, username string) (*entity.Customer, error) {
	var customer entity.Customer
	query := `SELECT * FROM customers WHERE username = ? LIMIT 1`

	err := repo.db.GetContext(ctx, &customer, query, username)
	if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            // No customer found
            return nil, nil
        }
        return nil, err
    }

    return &customer, nil
}

func (repo CustomerRepository) RegisterCommand(ctx context.Context, registerRequest model.RegisterRequest) error {
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

func (repo CustomerRepository) LoginCommand(c context.Context, loginRequest model.LoginRequest) (entity.Customer, error) {
	var customer entity.Customer
	query := "SELECT id, username, password FROM customers WHERE username = ?"
	err := repo.db.GetContext(c, &customer, query, loginRequest.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Customer{}, errors.New(httpcommon.ErrorMessage.BadCredential)
		}
		return entity.Customer{}, err
	}

	if customer.Password != loginRequest.Password {
		return entity.Customer{}, errors.New(httpcommon.ErrorMessage.BadCredential)
	}

	return customer, nil
}

func (repo CustomerRepository) UpdateRefreshToken(ctx context.Context, customerId int64, token string) error {
	query := `UPDATE customers SET refresh_token = :refresh_token WHERE id = :id`

	_, err := repo.db.NamedExecContext(ctx, query, map[string]interface{}{
		"refresh_token": token,
		"id":            customerId,
	})
	if err != nil {
		return err
	}

	return nil
}

func (repo CustomerRepository) ValidateRefreshToken(ctx context.Context, customerId int64, token string) bool {
	query := `SELECT refresh_token FROM customers WHERE id = ?`

	var storedToken string
	err := repo.db.GetContext(ctx, &storedToken, query, customerId)
	if err != nil {
		return false
	}

	return storedToken == token
}
