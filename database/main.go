package database

import (
	"database/sql"
	"log"
	"smartjobsolutions/types"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() *sql.DB {
	connStr := "user=raphaelndonga dbname=smartjobsolutions sslmode=disable"
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	db = conn
	return db
}

func GetDB() *sql.DB {
	return db
}

func AddUser(db *sql.DB, userDetails types.UserDetailsDB) error {
	pingErr := db.Ping()
	if pingErr != nil {
		return pingErr
	}
	query := `
		INSERT INTO userDetails (username, email, hashedpassword, location, usertype) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5
		)
	`

	_, queryErr := db.Exec(query, userDetails.Username, userDetails.Email, userDetails.HashedPassword, userDetails.Location, userDetails.UserType)

	if queryErr != nil {
		return queryErr
	}
	return nil
}

func GetUserByEmail(db *sql.DB, email string) (types.UserDetailsDB, error) {
	userDetails := types.UserDetailsDB{}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	query := `
		SELECT * FROM userdetails WHERE email = $1	
	`

	rows, err := db.Query(query, email)
	if err != nil {
		return userDetails, err
	}
	for rows.Next() {
		err := rows.Scan(
			&userDetails.Username,
			&userDetails.Email,
			&userDetails.HashedPassword,
			&userDetails.UserType,
		)
		if err != nil {
			return userDetails, err
		}
	}
	return userDetails, err
}
