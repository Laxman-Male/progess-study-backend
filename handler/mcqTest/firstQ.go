package mcqtest

import (
	"encoding/json"
	"net/http"
	"progress_tracker/handler"
	mcqtest "progress_tracker/service/mcqTest"
	"strings"
)

type FirstMCQ struct {
	MCQID       int    `json:"mcqID"`
	MCQ         string `json:"mcq"`
	Options     string `json:"options"`
	IsAttempted bool   `json:"isAttempted"`
	// Done is true once every MCQ for the plan has been attempted.
	Done      bool `json:"done"`
	Total     int  `json:"total"`
	Attempted int  `json:"attempted"`
}

func FirstQuestion(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		title := r.URL.Query().Get("Title")
		title = strings.TrimSpace(title)
		title = strings.Trim(title, `"`)

		email, ok := handler.GetUserIDFromContext(r.Context())
		if !ok {
			return
		}

		Fmcq, counts, err := mcqtest.FirstQuestionService(title, email)
		if err != nil {
			http.Error(w, "error getting first question", http.StatusInternalServerError)
			return
		}

		response := FirstMCQ{
			MCQID:       Fmcq.MCQID,
			MCQ:         Fmcq.MCQ,
			Options:     Fmcq.Options,
			IsAttempted: Fmcq.IsAttempted,
			Done:        Fmcq.MCQ == "",
			Total:       counts.Total,
			Attempted:   counts.Attempted,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}

}
