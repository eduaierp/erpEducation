package repository

import (
	"auth-service/model"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	user := &model.User{}
	err := r.DB.QueryRow("SELECT id, username, password, role FROM users WHERE username = $1", username).
		Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}
