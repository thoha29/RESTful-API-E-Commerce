package helpers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userId int, username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userId
	claims["user_name"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	secret := "mysecretkey"
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
