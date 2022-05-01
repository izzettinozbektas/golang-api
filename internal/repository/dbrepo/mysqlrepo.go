package dbrepo

import (
	"context"
	"github.com/izzettinozbektas/golang-api/internal/models"
	"log"
	"time"
)

func (m *mysqlDBRepo) UserCreate(res models.User) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into users (first_name, last_name, email, password, access_level, created_at, updated_at) 
			values (?, ?, ?, ?, ?, ?, ?);`

	_, err := m.DB.ExecContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Password,
		res.AccessLevel,
		res.CreatedAt,
		res.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
func (m *mysqlDBRepo) Users() ([]models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var users []models.User

	query := `select id, first_name, last_name, email, access_level, created_at, updated_at from users order by created_at asc`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var i models.User
		err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.AccessLevel,
			&i.CreatedAt,
			&i.UpdatedAt,
		)
		if err != nil {
			return users, err
		}

		users = append(users, i)
	}
	if err = rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}
func (m *mysqlDBRepo) UserUpdate(id int, res models.User) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update users set first_name = ?, last_name = ?, email = ?, password = ?, updated_at = ? where id = ?`
	_, err := m.DB.ExecContext(ctx, query,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Password,
		res.UpdatedAt,
		id,
	)
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}
