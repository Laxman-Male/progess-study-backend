package handler

import (
	"fmt"
	"net/http"
	"progress_tracker/database"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	// "golang.org/x/oauth2/authhandler"
)

func HandleLoginState(w http.ResponseWriter, r *http.Request) {

	// if r.Method==http.MethodGet{
	// name := r.URL.Query().Get("name")
	// email := r.URL.Query().Get("email")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	authHeader := r.Header.Get("Authorization")
	tokenStr := strings.TrimPrefix(authHeader, "Bearer")
	claims := jwt.MapClaims{}
	token, _ := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil
	})
	fmt.Println("token in login state", token)
	userId := claims["userId"].(string)
	_, err := database.DB.Exec(`update register set isLogin=1 where id= $1`, userId)
	if err != nil {
		http.Error(w, "could not update login state", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "User login state updated")
	// }

}
