package server

import "user/service"

type HTTPServer struct {
	Service *service.UserService
}

type GRPCServer struct {
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{
		Service: service.NewUserService(),
	}
}
