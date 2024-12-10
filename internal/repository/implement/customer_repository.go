package repositoryimplement

import (
	"context"
	"database/sql"
	"errors"

	"github.com/VuKhoa23/advanced-web-be/internal/database"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
	"github.com/VuKhoa23/advanced-web-be/internal/repository"
	"github.com/jmoiron/sqlx"
)

type CustomerRepository struct {
	db *sqlx.DB
}

func NewCustomerRepository(db database.Db) repository.CustomerRepository {
	return &CustomerRepository{db: db}
}

func (repo *CustomerRepository) CreateCommand(ctx context.Context, customer *entity.Customer) error {
	// Insert the new customer
	insertQuery := `INSERT INTO customers(email, phone_number, password) VALUES (:email, :phone_number, :password)`
	_, err := repo.db.NamedExecContext(ctx, insertQuery, customer)
	if err != nil {
		return err
	}
	return nil
}

func (repo *CustomerRepository) GetOneByEmailQuery(ctx context.Context, email string) (*entity.Customer, error) {
	var customer entity.Customer
	query := "SELECT * FROM customers WHERE email = ?"
	err := repo.db.QueryRowxContext(ctx, query, email).StructScan(&customer)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &customer, nil
}
