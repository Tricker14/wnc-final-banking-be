package entity

import "time"

type Account struct {
	ID         int64      `db:"id" json:"id"`
	CustomerID int64      `db:"customer_id" json:"customerId"`
	Number     string     `db:"number" json:"number"`
	Balance    int64      `db:"balance" json:"balance"`
	CreatedAt  *time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt  *time.Time `db:"updated_at" json:"updatedAt"`
	DeleteddAt *time.Time `db:"deleted_at" json:"deleteddAt"`
}
