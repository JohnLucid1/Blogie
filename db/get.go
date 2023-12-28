package db

import (
	"database/sql"
	"blog/structs"
)


func GetPostByID()  {}
func GetAuthorByID() {}
func GetCommentByID(){}


func Login(db *sql.DB, usernames string, password string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM Author WHERE username = $1 AND password = $2)"
	hashpass, err := HashPassword(password); if err != nil {
		return false, err
	}
	err = db.QueryRow(query, usernames, hashpass).Scan(&exists);
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetPassFromUsername(db *sql.DB, username  string) (string, error) {
	var password string
	query := "SELECT password FROM Author WHERE username = $1"
	err := db.QueryRow(query, username).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

// Dont need to verify token
func GetPostById(db *sql.DB, post_id int)  (string,error) {
	var post structs.Post

}
