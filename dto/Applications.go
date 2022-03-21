package dto

import (
	"database/sql"
	"time"
)

type Applications struct {
	AppID         int            `db:"app_id"`
	AppName       string         `db:"app_name"`
	Reason        sql.NullString `db:"reason"`
	EmailIDUsed   string         `db:"email_id_used"`
	BrowserStored string         `db:"browser_stored"`
	Password      string         `db:"password"`
	CreatedAt     time.Time      `db:"created_at"`
}
