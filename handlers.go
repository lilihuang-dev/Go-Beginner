package main

import (
	"log"
	"net/http"
	"encoding/json"
)

// RootHandler handles requests to the root URL
func RootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("RootHandler called")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the root path!"))
}


// LoginHandler handles requests to the login URL
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("LoginHandler called")
	// Read and parse the request body
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request payload"))
		return
	}

	// Placeholder for credential validation
	if creds.Username == "admin" && creds.Password == "password" {
		// Generate JWT token (omitted for brevity)
		token, err := GenerateJWT(creds.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(token))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid credentials"))
	}
}
