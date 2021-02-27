package internal

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	EXPIRED_TOKEN = errors.New("expired token")
	ILLEGAL_TOKEN = errors.New("illegal token")
)

type MyCustomClaims struct {
	jwt.StandardClaims
	Account string `json:"account"` // 账号
}

func CreateToken(sighKey, account string) string {
	claims := MyCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "lqh",
		},
		account,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	res, err := token.SignedString([]byte(sighKey))
	if err != nil {
		panic("jwt sigh error: " + err.Error())
	}
	return res
}

func ValidToken(sighKey, tokenString string) (string, error) {
	return at(time.Unix(0, 0), func() (string, error) {
		token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(sighKey), nil
		})
		if err != nil {
			return "", err
		}

		if claims, ok := token.Claims.(*MyCustomClaims); ok {
			if token.Valid {
				return claims.Account, nil
			}
			return "", EXPIRED_TOKEN
		} else {
			return "", ILLEGAL_TOKEN
		}
	})
}

func at(t time.Time, f func() (string, error)) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	defer func() {
		jwt.TimeFunc = time.Now
	}()
	return f()
}
