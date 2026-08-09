package mcqtest

import (
	"encoding/json"
	"net/http"
	"progress_tracker/handler"
	mcqtest "progress_tracker/service/mcqTest"
	"strings"
)

// QuizReview returns all MCQs for the user's plan with the selected option
// and correctness, so the frontend can show a colored correct/incorrect review.
func QuizReview(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		title := r.URL.Query().Get("Title")
		title = strings.TrimSpace(title)
		title = strings.Trim(title, `"`)

		email, ok := handler.GetUserIDFromContext(r.Context())
		if !ok {
			return
		}

		review, err := mcqtest.GetPlanReviewService(title, email)
		if err != nil {
			http.Error(w, "error getting quiz review", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(review)
	}
}
