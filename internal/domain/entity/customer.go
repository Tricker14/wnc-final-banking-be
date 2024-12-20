package entity

import "time"

type Customer struct {
	ID          int64      `db:"id" json:"id"`
	Email       string     `db:"email" json:"email"`
	Name        string     `db:"name" json:"name"`
	PhoneNumber string     `db:"phone_number" json:"phoneNumber"`
	Password    string     `db:"password" json:"password"`
	CreatedAt   *time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updatedAt"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deletedAt"`
}
