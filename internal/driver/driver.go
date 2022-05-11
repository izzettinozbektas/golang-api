package driver

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// DB ...
type DB struct {
	SQL *sql.DB
	// Mgo *mgo.database
}

// DBConn ...
var dbConn = &DB{}

// ConnectSQL ...
func ConnectSQL() (*DB, error) {

	//dbName := os.Getenv("DB_NAME")
	//dbPass := os.Getenv("DB_PASS")
	//dbHost := os.Getenv("DB_HOST")
	//dbPort := os.Getenv("DB_PORT")
	// geçici olarak basit kullanım
	dbName := "golang-db"
	dbUname := "golang"
	dbPass := "golangpass"
	dbHost := "app-mysql" // mysql container name olmalı

	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?parseTime=true",
		dbUname,
		dbPass,
		dbHost,
		dbName,
	)
	d, err := sql.Open("mysql", dbSource)
	if err != nil {
		panic(err)
	}

	dbConn.SQL = d
	return dbConn, err
}

// ConnectMgo ....
func ConnectMgo(host, port, uname, pass string) error {

	return nil
}
