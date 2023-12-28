package db

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "postgres"
)

func InitDatabase() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}



func get_id_by_username(db *sql.DB, username string) (int, error) {
	var id int
	sqlStatement := "SELECT author_id FROM Author WHERE username = $1"

	err := db.QueryRow(sqlStatement, username).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func check_admin_by_id(db *sql.DB, id int) (bool, error) {
	var is_admin bool

	sqlStatement := "SELECT is_admin FROM Author WHERE author_id = $1"

	err := db.QueryRow(sqlStatement, id).Scan(&is_admin)
	if err != nil {
		return false, err
	}
	return is_admin, nil
}
