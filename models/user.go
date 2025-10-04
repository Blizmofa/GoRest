package models

import (
	"errors"

	"example.com/go-rest/db"
	"example.com/go-rest/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

const (
	sqlInsertUser        = `INSERT INTO users(email, password) VALUES (?, ?)`
	sqlSelectUserByEmail = `SELECT id, password FROM users WHERE email = ?`
)

func (user *User) Save() error {

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	result, err := db.ExecStatement(sqlInsertUser, user.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	user.ID = id
	return err
}

func (user *User) ValidateCredentials() error {

	row := db.DB.QueryRow(sqlSelectUserByEmail, user.Email)

	var retrievedPassword string
	err := row.Scan(&user.ID, &retrievedPassword)

	if err != nil {
		return errors.New("Credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}

	return nil
}
