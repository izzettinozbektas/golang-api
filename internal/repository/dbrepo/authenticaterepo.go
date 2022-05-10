package dbrepo

import (
	"context"
	"github.com/izzettinozbektas/golang-api/internal/helpers"
	"github.com/izzettinozbektas/golang-api/internal/models"
	"time"
)

func (m *mysqlDBRepo) Login(res models.Authentication) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user models.User

	query := `select id, first_name, last_name, email, password, access_level, created_at, updated_at from users where email = ?`
	row := m.DB.QueryRowContext(ctx, query, res.Email)
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.AccessLevel,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (m *mysqlDBRepo) LoginCreate(token string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tokenFields, _ := helpers.DecodeJWT(token)
	expDate := tokenFields["exp"]

	stmt := `insert into user_access_tokens (token, exp_date,created_at) 
			values (?, ?, ?);`

	_, err := m.DB.ExecContext(ctx, stmt,
		token,
		expDate,
		time.Now().Format("2006-01-02 15:04:05"),
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (m *mysqlDBRepo) TokenControl(token string) (models.UserAcccessInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var userAInfo models.UserAcccessInfo
	query := `select token, exp_date, created_at from user_access_tokens where token = ?`
	row := m.DB.QueryRowContext(ctx, query, token)
	err := row.Scan(
		&userAInfo.Token,
		&userAInfo.ExpDate,
		&userAInfo.CreatedAt,
	)
	if err != nil {
		return userAInfo, err
	}
	return userAInfo, nil
}
