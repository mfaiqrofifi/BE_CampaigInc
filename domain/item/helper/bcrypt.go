package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(dt string) (error, string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dt), bcrypt.DefaultCost)
	if err != nil {
		return err, ""
	}
	return nil, string(hashedPassword)
}

func CheckPassword(hasshedPw, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hasshedPw), []byte(password))
}
