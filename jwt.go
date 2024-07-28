package main

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

// This is a secret key used to sign the JWT. Hard-code for testing
var jwtKey = []byte("my_secret_key")

// GenerateJWT generates a JWT for a given username
func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		}
	}
	// This creates a new JWT token with the specified signing method 
	// (jwt.SigningMethodHS256, which is HMAC-SHA256) and the claims we defined earlier.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// This signs the token with the secret key (jwtKey) and converts it to a string. 
	// The resulting token string is stored in tokenString.
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWT validates the given JWT token string and returns the claims if valid
func ValidateJWT(tokenString string) (*Claims, error) {
    // Step 1: Create a new instance of Claims to hold the token's information
    claims := &Claims{}
    
	// Step 2: Parse the token string
    // This converts the token string into a structured format and fills the claims
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        // Provide the key used to sign the token
        return jwtKey, nil
    })

    // Step 3: Check for errors during parsing the token
    if err != nil {
        return nil, err
    }

    // Step 4: Check if the token is not valid
    if !token.Valid {
        return nil, errors.New("invalid token")
    }

    // Step 5: If everything is okay, return the claims and nil error
    return claims, nil
}
