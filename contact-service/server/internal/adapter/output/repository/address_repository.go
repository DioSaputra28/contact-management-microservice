package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/DioSaputra28/contact-management-microservice/contact-service/server/internal/application/domain"
)

type AddressRepository struct {
	db *sql.DB
}

func NewAddressRepository(db *sql.DB) *AddressRepository {
	return &AddressRepository{db: db}
}

func (ar *AddressRepository) CreateAddress(contactId int64, street, city, state, zipCode, country string) (*domain.Address, error) {
	query := "INSERT INTO addresses (contact_id, street, city, state, zip_code, country) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := ar.db.Exec(query, contactId, street, city, state, zipCode, country)
	if err != nil {
		return nil, err
	}

	addressID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	address := &domain.Address{
		AddressId: int(addressID),
		ContactId: int(contactId),
		Street:    street,
		City:      city,
		State:     state,
		ZipCode:   zipCode,
		Country:   country,
		CreatedAt: &now,
	}

	return address, nil
}

func (ar *AddressRepository) UpdateAddress(contactId, addressId int64, street, city, state, zipCode, country string) (*domain.Address, error) {
	// Build dynamic update query based on non-empty fields
	query := "UPDATE addresses SET "
	args := []interface{}{}
	updates := []string{}

	if street != "" {
		updates = append(updates, "street = ?")
		args = append(args, street)
	}
	if city != "" {
		updates = append(updates, "city = ?")
		args = append(args, city)
	}
	if state != "" {
		updates = append(updates, "state = ?")
		args = append(args, state)
	}
	if zipCode != "" {
		updates = append(updates, "zip_code = ?")
		args = append(args, zipCode)
	}
	if country != "" {
		updates = append(updates, "country = ?")
		args = append(args, country)
	}

	if len(updates) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query += updates[0]
	for i := 1; i < len(updates); i++ {
		query += ", " + updates[i]
	}
	query += " WHERE address_id = ? AND contact_id = ?"
	args = append(args, addressId, contactId)

	_, err := ar.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	// Fetch updated address
	return ar.GetAddressById(contactId, addressId)
}

func (ar *AddressRepository) DeleteAddress(contactId, addressId int64) (*domain.Address, error) {
	// Get address first before deleting
	address, err := ar.GetAddressById(contactId, addressId)
	if err != nil {
		return nil, err
	}

	deleteQuery := "DELETE FROM addresses WHERE address_id = ? AND contact_id = ?"
	_, err = ar.db.Exec(deleteQuery, addressId, contactId)
	if err != nil {
		return nil, err
	}

	return address, nil
}

func (ar *AddressRepository) GetAddressById(contactId, addressId int64) (*domain.Address, error) {
	address := &domain.Address{}
	query := "SELECT address_id, contact_id, street, city, state, zip_code, country, created_at FROM addresses WHERE address_id = ? AND contact_id = ?"
	err := ar.db.QueryRow(query, addressId, contactId).Scan(
		&address.AddressId,
		&address.ContactId,
		&address.Street,
		&address.City,
		&address.State,
		&address.ZipCode,
		&address.Country,
		&address.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return address, nil
}

func (ar *AddressRepository) GetAddresses(contactId int64) ([]*domain.Address, error) {
	var addresses []*domain.Address

	query := `
		SELECT 
			address_id, 
			contact_id, 
			street, 
			city, 
			state, 
			zip_code, 
			country, 
			created_at
		FROM addresses
		WHERE contact_id = ?
		ORDER BY created_at DESC
	`

	rows, err := ar.db.Query(query, contactId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		address := &domain.Address{}
		err := rows.Scan(
			&address.AddressId,
			&address.ContactId,
			&address.Street,
			&address.City,
			&address.State,
			&address.ZipCode,
			&address.Country,
			&address.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}

	return addresses, nil
}
