package authorization

import (
	"devbook-api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userId uint64) (string, error) {

	permissions := jwt.MapClaims{}

	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString(config.SecretKey)
}

func ValidateToken(r *http.Request) error {

	tokenString := getToken(r)

	token, err := jwt.Parse(tokenString, getValidationKey)

	if err != nil {

		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		return nil
	}

	return errors.New("invalid token")
}

func getToken(r *http.Request) string {

	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {

		return strings.Split(token, " ")[1]
	}

	return ""
}

func getValidationKey(token *jwt.Token) (interface{}, error) {

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signature method. %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
