package models

import (
	"api/utils"
	"fmt"
	"time"
)

type User struct {
	Id           int       `json:"id"`
	UserName     string    `json:"user_name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
}

func NewUser(username, email, password string) *User {
	passwordHash, err := utils.HashPassword(password)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &User{
		UserName:     username,
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now().UTC(),
	}
}

func ValidatePassword(password string, user *User) bool {
	return utils.CheckPasswordHash(password, user.PasswordHash)
}

type CreateUserRequest struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserId int    `json:"user_id"`
	Token  string `json:"token"`
}
