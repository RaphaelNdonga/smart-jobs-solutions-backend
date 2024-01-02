package database

import "database/sql"

func LikePost(db *sql.DB, userId string, postId string) error {
	query := `
		INSERT INTO likes VALUES ($1, $2)
	`
	_, err := db.Exec(query, postId, userId)
	return err
}

func GetLikes(db *sql.DB, postId string) ([]string, error) {
	query := `
		SELECT user_id FROM likes WHERE post_id = $1	
	`
	rows, err := db.Query(query, postId)
	var users []string
	for rows.Next() {
		var user string
		rows.Scan(&user)
		users = append(users, user)
	}
	return users, err
}
