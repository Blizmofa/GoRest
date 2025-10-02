package models

import (
	"example.com/go-rest/db"
	"example.com/go-rest/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	result, err := statement.Exec(user.Email, hashedPassword)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	user.ID = id
	return err
}
