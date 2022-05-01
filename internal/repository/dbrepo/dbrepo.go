package dbrepo

import (
	"database/sql"
	"github.com/izzettinozbektas/golang-api/internal/repository"
)

type mysqlDBRepo struct {
	DB *sql.DB
}

func NewMysqlRepo(conn *sql.DB) repository.DatabaseRepo {
	return &mysqlDBRepo{
		DB: conn,
	}
}
