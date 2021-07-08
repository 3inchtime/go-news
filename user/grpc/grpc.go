package grpc

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	pb "user/proto"
	"user/server"
)

func Init(register *registry.Registry) micro.Service {
	microService := micro.NewService(
		micro.Name("user-grpc"),
		micro.Address(":18002"),
		micro.Registry(*register),
	)

	userGRPCServer := server.NewGRPCServer()
	err := pb.RegisterUserHandler(microService.Server(), userGRPCServer)

	if err != nil {
		panic(err)
	}

	return microService
}
