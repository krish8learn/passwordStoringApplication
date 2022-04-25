package dto

import (
	"database/sql"
	"errors"
	"time"
)

type Browsers struct {
	BrowserName    string       `db:"browser_name"`
	AccountCreated sql.NullBool `db:"account_created"`
	EmailIDUsed    string       `db:"email_id_used"`
	CreatedAt      time.Time    `db:"created_at"`
}

func SaveBrowser(browserName, emailIdUsed string, accountCreated bool) (*Browsers, error) {
	browser := Browsers{
		BrowserName: browserName,
		AccountCreated: sql.NullBool{
			Bool:  accountCreated,
			Valid: true,
		},
		EmailIDUsed: emailIdUsed,
	}
	result := DB.Create(&browser)

	if result.Error != nil {
		return nil, result.Error
	}
	return &browser, nil
}

func GetAllBrowsers() (*[]Browsers, error) {
	var browsers []Browsers
	// Get all records
	result := DB.Find(&browsers)
	// SELECT * FROM users;
	if result.Error != nil {
		return nil, result.Error
	}

	return &browsers, nil
}

func RemoveBrowser(browerName string) error {
	result := DB.Where("browser_name = ?", browerName).Delete(&Emails{})
	// result := DB.Delete(&Emails{}, browerName)
	if result.RowsAffected == 0 {
		return errors.New("Not Record found")
	}
	if result.Error != nil {
		return result.Error
	}

	return nil
}
