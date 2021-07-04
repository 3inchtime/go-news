package http

import (
	"github.com/gin-gonic/gin"
	"user/server"
)

func Init() *gin.Engine {
	gin.SetMode(gin.DebugMode)

	userServer := server.NewServer()

	r := gin.New()

	userGroup := r.Group("/user")

	userGroup.POST("/reg", userServer.CreateNewUser)
	userGroup.POST("/login", userServer.UserLogin)
	userGroup.POST("/modify", userServer.ModifyUser)

	return r
}