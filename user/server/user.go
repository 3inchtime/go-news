package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
	"user/model"
	"user/utils"
)

func (us *UserServer) CreateNewUser(c *gin.Context) {
	var ua model.UserAccountInfo
	if err := c.ShouldBindJSON(&ua); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	userID, err := utils.GenUUID()
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	createTime := time.Now().Unix()
	ua.UserID = userID
	ua.CreateTime = createTime
	ua.UpdateTime = createTime

	var u model.User
	u.UserID = userID
	u.CreateTime = ua.CreateTime
	u.UpdateTime = ua.CreateTime

	err = us.Service.CreateNewUser(&u, &ua)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("New UserID: %s", userID)
	c.JSON(http.StatusOK, gin.H{"UserID": userID})
	return
}

func (us *UserServer) UserLogin(c *gin.Context)  {
	var ua model.UserAccountInfo
	if err := c.ShouldBindJSON(&ua); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	userID, err := us.Service.CheckUser(&ua)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusForbidden,
			gin.H{"error": "Login Failed"})
		return
	} else if userID == "" {
		c.AbortWithStatusJSON(
			http.StatusForbidden,
			gin.H{"error": "Login Failed"})
		return
	} else {
		c.AbortWithStatusJSON(
			http.StatusOK,
			gin.H{"user_id": userID})
		return
	}
}