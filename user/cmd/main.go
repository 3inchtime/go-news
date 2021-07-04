package main

import (
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/consul/v2"
	"os"
	"os/signal"
	"syscall"
	"time"
	"user/http"
)

func main() {
	httpServer := http.Init()

	consulReg := consul.NewRegistry(
		registry.Addrs("192.168.1.103"),
	)

	microService := web.NewService(
		web.Name("go-news-user-http"),
		web.RegisterTTL(time.Second*30),      //设置注册服务的过期时间
		web.RegisterInterval(time.Second*20), //设置间隔多久再次注册服务
		web.Address(":18001"),
		web.Handler(httpServer),
		web.Registry(consulReg),
	)

	microService.Init()
	err := microService.Run()
	if err != nil {
		panic("micro server register error!")
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
