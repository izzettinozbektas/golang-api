package driver

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}

	dbConn.SQL = d
	return dbConn, err
}
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
