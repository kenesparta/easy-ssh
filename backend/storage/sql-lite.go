package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func SqlLite(dbSource string) (*sql.DB, error) {
	return sql.Open("sqlite3", dbSource)
}
