package handler

import (
	"context"
	"fmt"
	"net/http"
	allstruct "progress_tracker/allStruct"
	"strings"
)

type userIDKey string

const contextUser userIDKey = "userID"

// CorsMiddleware provides standard CORS headers for all routes.
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

		// Handle preflight OPTIONS request globally for ALL endpoints
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// JWTAuthMiddleware validates the JWT from the Authorization header.
func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//get authorization header
		authoHeader := r.Header.Get("Authorization")
		if authoHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}
		//check for bearer prefix
		if !strings.HasPrefix(authoHeader, "Bearer ") {
			http.Error(w, "Invalid Authorization header format must be 'bearer token'", http.StatusUnauthorized)
			return
		}
		//Extract the token string
		tokenString := strings.TrimPrefix(authoHeader, "Bearer ")

		//validate the token using allstruct.validateJWT function
		userID, err := allstruct.ValidateJWT(tokenString)
		if err != nil {
			fmt.Printf("JWT Validation Error: %v\n", err) // Log the actual error for debugging
			http.Error(w, fmt.Sprintf("--Unauthorized: %s", err.Error()), http.StatusUnauthorized)
			return
		}
		// If the token is valid, store the userId in the request's context
		// This makes the userId available to the actual handler function
		c := context.WithValue(r.Context(), contextUser, userID)
		//this c value will get to the next calling function
		//r.WithContext(c) creates a new request with updated context c.
		//next.ServeHTTP passes control to the next handler with this enriched request.
		//like this  middleware shares data or signals with downstream handlers.
		next.ServeHTTP(w, r.WithContext(c))

	})
}

// GetUserIDFromContext is a helper to easily retrieve the userId from the request context.
func GetUserIDFromContext(ctx context.Context) (string, bool) {
	userEmail, ok := ctx.Value(contextUser).(string)
	fmt.Println("user id in jwr atuth middle", userEmail)
	return userEmail, ok
}
