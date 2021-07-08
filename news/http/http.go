package http

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"net/http"
	"news/middleware"
)

func InitRouter(register *registry.Registry) *gin.Engine {
	internalService := micro.NewService(
		micro.Name("internal-news"),
		micro.Registry(*register),
	)
	authService := middleware.AuthService{}
	authService.Init(internalService)

	gin.SetMode(gin.DebugMode)

	r := gin.New()
	r.Use(authService.TokenCheck())

	newsGroup := r.Group("/news")

	newsGroup.GET("/list", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"user_id": "",
		})
		return
	})

	return r
}
