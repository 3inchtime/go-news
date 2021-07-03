package service

import (
	"github.com/sirupsen/logrus"
	"user/model"
)


func (us *UserService) CreateNewUser (u *model.User, ua *model.UserAccountInfo) error {
	err := us.dao.CreateUser(u, ua)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	return nil
}

func (us *UserService) CheckUser(ua *model.UserAccountInfo) (string, error) {
	userID, err := us.dao.CheckUserPwd(ua.Account, ua.Password)
	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}
	return userID, nil
}
