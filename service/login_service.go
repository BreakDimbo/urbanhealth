package service

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		log.Printf("decode creds json error: %s\n", err)
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	token, err := authenticate(creds)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		log.Printf("Invalid credentials: %s", err)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func authenticate(creds Credentials) (string, error) {
	if creds.Password != "abc" {
		return "", errors.New("invalid credentials")
	}
	return "token", nil
}
