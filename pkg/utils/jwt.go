package utils

import (
	"errors"
	"fmt"
	"order-food-service/entitys"
	"time"

	"github.com/golang-jwt/jwt"
)

type Users struct {
	entitys.User
}

type CustomClaims struct {
	Id       int
	UserName string
	Password string
	jwt.StandardClaims
}

var MySecret = []byte("密钥")

// 创建 Token
func GenToken(id int, userName string, password string) (string, error) {
	claim := CustomClaims{
		Id:       id,
		UserName: userName,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(), //60分钟后过期
			// 签名时间
			IssuedAt: time.Now().Unix(),
			Issuer:   "test", //签发人
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}

// 解析 token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		fmt.Println(" token parse err:", err)
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// 刷新 Token
func RefreshToken(tokenStr string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(time.Minute * 60).Unix()
		return GenToken(claims.Id, claims.UserName, claims.Password)
	}
	error := errors.New("Cloudn't handle this token")
	return "", error
}
