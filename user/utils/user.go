package utils

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func GenUUID() (string, error) {
	userID, err := uuid.NewUUID()
	if err != nil {
		logrus.Errorf("Gen UUID Error: %s\n", err.Error())
		return "", err
	}
	return userID.String(), nil
}


func HashAndSalt(pwdStr string) (pwdHash string, err error) {
	pwd := []byte(pwdStr)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return
	}
	pwdHash = string(hash)
	return
}

func ComparePasswords(hashedPwd string, plainPwd string) bool {
	logrus.Info("Checking password")
	byteHash := []byte(hashedPwd)
	bytePwd := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	if err != nil {
		return false
	}
	return true
}
