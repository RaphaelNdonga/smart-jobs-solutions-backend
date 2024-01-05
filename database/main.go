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

func AddUser(db *sql.DB, userDetails types.UserDetailsDB) (string, error) {
	pingErr := db.Ping()
	if pingErr != nil {
		return "", pingErr
	}
	query := `
		INSERT INTO userDetails (username, email, hashedpassword, location, usertype) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5
		)
		RETURNING id;
	`
	var lastInsertId string

	queryErr := db.QueryRow(query, userDetails.Username, userDetails.Email, userDetails.HashedPassword, userDetails.Location, userDetails.UserType).Scan(&lastInsertId)
	if queryErr != nil {
		return "", queryErr
	}
	return lastInsertId, nil
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
			&userDetails.Id,
			&userDetails.Username,
			&userDetails.Email,
			&userDetails.HashedPassword,
			&userDetails.UserType,
			&userDetails.Location,
		)
		if err != nil {
			return userDetails, err
		}
	}
	return userDetails, err
}

func GetUserById(db *sql.DB, userId string) (types.UserDetailsDB, error) {
	if err := db.Ping(); err != nil {
		return types.UserDetailsDB{}, err
	}
	query := `
		SELECT * FROM userdetails WHERE id = $1	
	`
	rows, err := db.Query(query, userId)
	if err != nil {
		return types.UserDetailsDB{}, err
	}
	var userdetails types.UserDetailsDB

	for rows.Next() {
		err := rows.Scan(&userdetails.Id, &userdetails.Username, &userdetails.Email, &userdetails.HashedPassword, &userdetails.UserType, &userdetails.Location)
		if err != nil {
			return types.UserDetailsDB{}, err
		}
	}
	return userdetails, nil
}

func GetServices(db *sql.DB) ([]types.Service, error) {
	var services []types.Service
	if err := db.Ping(); err != nil {
		return services, err
	}
	query := `
		SELECT * FROM services	
	`
	rows, err := db.Query(query)
	if err != nil {
		return services, err
	}

	for rows.Next() {
		var service types.Service
		err := rows.Scan(
			&service.Key_Service,
		)
		services = append(services, service)
		if err != nil {
			return []types.Service{}, err
		}
	}
	return services, nil
}
