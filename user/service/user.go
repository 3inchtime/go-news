package service

import (
	"github.com/sirupsen/logrus"
	"user/model"
)

func (us *UserService) CreateNewUser (user *model.User) (int, error) {
	UserID, err := us.dao.CreateUser(user)
	if err != nil {
		logrus.Error(err.Error())
		return 0, err
	}
	return UserID, nil
}
