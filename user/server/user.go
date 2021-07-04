package server

import (
	"errors"
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
	ua.Password, _ = utils.HashAndSalt(ua.Password)

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

func (us *UserServer) UserLogin(c *gin.Context) {
	var ua model.UserAccountInfo
	if err := c.ShouldBindJSON(&ua); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	logrus.Infof("Login request: %+v", ua)

	accountInfo, err := us.Service.CheckUser(&ua)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusForbidden,
			gin.H{"error": "Login Failed"})
		return
	} else if accountInfo == nil {
		c.AbortWithStatusJSON(
			http.StatusForbidden,
			gin.H{"error": "Login Failed"})
		return
	} else {
		check := utils.ComparePasswords(accountInfo.Password, ua.Password)
		if check {
			logrus.Infof("check userID: %s ok", accountInfo.UserID)
			user, _ := us.Service.QueryUserInfo(accountInfo.UserID)
			jwt, _ := utils.GenJWT(user)
			c.Header("Authorization", jwt)
			c.AbortWithStatusJSON(
				http.StatusOK,
				gin.H{
					"user_id": accountInfo.UserID,
				})
			return
		} else {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				gin.H{"error": "Login Failed"})
			return
		}
	}
}

func (us *UserServer) ModifyUser(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": errors.New("no authorization"),
		})
		return
	}
	logrus.Infof("parse jwt token: %s\n", authHeader)
	ua, err := utils.ParseJWT(authHeader)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": errors.New("authorization parse failed"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"UserID": ua.UserID})
	return
}
