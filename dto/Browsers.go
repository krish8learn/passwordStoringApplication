package dto

import (
	"database/sql"
	"time"
)

type Browsers struct {
	BrowserName    string       `db:"browser_name"`
	AccountCreated sql.NullBool `db:"account_created"`
	EmailIDUsed    string       `db:"email_id_used"`
	CreatedAt      time.Time    `db:"created_at"`
}
