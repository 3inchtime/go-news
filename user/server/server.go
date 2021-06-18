package server

import "user/service"

type UserServer struct {
	Service *service.UserService
}

func NewServer() *UserServer {
	return &UserServer{
		Service: service.NewUserService(),
	}
}