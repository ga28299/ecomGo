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
	Product_ID    uuid.UUID `json:"productId"`
	Product_Name  *string   `json:"productName"`
	Product_Price *int      `json:"price"`
}

type Address struct {
	Address_ID uuid.UUID `db:"address_id" json:"addressId"`
	House      *string   `db:"house" json:"house"`
	Street     *string   `db:"street" json:"street"`
	City       *string   `db:"city" json:"city"`
	State      *string   `db:"state" json:"state"`
	Country    *string   `db:"country" json:"country"`
	Pincode    *string   `db:"pincode" json:"pincode"`
}

type Order struct {
	Order_ID       uuid.UUID `db:"order_id" json:"orderId"`
	Order_Cart     []Product `json:"orderCart"`
	Order_Address  Address   `json:"orderAddress"`
	Price          int       `json:"price"`
	Discout        *int      `json:"discount"`
	Payment_Method Payment   `json:"paymentMethod"`
}

type Payment struct {
	Digital bool `json:"digital"`
	Cash    bool `json:"cash"`
}
