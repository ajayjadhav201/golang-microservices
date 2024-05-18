package common

import "golang.org/x/crypto/bcrypt"

func Encrypt(s string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	return string(bytes) //,nil
}

func Compare(s, encrypted string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(s))
	return err == nil
}
