package DB

import (
	"database/sql"
	"fmt"
	"go_project/stream"
)

func Connect(host, port, user, password, dbname, sslmode string) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)
	db, err := sql.Open("postgres", connStr)
	stream.Catch(err)
	//defer db.Close()
	return db, err
}
