package Entities

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
}

func (u *User) HashPassword() {
	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err.Error())
	}
	u.Password = string(password)
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func (u *User) HidePassword() {
	u.Password = ""
}
