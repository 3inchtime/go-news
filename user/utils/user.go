package utils

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func GenUUID() (string, error) {
	userID, err := uuid.NewUUID()
	if err != nil {
		logrus.Errorf("Gen UUID Error: %s\n", err.Error())
		return "", err
	}
	return userID.String(), nil
}
