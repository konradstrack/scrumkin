package db

import (
	"database/sql"

	// import MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

// DB represents a database handle
type DB struct {
	*sql.DB
}

// New connects to the database specified by the Data Source Name and returns a database handle
func New(dsn string) (*DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
