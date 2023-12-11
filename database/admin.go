package database

import (
	"database/sql"
	"log"
)

func AddService(db *sql.DB, key_service string) error {
	if err := db.Ping(); err != nil {
		log.Print("AddService error pinging db: ", err)
		return err
	}

	query := `
		INSERT INTO services VALUES ($1)	
	`

	_, err := db.Exec(query, key_service)

	if err != nil {
		log.Print("AddService error executing query: ", err)
		return err
	}

	return nil
}
