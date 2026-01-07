package domain

import "time"

type Address struct {
	AddressId int `db:"address_id" json:"address_id"`
	ContactId int `db:"contact_id" json:"contact_id"`
	Street    string `db:"street" json:"street"`
	City      string `db:"city" json:"city"`
	State     string `db:"state" json:"state"`
	ZipCode   string `db:"zip_code" json:"zip_code"`
	Country   string `db:"country" json:"country"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
}