package forum

import "golang.org/x/crypto/bcrypt"

//This adds encryption to passwords inputted by users in order to add a layer of security
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hash), err
}
