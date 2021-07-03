package utils

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	"user/model"
)

var SecretKey = "3inchtime"

func NewClaims(userID, account string){

}

func GenJWT(u *model.User) (string, error) {
	claims := model.UserLoginClaims{
		UserID:   u.UserID,
		UserName: u.UserName,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 签名过期时间
			Issuer:    "go-news",                    // 签名颁发者
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SecretKey))
}