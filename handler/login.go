package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/go-embed-go-web/model"
	"github.com/go-embed-go-web/repository"
	"github.com/go-embed-go-web/service"
)

func Login(db *sql.DB) {
	//input
	users := model.User{}
	file, err := os.Open("body.json")

	if err != nil {
		fmt.Println("Error: ", err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&users)
	if err != nil && err != io.EOF {
		fmt.Println("error decoding JSON: ", err)
	}

	// proses
	repo := repository.NewUserRepository(db)
	userService := service.NewUserService(repo)

	user, err := userService.LoginService(users)

	if err != nil {
		fmt.Println("err:", err)
		response := model.Response{
			StatusCode: 404,
			Message:    "Account not found",
			Data:       nil,
		}
		jsonData, err := json.MarshalIndent(response, "", "")
		if err != nil {
			fmt.Println("err: ", err)
		}
		fmt.Println(string(jsonData))
	} else {
		response := model.Response{
			StatusCode: 200,
			Message:    "login success",
			Data:       user,
		}
		jsonData, err := json.MarshalIndent(response, "", "")

		if err != nil {
			fmt.Println("err :", err)
		}

		fmt.Println(string(jsonData))
	}
}
