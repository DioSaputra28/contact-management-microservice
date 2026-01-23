package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/DioSaputra28/contact-management-microservice/contact-service/server/internal/application/domain"
)

type ContactRepository struct {
	db *sql.DB
}

func NewContactRepository(db *sql.DB) *ContactRepository {
	return &ContactRepository{db: db}
}

func (cr *ContactRepository) CreateContact(userId string, firstName, lastName, email, phone string) (*domain.Contact, error) {
	query := "INSERT INTO contacts (user_id, first_name, last_name, email, phone) VALUES (?, ?, ?, ?, ?)"
	result, err := cr.db.Exec(query, userId, firstName, lastName, email, phone)
	if err != nil {
		return nil, err
	}

	contactID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	contact := &domain.Contact{
		ContactId: int(contactID),
		UserId:    userId,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
		CreatedAt: &now,
	}

	return contact, nil
}

func (cr *ContactRepository) UpdateContact(userId string, contactId int64, firstName, lastName, email, phone string) (*domain.Contact, error) {
	// Build dynamic update query based on non-empty fields
	query := "UPDATE contacts SET "
	args := []interface{}{}
	updates := []string{}

	if firstName != "" {
		updates = append(updates, "first_name = ?")
		args = append(args, firstName)
	}
	if lastName != "" {
		updates = append(updates, "last_name = ?")
		args = append(args, lastName)
	}
	if email != "" {
		updates = append(updates, "email = ?")
		args = append(args, email)
	}
	if phone != "" {
		updates = append(updates, "phone = ?")
		args = append(args, phone)
	}

	if len(updates) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query += updates[0]
	for i := 1; i < len(updates); i++ {
		query += ", " + updates[i]
	}
	query += " WHERE contact_id = ? AND user_id = ?"
	args = append(args, contactId, userId)

	_, err := cr.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	// Fetch updated contact
	return cr.GetContactById(userId, contactId)
}

func (cr *ContactRepository) DeleteContact(userId string, contactId int64) (*domain.Contact, error) {
	// Get contact first before deleting
	contact, err := cr.GetContactById(userId, contactId)
	if err != nil {
		return nil, err
	}

	deleteQuery := "DELETE FROM contacts WHERE contact_id = ? AND user_id = ?"
	_, err = cr.db.Exec(deleteQuery, contactId, userId)
	if err != nil {
		return nil, err
	}

	return contact, nil
}

func (cr *ContactRepository) GetContactById(userId string, contactId int64) (*domain.Contact, error) {
	contact := &domain.Contact{}
	query := "SELECT contact_id, user_id, first_name, last_name, email, phone, created_at FROM contacts WHERE contact_id = ? AND user_id = ?"
	err := cr.db.QueryRow(query, contactId, userId).Scan(
		&contact.ContactId,
		&contact.UserId,
		&contact.FirstName,
		&contact.LastName,
		&contact.Email,
		&contact.Phone,
		&contact.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return contact, nil
}

func (cr *ContactRepository) GetContacts(userId string, page, limit int, search string) ([]*domain.Contact, *domain.ContactPagination, error) {
	var contacts []*domain.Contact
	var totalData int64

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	query := `
		SELECT 
			contact_id, 
			user_id, 
			first_name, 
			last_name, 
			email, 
			phone, 
			created_at,
			COUNT(*) OVER() as total_count
		FROM contacts
		WHERE user_id = ?
	`

	var args []any
	args = append(args, userId)

	if search != "" {
		query += " AND (first_name LIKE ? OR last_name LIKE ? OR email LIKE ?)"
		searchPattern := "%" + search + "%"
		args = append(args, searchPattern, searchPattern, searchPattern)
	}

	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, (page-1)*limit)

	fmt.Println("Sampai 163")
	rows, err := cr.db.Query(query, args...)
	fmt.Println(err)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		contact := &domain.Contact{}
		err := rows.Scan(
			&contact.ContactId,
			&contact.UserId,
			&contact.FirstName,
			&contact.LastName,
			&contact.Email,
			&contact.Phone,
			&contact.CreatedAt,
			&totalData,
		)
		if err != nil {
			return nil, nil, err
		}
		contacts = append(contacts, contact)
	}

	if len(contacts) == 0 && totalData == 0 {
		countQuery := "SELECT COUNT(*) FROM contacts WHERE user_id = ?"
		var countArgs []any
		countArgs = append(countArgs, userId)

		if search != "" {
			countQuery += " AND (first_name LIKE ? OR last_name LIKE ? OR email LIKE ?)"
			searchPattern := "%" + search + "%"
			countArgs = append(countArgs, searchPattern, searchPattern, searchPattern)
		}

		err := cr.db.QueryRow(countQuery, countArgs...).Scan(&totalData)
		if err != nil {
			return nil, nil, err
		}
	}

	totalPage := int64(0)
	if totalData > 0 {
		totalPage = totalData / int64(limit)
		if totalData%int64(limit) > 0 {
			totalPage++
		}
	}

	return contacts, &domain.ContactPagination{
		TotalData:   totalData,
		CurrentPage: int64(page),
		PageSize:    int64(limit),
		TotalPage:   totalPage,
	}, nil
}
