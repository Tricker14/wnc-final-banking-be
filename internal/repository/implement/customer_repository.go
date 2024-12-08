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

func (repo *CustomerRepository) RegisterCommand(ctx context.Context, registerRequest model.RegisterRequest) error {
	// Check if email already exists
	var existingCustomer entity.Customer
	query := "SELECT id FROM customers WHERE email = ?"
	err := repo.db.GetContext(ctx, &existingCustomer, query, registerRequest.Email)
	if err != nil && err.Error() != httpcommon.ErrorMessage.SqlxNoRow {
		return err
	}
	if err == nil {
		return errors.New("email already exists")
	}

	// Insert the new customer
	insertQuery := `INSERT INTO customers(email, phone, password) VALUES (:email, :phone, :password)`
	_, err = repo.db.NamedExecContext(ctx, insertQuery, registerRequest)
	if err != nil {
		return err
	}

	return nil
}

func (repo *CustomerRepository) LoginCommand(c context.Context, loginRequest model.LoginRequest) (*entity.Customer, error) {
	var customer entity.Customer
	query := "SELECT id, email, password FROM customers WHERE email = ?"
	err := repo.db.GetContext(c, &customer, query, loginRequest.Email)
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
