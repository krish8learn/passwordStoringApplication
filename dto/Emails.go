package dto

import (
	"database/sql"
	"errors"
	"fmt"
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

func GetAllEmails() (*[]Emails, error) {
	var emails []Emails
	// Get all records
	result := DB.Find(&emails)
	// SELECT * FROM users;
	if result.Error != nil {
		return nil, result.Error
	}

	return &emails, nil
}

func RemovePassword(emailId string) error {
	result := DB.Where("email_id = ?", emailId).Delete(&Emails{})
	// result := DB.Delete(&Emails{}, emailId)
	if result.RowsAffected == 0 {
		return errors.New("Not Record found")
	}
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdatePasswordReason(emailId, domainName, password, reason string) (*Emails, error) {
	email := Emails{
		EmailID:    emailId,
		DomainName: domainName,
		Password:   password,
		Reason:     sql.NullString{String: reason, Valid: true},
	}
	result := DB.Where("email_id = ?", emailId).Save(&email)

	fmt.Println(result.Error)
	if result.Error != nil {
		return nil, result.Error
	}

	return &email, nil
}
