package database

import (
	"database/sql"
	"smartjobsolutions/types"
)

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

func ClientPost(db *sql.DB, clientPost types.ClientPostJSON) (types.ClientPostResponse, error) {
	if err := db.Ping(); err != nil {
		return types.ClientPostResponse{}, err
	}
	query := `
		INSERT INTO clientposts VALUES (
			$1,
			NOW(),
			$2
		)	
		RETURNING id, post, created_at
	`
	var clientPostResponse types.ClientPostResponse
	err := db.QueryRow(query, clientPost.Id, clientPost.Post).Scan(&clientPostResponse.Id, &clientPostResponse.Post, &clientPostResponse.Timestamp)

	if err != nil {
		return types.ClientPostResponse{}, err
	}
	return clientPostResponse, err
}
