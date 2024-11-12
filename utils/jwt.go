package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	secretKey = []byte("your-secret-key")
)

// var (
// 	secretKey string = "secret-key"
// )

// Function to create JWT tokens with claims
func GenerateToken(email string, usrId int64) (string, error) {

	//  part1 Header.payload
	//// Create a new JWT token with claims
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": usrId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	// part3 signing
	tokenstring, err := tokenClaims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	// header.payload(claims).sign(encodedheader+encodedpayload+secretkey)
	return tokenstring, nil

}
