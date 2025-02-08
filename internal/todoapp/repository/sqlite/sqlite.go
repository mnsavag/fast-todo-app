package sqlite

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func NewSqliteConn(storagePath string) (*sqlx.DB, error) {
	conn, err := sqlx.Connect("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("cant create new sqlite conn, %w", err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("cant create new sqlite conn, %w", err)
	}

	return conn, nil
}
