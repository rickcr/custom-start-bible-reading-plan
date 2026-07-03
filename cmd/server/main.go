package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// corsMiddleware adds CORS headers to allow the Next.js frontend to call the API
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // The Next.js default port
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

		// Handle preflight OPTIONS request
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// UserResponse is the expected JSON response format
type UserResponse struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}

// userHandler processes the JWT from the Authorization header
func userHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Get the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
		return
	}

	// 2. Extract the token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// 3. A JWT has 3 parts separated by dots: header.payload.signature
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		http.Error(w, "Invalid JWT format", http.StatusBadRequest)
		return
	}

	// 4. Decode the payload (the middle part)
	// JWT uses base64url encoding, which doesn't always have standard base64 padding
	payloadStr := parts[1]
	// Add padding if necessary for Go's standard base64 decoder
	if pad := len(payloadStr) % 4; pad != 0 {
		payloadStr += strings.Repeat("=", 4-pad)
	}

	payloadBytes, err := base64.URLEncoding.DecodeString(payloadStr)
	if err != nil {
		log.Printf("Failed to decode token payload: %v", err)
		http.Error(w, "Failed to decode token", http.StatusBadRequest)
		return
	}

	// 5. Unmarshal the JSON payload to get the email
	var payload map[string]interface{}
	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
		log.Printf("Failed to unmarshal JSON payload: %v", err)
		http.Error(w, "Invalid token payload", http.StatusBadRequest)
		return
	}

	// 6. Extract the email field
	email, ok := payload["email"].(string)
	if !ok || email == "" {
		http.Error(w, "Email not found in token", http.StatusBadRequest)
		return
	}

	// 7. Return the response
	response := UserResponse{
		Email:   email,
		Message: fmt.Sprintf("Hello %s, your token was successfully processed by the Go backend!", email),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/user", userHandler)

	// Wrap the mux with the CORS middleware
	handler := corsMiddleware(mux)

	port := ":8080"
	fmt.Printf("Starting Go backend server on %s...\n", port)
	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
