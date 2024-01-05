package database

import (
	"database/sql"
	"smartjobsolutions/types"
)

func LikePost(db *sql.DB, userId string, postId string) error {
	query := `
		INSERT INTO likes VALUES ($1, $2)
	`
	_, err := db.Exec(query, postId, userId)
	return err
}

func UnlikePost(db *sql.DB, userId string, postId string) error {
	query := `
		DELETE FROM likes WHERE post_id = $1 AND user_id = $2	
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

func CommentPost(db *sql.DB, postId string, userId string, comment string) error {
	query := `
		INSERT INTO comments (post_id, user_id, comment, created_at) VALUES (
			$1, $2, $3, NOW()
		)	
	`
	_, err := db.Exec(query, postId, userId, comment)
	return err
}

func GetComments(db *sql.DB, postId string) ([]types.CommentResponse, error) {
	query := `
	SELECT userdetails.username, comments.comment, comments.created_at FROM comments INNER JOIN userdetails ON userdetails.id = comments.user_id WHERE comments.post_id = $1
	`
	rows, err := db.Query(query, postId)
	if err != nil {
		return []types.CommentResponse{}, err
	}
	var comments []types.CommentResponse
	for rows.Next() {
		var comment types.CommentResponse
		rows.Scan(&comment.Username, &comment.Comment, &comment.CreatedAt)
		comments = append(comments, comment)
	}
	return comments, nil
}

func GetUserPosts(db *sql.DB, userId string) ([]types.PostResponse, error) {
	query := `
		SELECT posts.post_id, userdetails.username, posts.post, posts.created_at, userdetails.location, posts.service FROM posts INNER JOIN userdetails ON userdetails.id = posts.user_id WHERE posts.user_id = $1; 
	`
	rows, err := db.Query(query, userId)

	if err != nil {
		return []types.PostResponse{}, err
	}

	var posts []types.PostResponse

	for rows.Next() {
		var post types.PostResponse
		rows.Scan(&post.Id, &post.Username, &post.Post, &post.CreatedAt, &post.Location, &post.Service)
		posts = append(posts, post)
	}
	return posts, nil
}
