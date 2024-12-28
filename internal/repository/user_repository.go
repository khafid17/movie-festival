package repository

import (
	"database/sql"
	"movie-festival/internal/entity"
)

type UserRepository interface {
	Create(user *entity.User) (int64, error)
	FindByUsername(username string) (*entity.User, error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(user *entity.User) (int64, error) {
	query := `INSERT INTO users (username, password) VALUES (?, ?)`
	result, err := r.db.Exec(query, user.Username, user.Password)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *UserRepositoryImpl) FindByUsername(username string) (*entity.User, error) {
	query := `SELECT id, username, password FROM users WHERE username = ?`
	row := r.db.QueryRow(query, username)
	user := &entity.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
