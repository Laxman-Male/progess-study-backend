// service/user.go
package service

import (
	"fmt"
	allstruct "progress_tracker/allStruct"
	"progress_tracker/repo"
)

// Import the errors package for creating new errors
// allstruct "progress_tracker/allStruct" // Adjust this import path if your 'allStruct' package is named differently or located elsewhere

// User struct definition. It's good practice to have this in a central
// 'allstruct' or 'model' package if it's used across multiple packages.
// For now, I'll place it here for demonstration, assuming it's also in allstruct.
type User struct {
	ID    string `json:"userId"`
	Name  string `json:"name"`
	Email string `json:"email"`
	// Password string `json:"-"` // Use `json:"-"` to prevent marshalling password into JSON
}

// Mock Database: In a real application, this would be your connection to a database
// (e.g., PostgreSQL, MongoDB) and you'd have functions interacting with it.
// For now, it's an in-memory map.
// var usersDB = make(map[string]User) // Key: UserID, Value: User struct

// generateUserID creates a simple, unique ID for mock users.
// func generateUserID() string {
// 	return fmt.Sprintf("user_%d", time.Now().UnixNano())
// }

// ValidateReg simulates user registration logic.
// It checks for duplicate emails and "saves" the user to the mock DB.
// In a real app: Hash the password using bcrypt or similar before storing!
// func ValidateReg(reg allstruct.RegInfo) (string, error) {
// 	for _, user := range usersDB {
// 		if user.Email == reg.Email {
// 			return "", errors.New("email already registered")
// 		}
// 	}

// 	newUserID := generateUserID()
// 	newUser := User{
// 		ID:       newUserID,
// 		Name:     reg.Name,
// 		Email:    reg.Email,
// 		Password: reg.Password, // WARNING: Hash this in production!
// 	}
// 	usersDB[newUserID] = newUser // Save to mock DB
// 	fmt.Printf("Service: Registered new user: %s (ID: %s)\n", newUser.Email, newUser.ID)
// 	return newUserID, nil
// }

// ValidateLogin simulates user login logic.
// It checks if the email exists and if the password matches.
// func ValidateLogin(loginInfo allstruct.LoginFromUser) (string, error) {
// 	for _, user := range usersDB {
// 		if user.Email == loginInfo.Email {
// 			// WARNING: In production, compare hashed passwords using bcrypt!
// 			if user.Password == loginInfo.Password { // Mock password comparison
// 				fmt.Printf("Service: User %s logged in successfully.\n", user.Email)
// 				return user.ID, nil // Return the userID on successful login
// 			}
// 			return "", errors.New("invalid password")
// 		}
// 	}
// 	return "", errors.New("user not found")
// }

// GetUserByID retrieves a user's details from the mock database by their UserID.
func GetUserByID(userID string) (allstruct.GetProfileDetails, error) {
	user, err := repo.Profile(userID)
	if err != nil {
		fmt.Println("error in user.go service")
		return allstruct.GetProfileDetails{}, err
	}
	fmt.Println("in user.go", user)
	// if !ok {
	// 	return User{}, errors.New("user not found")
	// }
	// IMPORTANT: Do NOT return the password field in a real API response.
	// We've used `json:"-"` tag in the User struct, but explicitly zeroing it out
	// here ensures it's never accidentally exposed in Go code either.
	// user.Password = ""
	return user, nil
}

// LogOut (for JWT, this is usually just client-side token deletion.
// If you implement a server-side token blacklist, this is where it'd go.)
// func LogOut(email string, password string) error {
// 	// For this mock, no server-side action for logout is performed for JWTs.
// 	// Logout is primarily handled by the client removing the token.
// 	fmt.Printf("Service: Mock Logout called for user: %s\n", email)
// 	return nil // Simulate successful logout
// }

// Assuming your allstruct.RegInfo and allstruct.LoginFromUser are defined here
// or in allstruct/models.go
// Example (if they are not already defined in allstruct/models.go):
/*
type RegInfo struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginFromUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
*/
