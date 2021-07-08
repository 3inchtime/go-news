package main

import (
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"news/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("192.168.1.103"),
	)

	newsService := service.NewNewsService(&consulReg)

	newsService.Init()
	newsService.Run()
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
