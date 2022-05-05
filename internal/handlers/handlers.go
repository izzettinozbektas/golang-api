package handlers

import (
	"github.com/izzettinozbektas/golang-api/internal/driver"
	"github.com/izzettinozbektas/golang-api/internal/repository"
	"github.com/izzettinozbektas/golang-api/internal/repository/dbrepo"
)

func ConnectionToDB(db *driver.DB) *Repository {
	return &Repository{
		DB: dbrepo.NewMysqlRepo(db.SQL),
	}
}

type Repository struct {
	DB repository.DatabaseRepo
}
