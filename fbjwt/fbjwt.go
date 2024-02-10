package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {
	// Initialize the Firebase app
	opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase app: %v", err)
	}

	// Create a Firebase Auth client
	auth, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Failed to create Firebase Auth client: %v", err)
	}

	// Generate a custom token
	uid := "user123"
	claims := map[string]interface{}{
		"admin": true,
	}
	token, err := auth.CustomTokenWithClaims(context.Background(), uid, claims)
	if err != nil {
		log.Fatalf("Failed to generate custom token: %v", err)
	}

	fmt.Println("Custom token:", token)
}
