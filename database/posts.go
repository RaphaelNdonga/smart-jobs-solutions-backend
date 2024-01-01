package database

import "database/sql"

func LikePost(db *sql.DB, userId string, postId string) error {
	query := `
		UPDATE posts SET liked_by = ARRAY_APPEND(liked_by, $1) WHERE id = $2	
	`
	_, err := db.Exec(query, userId, postId)
	return err
}
