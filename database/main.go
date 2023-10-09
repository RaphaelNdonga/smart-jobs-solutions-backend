package database

import (
	"database/sql"
	"log"
	"smartjobsolutions/types"
)

var db *sql.DB

func InitDB() {
	connStr := "user=raphaelndonga dbname=smartjobsolutions sslmode=disable"
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	db = conn
}

func AddUser(userDetails types.UserDetailsDB) {
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	query := `
		INSERT INTO userDetails VALUES (
			$1,
			$2,
			$3,
			$4
		)
	`

	_, queryErr := db.Query(query, userDetails.UserType, userDetails.Email, userDetails.HashedPassword, userDetails.UserType)

	if queryErr != nil {
		log.Fatal(queryErr)
	}
}
