package credential

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"strconv"
	"time"
)

var ErrInvalidToken = errors.New("invalid token")
var ErrTokenExpired = errors.New("token expired")

const tokenLifetime int64 = 1.8e+3

type credential interface {
	sign(privateKey string) (string, error)
	validate()
}

type JwtFactory struct {
}

func (jwtF JwtFactory) Sign(data map[string]interface{}, privateKey string) (jwtHash string, err error) {
	epochTime := time.Now().Unix()
	data["time"] = epochTime
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(data))

	return token.SignedString([]byte(privateKey))
}

func (jwtF JwtFactory) Validate(tokenString string, privateKey string) error {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(privateKey), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return ErrInvalidToken
	}

	tokenTimeString := fmt.Sprintf("%v", claims["time"])
	tokenTime, _ := strconv.ParseFloat(tokenTimeString, 64)
	currentTime := time.Now().Unix()

	if currentTime-int64(tokenTime) > tokenLifetime {
		return ErrTokenExpired
	}

	return nil
}
