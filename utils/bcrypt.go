package utils

import "golang.org/x/crypto/bcrypt"

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CompareHash(password string, hashedText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedText), []byte(password))
	return err == nil
}
