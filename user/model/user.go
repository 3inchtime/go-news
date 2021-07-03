package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	UserID     string
	UserName   string
	Telephone  string
	Email      string
	Age        int
	CreateTime int64
	UpdateTime int64
}

type UserAccountInfo struct {
	UserID     string
	Account    string
	Password   string
	CreateTime int64
	UpdateTime int64
}

type UserLoginClaims struct {
	UserID         string
	UserName    string
	StandardClaims jwt.StandardClaims
}

func (u UserLoginClaims) Valid() error {
	panic("implement me")
}


