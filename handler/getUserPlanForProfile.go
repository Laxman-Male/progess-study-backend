package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "progress_tracker/handler"
	"progress_tracker/service"
)

type Response struct {
	Count int      `json:"count"`
	Title []string `json:"title"`
	// Completed[i] is true once plan Title[i] has all weeks elapsed and every MCQ attempted.
	Completed []bool `json:"completed"`
}

func GetPlanForProfile(w http.ResponseWriter, r *http.Request) {
	email, ok := GetUserIDFromContext(r.Context())
	if !ok {
		fmt.Println("error in getting email for context")
		return
	}
	fmt.Println("email to get plan", email)
	count, title, completed, err := service.GetUserPlan(email)
	if err != nil {
		http.Error(w, "error getting plans", http.StatusInternalServerError)
		return
	}
	fmt.Println(count, "------------------")
	fmt.Println("title----------", title)
	response := Response{
		Count:     count,
		Title:     title,
		Completed: completed,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
