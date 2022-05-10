package repository

import (
	"github.com/izzettinozbektas/golang-api/internal/models"
)

type DatabaseRepo interface {
	UserCreate(res models.User) error
	Users() ([]models.User, error)
	UserUpdate(id int, res models.User) error
	User(id int) (models.User, error)
	UserDelete(id int) (bool, error)
	Login(res models.Authentication) (models.User, error)
	LoginCreate(token string) (bool, error)
	TokenControl(token string) (models.UserAcccessInfo, error)
}
