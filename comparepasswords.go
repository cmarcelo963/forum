package forum

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	//Compares the encrypted password from user input as well as the one in the database without deciphering it for security
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
