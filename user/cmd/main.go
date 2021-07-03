package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
	"user/server"
)

func main() {
	gin.SetMode(gin.DebugMode)

	userServer := server.NewServer()

	r := gin.New()

	userGroup := r.Group("/user")

	userGroup.POST("/reg", userServer.CreateNewUser)
	userGroup.POST("/login", userServer.UserLogin)
	userGroup.POST("/modify", userServer.ModifyUser)

	err := r.Run(":8888")

	if err != nil {
		panic("User Server Start Error!")
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
