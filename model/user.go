package model

type User struct {
	ID       int
	Username string
	Password string
	Role     string
}

type Response struct {
	StatusCode int
	Message    string
	Data       interface{}
}
