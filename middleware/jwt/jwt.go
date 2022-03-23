package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"server/middleware/config"
)

type MyClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var (
	_secret = os.Getenv("TOKEN_SECURITY_KEY")
)

func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(*jwt.Token) (i interface{}, err error) {
		return []byte(_secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func CreateToken(email string) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(config.TokenExpired())).Unix(),
		},
	})
	token, err := at.SignedString([]byte(_secret))
	if err != nil {
		return "", err
	}
	return token, nil
}
