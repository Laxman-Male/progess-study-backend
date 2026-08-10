package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	allstruct "progress_tracker/allStruct"

	// "progress_tracker/handler"

	"progress_tracker/service"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		var reg allstruct.RegInfo
		err := json.NewDecoder(r.Body).Decode(&reg)
		fmt.Println("enter into handler")
		if err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}
		userId, err := service.ValidateReg(reg)
		if err != nil {
			http.Error(w, "unable to insert value", http.StatusBadRequest)
			return

		}
		token, err := allstruct.GenerateJWT(userId)
		if err != nil {
			http.Error(w, "failed to generate token", http.StatusInternalServerError)
			return
		}
		// ✅ Respond with success message
		response := map[string]string{
			"message": "Registration successful!",
			"userId":  userId,
			"token":   token,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

		return
	}
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		// var isLogin=false
		var loginInfoUser allstruct.LoginFromUser
		err := json.NewDecoder(r.Body).Decode(&loginInfoUser)
		fmt.Println("login user info handler", loginInfoUser)
		if err != nil {
			http.Error(w, "invalid data recieve from front end ", http.StatusBadRequest)
			return
		}
		UserLoggedIn, err := service.ValidateLogin(loginInfoUser)
		if err != nil {
			http.Error(w, "check name ans password", http.StatusBadRequest)
			return

		}
		if UserLoggedIn != "" {
			fmt.Println("user id is =", UserLoggedIn)

		} else {
			fmt.Println("user id is blank =", UserLoggedIn)
		}
		token, err := allstruct.GenerateJWT(loginInfoUser.Email)
		if err != nil {
			fmt.Println("err in generate jwt")
		}
		response := map[string]string{
			"message": "login successful!",
			"userId":  UserLoggedIn,
			"token":   token,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

	}
}

func HandleLoggedOut(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var ForLogOut allstruct.LoginFromUser
		err := json.NewDecoder(r.Body).Decode(&ForLogOut)
		if err != nil {
			fmt.Println("error in decoding log out")
		}
		err = service.LogOut(ForLogOut.Email, ForLogOut.Name)
		if err != nil {

		}
	}
}
