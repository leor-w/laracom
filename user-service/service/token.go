package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/leor-w/laracom/user-service/model"
	"github.com/leor-w/laracom/user-service/repo"
)

var (
	key = []byte("laracomUserTokenKeySecret")
)

type CustomClaims struct {
	User *model.User
	jwt.StandardClaims
}

type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *model.User) (string, error)
}

type TokenService struct {
	Repo repo.UserRepositoryInterface
}

func (srv *TokenService) Decode(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func (srv *TokenService) Encode(user *model.User) (string, error) {
	expireToken := time.Now().Add(time.Hour * 72).Unix()

	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "laracom.user.service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}
