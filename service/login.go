// service/login_service.go
package service

import (
	"errors"
)

type LoginService interface {
	Login(username, password string) error
	Logout()
	IsLoggedIn() bool
}

type loginService struct {
	loggedIn bool
}

func NewLoginService() LoginService {
	return &loginService{}
}

func (s *loginService) Login(username, password string) error {
	if username == "jefri" && password == "123" {
		s.loggedIn = true
		return nil
	}
	return errors.New("invalid credentials")
}

func (s *loginService) Logout() {
	s.loggedIn = false
}

func (s *loginService) IsLoggedIn() bool {
	return s.loggedIn
}
