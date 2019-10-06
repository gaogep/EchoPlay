package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gaogep/EchoPlay/settings"
	"time"
)

var jwtSecret = []byte(settings.GlobalConf["JWTSECRET"].(string))

type Claims struct {
	NickName string `json:"nickname"`
	jwt.StandardClaims
}

func GenerateToken(nickname string, passwd string) (string, error) {
	nowTime := time.Now()
	expTime := nowTime.Add(time.Hour * 72)

	claims := Claims{
		NickName: nickname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
			Issuer:    "echo_play",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
