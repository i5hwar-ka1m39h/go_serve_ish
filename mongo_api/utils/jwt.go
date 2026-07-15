package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret []byte
var jwt_string = "JWT_SECRET"

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateAccessToken(email string) (string, error) {

	secret = []byte(os.Getenv(jwt_string))
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}

func CreateRefreshToken(email string) (string, error) {
	secret = []byte(os.Getenv(jwt_string))
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}

func ValidateToken(tokenString string)(*Claims, error){
	jwt.ParseWithClaims(tokenString, &Claims{}, 
		func(token *jwt.Token)(interface{}, error){
if _, ok := token.Method.(*jwt.SigningMethodHS256); !ok{
				return  nil, fmt.Errorf("token authentication fucked %v", token.Header["alg"] )
			}
			return  
		}
}

