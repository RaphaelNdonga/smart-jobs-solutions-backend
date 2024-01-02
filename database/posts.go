package database

import "database/sql"

func LikePost(db *sql.DB, userId string, postId string) error {
	query := `
		INSERT INTO likes VALUES ($1, $2)
	`
	_, err := db.Exec(query, postId, userId)
	return err
}
