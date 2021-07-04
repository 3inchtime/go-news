package main

import (
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"os"
	"os/signal"
	"syscall"
	"user/grpc"
	"user/http"
)

func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("192.168.1.103"),
	)

	httpService := http.Init(&consulReg)
	httpService.Init()

	grpcService := grpc.Init(&consulReg)
	grpcService.Init()

	err := grpcService.Run()
	if err != nil {
		panic("user http server start error!")
	}
	err = httpService.Run()
	if err != nil {
		panic("user http server start error!")
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		switch <-c {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			return
		case syscall.SIGHUP:
		}
	}
}
