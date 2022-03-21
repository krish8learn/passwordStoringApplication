package dto

import (
	"database/sql"
	"time"
)

type Emails struct {
	EmailID    string         `db:"email_id"`
	DomainName string         `db:"domain_name"`
	Password   string         `db:"password"`
	Reason     sql.NullString `db:"reason"`
	CreatedAt  time.Time      `db:"created_at"`
}

func CreateEmail(emailId, domainName, password, reason string) (*Emails, error) {
	email := Emails{
		EmailID:    emailId,
		DomainName: domainName,
		Password:   password,
		Reason:     sql.NullString{String: reason, Valid: true},
	}
	result := DB.Create(&email)

	if result.Error != nil {
		return nil, result.Error
	}
	return &email, nil
}

func GetEmail(emailId string) (*Emails, error) {
	var email Emails
	result := DB.Where("email_id = ?", emailId).First(&email)

	if result.Error != nil {
		return nil, result.Error
	}

	return &email, nil
}
