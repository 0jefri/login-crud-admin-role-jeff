package handler

import (
	"database/sql"
	"fmt"
)

func Logout(db *sql.DB) {
	fmt.Println("Logged out successfully.")
}
