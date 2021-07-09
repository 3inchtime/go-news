package grpc

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/sirupsen/logrus"
	pb "user/proto"
	"user/server"
	"user/utils"
)

var JaegerAddr = "192.168.1.103:6831"

func Init(register *registry.Registry) micro.Service {
	jaegerTracer, _, err := utils.NewJaegerTracer("user-grpc", JaegerAddr)
	if err != nil {
		logrus.Fatal(err)
	}

	microService := micro.NewService(
		micro.Name("user-grpc"),
		micro.Address(":18002"),
		micro.Registry(*register),
		micro.WrapHandler(opentracing.NewHandlerWrapper(jaegerTracer)),
	)

	userGRPCServer := server.NewGRPCServer()
	err = pb.RegisterUserHandler(microService.Server(), userGRPCServer)

	if err != nil {
		panic(err)
	}

	return microService
}
