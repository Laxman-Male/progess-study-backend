package handler

import (
	"encoding/json"
	"fmt"

	// "fmt"
	"net/http"
	// allstruct "progress_tracker/allStruct"
	"progress_tracker/service"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	// if r.Method == http.MethodGet {
	// 	// will come back here once i design the login pop up
	// 	// email := r.URL.Query().Get("email")
	// 	// name := r.URL.Query().Get("name")

	// 	if email == "" || name == "" {
	// 		http.Error(w, "required field", http.StatusBadRequest)
	// 	}
	// 	var userProfileD allstruct.GetProfileDetails
	// 	userProfileD.Email = email
	// 	userProfileD.Name = name
	// 	profile, err := service.ProfileDetails(userProfileD)
	// 	if err != nil {
	// 		fmt.Println("error in handler")
	// 	}
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(profile)
	// }
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// The JWTAuthMiddleware has already run and stored the userId in the request context.
	// We retrieve it here.
	userEmail, ok := GetUserIDFromContext(r.Context())
	if !ok {
		// This should theoretically not happen if middleware is correctly applied,
		// but it's a good safeguard.
		http.Error(w, "User ID not found in context (middleware error)", http.StatusInternalServerError)
		return
	}

	// Now, use this validated userId to fetch the user's data from your service/database.
	// Assuming service.GetUserByID is a function that retrieves user details from your DB.
	fmt.Println("in getUserid frm context", userEmail)
	user, err := service.GetUserByID(userEmail) // You'll need to implement this in your service package
	if err != nil {
		http.Error(w, "--User not found or database error--", http.StatusNotFound)
		return
	}

	// Prepare and send the user profile data as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user) // 'user' should be a struct that json.Encoder can handle (e.g., allstruct.User)

}
