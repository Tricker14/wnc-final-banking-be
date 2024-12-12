package repositoryimplement

import (
	"context"

	httpcommon "github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/domain/http_common"

	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/database"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/domain/entity"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/repository"
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
		if err.Error() == httpcommon.ErrorMessage.SqlxNoRow {
			return nil, nil
		}
		return nil, err
	}
	return &customer, nil
}

func (repo *CustomerRepository) GetOneByIdQuery(ctx context.Context, id int64) (*entity.Customer, error) {
	var customer entity.Customer
	query := "SELECT * FROM customers WHERE id = ?"
	err := repo.db.QueryRowxContext(ctx, query, id).StructScan(&customer)
	if err != nil {
		if err.Error() == httpcommon.ErrorMessage.SqlxNoRow {
			return nil, nil
		}
		return nil, err
	}
	return &customer, nil
}

func (repo *CustomerRepository) GetIdByEmailQuery(ctx context.Context, email string) (int64, error) {
	var customer entity.Customer
	query := "SELECT * FROM customers WHERE email = ?"
	err := repo.db.QueryRowxContext(ctx, query, email).StructScan(&customer)
	if err != nil {
		if err.Error() == httpcommon.ErrorMessage.SqlxNoRow {
			return 0, nil
		}
		return 0, err
	}
	return customer.ID, nil
}

func (repo *CustomerRepository) UpdatePasswordByIdQuery(ctx context.Context, id int64, password string) error {
	query := "UPDATE customers SET password = ? WHERE id = ?"
	_, err := repo.db.ExecContext(ctx, query, password, id)
	if err != nil {
		return err
	}

	return nil
}