package utils

import (
	"errors"
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

// Function to verify JWT tokens
func VerfiyToken(tokenString string) (int64, error) {

	// Parse the token with the secret key
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		//  check about SigningMethod  algo method that already impl in generate token
		//  by type checking
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected SigningMethod")
		}

		return secretKey, nil
	})
	// Check for verification errors
	if err != nil {
		return 0, errors.New("Could not parse token ")
	}

	// extract data from parsed token
	// based on algo and secret key .....
	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, errors.New("invalid token ")
	}

	// extract data from parsed token
	// cause we store claims in generate token as jwt.MapClaims => validate it by type checking  => type MapClaims map[string]interface{}
	mapClaimsData, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims  ")
	}

	//extract data claims
	// email := mapClaimsData["email"]
	userid := int64(mapClaimsData["userId"].(float64))
	return userid, nil

}
