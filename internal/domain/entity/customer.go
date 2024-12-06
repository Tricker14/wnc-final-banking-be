package entity

import "time"

type Customer struct {
	ID        int64      `db:"id" json:"id"`
	Username  string     `db:"username" json:"username"`
	Email     string     `db:"email" json:"email"`
	Phone     string     `db:"phone" json:"phone"`
	Password  string     `db:"password" json:"password"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}