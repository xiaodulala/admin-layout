package utils

import "golang.org/x/crypto/bcrypt"

// Encrypt 加密密码
func Encrypt(source string) (string, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashBytes), err
}


// Compare 比较密码
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
