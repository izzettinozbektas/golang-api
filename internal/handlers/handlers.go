package handlers

import (
	"github.com/izzettinozbektas/golang-api/internal/driver"
	"github.com/izzettinozbektas/golang-api/internal/repository"
	"github.com/izzettinozbektas/golang-api/internal/repository/dbrepo"
	"github.com/izzettinozbektas/golang-api/internal/repository/filerepo"
)

var Repo *Repository

type Repository struct {
	DB repository.DatabaseRepo
	RQ repository.FileRepo
}

func NewRepo(db *driver.DB) *Repository {
	return &Repository{
		DB: dbrepo.NewMysqlRepo(db.SQL),
		RQ: filerepo.NewFileRepo(),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
