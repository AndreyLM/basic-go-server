package jwt

import (
	"crypto/rsa"
	"errors"
	"io/ioutil"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	// HoursInDay - hours in day
	HoursInDay = 24
	// DaysInWeek - days in week
	DaysInWeek = 7
)

var (
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
)

// Init - init jwt
func Init(PrivateKeyPath, PublicKeyPath string) error {
	signBytes, err := ioutil.ReadFile(PrivateKeyPath)
	if err != nil {
		return err
	}
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return err
	}
	verifyBytes, err := ioutil.ReadFile(PublicKeyPath)
	if err != nil {
		return err
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return err
	}

	return nil
}

// GetToken - gets token
func GetToken(id int64) string {
	token := jwt.New(jwt.SigningMethodRS512)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * HoursInDay * DaysInWeek).Unix()
	claims["iat"] = time.Now().Unix()
	claims["id"] = id
	token.Claims = claims

	tokenString, _ := token.SignedString(signKey)

	return tokenString
}

// IsTokenValid - checks if toket is valid
func IsTokenValid(val string) (int64, error) {
	token, err := jwt.Parse(val, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})

	switch err.(type) {
	case nil:
		if !token.Valid {
			return 0, errors.New("Token is invalid")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return 0, errors.New("Token is invalid")
		}
		return int64(claims["id"].(float64)), nil
	case *jwt.ValidationError:
		vErr := err.(*jwt.ValidationError)
		switch vErr.Errors {
		case jwt.ValidationErrorExpired:
			return 0, errors.New("Token Expired, get a new one")
		default:
			log.Println(vErr)
			return 0, errors.New("Error while parsing token")
		}
	default:
		return 0, errors.New("Unable to parse token")
	}
}
