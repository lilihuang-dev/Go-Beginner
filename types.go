package main

import "github.com/dgrijalva/jwt-go"

// Credentials struct to hold username and password
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claims struct to hold the JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}