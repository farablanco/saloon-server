package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func hashAndSalt(pwd []byte) (string, error) {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash), err
}
