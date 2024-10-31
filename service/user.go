package service

import (
	"fmt"

	"github.com/go-embed-go-web/model"
	"github.com/go-embed-go-web/repository"
)

type UserService interface {
	RegisterNewUser(payload model.User) error
	LoginService(payload model.User) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func (u *userService) RegisterNewUser(payload model.User) error {
	if payload.ID == 0 || payload.Username == "" || payload.Password == "" || payload.Role == "" {
		return fmt.Errorf("all payload is required")
	}

	err := u.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create user: %s", err)
	}
	return nil
	// userRepo := repository.NewUserRepository(&payload)
}

func (u *userService) LoginService(payload model.User) (*model.User, error) {
	if payload.Username == "" || payload.Password == "" {
		return nil, fmt.Errorf("username, password is required")
	}

	users, err := u.repo.GetUserLogin(payload)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func NewUserService(usrRepo repository.UserRepository) UserService {
	return &userService{repo: usrRepo}
}
