package repository

import (
	"fitter-srv/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, surname, email, avatar, description, login, password_hash, locked, is_trainer, is_admin, is_moder) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Name, user.Surname, user.Email, user.Avatar, user.Description, user.Login, user.Password, user.Locked, user.IsTrainer, user.IsAdmin, user.IsModer)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(login, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 AND password_hash=$2", userTable)
	err := r.db.Get(&user, query, login, password)
	return user, err
}
