package main

import (
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"user/grpc"
	"user/http"
)

func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("192.168.1.103"),
	)

	httpService := http.Init(&consulReg)
	err := httpService.Init()
	if err != nil {
		panic("user http server init error!")
	}

	grpcService := grpc.Init(&consulReg)
	grpcService.Init()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		err := grpcService.Run()
		if err != nil {
			panic("user grpc server start error!")
		}
	}()

	go func() {
		err := httpService.Run()
		if err != nil {
			panic("user http server start error!")
		}
	}()

	wg.Wait()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		switch <-c {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			wg.Done()
			return
		case syscall.SIGHUP:
		}
	}
}
