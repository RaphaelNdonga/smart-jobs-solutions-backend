package database

import (
	"database/sql"
	"log"
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

func ClientPost(db *sql.DB, clientPost types.PostJSON) error {
	if err := db.Ping(); err != nil {
		return err
	}
	log.Print("client post: ", clientPost)
	query := `
		INSERT INTO clientposts VALUES (
			$1,
			NOW(),
			$2,
			$3
		)	
	`
	_, err := db.Exec(query, clientPost.Id, clientPost.Post, clientPost.Service)

	return err
}

func GetClientPosts(db *sql.DB) ([]types.ClientPostResponse, error) {
	if err := db.Ping(); err != nil {
		return []types.ClientPostResponse{}, err
	}
	query := `
	SELECT userdetails.username, clientposts.post, clientposts.created_at, userdetails.location FROM clientposts INNER JOIN userdetails ON userdetails.id = clientposts.id
	`

	rows, err := db.Query(query)
	if err != nil {
		return []types.ClientPostResponse{}, err
	}
	var clientPostResponses []types.ClientPostResponse
	for rows.Next() {
		var clientPostResponse types.ClientPostResponse
		rows.Scan(&clientPostResponse.Username, &clientPostResponse.Post, &clientPostResponse.CreatedAt, &clientPostResponse.Location)
		clientPostResponses = append(clientPostResponses, clientPostResponse)
	}
	return clientPostResponses, nil
}
