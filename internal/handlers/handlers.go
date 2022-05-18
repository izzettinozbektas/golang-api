package handlers

import (
	"github.com/izzettinozbektas/golang-api/internal/driver"
	"github.com/izzettinozbektas/golang-api/internal/repository"
	"github.com/izzettinozbektas/golang-api/internal/repository/dbrepo"
)

var Repo *Repository

type Repository struct {
	DB repository.DatabaseRepo
}

func NewRepo(db *driver.DB) *Repository {
	return &Repository{
		DB: dbrepo.NewMysqlRepo(db.SQL),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
