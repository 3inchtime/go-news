package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"user/model"
)

func (us *UserServer) CreateNewUser(c *gin.Context) {
	var ub model.UserBaseInfo
	if err := c.ShouldBindJSON(&ub); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	var user model.User
	user.Account = ub.Account
	user.UserName = ub.UserName
	user.HashPassword = ub.HashPassword

	UserID, err := us.Service.CreateNewUser(&user)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("New UserID: %d", UserID)
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
	return
}
