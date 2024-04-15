package utils

import (
	"awesomeProject/internal/domain"
	"crypto/rand"
	"crypto/sha256"
	"go.uber.org/zap"
)

func SaltPassword(password string, salt string) string {
	saltedPassword := password + salt

	hashedPassword := hash(saltedPassword)

	return hashedPassword
}

func CreateSalt() (string, error) {
	salt := make([]byte, 8)

	_, err := rand.Read(salt)

	if err != nil {
		zap.L().Error("failed to generate salt", zap.Error(err))

		return "", err
	}

	return string(salt), err
}

func hash(s string) string {
	h := sha256.New()

	h.Write([]byte(s))

	hashedStr := h.Sum(nil)

	return string(hashedStr)
}

func ComparePasswords(password, salt, hashedPassword string) error {
	userPassword := hash(password + salt)

	if userPassword != hashedPassword {
		zap.L().Error("passwords do not match")

		return domain.ErrInvalidCredentials
	}

	return nil
}
