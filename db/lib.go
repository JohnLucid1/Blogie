package db

import (
	"blog/structs"
	"database/sql"
	"errors"
)


func GetAllUsers(db *sql.DB, user_id int) ([]structs.Author, error) {
	admin, err := check_admin_by_id(db, user_id); if err != nil {
		return []structs.Author{}, err
	}
	if admin {

	var Authors []structs.Author
	rows, err := db.Query("SELECT * FROM Author"); if err != nil {
		return []structs.Author{}, err
	}

	for rows.Next() {
		var author structs.Author
		err := rows.Scan(
			&author.AuthorID,
			&author.Is_admin,
			&author.Username,
			&author.Password,
			&author.Picture_path,
		)
		if err != nil {
			return []structs.Author{}, err
		}
		Authors = append(Authors, author)
	}

		return Authors, nil
	}
	return []structs.Author{}, errors.New("User is not admin")
}
