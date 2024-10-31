package repository

import (
	"database/sql"

	"github.com/go-embed-go-web/model"
)

type UserRepository interface {
	BaseRepository[model.User]
	GetUserLogin(payload model.User) (*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

// List implements UserRepository.
// func (u *userRepository) List() ([]model.User, error) {
// 	query := `SELECT id, username, password, email FROM customers`
// 	rows, err := u.db.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var users
// }

func (u *userRepository) Create(payload model.User) error {
	query := "INSERT INTO User(username, password, role) VALUES($1, $2, $3)"
	_, err := u.db.Exec(query, payload.Username, payload.Password, payload.Role)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) GetUserLogin(payload model.User) (*model.User, error) {
	query := `SELECT id, username, password FROM users WHERE username=$1 AND password=$2`
	var userRespon model.User
	err := u.db.QueryRow(query, payload.Username, payload.Password).Scan(&userRespon.ID, &userRespon.Username, &userRespon.Password)
	if err != nil {
		return nil, err
	}
	return &userRespon, nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
