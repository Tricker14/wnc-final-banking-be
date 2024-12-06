package entity

import "time"

type RefreshToken struct {
	ID        int64      `db:"id" json:"id"`
	CustomerID  int64     `db:"customer_id" json:"customerId"`
	Value string `db:"value" json:"value"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}