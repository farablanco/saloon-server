package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	//Id        int    `gorm:"AUTO_INCREMENT"`
	Firstname string `sql:"size:100"`
	Lastname  string `sql:"size:100"`
	Password  string `sql:"size:200"`
	Email     string `gorm:"unique_index;not null"`
	IsAdmin   bool   `sql:"size:1"`
}

func (u *User) HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}

func (u *User) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
	return err == nil
}
