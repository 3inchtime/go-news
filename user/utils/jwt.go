package utils

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"time"
	"user/model"
)

var SecretKey = []byte("go-news")

func GenJWT(u *model.User) (string, error) {
	logrus.Infof("gen jwt token info: %+v", *u)
	claims := model.UserLoginClaims{
		UserID:   u.UserID,
		UserName: u.UserName,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000, // 签名生效时间
			ExpiresAt: time.Now().Unix() + 3600, // 签名过期时间
			Issuer:    "go-news",                // 签名颁发者
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}

func ParseJWT(jwtToken string) (*model.UserLoginClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(jwtToken, &model.UserLoginClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		logrus.Errorf("parse jwt error: %s\n", err.Error())
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*model.UserLoginClaims); ok && tokenClaims.Valid {
			logrus.Infof("jwt token parse result: %+v", *claims)
			return claims, nil
		}
	}
	return nil, err
}
