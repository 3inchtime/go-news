package middleware

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/sirupsen/logrus"
	"net/http"
	pb "news/proto"
)

type AuthService struct {
	auth pb.UserService
}

func (a *AuthService)Init(s micro.Service) {
	authService := pb.NewUserService("user-grpc", s.Client())
	a.auth = authService
}

func (a *AuthService)TokenCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": errors.New("no authorization"),
			})
			return
		}
		req := &pb.TokenCheckRequest{
			JwtToken: authHeader,
		}
		ctx := context.Background()
		res, err := a.auth.TokenCheck(ctx, req)
		if err != nil {
			logrus.Errorf(err.Error())
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": errors.New("token check error"),
			})
			return
		}
		logrus.Infof("token check res: %+v", res)
		c.Set("UserID", res.UserId)
		c.Set("UserName", res.UserName)
		c.Next()
	}
}