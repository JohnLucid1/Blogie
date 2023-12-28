package db

import (
	"blog/structs"
	"database/sql"

	"github.com/labstack/gommon/log"
	pq "github.com/lib/pq"
)


func CreateAuthor(db *sql.DB, newAuthor structs.Author) error {
	stmt, err := db.Prepare(`
		INSERT INTO Author (is_admin, username, password, picture)
		VALUES ($1, $2, $3, $4)
	`)
	if err != nil {
		return err
	}

	hashed, err := HashPassword(newAuthor.Password)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		newAuthor.Is_admin,
		newAuthor.Username,
		hashed,
		newAuthor.Picture_path,
	)
	if err != nil {
		return err
	}
	return nil
}


func CreatePost(db *sql.DB, newPost structs.Post, username string) error {
	author_id, err := get_id_by_username(db, username)
	if err != nil {
		log.Error("Can't get author id by username")
		return err
	}

	stmt, err := db.Prepare(`
		INSERT INTO Post (title, subtitle, categories, author_id, content)
		VALUES ($1, $2, $3, $4, $5)
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		newPost.Title,
		newPost.Subtitle,
		pq.StringArray(newPost.Categories),
		author_id,
		newPost.Content,
	)
	if err != nil {
		return err
	}

	return nil
}

func CreateComment(db *sql.DB, newComment structs.Comment) error {
	stmt, err := db.Prepare(`
		INSERT INTO Comment (author_id, foreign_post_id, comment_content)
		VALUES ($1, $2, $3)
	`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		newComment.Author_id,
		newComment.Foreign_post_id,
		newComment.Content,
	)
	if err != nil {
		return err
	}
	return nil
}
