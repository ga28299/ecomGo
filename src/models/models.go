package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `db:"id" json:"id"`
	FirstName    *string   `db:"first_name" json:"firstName" form:"firstName"`
	LastName     *string   `db:"last_name" json:"lastName" form:"lastName"`
	Email        *string   `db:"email" json:"email" form:"email"`
	Password     *string   `db:"password" json:"password" form:"password"`
	Phone        *string   `db:"phone" json:"phone" form:"phone"`
	Token        *string   `db:"token" json:"token"`
	RefreshToken *string   `db:"refresh_token" json:"refreshToken"`
	CreatedAt    time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt    time.Time `db:"updated_at" json:"updatedAt"`
	UserID       string    `db:"user_id" json:"userId"`
	UserCart     []Product `json:"userCart"`
	Addresses    []Address `json:"addresses"`
	Orders       []Order   `json:"orders"`
}

type Product struct {
}

type Address struct {
}

type Order struct {
}
