package database

import (
	"database/sql"
)

func AddService(db *sql.DB, key_service string) error {
	if err := db.Ping(); err != nil {
		return err
	}

	query := `
		INSERT INTO services VALUES ($1)	
	`

	_, err := db.Exec(query, key_service)

	if err != nil {
		return err
	}

	return nil
}
