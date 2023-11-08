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

func AddServiceProvider(db *sql.DB, serviceProvider types.ServiceProvider) error {
	pingErr := db.Ping()
	if pingErr != nil {
		return pingErr
	}
	query := `
		INSERT INTO serviceProvider VALUES (
			$1,
			$2,
			$3
		)	
	`
	_, err := db.Exec(query, serviceProvider.Id, serviceProvider.Service, serviceProvider.Description)
	if err != nil {
		return err
	}
	return nil
}

func AddClient(db *sql.DB, client types.Client) error {
	pingErr := db.Ping()
	if pingErr != nil {
		return pingErr
	}
	query := `
		INSERT INTO client VALUES (
			$1,
			$2
		)	
	`
	_, err := db.Exec(query, client.Id, client.Service)
	if err != nil {
		return err
	}
	return nil
}

func GetServiceProviders(db *sql.DB) ([]types.ServiceProvider, error) {
	pingErr := db.Ping()

	if pingErr != nil {
		return []types.ServiceProvider{}, pingErr
	}

	query := `
		SELECT * FROM serviceprovider 
	`
	rows, err := db.Query(query)
	if err != nil {
		return []types.ServiceProvider{}, err
	}
	var serviceProviderList []types.ServiceProvider
	for rows.Next() {
		var serviceProvider types.ServiceProvider
		if err := rows.Scan(&serviceProvider.Id, &serviceProvider.Description, &serviceProvider.Service); err != nil {
			return []types.ServiceProvider{}, err
		}
		serviceProviderList = append(serviceProviderList, serviceProvider)
	}
	return serviceProviderList, nil
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
	var id string
	for rows.Next() {
		err := rows.Scan(
			&id,
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
