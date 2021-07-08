package service

import (
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"news/http"
	"time"
)


func NewNewsService(register *registry.Registry) web.Service {

	router := http.InitRouter(register)

	microService := web.NewService(
		web.Name("news-http"),
		web.RegisterTTL(time.Second*30),      //设置注册服务的过期时间
		web.RegisterInterval(time.Second*20), //设置间隔多久再次注册服务
		web.Address(":18003"),
		web.Handler(router),
		web.Registry(*register),
	)

	return microService
}