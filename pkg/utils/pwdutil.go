package utils

import "golang.org/x/crypto/bcrypt"

// Encrypt 加密密码
func Encrypt(source string) (string, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashBytes), err
}

// CompareHashAndPassword 比较密码
func CompareHashAndPassword(e string, p string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {
		return false, err
	}
	return true, nil
}
