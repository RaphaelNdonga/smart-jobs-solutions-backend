package database

import (
	"database/sql"
	"log"
	"smartjobsolutions/types"
)

func GetClient(db *sql.DB, userId string) (types.Client, error) {
	var client types.Client
	if err := db.Ping(); err != nil {
		return types.Client{}, err
	}
	query := `
		SELECT * FROM client WHERE id = $1	
	`
	err := db.QueryRow(query, userId).Scan(&client.Id, &client.Service)
	if err != nil {
		return types.Client{}, err
	}
	return client, nil
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

func ClientPost(db *sql.DB, clientPost types.PostJSON) error {
	if err := db.Ping(); err != nil {
		return err
	}
	log.Print("client post: ", clientPost)
	query := `
		INSERT INTO posts (id, created_at, post, service, user_type) VALUES (
			$1,
			NOW(),
			$2,
			$3,
			'client'
		)	
	`
	_, err := db.Exec(query, clientPost.Id, clientPost.Post, clientPost.Service)

	return err
}

func GetClientPosts(db *sql.DB, service string) ([]types.PostResponse, error) {
	if err := db.Ping(); err != nil {
		return []types.PostResponse{}, err
	}
	var query string
	var rows *sql.Rows
	var err error

	if service == "" {
		query = `
			SELECT userdetails.username, posts.post, posts.created_at, userdetails.location, posts.service FROM posts INNER JOIN userdetails ON userdetails.id = posts.id; 
		`
		rows, err = db.Query(query)
	} else {
		query = `
			SELECT userdetails.username, posts.post, posts.created_at, userdetails.location, posts.service FROM posts INNER JOIN userdetails ON userdetails.id = posts.id WHERE service = $1 AND user_type = 'client';
		`
		rows, err = db.Query(query, service)
	}
	if err != nil {
		return []types.PostResponse{}, err
	}
	var clientPostResponses []types.PostResponse
	for rows.Next() {
		var clientPostResponse types.PostResponse
		rows.Scan(&clientPostResponse.Username, &clientPostResponse.Post, &clientPostResponse.CreatedAt, &clientPostResponse.Location, &clientPostResponse.Service)
		clientPostResponses = append(clientPostResponses, clientPostResponse)
	}
	log.Print("clientpost responses: ", clientPostResponses)
	return clientPostResponses, nil
}
