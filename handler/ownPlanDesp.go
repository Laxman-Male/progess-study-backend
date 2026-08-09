package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"progress_tracker/service"
	"strconv"
	"strings"
)

type Count struct {
	Count int    `db:"weekCount" json:"count"`
	Title string `db:"title" json:"title"`
}
type WeekCompletionStatus struct {
	CompletedWeeks int  `json:"completedWeeks"`
	TotalWeeks     int  `json:"totalWeeks"`
	TestEnabled    bool `json:"testEnabled"`
}

func GetOwnPlanDescription(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		title := r.URL.Query().Get("title")
		fmt.Println("-------title-----", title)
		lowerTitle := strings.ToLower(title)
		userEmail, ok := GetUserIDFromContext(r.Context())
		if !ok {
			return
		}
		plan, err := service.OwnPlanDescriptionService(lowerTitle, userEmail)
		if err != nil {
			return
		}
		var raw json.RawMessage

		err = json.Unmarshal([]byte(plan), &raw)
		if err != nil {
			http.Error(w, "failed to parse", http.StatusInternalServerError)
			return
		}
		fmt.Println("plan-------------------in handler", plan)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// json.NewEncoder(w).Encode(plan)
		w.Write([]byte(raw))
	}
}

func GetWeekPlanViewPlan(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		title := r.URL.Query().Get("title")
		fmt.Println("-------title-----", title)
		lowerTitle := strings.ToLower(title)
		userEmail, ok := GetUserIDFromContext(r.Context())
		if !ok {
			return
		}

		count, err := service.GetWeekPlanViewPlan(userEmail, lowerTitle)
		if err != nil {
			http.Error(w, "not getting count", http.StatusBadRequest)
		}
		ct := Count{
			Count: count,
			Title: title,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ct)

	}
}

func WeekCompleted(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		nowTimeStr := r.URL.Query().Get("nowTime")
		nowTime, err := strconv.ParseInt(nowTimeStr, 10, 64)
		if err != nil {
			http.Error(w, "error in converting nowTime into number", http.StatusBadRequest)
			return
		}

		title := r.URL.Query().Get("title")

		userEmail, ok := GetUserIDFromContext(r.Context())
		if !ok {
			return
		}

		completedWeeks, totalWeeks, err := service.WeekCompletedService(title, userEmail, nowTime)
		if err != nil {
			fmt.Println("error in getting week completion status")
			http.Error(w, "error in getting week completion status", http.StatusInternalServerError)
			return
		}
		status := WeekCompletionStatus{
			CompletedWeeks: completedWeeks,
			TotalWeeks:     totalWeeks,
			TestEnabled:    totalWeeks > 0 && completedWeeks >= totalWeeks,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(status)

	}
}
