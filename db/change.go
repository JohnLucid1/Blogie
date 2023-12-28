package db

import (
	"database/sql"
	"time"
)

func CreateEdit(db *sql.DB, edit_content string, post_id int) error {
	stmt, err := db.Prepare(`
		UPDATE Post SET edit = $1, edit_created = $2 WHERE post_id = $3
	`)
	if err != nil {
		return err
	}

	current_time := time.Now()
	_, err = stmt.Exec(
		edit_content,
		current_time,
		post_id,
	)

	if err != nil {
		return err
	}
	return nil
}

func IncreaseLikeCount(postID int, db *sql.DB) error {
	sqlStatement := `
		UPDATE Post
		SET likes = likes + 1
		WHERE post_id = $1
		RETURNING likes
	`

	var newLikeCount int
	err := db.QueryRow(sqlStatement, postID).Scan(&newLikeCount)
	if err != nil {
		return err
	}

	return nil
}

func DecreaseDislikeCount(postID int, db *sql.DB) error {
	sqlStatement := `
		UPDATE Post
		SET dislikes = dislikes + 1
		WHERE post_id = $1
		RETURNING dislikes
	`

	var newDislikeCount int
	err := db.QueryRow(sqlStatement, postID).Scan(&newDislikeCount)
	if err != nil {
		return err
	}

	return nil
}
