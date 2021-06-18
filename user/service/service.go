package service

import (
	"user/dao"
)

type UserService struct {
	dao *dao.Dao
}

func NewUserService() *UserService {
	return &UserService{
		dao: dao.NewDao(),
	}
}