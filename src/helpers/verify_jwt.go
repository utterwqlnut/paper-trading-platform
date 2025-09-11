package helpers
/*
JWT verification and creation
*/

import (
	"os"
	"time"
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Username string
	jwt.RegisteredClaims
}

func Creator() func(user string) (string, error) {
	mySignKey := []byte(os.Getenv("SIGN_KEY"))

	return func(user string) (string, error) {
		claims := CustomClaims {
			user,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(24*time.Hour)),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(mySignKey)

		return ss,err
	}
}

func Verifier() func(tokenStr string) (string, error) {
	mySignKey := []byte(os.Getenv("SIGN_KEY"))	
	
	return func(user string) (string, error) {
		token, err := jwt.ParseWithClaims(user, &CustomClaims{}, func(token *jwt.Token) (any, error) {
			return mySignKey, nil
		})
		if err != nil {
			return "", errors.New("Invalid Token Could Not Process Request")
		} else if claims, ok := token.Claims.(*CustomClaims); ok {
			return claims.Username, nil;
		} else {
			return "", errors.New("Invalid Token Could Not Process Request")
		}
	}
}