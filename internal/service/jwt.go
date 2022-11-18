package service

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type Claims struct {
	Username string
	Role     string
	jwt.StandardClaims
}

var JWTKey = []byte("the_secret_key")

func (a *authService) generateToken(username string, role string) (string, error) {
	expiredTime := time.Now().Add(time.Hour * 24).Unix()

	claims := &Claims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(JWTKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
