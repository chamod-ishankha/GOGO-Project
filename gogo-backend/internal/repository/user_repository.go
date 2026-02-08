package repository

import (
	"github.com/chamod-ishankha/gogo-project/gogo-backend/internal/model"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB *sqlx.DB
}

func (r *UserRepository) CreateUser(user *model.Register) error {
	_, err := r.DB.Exec(
		"INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4)",
		user.Name, user.Email, user.Password, user.Role,
	)
	return err
}

func (r *UserRepository) GetUserByEmail(email string) (*model.Register, error) {
	var user model.Register
	query := "SELECT id, name, email, password, role FROM users WHERE email=$1"
	err := r.DB.Get(&user, query, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByID(id int64) (*model.User, error) {
	var user model.User
	query := `
		SELECT id, name, email, role
		FROM gogo.users
		WHERE id = $1
	`
	err := r.DB.Get(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(id int64, name, email string) error {
	query := `
		UPDATE gogo.users
		SET name = $1, email = $2, updated_at = NOW()
		WHERE id = $3
	`
	_, err := r.DB.Exec(query, name, email, id)
	return err
}
