package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"user/server"
)

func main() {
	gin.SetMode(gin.DebugMode)

	userServer := server.NewServer()

	r := gin.New()

	r.POST("/user/reg", userServer.CreateNewUser)
	err := r.Run(":8888")

	logrus.Info("User Server Runner Start")
	if err != nil {
		logrus.Error(err)
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