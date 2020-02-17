package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var secret = []byte("secret")

// GenerateToken provides a JWT token containing the user id
func GenerateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
	})

	tokenStr, err := token.SignedString([]byte(secret))

	return tokenStr, err
}

// CheckJWT checks that a JWT is valid
func CheckJWT(tokenStr string) error {
	_, err := parseJWT(tokenStr)
	return err
}

func parseJWT(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Incorrect signing method")
		}

		return secret, nil
	})

	return token, err
}

// ParseJWTClaims returns the fully parsed token claims
func ParseJWTClaims(tokenStr string) (jwt.MapClaims, error) {
	token, err := parseJWT(tokenStr)

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("Token was invalid")
}

// TokenFromHeader returns the token from the request
func TokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", nil
	}

	authSplit := strings.Split(authHeader, " ")
	if len(authSplit) != 2 || strings.ToLower(authSplit[0]) != "bearer" {
		return "", errors.New("Authorization header format should be Bearer {token}")
	}

	return authSplit[1], nil
}
