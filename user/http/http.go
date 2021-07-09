package http

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"time"
	"user/server"
)


func Init(register *registry.Registry) web.Service {
	gin.SetMode(gin.DebugMode)

	userServer := server.NewHTTPServer()

	r := gin.New()

	userGroup := r.Group("/user")

	userGroup.POST("/reg", userServer.CreateNewUser)
	userGroup.POST("/login", userServer.UserLogin)
	userGroup.POST("/modify", userServer.ModifyUser)

	microService := web.NewService(
		web.Name("user-http"),
		web.RegisterTTL(time.Second*30),      //设置注册服务的过期时间
		web.RegisterInterval(time.Second*20), //设置间隔多久再次注册服务
		web.Address(":18001"),
		web.Handler(r),
		web.Registry(*register),
	)

	return microService
}
