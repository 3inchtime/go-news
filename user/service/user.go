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

func (us *UserService) CheckUser(ua *model.UserAccountInfo) (*model.UserAccountInfo, error) {
	accountInfo, err := us.dao.CheckUserPwd(ua.Account)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	logrus.Infof("Query account result: %+v",*accountInfo)
	return accountInfo, nil
}

func (us *UserService) QueryUserInfo (userID string) (*model.User, error) {
	userInfo, err := us.dao.QueryUserInfo(userID)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	logrus.Infof("Query user result: %+v",*userInfo)
	return userInfo, nil
}
