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
	DB.Create(&email)

	if DB.Error != nil {
		return nil, DB.Error
	}
	return &email, nil
}
