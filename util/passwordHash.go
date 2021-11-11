package util

import "golang.org/x/crypto/bcrypt"

func PasswordHasher(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	return string(hash), err
}

//match the actual password with the hashed password
