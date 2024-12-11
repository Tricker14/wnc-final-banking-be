package repositoryimplement

import (
	"context"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/database"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/domain/entity"
	httpcommon "github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/domain/http_common"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/repository"
	"github.com/jmoiron/sqlx"
)

type AccountRepository struct {
	db *sqlx.DB
}

func NewAccountRepository(db database.Db) repository.AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (repo *AccountRepository) UpdateCommand(ctx context.Context, account entity.Account) error {
	query := `
	UPDATE accounts
	SET balance = :balance
	WHERE number = :number
	`
	_, err := repo.db.NamedExecContext(ctx, query, account)
	return err
}

func (repo *AccountRepository) GetOneByNumberQuery(ctx context.Context, number string) (*entity.Account, error) {
	var account entity.Account
	query := "SELECT * FROM accounts WHERE number = ?"
	err := repo.db.QueryRowxContext(ctx, query, number).StructScan(&account)
	if err != nil {
		if err.Error() == httpcommon.ErrorMessage.SqlxNoRow {
			return nil, nil
		}
		return nil, err
	}
	return &account, nil
}

func (repo *AccountRepository) GetOneByCustomerIdQuery(ctx context.Context, customerId int64) (*entity.Account, error) {
	var account entity.Account
	query := "SELECT * FROM accounts WHERE customer_id = ?"
	err := repo.db.QueryRowxContext(ctx, query, customerId).StructScan(&account)
	if err != nil {
		if err.Error() == httpcommon.ErrorMessage.SqlxNoRow {
			return nil, nil
		}
		return nil, err
	}
	return &account, nil
}
