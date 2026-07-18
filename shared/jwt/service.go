package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("ajhdaudhsauhdnajdhsaudha")

func CreateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userID,
		"exp":    time.Now().Add(1 * time.Hour).Unix(),
		"lat":    time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return signed, nil
}
