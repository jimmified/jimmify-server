package auth

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//JPASS server password
var JPASS string

//JKEY encryption key
var JKEY string

//Init get the password and key from environment variables
func Init() {
	JPASS = os.Getenv("JPASS")
	if JPASS == "" {
		log.Fatal("JPASS Password environment variable is not set")
	}
	JKEY = os.Getenv("JKEY")
	if JKEY == "" {
		log.Fatal("JKEY Encryption key is not set")
	}
}

//CheckToken verfies a token
func CheckToken(tokenString string) (string, error) {
	//check for valid jwt syntax
	match, _ := regexp.MatchString("\\w+\\.\\w+\\.\\w+", tokenString)
	if match != true {
		return "", errors.New("Invalid Token")
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JKEY), nil
	})

	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		return claims["user"].(string), nil
	}
	return "", errors.New("Invalid Token")
}

//CreateToken genreate a new token
func CreateToken(username string) (string, error) {

	// Create the short-term token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": username,
		"exp":  time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(JKEY))

	return tokenString, err
}
