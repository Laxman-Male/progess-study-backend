package mcqtest

import (
	"encoding/json"
	"net/http"
	allstruct "progress_tracker/allStruct"
	"progress_tracker/handler"
	mcqtest "progress_tracker/service/mcqTest"
	"strings"
)

type Response struct {
	MCQID       int    `json:"mcqID"`
	MCQ         string `json:"mcq"`
	Options     string `json:"options"`
	IsAttempted bool   `json:"isAttempted"`
	// Done is true once every MCQ for the plan has been attempted.
	Done      bool `json:"done"`
	Total     int  `json:"total"`
	Attempted int  `json:"attempted"`
}

func SubmitQuestion(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		var optTitlee allstruct.OptTitle
		var mcq allstruct.MCQFormat
		err := json.NewDecoder(r.Body).Decode(&optTitlee)
		if err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		optTitlee.Title = strings.TrimSpace(optTitlee.Title)
		optTitlee.Title = strings.Trim(optTitlee.Title, `"`)
		email, ok := handler.GetUserIDFromContext(r.Context())
		if !ok {
			return
		}

		mcq, counts, err := mcqtest.SubmitQuestionService(optTitlee.Option, email, optTitlee.Title, optTitlee.MCQID)
		if err != nil {
			http.Error(w, "error submitting answer", http.StatusInternalServerError)
			return
		}

		response := Response{
			MCQID:       mcq.MCQID,
			MCQ:         mcq.MCQ,
			Options:     mcq.Options,
			IsAttempted: mcq.IsAttempted,
			Done:        mcq.MCQ == "",
			Total:       counts.Total,
			Attempted:   counts.Attempted,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

	}
}
