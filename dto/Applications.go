package dto

import (
	"database/sql"
	"errors"
	"time"
)

type Applications struct {
	AppID         int            `db:"app_id" gorm:"primaryKey"`
	AppName       string         `db:"app_name"`
	Reason        sql.NullString `db:"reason"`
	EmailIDUsed   string         `db:"email_id_used"`
	BrowserStored string         `db:"browser_stored"`
	Password      string         `db:"password"`
	CreatedAt     time.Time      `db:"created_at"`
}

func CreateApp(arg Applications) (*Applications, error) {
	app := Applications{
		AppName:       arg.AppName,
		Reason:        arg.Reason,
		EmailIDUsed:   arg.EmailIDUsed,
		BrowserStored: arg.BrowserStored,
		Password:      arg.Password,
	}
	// fmt.Println(app.AppID)
	result := DB.Create(&app)
	// fmt.Println(app.AppID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &app, nil
}

func GetAllApps() (*[]Applications, error) {
	var apps []Applications
	// Get all records
	result := DB.Find(&apps)
	// SELECT * FROM users;
	if result.Error != nil {
		return nil, result.Error
	}

	return &apps, nil
}

func GetApp(appName string) (*Applications, error) {
	var app Applications
	result := DB.Where("app_name = ?", appName).First(&app)

	if result.Error != nil {
		return nil, result.Error
	}

	return &app, nil
}

func RemoveApp(appName string) error {
	result := DB.Where("email_id = ?", appName).Delete(&Applications{})
	// result := DB.Delete(&Applications{}, appName)
	if result.RowsAffected == 0 {
		return errors.New("Not Record found")
	}
	if result.Error != nil {
		return result.Error
	}

	return nil
}
